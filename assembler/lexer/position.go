package lexer

import "fmt"

type Position struct {
	line   int
	column int
}

func (p *Position) Format() string {
	return fmt.Sprintf("%d:%d", p.line, p.column)
}

func (p *Position) NextColumn() {
	p.column++
}

func (p *Position) PreviousColumn() {
	p.column--
}

func (p *Position) NextLine() {
	p.line++
	p.column = 1
}

func (p *Position) PreviousLine() {
	p.line--
	p.column = 1
}
