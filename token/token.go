package token

// this allows us to use many different types of tokens
// without having to define a new type for each one
// we can use a string to represent the token type
// not ideal, but will revisit later
type TokenType string

// TODO: future improvement:
// also track line number and column number
// or file name and position
// for better error reporting
// initialize the lexer with an io.Reader and the filename
type Token struct {
	Type    TokenType
	Literal string
}

// TODO: add other token types
const (
	// special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// literals
	IDENTIFIER = "IDENTIFIER" // function or variable name
	INT        = "INT"        // 123456...
	STRING = "STRING" // "foobar"

	// operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	EQ     = "=="
	NOT_EQ = "!="

	// delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LEFT_PAREN  = "("
	RIGHT_PAREN = ")"
	LEFT_BRACE  = "{"
	RIGHT_BRACE = "}"

	// keywords

	// Keywords
	FUNCTION = "FUNCTION"
	CONST    = "CONST"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"function": FUNCTION,
	"const":    CONST,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
