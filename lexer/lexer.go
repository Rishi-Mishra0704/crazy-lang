package lexer

import (
	"strings"
	"unicode"

	"github.com/Rishi-Mishra0704/crazy-lang/tokens"
)

type Token struct {
	Type  string
	Value string
}

func Lexer(line string) ([]Token, error) {
	var tokensList []Token
	line = strings.TrimSpace(line)
	if line == "" {
		return tokensList, nil
	}

	runes := []rune(line)
	i := 0
	for i < len(runes) {
		ch := runes[i]

		if unicode.IsSpace(ch) && ch != '\n' {
			i++
			continue
		}

		if ch == '\n' {
			tokensList = append(tokensList, Token{Type: tokens.NEWLINE, Value: "\n"})
			i++
			continue
		}

		// Handle Windows-style \r\n newlines
		if ch == '\r' {
			if i+1 < len(runes) && runes[i+1] == '\n' {
				tokensList = append(tokensList, Token{Type: tokens.NEWLINE, Value: "\r\n"})
				i += 2
			} else {
				tokensList = append(tokensList, Token{Type: tokens.NEWLINE, Value: "\r"})
				i++
			}
			continue
		}

		// comments
		if ch == '#' {
			tok, newI := readComment(runes, i)
			tokensList = append(tokensList, tok)
			i = newI
			continue
		}

		// multi-char ops
		if tok, newI, ok := readMultiCharOperator(runes, i); ok {
			tokensList = append(tokensList, tok)
			i = newI
			continue
		}

		// strings
		if ch == '"' || ch == '\'' {
			tok, newI, err := readString(runes, i)
			if err != nil {
				return nil, err
			}
			tokensList = append(tokensList, tok)
			i = newI
			continue
		}

		// identifiers / keywords
		if unicode.IsLetter(ch) || ch == '_' {
			tok, newI := readIdentifierOrKeyword(runes, i)
			tokensList = append(tokensList, tok)
			i = newI
			continue
		}

		// numbers
		if unicode.IsDigit(ch) {
			tok, newI := readNumber(runes, i)
			tokensList = append(tokensList, tok)
			i = newI
			continue
		}

		// single-char ops / delimiters
		if tok, ok := readSingleCharOperator(ch); ok {
			tokensList = append(tokensList, tok)
			i++
			continue
		}

		// unknown
		tokensList = append(tokensList, Token{Type: tokens.ILLEGAL, Value: string(ch)})
		i++

	}

	return tokensList, nil
}
