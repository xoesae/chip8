package parser

import (
	"reflect"
	"testing"

	"github.com/xoesae/chip8/assembler/token"
)

func TestIsOperand(t *testing.T) {
	cases := []struct {
		tkn      token.Token
		expected bool
	}{
		{token.Register{Value: "V1"}, true},
		{token.NumericLiteral{Value: 42}, true},
		{token.LabelOperand{Value: "num1"}, true},
		{token.Label{Value: "start"}, false},
		{token.Instruction{Value: "JP"}, false},
		{token.Directive{Value: "org"}, false},
		{token.EOF{}, false},
	}

	for _, c := range cases {
		if got := isOperand(c.tkn); got != c.expected {
			t.Errorf("%T: got %v, expected %v", c.tkn, got, c.expected)
		}
	}
}

func TestParseLabel(t *testing.T) {
	psr := NewParser([]token.Token{token.Label{Value: "start"}})

	expression, ok := psr.parseLabel()
	if !ok {
		t.Fatal("parseLabel: expected true, got false")
	}

	expected := token.Expression{token.Label{Value: "start"}}
	if !reflect.DeepEqual(expression, expected) {
		t.Errorf("parseLabel: got %#v, expected %#v", expression, expected)
	}
}

func TestParseOrgDirective(t *testing.T) {
	tokens := []token.Token{token.Directive{Value: "org"}, token.NumericLiteral{Value: 512}}
	psr := NewParser(tokens)

	expression, ok := psr.parseDirective()
	if !ok {
		t.Fatal("parseDirective: expected true, got false")
	}

	expected := token.Expression{token.Directive{Value: "org"}, token.NumericLiteral{Value: 512}}
	if !reflect.DeepEqual(expression, expected) {
		t.Errorf("parseDirective: got %#v, expected %#v", expression, expected)
	}
}

func TestParseDbDirective(t *testing.T) {
	tokens := []token.Token{token.Directive{Value: "db"}, token.NumericLiteral{Value: 512}, token.NumericLiteral{Value: 512}, token.NumericLiteral{Value: 512}}
	psr := NewParser(tokens)

	expression, ok := psr.parseDirective()
	if !ok {
		t.Fatal("parseDirective: expected true, got false")
	}

	expected := token.Expression{token.Directive{Value: "db"}, token.NumericLiteral{Value: 512}, token.NumericLiteral{Value: 512}, token.NumericLiteral{Value: 512}}
	if !reflect.DeepEqual(expression, expected) {
		t.Errorf("parseDirective: got %#v, expected %#v", expression, expected)
	}
}

func TestParseInstructionOperands(t *testing.T) {
	tokens := []token.Token{
		token.Instruction{Value: "LD"},
		token.Register{Value: "V1"},
		token.NumericLiteral{Value: 42},
		token.LabelOperand{Value: "num1"},
	}
	psr := NewParser(tokens)

	expression, ok := psr.parseInstruction()
	if !ok {
		t.Fatal("parseInstruction: expected true, got false")
	}

	expected := token.Expression{
		token.Instruction{Value: "LD"},
		token.Register{Value: "V1"},
		token.NumericLiteral{Value: 42},
		token.LabelOperand{Value: "num1"},
	}
	if !reflect.DeepEqual(expression, expected) {
		t.Errorf("parseInstruction: got %#v, expected %#v", expression, expected)
	}
}

func TestParseMixedTokens(t *testing.T) {
	tokens := []token.Token{
		token.Directive{Value: "org"}, token.NumericLiteral{Value: 512},
		token.Label{Value: "start"},
		token.Instruction{Value: "JP"}, token.LabelOperand{Value: "end"},
		token.EOF{},
	}
	psr := NewParser(tokens)

	expressions := psr.Parse()
	expected := []token.Expression{
		{token.Directive{Value: "org"}, token.NumericLiteral{Value: 512}},
		{token.Label{Value: "start"}},
		{token.Instruction{Value: "JP"}, token.LabelOperand{Value: "end"}},
	}
	if !reflect.DeepEqual(expressions, expected) {
		t.Errorf("Parse: got %#v, expected %#v", expressions, expected)
	}
}
