package ast

import (
	"bytes"
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

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
