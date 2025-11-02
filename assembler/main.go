package main

import (
	"fmt"
	"os"

	"github.com/xoesae/chip8/assembler/codegenerator"
	"github.com/xoesae/chip8/assembler/lexer"
	"github.com/xoesae/chip8/assembler/parser"
)

func mapToSlice(opcodes map[uint32]byte) []byte {
	var _max uint32
	for addr := range opcodes {
		if addr > _max {
			_max = addr
		}
	}

	bytes := make([]byte, _max+1)
	for addr, val := range opcodes {
		bytes[addr] = val
	}

	return bytes
}

func assemble(input string) []byte {
	// lexer
	lxr := lexer.NewLexer(input)
	tokens := lxr.Lex()

	// parser
	psr := parser.NewParser(tokens)
	expressions := psr.Parse()

	// code generator
	cg := codegenerator.NewCodeGenerator()
	opcodes := cg.Generate(expressions)

	return mapToSlice(opcodes)
}

func writeBinary(bin []byte, filename string) error {
	return os.WriteFile(filename, bin, 0644)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: compile <input> <output>")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	input, err := os.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	bin := assemble(string(input))
	err = writeBinary(bin, outputFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
