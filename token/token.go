package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// 表示遇到未知的词法单元或字符
	ILLEGAL = "ILLEGAL"
	// 表示文件结尾
	EOF = "EOF"

	// 标识符 + 字面量
	IDENT = "IDENT"
	INT   = "INT"

	// 运算符
	ASSIGN = "="
	PLUS   = "+"

	// 分隔符
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
