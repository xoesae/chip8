package codegenerator

import (
	"github.com/xoesae/chip8/assembler/lexer/token"
	"github.com/xoesae/chip8/assembler/parser"
)

func (c *CodeGenerator) processJumpInstruction(first byte, expression parser.Expression) {
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

func (c *CodeGenerator) processSEInstruction(literalOpcode byte, registerOpcode byte, expression parser.Expression) {
	// expression[1] = Vx
	// expression[2] = byte or Vy

	vx, ok := expression[1].(token.Register)
	if !ok {
		panic("Register expected")
	}

	switch operand := expression[2].(type) {
	case token.NumericLiteral:
		msb := literalOpcode | (vx.Value[1] - '0') // V2 -> 0x32
		lsb := byte(operand.Value & 0xFF)
		c.appendOpcode(msb, lsb)
	case token.Register:
		msb := registerOpcode | (vx.Value[1] - '0')
		lsb := (operand.Value[1] - '0') << 4
		c.appendOpcode(msb, lsb)
	default:
		panic("Register or NumericLiteral expected")
	}
}
