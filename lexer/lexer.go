// Package lexer 词法分析器
package lexer

type Lexer struct {
	input string

	// 所输入字符串中的当前位置（指向当前字符）
	position int

	// 所输入字符串中的当前读取位置（指向当前字符之后的一个字符）
	readPosition int

	// 当前正在查看的字符
	ch byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

// 读取 input 中的下一个字符，并前移其在 input 中的位置
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
