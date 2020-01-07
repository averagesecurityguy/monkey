package lexer

import (
	"testing"

	"token"
)

type TokenTest struct {
	expectedType    token.TokenType
	expectedLiteral rune
}

func TestNextToken(t *testing.T) {
	input := "=+(){},+😀\n+"

	tests := []TokenTest{
		{token.ASSIGN, '='},
		{token.PLUS, '+'},
		{token.LPAREN, '('},
		{token.RPAREN, ')'},
		{token.LBRACE, '{'},
		{token.RBRACE, '}'},
		{token.COMMA, ','},
		{token.SEMICOLON, ';'},
		{token.SMILEY, '😀'},
		{token.NEWLINE, '\n'},
		{token.ASSIGN, '='},
	}

	l := NewLexer(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Line %d Char: %d - Expected %q found %q", l.lineNumber, l.linePosition, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Line: %d Char: %d - Expected %q found %q", l.lineNumber, l.linePosition, tt.expectedLiteral, tok.Literal)
		}
	}
}
