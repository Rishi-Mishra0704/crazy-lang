package lexer

import (
	"fmt"
	"unicode"

	"github.com/Rishi-Mishra0704/crazy-lang/tokens"
)

// keyword lookup
var keywords = map[string]string{
	"let":  tokens.LET,
	"if":   tokens.IF,
	"else": tokens.ELSE,
	"for":  tokens.FOR,
	"fun":  tokens.FUN,
}

// multi-char operators
var multiCharOps = map[string]string{
	"==": tokens.EQ,
	"!=": tokens.NEQ,
	"<=": tokens.LTE,
	">=": tokens.GTE,
}

// single-char operators/delimiters
var singleCharOps = map[rune]string{
	'+': tokens.PLUS,
	'-': tokens.MINUS,
	'*': tokens.MULT,
	'/': tokens.DIV,
	'{': tokens.LBRACE,
	'}': tokens.RBRACE,
	'(': tokens.LPAREN,
	')': tokens.RPAREN,
	',': tokens.COMMA,
	';': tokens.SEMICOLON,
	'=': tokens.ASSIGN,
	'<': tokens.LT,
	'>': tokens.GT,
}

// readString handles both single and double-quoted strings
func readString(runes []rune, start int) (Token, int, error) {
	quote := runes[start]
	i := start + 1
	for i < len(runes) {
		if runes[i] == quote {
			value := string(runes[start+1 : i])
			return Token{Type: tokens.STRING, Value: value}, i + 1, nil
		}
		i++
	}
	return Token{}, i, fmt.Errorf("unterminated string starting at position %d", start)
}

// readIdentifierOrKeyword returns either an IDENT or a KEYWORD token
func readIdentifierOrKeyword(runes []rune, start int) (Token, int) {
	i := start
	for i < len(runes) && (unicode.IsLetter(runes[i]) || unicode.IsDigit(runes[i]) || runes[i] == '_') {
		i++
	}
	value := string(runes[start:i])
	if kwType, isKW := keywords[value]; isKW {
		return Token{Type: kwType, Value: value}, i
	}
	return Token{Type: tokens.IDENT, Value: value}, i
}

// readNumber parses integers and decimals
func readNumber(runes []rune, start int) (Token, int) {
	i := start
	hasDot := false
	for i < len(runes) {
		if runes[i] == '.' {
			if hasDot {
				break
			}
			hasDot = true
		} else if !unicode.IsDigit(runes[i]) {
			break
		}
		i++
	}
	return Token{Type: tokens.NUMBER, Value: string(runes[start:i])}, i
}

// readComment grabs everything from '#' to end of line
func readComment(runes []rune, start int) (Token, int) {
	i := start
	for i < len(runes) && runes[i] != '\n' && runes[i] != '\r' {
		i++
	}
	value := string(runes[start:i])
	return Token{Type: tokens.COMMENT, Value: value}, i
}

// readMultiCharOperator checks for operators like '==', '!=', '<=', '>='
func readMultiCharOperator(runes []rune, start int) (Token, int, bool) {
	if start+1 < len(runes) {
		twoChar := string(runes[start : start+2])
		if tokType, ok := multiCharOps[twoChar]; ok {
			return Token{Type: tokType, Value: twoChar}, start + 2, true
		}
	}
	return Token{}, start, false
}

// readSingleCharOperator returns single-character operator/delimiter tokens
func readSingleCharOperator(ch rune) (Token, bool) {
	if tokType, ok := singleCharOps[ch]; ok {
		return Token{Type: tokType, Value: string(ch)}, true
	}
	return Token{}, false
}
