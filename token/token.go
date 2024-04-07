package token

type Token struct {
	Type    Type
	Literal string
}

type Type string

const (
	EOF     = "EOF"
	ILLEGAL = "illegal"
	MESSAGE = "message"
	IDENT   = "identifier"
	LBRACE  = "{"
	RBRACE  = "}"
)

var keywords = map[string]Type{
	"message": MESSAGE,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
