package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var _token token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			_token = token.Token{Type: token.EQ, Literal: "=="}
		} else {
			_token = newToken(token.ASSIGN, l.char)
		}
	case ';':
		_token = newToken(token.SEMICOLON, l.char)
	case '(':
		_token = newToken(token.LPAREN, l.char)
	case ')':
		_token = newToken(token.RPAREN, l.char)
	case ',':
		_token = newToken(token.COMMA, l.char)
	case '+':
		_token = newToken(token.PLUS, l.char)
	case '-':
		_token = newToken(token.MINUS, l.char)
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			_token = token.Token{Type: token.NOT_EQ, Literal: "!="}
		} else {
			_token = newToken(token.EXCLAMATION, l.char)
		}
	case '*':
		_token = newToken(token.ASTERISK, l.char)
	case '/':
		_token = newToken(token.FOREWARD_SLASH, l.char)
	case '<':
		_token = newToken(token.LT, l.char)
	case '>':
		_token = newToken(token.GT, l.char)
	case '{':
		_token = newToken(token.LBRACE, l.char)
	case '}':
		_token = newToken(token.RBRACE, l.char)
	case 0:
		_token = newToken(token.EOF, l.char)
	default:
		if isLetter(l.char) {
			_token.Literal = l.readIdentifier()
			_token.Type = token.LookupIdent(_token.Literal)
			return _token
		} else if isDigit(l.char) {
			_token.Type = token.INT
			_token.Literal = l.readNumber()
			return _token
		} else {
			_token = newToken(token.ILLEGAL, l.char)
		}
	}

	l.readChar()

	return _token
}

func (l *Lexer) readIdentifier() string {

	startPosition := l.position

	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) readNumber() string {

	startPosition := l.position

	for isDigit(l.char) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	if tokenType == token.EOF {
		return token.Token{Type: tokenType, Literal: ""}
	}
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
