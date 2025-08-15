package main

import (
	"bufio"

	"log"
	"os"
	"strings"

	"github.com/Rishi-Mishra0704/crazy-lang/lexer"
	"github.com/Rishi-Mishra0704/crazy-lang/parser"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Invalid usage")
	}
	filename := os.Args[1]
	if !strings.HasSuffix(filename, ".crazy") {
		log.Fatal("Invalid file type: must end with .crazy")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens, err := lexer.Lexer(line)
		if err != nil {
			log.Fatal("Error lexing file:", err)
		}
		err = parser.Parse(tokens)
		if err != nil {
			log.Fatal("Error parsing file:", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}
}
