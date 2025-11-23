package parser

import (
	"reflect"
	"testing"
)

func TestTokenizerSimple(t *testing.T) {
	input := "3 + 5"
	expected := []Token{
		{Type: TokenNumber, Value: "3"},
		{Type: TokenOperator, Value: "+"},
		{Type: TokenNumber, Value: "5"},
	}

	tokens, err := Tokenize(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Fatalf("expected %v, got %v", expected, tokens)
	}
}

func TestTokenizerParentheses(t *testing.T) {
	input := "(10 - 2) * 3"
	expected := []Token{
		{Type: TokenLParen, Value: "("},
		{Type: TokenNumber, Value: "10"},
		{Type: TokenOperator, Value: "-"},
		{Type: TokenNumber, Value: "2"},
		{Type: TokenRParen, Value: ")"},
		{Type: TokenOperator, Value: "*"},
		{Type: TokenNumber, Value: "3"},
	}

	tokens, err := Tokenize(input)
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Fatalf("expected %v, got %v", expected, tokens)
	}
}

func TestTokenizerBracketsBraces(t *testing.T) {
	input := "[3 + {2 * (1 + 1)}]"
	expected := []Token{
		{Type: TokenLBracket, Value: "["},
		{Type: TokenNumber, Value: "3"},
		{Type: TokenOperator, Value: "+"},
		{Type: TokenLBrace, Value: "{"},
		{Type: TokenNumber, Value: "2"},
		{Type: TokenOperator, Value: "*"},
		{Type: TokenLParen, Value: "("},
		{Type: TokenNumber, Value: "1"},
		{Type: TokenOperator, Value: "+"},
		{Type: TokenNumber, Value: "1"},
		{Type: TokenRParen, Value: ")"},
		{Type: TokenRBrace, Value: "}"},
		{Type: TokenRBracket, Value: "]"},
	}

	tokens, err := Tokenize(input)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Fatalf("\nexpected:\n%v\ngot:\n%v", expected, tokens)
	}
}

func TestTokenizerDecimalNumbers(t *testing.T) {
	input := "3.14 * 2.0"
	expected := []Token{
		{Type: TokenNumber, Value: "3.14"},
		{Type: TokenOperator, Value: "*"},
		{Type: TokenNumber, Value: "2.0"},
	}

	tokens, err := Tokenize(input)
	if err != nil {
		t.Fatalf("error: %v", err)
	}

	if !reflect.DeepEqual(tokens, expected) {
		t.Fatalf("expected %v, got %v", expected, tokens)
	}
}

func TestTokenizerInvalidNumber(t *testing.T) {
	input := "3.1.4"

	_, err := Tokenize(input)
	if err == nil {
		t.Fatalf("expected error for invalid number, but got none")
	}
}
