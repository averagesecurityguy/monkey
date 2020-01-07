package lexer

import (
    "token"
)

type Lexer struct {
    input []rune
    position int // current position in input (points to current char)
    readPosition int // current reading position in input (after current char)
    lineNumber int // current line number
    linePosition int // current position in the line
    ch rune // current char under examination
}

func NewLexer(input string) *Lexer {
    l := &Lexer{input: []rune(input)}
    l.readRune()

    return l
}

func (l *Lexer) readRune() {
    if l.readPosition >= len(l.input) {
        l.ch = 0
    } else {
        l.ch = l.input[l.readPosition]
    }

    if l.ch == '\n' {
        l.lineNumber += 1
        l.linePosition = 0
    }

    l.position = l.readPosition
    l.readPosition += 1
    l.linePosition += 1
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    switch l.ch {
    case '=':
        tok = token.Token{token.ASSIGN, l.ch}
    case ';':
        tok = token.Token{token.SEMICOLON, l.ch}
    case '(':
        tok = token.Token{token.LPAREN, l.ch}
    case ')':
        tok = token.Token{token.RPAREN, l.ch}
    case ',':
        tok = token.Token{token.COMMA, l.ch}
    case '+':
        tok = token.Token{token.PLUS, l.ch}
    case '{':
        tok = token.Token{token.LBRACE, l.ch}
    case '}':
        tok = token.Token{token.RBRACE, l.ch}
    case '😀':
        tok = token.Token{token.SMILEY, l.ch}
    case '\n':
        tok = token.Token{token.NEWLINE, l.ch}
    case '0':
        tok = token.Token{token.EOF, rune(0)}
    default:
        tok = token.Token{token.ILLEGAL, rune(0)}
    }

    l.readRune()

    return tok
}
