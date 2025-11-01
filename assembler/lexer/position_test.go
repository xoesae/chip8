package lexer

import (
	"testing"
)

func TestPositionFormat(t *testing.T) {
	p := Position{line: 3, column: 7}
	got := p.format()
	want := "3:7"
	if got != want {
		t.Errorf("format() = %s; want %s", got, want)
	}
}

func TestPositionNextColumn(t *testing.T) {
	p := Position{line: 1, column: 1}
	p.nextColumn()
	if p.column != 2 {
		t.Errorf("nextColumn() column = %d; want 2", p.column)
	}
}

func TestPositionPreviousColumn(t *testing.T) {
	p := Position{line: 1, column: 3}
	p.previousColumn()
	if p.column != 2 {
		t.Errorf("previousColumn() column = %d; want 2", p.column)
	}
}

func TestPositionNextLine(t *testing.T) {
	p := Position{line: 5, column: 4}
	p.nextLine()
	if p.line != 6 || p.column != 1 {
		t.Errorf("nextLine() = %d:%d; want 6:1", p.line, p.column)
	}
}

func TestPositionPreviousLine(t *testing.T) {
	p := Position{line: 5, column: 8}
	p.previousLine()
	if p.line != 4 || p.column != 1 {
		t.Errorf("previousLine() = %d:%d; want 4:1", p.line, p.column)
	}
}
