package parser

import (
	"testing"
)

func evalExpression(t *testing.T, expr string) float64 {
	tokens, err := Tokenize(expr)
	if err != nil {
		t.Fatalf("error in tokenization: %v", err)
	}

	rpn, err := ShuntingYard(tokens)
	if err != nil {
		t.Fatalf("error in shunting yard: %v", err)
	}

	ast, err := BuildAST(rpn)
	if err != nil {
		t.Fatalf("error building AST: %v", err)
	}

	result, err := EvalASTLocal(ast)
	if err != nil {
		t.Fatalf("error in eval: %v", err)
	}

	return result
}

func TestEvalSimpleAdd(t *testing.T) {
	got := evalExpression(t, "3 + 4")
	expected := 7.0

	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestEvalSimpleMul(t *testing.T) {
	got := evalExpression(t, "2 * 5")
	expected := 10.0

	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestEvalOrderOfOperations(t *testing.T) {
	// 3 + 2 * 5 = 3 + 10 = 13
	got := evalExpression(t, "3 + 2 * 5")
	expected := 13.0

	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestEvalParentheses(t *testing.T) {
	// (3 + 2) * 5 = 5 * 5 = 25
	got := evalExpression(t, "(3 + 2) * 5")
	expected := 25.0

	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestEvalNestedGroups(t *testing.T) {
	// (10 - [2 + {3 * (2 + 1)}]) = (10 - [2 + {3*3}]) = (10 - [2 + 9]) = (10 - 11) = -1
	got := evalExpression(t, "(10 - [2 + {3 * (2 + 1)}])")
	expected := -1.0

	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestEvalDivision(t *testing.T) {
	got := evalExpression(t, "10 / 2")
	expected := 5.0

	if got != expected {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

func TestEvalDivisionByZero(t *testing.T) {
	tokens, _ := Tokenize("10 / 0")
	rpn, _ := ShuntingYard(tokens)
	ast, _ := BuildAST(rpn)

	_, err := EvalASTLocal(ast)
	if err == nil {
		t.Fatalf("expected division by zero error, but got none")
	}
}
