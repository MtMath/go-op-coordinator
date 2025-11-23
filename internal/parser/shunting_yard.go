package parser

import (
    "fmt"
)

var precedence = map[string]int{
    "+": 1,
    "-": 1,
    "*": 2,
    "/": 2,
}

func ShuntingYard(tokens []Token) ([]Token, error) {
    output := []Token{}
    stack := []Token{}

    for _, tok := range tokens {
        switch tok.Type {

        case TokenNumber:
            output = append(output, tok)

        case TokenOperator:
            for len(stack) > 0 {
                top := stack[len(stack)-1]
                if top.Type == TokenOperator && precedence[top.Value] >= precedence[tok.Value] {
                    output = append(output, top)
                    stack = stack[:len(stack)-1]
                } else {
                    break
                }
            }
            stack = append(stack, tok)

        case TokenLParen, TokenLBracket, TokenLBrace:
            stack = append(stack, tok)

        case TokenRParen, TokenRBracket, TokenRBrace:
            matchErr := fmt.Errorf("agrupamento inválido")
            var open TokenType

            if tok.Type == TokenRParen {
                open = TokenLParen
            } else if tok.Type == TokenRBracket {
                open = TokenLBracket
            } else {
                open = TokenLBrace
            }

            found := false
            for len(stack) > 0 {
                top := stack[len(stack)-1]
                stack = stack[:len(stack)-1]

                if top.Type == open {
                    found = true
                    break
                }

                output = append(output, top)
            }

            if !found {
                return nil, matchErr
            }
        }
    }

    for len(stack) > 0 {
        if stack[len(stack)-1].Type >= TokenLParen {
            return nil, fmt.Errorf("agrupamento não fechado")
        }
        output = append(output, stack[len(stack)-1])
        stack = stack[:len(stack)-1]
    }

    return output, nil
}
