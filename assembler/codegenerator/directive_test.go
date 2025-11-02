package codegenerator

import (
	"testing"

	"github.com/xoesae/chip8/assembler/token"
)

func TestProcessDirectiveOrg(t *testing.T) {
	cg := NewCodeGenerator()

	expression := token.Expression{
		token.Directive{Value: string(token.Org)},
		token.NumericLiteral{Value: 0x400},
	}

	cg.processDirective(expression)

	if cg.addressCounter.pos != 0x400 {
		t.Errorf("org: pos = 0x%X, expected 0x400", cg.addressCounter.pos)
	}
}

func TestProcessDirectiveDb(t *testing.T) {
	cg := NewCodeGenerator()

	expression := token.Expression{
		token.Directive{Value: string(token.Db)},
		token.NumericLiteral{Value: 0x20},
		token.NumericLiteral{Value: 0x60},
		token.NumericLiteral{Value: 0xA0},
	}

	cg.processDirective(expression)

	expected := map[uint32]byte{
		0x0: 0x20,
		0x1: 0x60,
		0x2: 0xA0,
	}

	for addr, val := range expected {
		got := cg.opcodes[addr]
		if got != val {
			t.Errorf("db: wrote 0x%X in 0x%X, expected 0x%X", got, addr, val)
		}
	}

	if cg.addressCounter.pos != 0x3 {
		t.Errorf("db: pos = 0x%X, expected 0x3", cg.addressCounter.pos)
	}
}
