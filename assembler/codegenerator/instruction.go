package codegenerator

import (
	"github.com/xoesae/chip8/assembler/token"
)

type InstructionGenerator interface {
	generate() *OpCode
}

type CLS struct{}

func (c CLS) generate() *OpCode {
	return NewOpCode([2]byte{0x00, 0xE0})
}

type RET struct{}

func (r RET) generate() *OpCode {
	return NewOpCode([2]byte{0x00, 0xEE})
}

type JP struct {
	expression token.Expression
	labels     LabelMap
}

func (j JP) generate() *OpCode {
	var addr token.Token
	var prefix byte
	var nnn uint16

	switch len(j.expression) {
	case 2:
		// JP addr (1NNN)
		addr = j.expression[1]
		prefix = byte(0x10)
	case 3:
		// JP V0, addr (BNNN)
		addr = j.expression[2]
		prefix = byte(0xB0)
	default:
		panic("invalid argument to jp instruction")
	}

	switch addr.(type) {
	case token.LabelOperand:
		label, _ := addr.(token.LabelOperand)
		labelAddress, exists := j.labels[label.Value]
		if !exists {
			panic("define label before jump it. label: " + label.Value)
		}

		nnn = uint16(labelAddress)
	case token.NumericLiteral:
		literal, _ := addr.(token.NumericLiteral)
		nnn = uint16(literal.Value)
	default:
		panic("invalid argument to jp instruction")
	}

	return NewOpCodePNNN(prefix, nnn)
}

type CALL struct {
	expression token.Expression
	labels     LabelMap
}

func (c CALL) generate() *OpCode {
	// CALL addr (2NNN)
	addr := c.expression[1]
	prefix := byte(0x20)
	var nnn uint16

	switch addr.(type) {
	case token.LabelOperand:
		label, _ := addr.(token.LabelOperand)
		labelAddress, exists := c.labels[label.Value]
		if !exists {
			panic("define label before call it. label: " + label.Value)
		}

		nnn = uint16(labelAddress)
	case token.NumericLiteral:
		literal, _ := addr.(token.NumericLiteral)
		nnn = uint16(literal.Value)
	default:
		panic("invalid argument to call instruction")
	}

	return NewOpCodePNNN(prefix, nnn)
}

type SE struct {
	expression token.Expression
}

func (s SE) generate() *OpCode {
	operand := s.expression[2]
	vx, ok := s.expression[1].(token.Register)
	if !ok {
		panic("Register expected for SE instruction")
	}
	x := vx.Value[1]

	switch operand.(type) {
	case token.NumericLiteral:
		// SE Vx, NN (3XNN)
		literal, _ := operand.(token.NumericLiteral)

		return NewOpCodePXNN(byte(0x30), x, byte(literal.Value))

	case token.Register:
		// SE Vx, Vy (5XY0)
		vy, _ := operand.(token.Register)
		y := vy.Value[1]

		return NewOpCodePXYS(byte(0x50), x, y, 0x00)
	default:
		panic("Register or NumericLiteral expected")
	}
}

type SNE struct {
	expression token.Expression
}

func (s SNE) generate() *OpCode {
	operand := s.expression[2]
	vx, ok := s.expression[1].(token.Register)
	if !ok {
		panic("Register expected for SE instruction")
	}
	x := vx.Value[1]

	switch operand.(type) {
	case token.NumericLiteral:
		// SNE Vx, NN (4XNN)
		literal, _ := operand.(token.NumericLiteral)

		return NewOpCodePXNN(byte(0x40), x, byte(literal.Value))

	case token.Register:
		// SNE Vx, Vy (9XY0)
		vy, _ := operand.(token.Register)
		y := vy.Value[1]

		return NewOpCodePXYS(byte(0x90), x, y, 0x00)
	default:
		panic("Register or NumericLiteral expected")
	}
}

type LD struct {
	expression token.Expression
}

