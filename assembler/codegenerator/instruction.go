package codegenerator

import (
	"github.com/xoesae/chip8/assembler/lexer/token"
	"github.com/xoesae/chip8/assembler/parser"
)

func (c *CodeGenerator) processNNNInstruction(first byte, expression parser.Expression) {
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

func (c *CodeGenerator) processSEInstruction(first byte, expression parser.Expression) {
	// expression[1] = Vx
	// expression[2] = byte or Vy
	registerOpcode := byte(0x50)

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("Register expected")
	}

	switch operand := expression[2].(type) {
	case token.NumericLiteral:
		// SE Vx, NN (3XNN)
		msb := first | (vx.Value[1] - '0') // V2 -> 0x32
		lsb := byte(operand.Value & 0xFF)
		c.appendOpcode(msb, lsb)
	case token.Register:
		// SE Vx, Vy (5XY0)
		msb := registerOpcode | (vx.Value[1] - '0')
		lsb := (operand.Value[1] - '0') << 4
		c.appendOpcode(msb, lsb)
	default:
		panic("Register or NumericLiteral expected")
	}
}

func (c *CodeGenerator) processLDInstruction(expression parser.Expression) {
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

func (c *CodeGenerator) processADDInstruction(expression parser.Expression) {
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

func (c *CodeGenerator) processSUBInstruction(expression parser.Expression) {
	// SUB Vx, Vy | Vx -= Vy | 8XY5

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("invalid ADD instruction")
	}
	if vx.Value[0] != 'V' {
		panic("invalid ADD instruction")
	}

	vy, ok := expression[2].(token.Register)
	if !ok {
		panic("invalid ADD instruction")
	}
	if vy.Value[0] != 'V' {
		panic("invalid ADD instruction")
	}

	x := vx.Value[1]
	y := vy.Value[1]
	msb := 0x80 | x
	lsb := (y << 4) | 0x5
	c.appendOpcode(msb, lsb)
}
