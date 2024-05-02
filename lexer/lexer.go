// Package lexer 词法分析器
package lexer

import "monkey/token"

type Lexer struct {
	input string

	// 所输入字符串中的当前位置（指向当前字符）
	position int

	// 所输入字符串中的当前读取位置（指向当前字符之后的一个字符）
	readPosition int

	// 当前正在查看的字符
	ch byte
}

// New 创建词法分析器 Lexer 实例
func New(input string) *Lexer {
	// 初始化 Lexer
	l := &Lexer{input: input}
	l.readChar()

	return l
}

// 读取 input 中的下一个字符，并前移其在 input 中的位置
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		// 如果读取到字符串尾部，设置 l.ch 为 0
		// 这其实是 NUL 字符的 ASCII 编码
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

// 读入下一个 token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifer()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

// 读取标识符
func (l *Lexer) readIdentifer() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// 读取数字
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// 跳过空白符
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 判断字符是否是字母
func isLetter(ch byte) bool {
	// _ 也会被视为字母
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 判断字符是否是数字
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// 创建一个 Token 实例
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
