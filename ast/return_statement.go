package ast

import (
	"monkey/token"
)

// ReturnStatement LetStatement Let 语句
type ReturnStatement struct {
	Token       token.Token // token.LET 词法单元
	ReturnValue Expression
}

// 语句节点
func (rs *ReturnStatement) statementNode() {
	// todo
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
