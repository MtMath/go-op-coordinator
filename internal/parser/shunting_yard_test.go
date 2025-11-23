package parser

import (
	"reflect"
	"testing"
)

func TestShuntingYardSimple(t *testing.T) {
	tokens, _ := Tokenize("3 + 4")
	rpn, err := ShuntingYard(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []Token{
		{Type: TokenNumber, Value: "3"},
		{Type: TokenNumber, Value: "4"},
		{Type: TokenOperator, Value: "+"},
	}

	if !reflect.DeepEqual(rpn, expected) {
		t.Fatalf("expected %v, got %v", expected, rpn)
	}
}

func TestShuntingYardPrecedence(t *testing.T) {
	tokens, _ := Tokenize("3 + 4 * 2")
	rpn, err := ShuntingYard(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []Token{
		{Type: TokenNumber, Value: "3"},
		{Type: TokenNumber, Value: "4"},
		{Type: TokenNumber, Value: "2"},
		{Type: TokenOperator, Value: "*"},
		{Type: TokenOperator, Value: "+"},
	}

	if !reflect.DeepEqual(rpn, expected) {
		t.Fatalf("expected %v, got %v", expected, rpn)
	}
}

func TestShuntingYardParentheses(t *testing.T) {
	tokens, _ := Tokenize("(3 + 4) * 2")
	rpn, err := ShuntingYard(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []Token{
		{Type: TokenNumber, Value: "3"},
		{Type: TokenNumber, Value: "4"},
		{Type: TokenOperator, Value: "+"},
		{Type: TokenNumber, Value: "2"},
		{Type: TokenOperator, Value: "*"},
	}

	if !reflect.DeepEqual(rpn, expected) {
		t.Fatalf("expected %v, got %v", expected, rpn)
	}
}

func TestShuntingYardBracesAndBrackets(t *testing.T) {
	tokens, _ := Tokenize("[3 + {4 * (2 - 1)}]")
	rpn, err := ShuntingYard(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []Token{
		{Type: TokenNumber, Value: "3"},
		{Type: TokenNumber, Value: "4"},
		{Type: TokenNumber, Value: "2"},
		{Type: TokenNumber, Value: "1"},
		{Type: TokenOperator, Value: "-"},
		{Type: TokenOperator, Value: "*"},
		{Type: TokenOperator, Value: "+"},
	}

	if !reflect.DeepEqual(rpn, expected) {
		t.Fatalf("\nexpected:\n%v\ngot:\n%v", expected, rpn)
	}
}
