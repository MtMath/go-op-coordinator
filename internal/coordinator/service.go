package coordinator

import (
	"context"
	"fmt"
	"notask/op-coordinator/internal/clients"
	"notask/op-coordinator/internal/dispatcher"
	"notask/op-coordinator/internal/parser"
	"sync"

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

	var left, right float64
	var errLeft, errRight error
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		left, errLeft = s.evalDistributed(node.Left)
	}()

	go func() {
		defer wg.Done()
		right, errRight = s.evalDistributed(node.Right)
	}()

	wg.Wait()

	if errLeft != nil {
		return 0, errLeft
	}

	if errRight != nil {
		return 0, errRight
	}

	result, err := s.Dispatcher.Dispatch(node.Value, left, right)
	if err != nil {
		return 0, fmt.Errorf("error calling dispatcher: %w", err)
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
