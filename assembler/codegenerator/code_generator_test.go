package codegenerator

import (
	"testing"

	"github.com/xoesae/chip8/assembler/token"
)

func TestNewCodeGenerator(t *testing.T) {
	cg := NewCodeGenerator()

	if cg.addressCounter.pos != 0 {
		t.Errorf("pos = %d, expected 0", cg.addressCounter.pos)
	}

	if cg.labels == nil {
		t.Errorf("labels not initialized")
	}

	if cg.opcodes == nil {
		t.Errorf("opcodes not initialized")
	}
}

func TestAppendOpcodeAdvancesAddress(t *testing.T) {
	cg := NewCodeGenerator()

	opcode := &OpCode{Bytes: [2]byte{0xAA, 0xBB}}
	cg.appendOpcode(opcode)

	if cg.opcodes[0] != 0xAA {
		t.Errorf("MSB not set on addr 0, got 0x%X", cg.opcodes[0])
	}
	if cg.opcodes[1] != 0xBB {
		t.Errorf("LSB not set on addr 1, got 0x%X", cg.opcodes[1])
	}
	if cg.addressCounter.pos != 2 {
		t.Errorf("addressCounter.pos = %d, expected 2", cg.addressCounter.pos)
	}
}

func TestMustAsPanicsWrongType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("mustAs not panic")
		}
	}()
	mustAs[token.Register](token.NumericLiteral{Value: 5})
}

func TestGenerateBasicInstructions(t *testing.T) {
	cg := NewCodeGenerator()

	expressions := []token.Expression{
		{token.Instruction{Value: "CLS"}},
		{token.Instruction{Value: "RET"}},
		{token.Instruction{Value: "SE"}, token.Register{Value: "V2"}, token.NumericLiteral{Value: 0xAB}},
	}

	opcodes := cg.Generate(expressions)

	expected := map[uint32]byte{0: 0x00, 1: 0xE0, 2: 0x00, 3: 0xEE, 4: 0x32, 5: 0xAB}
	for addr, val := range expected {
		got, ok := opcodes[addr]
		if !ok {
			t.Errorf("opcode not found on 0x%X", addr)
		}

		if got != val {
			t.Errorf("opcode on addr 0x%X: got 0x%X, expected 0x%X", addr, got, val)
		}
	}
}

func TestProcessLabelPanicsRepeated(t *testing.T) {
	cg := NewCodeGenerator()
	cg.labels["start"] = 100
	expression := []token.Expression{
		{token.Label{Value: "start"}},
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("processLabel not panic")
		}
	}()

	cg.Generate(expression)
}
