package coordinator

import (
	"context"
	"fmt"
	"notask/op-coordinator/internal/clients"
	"notask/op-coordinator/internal/dispatcher"
	"notask/op-coordinator/internal/parser"

	coordpb "notask/op-coordinator/api/coordpb"
)

type CoordinatorService struct {
	coordpb.UnimplementedCoordinatorServiceServer

	AddClient *clients.AddClient
	SubClient *clients.SubClient
	MulClient *clients.MulClient
	DivClient *clients.DivClient

	Dispatcher *dispatcher.Dispatcher
}

func NewCoordinatorService(addAddr, subAddr, mulAddr, divAddr string) *CoordinatorService {

	disp := dispatcher.NewDispatcher(
		addAddr,
		subAddr,
		mulAddr,
		divAddr,
	)

	return &CoordinatorService{
		Dispatcher: disp,
	}
}

func (s *CoordinatorService) evalDistributed(node *parser.ASTNode) (float64, error) {
	if node.Type == parser.NumberNode {
		return node.EvalLiteral()
	}

	left, err := s.evalDistributed(node.Left)
	if err != nil {
		return 0, err
	}

	right, err := s.evalDistributed(node.Right)
	if err != nil {
		return 0, err
	}

	result, err := s.Dispatcher.Dispatch(node.Value, left, right)
	if err != nil {
		return 0, fmt.Errorf("erro ao chamar dispatcher: %w", err)
	}

	return result, nil
}

func (s *CoordinatorService) Evaluate(ctx context.Context, req *coordpb.EvaluateRequest) (*coordpb.EvaluateResponse, error) {

	tokens, err := parser.Tokenize(req.Expression)
	if err != nil {
		return nil, fmt.Errorf("tokenizer: %w", err)
	}

	rpn, err := parser.ShuntingYard(tokens)
	if err != nil {
		return nil, fmt.Errorf("shunting-yard: %w", err)
	}

	ast, err := parser.BuildAST(rpn)
	if err != nil {
		return nil, fmt.Errorf("AST: %w", err)
	}

	result, err := s.evalDistributed(ast)
	if err != nil {
		return nil, fmt.Errorf("eval: %w", err)
	}

	rpnStr := ""
	for _, t := range rpn {
		rpnStr += t.Value + " "
	}

	return &coordpb.EvaluateResponse{
		Result: result,
		Rpn:    rpnStr,
	}, nil
}
