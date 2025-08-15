package parser

import (
	"fmt"

	"github.com/Rishi-Mishra0704/crazy-lang/lexer"
)

func Parse(tokens []lexer.Token) error {
	fmt.Println(tokens[0].Value)
	return nil
}
