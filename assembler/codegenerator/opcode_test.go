package codegenerator

import (
	"testing"
)

func TestNewOpCode(t *testing.T) {
	opcode := NewOpCode([2]byte{0x12, 0x34})
	expected := [2]byte{0x12, 0x34}
	if opcode.Bytes != expected {
		t.Errorf("NewOpCode([2]byte{0x12, 0x34}) = %v; expected %v", opcode.Bytes, expected)
	}
}

func TestNewOpCodePNNN(t *testing.T) {
	opcode := NewOpCodePNNN(0x1, 0x234)
	expected := [2]byte{0x12, 0x34}
	if opcode.Bytes != expected {
		t.Errorf("NewOpCodePNNN(0x1, 0x234) = %v; expected %v", opcode.Bytes, expected)
	}

	opcode = NewOpCodePNNN(0xA, 0xDEA)
	expected = [2]byte{0xAD, 0xEA}
	if opcode.Bytes != expected {
		t.Errorf("NewOpCodePNNN(0xA, 0xDEA) = %v; expected %v", opcode.Bytes, expected)
	}
}

func TestNewOpCodePXNN(t *testing.T) {
	opcode := NewOpCodePXNN(0x6, 0xF, 0xAB)
	expected := [2]byte{0x6F, 0xAB}
	if opcode.Bytes != expected {
		t.Errorf("NewOpCodePXNN(0x6, 0xF, 0xAB) = %v; expected %v", opcode.Bytes, expected)
	}

	opcode = NewOpCodePXNN(0xC, 0xA, 0xFE)
	expected = [2]byte{0xCA, 0xFE}
	if opcode.Bytes != expected {
		t.Errorf("NewOpCodePXNN(0xC, 0xA, 0xFE) = %v; expected %v", opcode.Bytes, expected)
	}
}

func TestNewOpCodePXYS(t *testing.T) {
	opcode := NewOpCodePXYS(0x8, 0x2, 0x4, 0xE)
	expected := [2]byte{0x82, 0x4E}
	if opcode.Bytes != expected {
		t.Errorf("NewOpCodePXYS(0x8, 0x2, 0x4, 0xE) = %v; expected %v", opcode.Bytes, expected)
	}

	opcode = NewOpCodePXYS(0x9, 0x3, 0x0, 0x0)
	expected = [2]byte{0x93, 0x00}
	if opcode.Bytes != expected {
		t.Errorf("NewOpCodePXYS(0x9, 0x3, 0x0, 0x0) = %v; expected %v", opcode.Bytes, expected)
	}
}

func TestNewOpCodePXSS(t *testing.T) {
	opcode := NewOpCodePXSS(0xE, 0x3, 0x9E)
	expected := [2]byte{0xE3, 0x9E}
	if opcode.Bytes != expected {
		t.Errorf("NewOpCodePXSS(0xE, 0x3, 0x9E) = %v; expected %v", opcode.Bytes, expected)
	}

	opcode = NewOpCodePXSS(0xF, 0x0, 0x65)
	expected = [2]byte{0xF0, 0x65}
	if opcode.Bytes != expected {
		t.Errorf("NewOpCodePXSS(0xF, 0x0, 0x65) = %v; expected %v", opcode.Bytes, expected)
	}
}
