package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/xoesae/chip8/assembler/lexer"
	"github.com/xoesae/chip8/assembler/parser"
)

func main() {
	input, err := os.ReadFile("assembler/input.txt")
	if err != nil {
		panic(err)
	}

	// lexer
	lxr := lexer.NewLexer(strings.NewReader(string(input)))
	tokens := lxr.Lex()

	//for _, t := range tokens {
	//	fmt.Println(t.Format())
	//}

	// parser
	psr := parser.NewParser(tokens)
	expressions := psr.Parse()

	for _, e := range expressions {
		fmt.Println(e)
	}
}
