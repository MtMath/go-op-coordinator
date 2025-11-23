package parser

import (
	"testing"
)

func TestASTSimple(t *testing.T) {
	rpn := []Token{
		{Type: TokenNumber, Value: "3"},
		{Type: TokenNumber, Value: "4"},
		{Type: TokenOperator, Value: "+"},
	}

	ast, err := BuildAST(rpn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ast.Value != "+" {
		t.Fatalf("expected operator '+', got: %s", ast.Value)
	}

	if ast.Left.Value != "3" || ast.Right.Value != "4" {
		t.Fatalf("incorrect tree: %v", ast)
	}
}

func TestASTComplex(t *testing.T) {
	rpn := []Token{
		{Type: TokenNumber, Value: "3"},
		{Type: TokenNumber, Value: "4"},
		{Type: TokenOperator, Value: "+"},
		{Type: TokenNumber, Value: "2"},
		{Type: TokenOperator, Value: "*"},
	}

	ast, err := BuildAST(rpn)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if ast.Value != "*" {
		t.Fatalf("expected operator '*', got: %s", ast.Value)
	}

	if ast.Left.Value != "+" {
		t.Fatalf("expected operator '+', got: %s", ast.Left.Value)
	}

	if ast.Left.Left.Value != "3" || ast.Left.Right.Value != "4" {
		t.Fatalf("incorrect left subtree")
	}

	if ast.Right.Value != "2" {
		t.Fatalf("expected number '2', got: %s", ast.Right.Value)
	}
}

func TestASTInvalid(t *testing.T) {
	rpn := []Token{
		{Type: TokenOperator, Value: "+"},
	}

	_, err := BuildAST(rpn)
	if err == nil {
		t.Fatalf("expected error for invalid RPN, but got none")
	}
}
