package codegenerator

import (
	"testing"

	"github.com/xoesae/chip8/assembler/token"
)

func reg(name string) token.Register      { return token.Register{Value: name} }
func num(n uint32) token.NumericLiteral   { return token.NumericLiteral{Value: n} }
func lbl(name string) token.LabelOperand  { return token.LabelOperand{Value: name} }
func instr(name string) token.Instruction { return token.Instruction{Value: name} }

var labels = LabelMap{"loop": 0x567, "func": 0x444}

func TestCLSGenerate(t *testing.T) {
	expected := [2]byte{0x00, 0xE0}
	opcode := CLS{}.generate()
	if opcode.Bytes != expected {
		t.Errorf("CLS: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestRETGenerate(t *testing.T) {
	expected := [2]byte{0x00, 0xEE}
	opcode := RET{}.generate()
	if opcode.Bytes != expected {
		t.Errorf("RET: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestJPGenerateNumeric(t *testing.T) {
	expr := token.Expression{instr("JP"), num(0x234)}
	expected := [2]byte{0x12, 0x34}
	opcode := JP{expression: expr, labels: LabelMap{}}.generate()
	if opcode.Bytes != expected {
		t.Errorf("JP num: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestJPGenerateLabel(t *testing.T) {
	expr := token.Expression{instr("JP"), lbl("loop")}
	expected := [2]byte{0x15, 0x67}
	opcode := JP{expression: expr, labels: labels}.generate()
	if opcode.Bytes != expected {
		t.Errorf("JP label: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestJPGenerateV0Numeric(t *testing.T) {
	expr := token.Expression{instr("JP"), reg("V0"), num(0x222)}
	expected := [2]byte{0xB2, 0x22}
	opcode := JP{expression: expr, labels: LabelMap{}}.generate()
	if opcode.Bytes != expected {
		t.Errorf("JP V0, num: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestCALLGenerateNumeric(t *testing.T) {
	expr := token.Expression{instr("CALL"), num(0x789)}
	expected := [2]byte{0x27, 0x89}
	opcode := CALL{expression: expr, labels: LabelMap{}}.generate()
	if opcode.Bytes != expected {
		t.Errorf("CALL num: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestCALLGenerateLabel(t *testing.T) {
	expr := token.Expression{instr("CALL"), lbl("func")}
	expected := [2]byte{0x24, 0x44}
	opcode := CALL{expression: expr, labels: labels}.generate()
	if opcode.Bytes != expected {
		t.Errorf("CALL label: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSEGenerateNumeric(t *testing.T) {
	expr := token.Expression{instr("SE"), reg("V2"), num(0xAB)}
	expected := [2]byte{0x32, 0xAB}
	opcode := SE{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SE Vx, NN: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSEGenerateRegister(t *testing.T) {
	expr := token.Expression{instr("SE"), reg("V4"), reg("V5")}
	expected := [2]byte{0x54, 0x50}
	opcode := SE{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SE Vx, Vy: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSNEGenerateNumeric(t *testing.T) {
	expr := token.Expression{instr("SNE"), reg("V2"), num(0xCD)}
	expected := [2]byte{0x42, 0xCD}
	opcode := SNE{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SNE Vx, NN: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSNEGenerateRegister(t *testing.T) {
	expr := token.Expression{instr("SNE"), reg("V6"), reg("V7")}
	expected := [2]byte{0x96, 0x70}
	opcode := SNE{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SNE Vx, Vy: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestLDGenerateVxNN(t *testing.T) {
	expr := token.Expression{instr("LD"), reg("V8"), num(0x66)}
	expected := [2]byte{0x68, 0x66}
	opcode := LD{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("LD Vx,NN: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestLDGenerateVxVy(t *testing.T) {
	expr := token.Expression{instr("LD"), reg("V8"), reg("V9")}
	expected := [2]byte{0x88, 0x90}
	opcode := LD{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("LD Vx, Vy: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestLDGenerateIAddr(t *testing.T) {
	expr := token.Expression{instr("LD"), reg("I"), num(0x200)}
	expected := [2]byte{0xA2, 0x00}
	opcode := LD{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("LD I,addr: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestADDGenerateVxNN(t *testing.T) {
	expr := token.Expression{instr("ADD"), reg("V2"), num(0xEF)}
	expected := [2]byte{0x72, 0xEF}
	opcode := ADD{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("ADD Vx,NN: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestADDGenerateVxVy(t *testing.T) {
	expr := token.Expression{instr("ADD"), reg("V1"), reg("V3")}
	expected := [2]byte{0x81, 0x34}
	opcode := ADD{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("ADD Vx,Vy: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestADDGenerateIVx(t *testing.T) {
	expr := token.Expression{instr("ADD"), reg("I"), reg("V2")}
	expected := [2]byte{0xF2, 0x1E}
	opcode := ADD{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("ADD I,Vx: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSUBGenerate(t *testing.T) {
	expr := token.Expression{instr("SUB"), reg("V4"), reg("V5")}
	expected := [2]byte{0x84, 0x55}
	opcode := SUB{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SUB: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSUBNGenerate(t *testing.T) {
	expr := token.Expression{instr("SUBN"), reg("V6"), reg("V7")}
	expected := [2]byte{0x86, 0x77}
	opcode := SUBN{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SUBN: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestORGenerate(t *testing.T) {
	expr := token.Expression{instr("OR"), reg("V8"), reg("V9")}
	expected := [2]byte{0x88, 0x91}
	opcode := OR{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("OR: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestANDGenerate(t *testing.T) {
	expr := token.Expression{instr("AND"), reg("VA"), reg("VB")}
	expected := [2]byte{0x8A, 0xB2}
	opcode := AND{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("AND: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestXORGenerate(t *testing.T) {
	expr := token.Expression{instr("XOR"), reg("VC"), reg("VD")}
	expected := [2]byte{0x8C, 0xD3}
	opcode := XOR{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("XOR: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSHRGenerate(t *testing.T) {
	expr := token.Expression{instr("SHR"), reg("V2"), reg("V3")}
	expected := [2]byte{0x82, 0x36}
	opcode := SHR{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SHR: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSHLGenerate(t *testing.T) {
	expr := token.Expression{instr("SHL"), reg("V2"), reg("V3")}
	expected := [2]byte{0x82, 0x3E}
	opcode := SHL{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SHL: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestRNDGenerate(t *testing.T) {
	expr := token.Expression{instr("RND"), reg("V4"), num(0xAC)}
	expected := [2]byte{0xC4, 0xAC}
	opcode := RND{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("RND: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestDRWGenerate(t *testing.T) {
	expr := token.Expression{instr("DRW"), reg("V2"), reg("V3"), num(5)}
	expected := [2]byte{0xD2, 0x35}
	opcode := DRW{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("DRW: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSKPGenerate(t *testing.T) {
	expr := token.Expression{instr("SKP"), reg("V8")}
	expected := [2]byte{0xE8, 0x9E}
	opcode := SKP{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SKP: got %v, expected %v", opcode.Bytes, expected)
	}
}

func TestSKNPGenerate(t *testing.T) {
	expr := token.Expression{instr("SKNP"), reg("V7")}
	expected := [2]byte{0xE7, 0xA1}
	opcode := SKNP{expression: expr}.generate()
	if opcode.Bytes != expected {
		t.Errorf("SKNP: got %v, expected %v", opcode.Bytes, expected)
	}
}
