package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//식별자 + 리터럴
	IDENT = "IDENT"
	INT   = "INT"

	//연산자
	ASSIGN = "="
	PLUS   = "+"
	BANG   = "!"
	MINUS  = "-"
	SLASH  = "/"
	EQ     = "=="
	NOT_EQ = "!="

	LT = "<"
	GT = ">"

	//구분자
	COMMA     = ","
	SEMICOLON = ";"

	ASTERISK = "*"
	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"

	//키워드
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE  = "true"
	FALSE = "false"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
