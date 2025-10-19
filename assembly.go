package main

import (
	"fmt"
	"os"

	"github.com/xoesae/chip8/assembly/parser"
)

func main() {
	p := parser.Parser{}

	fmt.Printf("FROM: %s | TO: %s \n", os.Args[1], os.Args[2])

	p.ParseFile(os.Args[1], os.Args[2])
}
