package lexer

import (
	"reflect"
	"testing"

	"github.com/xoesae/chip8/assembler/token"
)

func TestLex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name:  "Empty input",
			input: "",
			expected: []token.Token{
				token.EOF{},
			},
		},
		{
			name:  "Instructions with comments",
			input: "LD I, #num1      ; Point I to the `1` sprite",
			expected: []token.Token{
				token.Instruction{Value: "LD"},
				token.Register{Value: "I"},
				token.LabelOperand{Value: "#num1"},
				token.EOF{},
			},
		},
		{
			name:  "Label and JP",
			input: "end         JP #end          ; loop forever",
			expected: []token.Token{
				token.Label{Value: "end"},
				token.Instruction{Value: "JP"},
				token.LabelOperand{Value: "#end"},
				token.EOF{},
			},
		},
		{
			name:  "Directives org and db",
			input: "org $200\nnum1\n            db $20 $60 $20 $20 $70",
			expected: []token.Token{
				token.Directive{Value: "org"},
				token.NumericLiteral{Value: 512},
				token.Label{Value: "num1"},
				token.Directive{Value: "db"},
				token.NumericLiteral{Value: 32},
				token.NumericLiteral{Value: 96},
				token.NumericLiteral{Value: 32},
				token.NumericLiteral{Value: 32},
				token.NumericLiteral{Value: 112},
				token.EOF{},
			},
		},
		{
			name:  "DRW Instruction",
			input: "DRW V0, V1, 5",
			expected: []token.Token{
				token.Instruction{Value: "DRW"},
				token.Register{Value: "V0"},
				token.Register{Value: "V1"},
				token.NumericLiteral{Value: 5},
				token.EOF{},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			lexer := NewLexer(tc.input)
			result := lexer.Lex()
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Test '%s' failed. Expected %+v, result %+v", tc.name, tc.expected, result)
			}
		})
	}
}
