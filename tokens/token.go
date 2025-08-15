package tokens

const (
	// Keywords
	LET    = "LET"
	IF     = "IF"
	ELSE   = "ELSE"
	FOR    = "FOR"
	FUN    = "FUN"
	RETURN = "RETURN"

	// Operators
	PLUS   = "+"
	MINUS  = "-"
	MULT   = "*"
	DIV    = "/"
	ASSIGN = "="
	EQ     = "=="
	NEQ    = "!="
	LT     = "<"
	GT     = ">"
	LTE    = "<="
	GTE    = ">="

	// Delimiters
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	COMMA     = ","
	SEMICOLON = ";"

	// Literals
	IDENT  = "IDENTIFIER"
	NUMBER = "NUMBER"
	STRING = "STRING"
	BOOL   = "BOOLEAN"

	// Others
	COMMENT = "COMMENT"
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
	NEWLINE = "NEWLINE"
)
