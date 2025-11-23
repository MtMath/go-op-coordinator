package parser

import (
	"fmt"
	"strconv"
)

func EvalASTLocal(n *ASTNode) (float64, error) {
	if n.Type == NumberNode {
		return strconv.ParseFloat(n.Value, 64)
	}

	left, err := EvalASTLocal(n.Left)
	if err != nil {
		return 0, err
	}

	right, err := EvalASTLocal(n.Right)
	if err != nil {
		return 0, err
	}

	switch n.Value {
	case "+":
		return left + right, nil
	case "-":
		return left - right, nil
	case "*":
		return left * right, nil
	case "/":
		if right == 0 {
			return 0, fmt.Errorf("divisão por zero")
		}
		return left / right, nil
	}

	return 0, fmt.Errorf("operador inválido: %s", n.Value)
}
