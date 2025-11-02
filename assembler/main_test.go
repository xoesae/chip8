package main

import (
	"os"
	"testing"
)

func TestMapToSlice(t *testing.T) {
	opcodes := map[uint32]byte{
		0: 0xAB,
		2: 0xCD,
		1: 0xEF,
	}
	out := mapToSlice(opcodes)
	expected := []byte{0xAB, 0xEF, 0xCD}

	if len(out) != len(expected) {
		t.Fatalf("mapToSlice: len = %d, expected %d", len(out), len(expected))
	}

	for i := range expected {
		if out[i] != expected[i] {
			t.Errorf("mapToSlice[%d] = 0x%X, expected 0x%X", i, out[i], expected[i])
		}
	}
}

func TestWriteBinary(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5}
	filename := "output.bin"

	defer os.Remove(filename)

	err := writeBinary(data, filename)
	if err != nil {
		t.Fatalf("error on write file: %v", err)
	}

	file, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("error on read file: %v", err)
	}

	if string(file) != string(data) {
		t.Errorf("wrong file contents")
	}
}

func TestAssembleMinimal(t *testing.T) {
	src := "CLS\nRET"
	bin := assemble(src)

	expected := []byte{0x00, 0xE0, 0x00, 0xEE}
	if len(bin) != len(expected) {
		t.Fatalf("assemble: len = %d, expected %d", len(bin), len(expected))
	}

	for i := range expected {
		if bin[i] != expected[i] {
			t.Errorf("assemble[%d] = 0x%X, expected 0x%X", i, bin[i], expected[i])
		}
	}
}
