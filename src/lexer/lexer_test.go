package lexer

import (
	"testing"

	"token"
)

type TokenTest struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNumbers(t *testing.T) {
	input := `
    let ten = 10 ;
    let tenhex = 0x0a ;
    let hex = 0xBB;
    `

	tests := []TokenTest{
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "tenhex"},
		{token.ASSIGN, "="},
		{token.HEX, "0a"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "hex"},
		{token.ASSIGN, "="},
		{token.HEX, "BB"},
		{token.SEMICOLON, ";"},
	}

	runTests(t, input, tests)
}

func TestSingleOperators(t *testing.T) {
	input := `!- / *5;
    5<10>5;
    `

	tests := []TokenTest{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}

	runTests(t, input, tests)
}

func TestDoubleOperators(t *testing.T) {
	input := `
    10 == 10;
    10 != 9;
    10 <= 10;
    10 >= 10;
    `

	tests := []TokenTest{
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.LT_EQ, "<="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.GT_EQ, ">="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
	}

	runTests(t, input, tests)
}

func TestStructure(t *testing.T) {
	input := `let five = 5;
    let ten = 10;
    let smile = 😀;

    let add = fn(x, y) {
        x + y;
    };

    let result = add(five, ten);

    if (5 < 10) {
        return true;
    } else {
        return false;
    }
    `

	tests := []TokenTest{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "smile"},
		{token.ASSIGN, "="},
		{token.SMILEY, "😀"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.EOF, "EOF"},
	}

	runTests(t, input, tests)
}

func runTests(t *testing.T, input string, tests []TokenTest) {

	l := NewLexer(input)

	for _, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Char: %d - Expected %q found %q", l.position, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Char: %d - Expected %q found %q", l.position, tt.expectedLiteral, tok.Literal)
		}
	}
}
