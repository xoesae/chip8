package codegenerator

import (
	"github.com/xoesae/chip8/assembler/token"
)

func (c *CodeGenerator) processNNNInstruction(first byte, expression token.Expression) {
	if lit, ok := expression[1].(token.NumericLiteral); ok {
		msb := first | byte((lit.Value&0xF00)>>8)
		lsb := byte(lit.Value & 0x0FF)
		c.appendOpcode(msb, lsb)
	} else if lbl, ok := expression[1].(token.LabelOperand); ok {
		addr, exists := c.labels[lbl.Value]
		if exists {
			msb := first | byte((addr&0xF00)>>8)
			lsb := byte(addr & 0xFF)
			c.appendOpcode(msb, lsb)
		} else {
			panic("invalid label operand: " + lbl.Value)
		}
	}
}

func (c *CodeGenerator) processJPInstruction(expression token.Expression) {
	if len(expression) == 2 {
		// JP addr | 1NNN
		if lit, ok := expression[1].(token.NumericLiteral); ok {
			msb := 0x10 | byte((lit.Value&0xF00)>>8)
			lsb := byte(lit.Value & 0x0FF)
			c.appendOpcode(msb, lsb)
		} else if lbl, ok := expression[1].(token.LabelOperand); ok {
			addr, exists := c.labels[lbl.Value]
			if exists {
				msb := 0x10 | byte((addr&0xF00)>>8)
				lsb := byte(addr & 0xFF)
				c.appendOpcode(msb, lsb)
			} else {
				panic("invalid label operand: " + lbl.Value)
			}
		}
	}

	if len(expression) == 3 {
		// JP V0, addr | BNNN
		if lit, ok := expression[2].(token.NumericLiteral); ok {
			msb := 0xB0 | byte((lit.Value&0xF00)>>8)
			lsb := byte(lit.Value & 0x0FF)
			c.appendOpcode(msb, lsb)
		} else if lbl, ok := expression[2].(token.LabelOperand); ok {
			addr, exists := c.labels[lbl.Value]
			if exists {
				msb := 0xB0 | byte((addr&0xF00)>>8)
				lsb := byte(addr & 0xFF)
				c.appendOpcode(msb, lsb)
			} else {
				panic("invalid label operand: " + lbl.Value)
			}
		}
	}

	panic("Invalid JP instruction")

}

func (c *CodeGenerator) processSEInstruction(first byte, expression token.Expression) {
	// expression[1] = Vx
	// expression[2] = byte or Vy

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("Register expected")
	}

	switch operand := expression[2].(type) {
	case token.NumericLiteral:
		// SE Vx, NN (3XNN)
		// SNE Vx, NN (4XNN)
		msb := first | (vx.Value[1] - '0') // V2 -> 0x32
		lsb := byte(operand.Value & 0xFF)
		c.appendOpcode(msb, lsb)
	case token.Register:
		// SE Vx, Vy (5XY0)
		// SNE Vx, Vy (9XY0)
		registerOpcode := byte(0x50)

		if first == 0x40 {
			registerOpcode = byte(0x90)
		}

		msb := registerOpcode | (vx.Value[1] - '0')
		lsb := (operand.Value[1] - '0') << 4
		c.appendOpcode(msb, lsb)
	default:
		panic("Register or NumericLiteral expected")
	}
}

