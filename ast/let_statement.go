package ast

import (
	"monkey/token"
)

// LetStatement Let 语句
type LetStatement struct {
	Token token.Token // token.LET 词法单元
	Name  *Identifier // 绑定的标识符
	Value Expression  // 产生值的表达式
}

// 语句节点
func (ls *LetStatement) statementNode() {
	// todo
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
