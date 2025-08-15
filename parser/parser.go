package parser

import (
	"fmt"

	"github.com/Rishi-Mishra0704/crazy-lang/lexer"
)

func Parse(tokens []lexer.Token) error {
	for _, t := range tokens {
		fmt.Printf("Token(Type=%s, Value=%s)\n", t.Type, t.Value)
	}
	return nil
}
