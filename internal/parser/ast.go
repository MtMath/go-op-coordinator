package parser

import (
	"fmt"
	"strconv"
)

type ASTNodeType int

const (
	NumberNode ASTNodeType = iota
	OperatorNode
)

type ASTNode struct {
	Type  ASTNodeType
	Value string

	Left  *ASTNode
	Right *ASTNode
}

func BuildAST(rpn []Token) (*ASTNode, error) {
	stack := []*ASTNode{}

	for _, tok := range rpn {
		switch tok.Type {

		case TokenNumber:
			stack = append(stack, &ASTNode{
				Type:  NumberNode,
				Value: tok.Value,
			})

		case TokenOperator:
			if len(stack) < 2 {
				return nil, fmt.Errorf("expressão inválida")
			}

			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			stack = append(stack, &ASTNode{
				Type:  OperatorNode,
				Value: tok.Value,
				Left:  left,
				Right: right,
			})
		}
	}

	if len(stack) != 1 {
		return nil, fmt.Errorf("expressão inválida")
	}

	return stack[0], nil
}

func (n *ASTNode) EvalLiteral() (float64, error) {
	if n.Type != NumberNode {
		return 0, fmt.Errorf("não é número")
	}
	return strconv.ParseFloat(n.Value, 64)
}