func (c *CodeGenerator) processLDInstruction(expression token.Expression) {
	destination := expression[1].(token.Register)
	origin := expression[2]

	// Dest is Vx
	if destination.Value[0] == 'V' {
		x := destination.Value[1]

		switch origin.(type) {
		case token.NumericLiteral:
			// LD Vx, byte (6XNN)
			msb := 0x60 | x
			lsb := byte(origin.(token.NumericLiteral).Value & 0xFF)
			c.appendOpcode(msb, lsb)
			return
		case token.Register:
			reg := origin.(token.Register).Value

			if reg[0] == 'V' {
				// LD Vx, Vy (8XY0)
				y := reg[1]
				msb := 0x80 | x
				lsb := y<<4 | 0x00
				c.appendOpcode(msb, lsb)
				return
			}

			if reg == string(token.DT) {
				// LD Vx, DT (FX07)
				msb := 0xF0 | x
				lsb := 0x07
				c.appendOpcode(msb, byte(lsb))
				return
			}

			if reg == string(token.K) {
				// LD Vx, K (FX0A)
				msb := 0xF0 | x
				lsb := 0x0A
				c.appendOpcode(msb, byte(lsb))
				return
			}

			if reg == string(token.VI) {
				// LD Vx, [I] (FX65)
				msb := 0xF0 | x
				lsb := 0x65
				c.appendOpcode(msb, byte(lsb))
				return
			}
		}
	}

	if destination.Value == string(token.I) {
		// LD I, addr (ANNN)
		addr, ok := origin.(token.NumericLiteral)
		if !ok {
			panic("invalid LD instruction")
		}

		msb := 0xA0 | byte((addr.Value>>8)&0x0F)
		lsb := byte(addr.Value & 0xFF)
		c.appendOpcode(msb, lsb)
		return
	}

	if destination.Value == string(token.DT) {
		// LD DT, Vx (FX15)
		vx, ok := origin.(token.Register)
		if !ok {
			panic("invalid LD instruction")
		}

		x := vx.Value[1]
		msb := 0xF0 | x
		lsb := 0x15
		c.appendOpcode(msb, byte(lsb))
		return
	}

	if destination.Value == string(token.ST) {
		// LD ST, Vx (FX18)
		vx, ok := origin.(token.Register)
		if !ok {
			panic("invalid LD instruction")
		}

		x := vx.Value[1]
		msb := 0xF0 | x
		lsb := 0x18
		c.appendOpcode(msb, byte(lsb))
		return
	}

	if destination.Value == string(token.F) {
		// LD F, Vx (FX29)
		vx, ok := origin.(token.Register)
		if !ok {
			panic("invalid LD instruction")
		}

		x := vx.Value[1]
		msb := 0xF0 | x
		lsb := 0x29
		c.appendOpcode(msb, byte(lsb))
		return

	}

	if destination.Value == string(token.B) {
		// LD B, Vx (FX33)
		vx, ok := origin.(token.Register)
		if !ok {
			panic("invalid LD instruction")
		}

		x := vx.Value[1]
		msb := 0xF0 | x
		lsb := 0x33
		c.appendOpcode(msb, byte(lsb))
		return

	}

	if destination.Value == string(token.VI) {
		// LD [I], Vx (FX55)
		vx, ok := origin.(token.Register)
		if !ok {
			panic("invalid LD instruction")
		}

		x := vx.Value[1]
		msb := 0xF0 | x
		lsb := 0x55
		c.appendOpcode(msb, byte(lsb))
		return
	}

	panic("Invalid LD instruction")
}

func (c *CodeGenerator) processADDInstruction(expression token.Expression) {
	destination := expression[1].(token.Register)

	if destination.Value[0] == 'V' {
		x := destination.Value[1] // V[x]

		if num, ok := expression[2].(token.NumericLiteral); ok {
			// ADD Vx, byte	| Vx += byte | 7XNN
			msb := 0x70 | x
			lsb := byte(num.Value)
			c.appendOpcode(msb, lsb)
		}

		if register, ok := expression[2].(token.Register); ok {
			// ADD Vx, Vy | Vx += Vy | 8XY4
			if register.Value[0] != 'V' {
				panic("invalid ADD instruction")
			}

			y := register.Value[1]

			msb := 0x80 | x
			lsb := (y << 4) | 0x4
			c.appendOpcode(msb, lsb)
		}
	}

	if destination.Value == string(token.I) {
		// ADD I, Vx | I += Vx | FX1E
		vx, ok := expression[2].(token.Register)
		if !ok {
			panic("invalid ADD instruction")
		}
		if vx.Value[0] != 'V' {
			panic("invalid ADD instruction")
		}

		x := vx.Value[1]

		msb := 0xF0 | x
		lsb := 0x1E
		c.appendOpcode(msb, byte(lsb))
	}

	panic("invalid ADD instruction")
}

