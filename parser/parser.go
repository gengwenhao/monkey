package parser

import (
	"monkey/lexer"
	"monkey/token"
)

// Parser 语法分析器
type Parser struct {
	l *lexer.Lexer // 指向词法分析器实例

	curToken  token.Token // 指向当前词法单元
	peekToken token.Token // 指向下一个词法单元
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// 读取两个词法单元，以设置 curToken 和 peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func parseProgram() {
}