func (l LD) generate() *OpCode {
	fxnnCases := map[string]byte{
		string(token.DT): 0x15,
		string(token.ST): 0x18,
		string(token.F):  0x29,
		string(token.B):  0x33,
		string(token.VI): 0x55,
	}

	destination := l.expression[1].(token.Register)
	origin := l.expression[2]

	// LD I, addr (ANNN)
	if destination.Value == string(token.I) {
		literal := mustAs[token.NumericLiteral](origin)

		return NewOpCodePNNN(0xA0, uint16(literal.Value))
	}

	// LD {DT, ST, F, B, [I]}, Vx (FXNN)
	if suffix, found := fxnnCases[destination.Value]; found {
		vx := mustAs[token.Register](origin)

		return NewOpCodePXNN(0xF0, vx.Value[1], suffix)
	}

	x := mustAs[token.Register](origin).Value[1]
	switch o := origin.(type) {
	case token.NumericLiteral:
		// LD Vx, byte (6XNN)
		return NewOpCodePXNN(0x60, x, byte(o.Value))
	case token.Register:
		y := o.Value
		switch {
		case y[0] == 'V':
			// LD Vx, Vy (8XY0)
			return NewOpCodePXYS(0x80, x, y[1], 0x00)
		case y == string(token.DT):
			// LD Vx, DT (FX07)
			return NewOpCodePXNN(0xF0, x, 0x07)
		case y == string(token.K):
			// LD Vx, K (FX0A)
			return NewOpCodePXNN(0xF0, x, 0x0A)
		case y == string(token.VI):
			// LD Vx, [I] (FX65)
			return NewOpCodePXNN(0xF0, x, 0x65)
		}
	}

	panic("Invalid LD instruction")
}

type ADD struct {
	expression token.Expression
}

func (a ADD) generate() *OpCode {
	destination := mustAs[token.Register](a.expression[1])

	if destination.Value[0] == 'V' {
		x := destination.Value[1]

		if num, ok := a.expression[2].(token.NumericLiteral); ok {
			// ADD Vx, byte (7XNN)
			return NewOpCodePXNN(0x70, x, byte(num.Value))
		}

		if register, ok := a.expression[2].(token.Register); ok {
			// ADD Vx, Vy (8XY4)
			if register.Value[0] != 'V' {
				panic("invalid ADD instruction")
			}
			y := register.Value[1]

			return NewOpCodePXYS(0x80, x, y, 0x04)
		}
	}

	if destination.Value == string(token.I) {
		// ADD I, Vx (FX1E)
		vx, ok := a.expression[2].(token.Register)
		if !ok {
			panic("invalid ADD instruction")
		}
		if vx.Value[0] != 'V' {
			panic("invalid ADD instruction")
		}

		x := vx.Value[1]

		return NewOpCodePXNN(0xF0, x, 0x1E)
	}

	panic("invalid ADD instruction")
}

type SUB struct {
	expression token.Expression
}

