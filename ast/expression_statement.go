package ast

import "monkey/token"

type ExpressionStatement struct {
	// 该表达式中的第一个词法单元
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {

}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
