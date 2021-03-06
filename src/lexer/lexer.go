package lexer

import (
	"fmt"
	"unicode"

	"token"
)

type Lexer struct {
	input        []rune
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           rune // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: []rune(input)}
	l.readRune()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case ';':
		tok = token.Token{token.SEMICOLON, string(l.ch)}
	case '(':
		tok = token.Token{token.LPAREN, string(l.ch)}
	case ')':
		tok = token.Token{token.RPAREN, string(l.ch)}
	case ',':
		tok = token.Token{token.COMMA, string(l.ch)}
	case '+':
		tok = token.Token{token.PLUS, string(l.ch)}
	case '{':
		tok = token.Token{token.LBRACE, string(l.ch)}
	case '}':
		tok = token.Token{token.RBRACE, string(l.ch)}
	case '-':
		tok = token.Token{token.MINUS, string(l.ch)}
	case '*':
		tok = token.Token{token.ASTERISK, string(l.ch)}
	case '/':
		tok = token.Token{token.SLASH, string(l.ch)}
	case '0':
		if l.isTwoCharOp() {
			l.readRune()
			hex := string(l.readHex())
			tok = token.Token{token.HEX, hex}
			return tok
		}
	case '=':
		if l.isTwoCharOp() {
			tok = token.Token{token.EQ, "=="}
		} else {
			tok = token.Token{token.ASSIGN, string(l.ch)}
		}
	case '!':
		if l.isTwoCharOp() {
			tok = token.Token{token.NOT_EQ, "!="}
		} else {
			tok = token.Token{token.BANG, string(l.ch)}
		}
	case '<':
		if l.isTwoCharOp() {
			tok = token.Token{token.LT_EQ, "<="}
		} else {
			tok = token.Token{token.LT, string(l.ch)}
		}
	case '>':
		if l.isTwoCharOp() {
			tok = token.Token{token.GT_EQ, ">="}
		} else {
			tok = token.Token{token.GT, string(l.ch)}
		}
	case '😀':
		tok = token.Token{token.SMILEY, string(l.ch)}
	case 0:
		tok = token.Token{token.EOF, "EOF"}
	default:
		switch {
		case unicode.IsLetter(l.ch):
			ident := string(l.readIdentifier())
			tok = token.Token{token.LookupIdent(ident), ident}
			return tok
		case unicode.IsDigit(l.ch):
			digit := string(l.readNumber())
			tok = token.Token{token.INT, digit}
			return tok
		default:
			tok = token.Token{token.ILLEGAL, string(l.ch)}
		}
	}

	l.readRune()

	return tok
}

func (l *Lexer) readRune() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekRune() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) isTwoCharOp() bool {
	op := []rune{l.ch, l.peekRune()}

	switch string(op) {
	case "==":
		l.readRune()
		return true
	case "!=":
		l.readRune()
		return true
	case "<=":
		l.readRune()
		return true
	case ">=":
		l.readRune()
		return true
	case "0x":
		l.readRune()
		return true
	default:
		return false
	}
}

func (l *Lexer) readIdentifier() []rune {
	position := l.position
	for unicode.IsLetter(l.ch) {
		l.readRune()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() []rune {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readRune()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readHex() []rune {
	position := l.position
	fmt.Println(string(l.ch))
	for unicode.IsDigit(l.ch) || unicode.ToLower(l.ch) == 'a' || unicode.ToLower(l.ch) == 'b' ||
		unicode.ToLower(l.ch) == 'c' || unicode.ToLower(l.ch) == 'd' || unicode.ToLower(l.ch) == 'e' ||
		unicode.ToLower(l.ch) == 'f' {
		l.readRune()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.ch) {
		l.readRune()
	}
}
