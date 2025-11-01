package lexer

import "fmt"

type Position struct {
	line   int
	column int
}

func (p *Position) format() string {
	return fmt.Sprintf("%d:%d", p.line, p.column)
}

func (p *Position) nextColumn() {
	p.column++
}

func (p *Position) previousColumn() {
	p.column--
}

func (p *Position) nextLine() {
	p.line++
	p.column = 1
}

func (p *Position) previousLine() {
	p.line--
	p.column = 1
}
