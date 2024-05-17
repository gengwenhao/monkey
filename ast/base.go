package ast

import (
	"bytes"
	"monkey/token"
)

// Node 是 Statement 和 Expression 必须实现的接口
type Node interface {
	// TokenLiteral 这个方法仅仅用于调试和测试
	TokenLiteral() string // TokenLiteral 存储了字符串源码（关联的词法单元字面量）

	String() string // String 方法既可以在调试时打印 AST 节点，也可以用来比较 AST 节点
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// Identifier 标识符
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

func (i *Identifier) String() string {
	return i.Value
}

// IntegerLiteral 数值字面量类型
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {
	//
}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}