func (c *CodeGenerator) processSUBInstruction(expression token.Expression) {
	// SUB Vx, Vy | Vx -= Vy | 8XY5

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid SUB instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid SUB instruction")
	}

	vy, ok := expression[2].(token.Register)
	if !ok {
		panic("invalid SUB instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid SUB instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	msb := 0x80 | x
	lsb := (y << 4) | 0x5
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processSUBNInstruction(expression token.Expression) {
	// SUBN Vx, Vy | Vx = Vy - Vx | 8XY7

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid SUBN instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid SUBN instruction")
	}

	vy, ok := expression[2].(token.Register)
	if !ok {
		panic("invalid SUBN instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid SUBN instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	msb := 0x80 | x
	lsb := (y << 4) | 0x7
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processORInstruction(expression token.Expression) {
	// OR Vx, Vy | 8XY1

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid OR instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid OR instruction")
	}

	vy, ok := expression[2].(token.Register)
	if !ok {
		panic("invalid OR instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid OR instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	msb := 0x80 | x
	lsb := (y << 4) | 0x1
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processANDInstruction(expression token.Expression) {
	// AND Vx, Vy | 8XY2

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid AND instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid AND instruction")
	}

	vy, ok := expression[2].(token.Register)
	if !ok {
		panic("invalid AND instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid AND instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	msb := 0x80 | x
	lsb := (y << 4) | 0x2
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processXORInstruction(expression token.Expression) {
	// XOR Vx, Vy | 8XY3

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid XOR instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid XOR instruction")
	}

	vy, ok := expression[2].(token.Register)
	if !ok {
		panic("invalid XOR instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid XOR instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	msb := 0x80 | x
	lsb := (y << 4) | 0x3
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processSHRInstruction(expression token.Expression) {
	// SHR Vx, Vy | 8XY6

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid SHR instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid SHR instruction")
	}

	vy, ok := expression[2].(token.Register)
	if !ok {
		panic("invalid SHR instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid SHR instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	msb := 0x80 | x
	lsb := (y << 4) | 0x6
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processSHLInstruction(expression token.Expression) {
	// SHL Vx, Vy | 8XYE

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid SHR instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid SHR instruction")
	}

	vy, ok := expression[2].(token.Register)
	if !ok {
		panic("invalid SHR instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid SHR instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	msb := 0x80 | x
	lsb := (y << 4) | 0xE
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processRNDInstruction(expression token.Expression) {
	// RND Vx, byte | CXNN

	vx, ok := expression[1].(token.Register)

	if !ok {
		panic("invalid RND instruction")
	}

	if vx.Value[0] != 'V' {
		panic("invalid RND instruction")
	}

	x := vx.Value[1]

	num, ok := expression[2].(token.NumericLiteral)

	if !ok {
		panic("invalid RND instruction")
	}

	msb := 0x80 | x
	lsb := byte(num.Value)
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processDRWInstruction(expression token.Expression) {
	// DRW Vx, Vy, nibble | DXYN

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid DRW instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid DRW instruction")
	}

	vy, ok := expression[2].(token.Register)
	if !ok {
		panic("invalid DRW instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid DRW instruction")
	}

	num, ok := expression[3].(token.NumericLiteral)
	if !ok {
		panic("invalid DRW instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	n := uint8(num.Value)

	msb := 0xD0 | x
	lsb := (y << 4) | n
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processSKPInstruction(expression token.Expression) {
	// SKP Vx | EX9E

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid SKP instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid SKP instruction")
	}

	x := vx.Value[1]

	msb := 0xE0 | x
	lsb := byte(0x9E)
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processSKPNInstruction(expression token.Expression) {
	// SKNP Vx | EXA1

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid SKP instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid SKP instruction")
	}

	x := vx.Value[1]

	msb := 0xE0 | x
	lsb := byte(0xA1)
	c.appendOpcode(msb, lsb)
}

func (c *CodeGenerator) processSKNPInstruction(expression token.Expression) {
	// SKNP Vx | EXA1

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid SKP instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid SKP instruction")
	}

	x := vx.Value[1]

	msb := 0xE0 | x
	lsb := byte(0xA1)
	c.appendOpcode(msb, lsb)
}
