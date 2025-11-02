package codegenerator

import (
	"testing"

	"github.com/xoesae/chip8/assembler/token"
)

func TestProcessLabel(t *testing.T) {
	cg := NewCodeGenerator()
	expression := token.Expression{token.Label{Value: "start"}}

	err := cg.processLabel(expression)
	if err != nil {
		t.Errorf("processLabel returned an error: %s", err)
	}

	addr, ok := cg.labels["start"]
	if !ok {
		t.Fatal("Label 'start' not found after processing label")
	}

	if addr != 0 {
		t.Errorf("Label address = 0x%X, expected 0x0", addr)
	}
}

func TestProcessLabelRepeated(t *testing.T) {
	cg := NewCodeGenerator()
	cg.labels["LOOP"] = 0x400
	expression := token.Expression{token.Label{Value: "LOOP"}}

	err := cg.processLabel(expression)

	if err == nil {
		t.Errorf("Repeated label should have failed")
	}
}
