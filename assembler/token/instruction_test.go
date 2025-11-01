package token

import (
	"testing"
)

func TestInstructionKind(t *testing.T) {
	i := Instruction{Value: "JP"}
	got := i.Kind()
	want := "Instruction"
	if got != want {
		t.Errorf("Kind() = %s; want %s", got, want)
	}
}

func TestInstructionFormat(t *testing.T) {
	i := Instruction{Value: "CALL"}
	got := i.Format()
	want := "Instruction(CALL)"
	if got != want {
		t.Errorf("Format() = %s; want %s", got, want)
	}
}

func TestIsInstructionTrue(t *testing.T) {
	opcodes := []string{
		"CLS", "RET", "JP", "CALL", "SE", "SNE", "LD", "ADD", "SUB",
		"SUBN", "OR", "AND", "XOR", "SHR", "SHL", "RND", "DRW",
		"SKP", "SKNP",
	}
	for _, op := range opcodes {
		if !IsInstruction(op) {
			t.Errorf("IsInstruction(%q) = false; want true", op)
		}
	}
}

func TestIsInstructionFalse(t *testing.T) {
	notOpcodes := []string{
		"jmp", "Load", "JPX", "set", "ADD1", "jump", "JPP", "", "XORR", "abcdef",
	}
	for _, word := range notOpcodes {
		if IsInstruction(word) {
			t.Errorf("IsInstruction(%q) = true; want false", word)
		}
	}
}
