package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// token types
const (
	ILLEGAL = "ILLEGAL" // トークンや文字が未知である場合
	EOF     = "EOF"     // end of file の略。入力の終わりを表す

	INDENT = "INDENT"
	INT    = "INT"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