func (s SUB) generate() *OpCode {
	// SUB Vx, Vy (8XY5)
	vx := mustAs[token.Register](s.expression[1])
	vy := mustAs[token.Register](s.expression[2])

	if vx.Value[0] != 'V' {
		panic("invalid SUB instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid SUB instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]

	return NewOpCodePXYS(0x80, x, y, 0x05)
}

type SUBN struct {
	expression token.Expression
}

func (s SUBN) generate() *OpCode {
	// SUBN Vx, Vy (8XY7)
	vx := mustAs[token.Register](s.expression[1])
	vy := mustAs[token.Register](s.expression[2])

	if vx.Value[0] != 'V' {
		panic("invalid SUBN instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid SUBN instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]

	return NewOpCodePXYS(0x80, x, y, 0x07)
}

type OR struct {
	expression token.Expression
}

func (o OR) generate() *OpCode {
	// OR Vx, Vy (8XY1)
	vx := mustAs[token.Register](o.expression[1])
	vy := mustAs[token.Register](o.expression[2])

	if vx.Value[0] != 'V' {
		panic("invalid OR instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid OR instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]

	return NewOpCodePXYS(0x80, x, y, 0x01)
}

type AND struct {
	expression token.Expression
}

func (a AND) generate() *OpCode {
	// AND Vx, Vy (8XY2)
	vx := mustAs[token.Register](a.expression[1])
	vy := mustAs[token.Register](a.expression[2])

	if vx.Value[0] != 'V' {
		panic("invalid AND instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid AND instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]

	return NewOpCodePXYS(0x80, x, y, 0x02)
}

type XOR struct {
	expression token.Expression
}

func (a XOR) generate() *OpCode {
	// XOR Vx, Vy (8XY3)
	vx := mustAs[token.Register](a.expression[1])
	vy := mustAs[token.Register](a.expression[2])

	if vx.Value[0] != 'V' {
		panic("invalid XOR instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid XOR instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]

	return NewOpCodePXYS(0x80, x, y, 0x03)
}

type SHR struct {
	expression token.Expression
}

func (s SHR) generate() *OpCode {
	// SHR Vx, Vy (8XY6)
	vx := mustAs[token.Register](s.expression[1])
	vy := mustAs[token.Register](s.expression[2])

	if vx.Value[0] != 'V' {
		panic("invalid SHR instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid SHR instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]

	return NewOpCodePXYS(0x80, x, y, 0x06)
}

type SHL struct {
	expression token.Expression
}

func (s SHL) generate() *OpCode {
	// SHL Vx, Vy (8XYE)
	vx := mustAs[token.Register](s.expression[1])
	vy := mustAs[token.Register](s.expression[2])

	if vx.Value[0] != 'V' {
		panic("invalid SHL instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid SHL instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]

	return NewOpCodePXYS(0x80, x, y, 0x0E)
}

type RND struct {
	expression token.Expression
}

func (r RND) generate() *OpCode {
	// RND Vx, byte (CXNN)
	vx := mustAs[token.Register](r.expression[1])
	literal := mustAs[token.NumericLiteral](r.expression[2])

	if vx.Value[0] != 'V' {
		panic("invalid RND instruction")
	}
	x := vx.Value[1]

	return NewOpCodePXSS(0xC0, x, byte(literal.Value))
}

type DRW struct {
	expression token.Expression
}

func (d DRW) generate() *OpCode {
	// DRW Vx, Vy, N (DXYN)
	vx := mustAs[token.Register](d.expression[1])
	vy := mustAs[token.Register](d.expression[2])
	literal := mustAs[token.NumericLiteral](d.expression[3])

	if vx.Value[0] != 'V' {
		panic("invalid DRW instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid DRW instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	n := uint8(literal.Value)

	return NewOpCodePXYS(0xD0, x, y, n)
}

type SKP struct {
	expression token.Expression
}

func (s SKP) generate() *OpCode {
	// SKP Vx (EX9E)
	vx := mustAs[token.Register](s.expression[1])

	if vx.Value[0] != 'V' {
		panic("invalid SKP instruction")
	}
	x := vx.Value[1]

	return NewOpCodePXNN(0xE0, x, 0x9E)
}

type SKNP struct {
	expression token.Expression
}

func (s SKNP) generate() *OpCode {
	// SKNP Vx (EXA1)
	vx := mustAs[token.Register](s.expression[1])

	if vx.Value[0] != 'V' {
		panic("invalid SKP instruction")
	}
	x := vx.Value[1]

	return NewOpCodePXNN(0xE0, x, 0xA1)
}

func (c *CodeGenerator) getInstructionGenerator(expression token.Expression) InstructionGenerator {
	instruction := expression[0].(token.Instruction)

	switch instruction.Value {
	case string(token.CLS):
		return CLS{}
	case string(token.RET):
		return RET{}
	case string(token.JP):
		return JP{
			expression: expression,
			labels:     c.labels,
		}
	case string(token.CALL):
		return CALL{
			expression: expression,
			labels:     c.labels,
		}
	case string(token.SE):
		return SE{
			expression: expression,
		}
	case string(token.SNE):
		return SNE{
			expression: expression,
		}
	case string(token.LD):
		return LD{
			expression: expression,
		}
	case string(token.ADD):
		return ADD{
			expression: expression,
		}
	case string(token.SUB):
		return SUB{
			expression: expression,
		}
	case string(token.SUBN):
		return SUBN{
			expression: expression,
		}
	case string(token.OR):
		return OR{
			expression: expression,
		}
	case string(token.AND):
		return AND{
			expression: expression,
		}
	case string(token.XOR):
		return XOR{
			expression: expression,
		}
	case string(token.SHR):
		return SHR{
			expression: expression,
		}
	case string(token.SHL):
		return SHL{
			expression: expression,
		}
	case string(token.RND):
		return RND{
			expression: expression,
		}
	case string(token.DRW):
		return DRW{
			expression: expression,
		}
	case string(token.SKP):
		return SKP{
			expression: expression,
		}
	case string(token.SKNP):
		return SKNP{
			expression: expression,
		}
	default:
		panic("invalid instruction")
	}
}
