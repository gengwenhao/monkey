package ast

import "monkey/token"

// Node 是 Statement 和 Expression 必须实现的接口
type Node interface {
	TokenLiteral() string
}

// Statement 语句
type Statement interface {
	Node
	statementNode()
}

// Expression 表达式
type Expression interface {
	Node
	expressionNode()
}

// Program 是语法分析器生成的每个 AST 的根节点
type Program struct {
	// 每个有效的 Monkey 程序都是一系列
	// 位于 Program.Statements 中的语句 Statement
	Statements []Statement
}

// TokenLiteral 词法单元的字面量
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement Let 语句
type LetStatement struct {
	Token token.Token // token.LET 词法单元
	Name  *Identifier
	Value Expression
}

// 语句节点
func (ls *LetStatement) statementNode() {
	// todo
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token // token.IDENT 词法单元
	Value string
}

func (i *Identifier) expressionNode() {
	// todo
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
