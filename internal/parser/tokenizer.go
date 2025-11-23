package parser

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	TokenNumber TokenType = iota
	TokenOperator
	TokenLParen   // (
	TokenRParen   // )
	TokenLBracket // [
	TokenRBracket // ]
	TokenLBrace   // {
	TokenRBrace   // }
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(expr string) ([]Token, error) {
	tokens := []Token{}
	runes := []rune(expr)
	n := len(runes)

	for i := 0; i < n; {
		ch := runes[i]

		if unicode.IsSpace(ch) {
			i++
			continue
		}

		if unicode.IsDigit(ch) || ch == '.' {
			start := i
			dotCount := 0

			for i < n && (unicode.IsDigit(runes[i]) || runes[i] == '.') {
				if runes[i] == '.' {
					dotCount++
					if dotCount > 1 {
						return nil, fmt.Errorf("número inválido na posição %d", i)
					}
				}
				i++
			}

			tokens = append(tokens, Token{
				Type:  TokenNumber,
				Value: string(runes[start:i]),
			})
			continue
		}

		switch ch {
		case '+', '-', '*', '/':
			tokens = append(tokens, Token{Type: TokenOperator, Value: string(ch)})
		case '(':
			tokens = append(tokens, Token{Type: TokenLParen, Value: "("})
		case ')':
			tokens = append(tokens, Token{Type: TokenRParen, Value: ")"})
		case '[':
			tokens = append(tokens, Token{Type: TokenLBracket, Value: "["})
		case ']':
			tokens = append(tokens, Token{Type: TokenRBracket, Value: "]"})
		case '{':
			tokens = append(tokens, Token{Type: TokenLBrace, Value: "{"})
		case '}':
			tokens = append(tokens, Token{Type: TokenRBrace, Value: "}"})
		default:
			return nil, fmt.Errorf("Invalid: %c", ch)
		}

		i++
	}

	return tokens, nil
}
