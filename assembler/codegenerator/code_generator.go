package codegenerator

import (
	"github.com/xoesae/chip8/assembler/token"
)

type LabelMap map[string]uint32

type AddressCounter struct {
	pos uint32
}

func (a *AddressCounter) Pos() uint32 {
	return a.pos
}

func (a *AddressCounter) SetPos(pos uint32) {
	a.pos = pos
}

func (a *AddressCounter) Advance() {
	a.pos++
}

type CodeGenerator struct {
	addressCounter AddressCounter
	labels         LabelMap
	opcodes        []byte
}

func (c *CodeGenerator) appendOpcode(msb byte, lsb byte) {
	c.opcodes[c.addressCounter.Pos()] = msb // set MSB
	c.addressCounter.Advance()

	c.opcodes[c.addressCounter.Pos()] = lsb // set LSB
	c.addressCounter.Advance()
}

func (c *CodeGenerator) processLabel(expression token.Expression) {
	label := expression[0].(token.Label)
	if _, exists := c.labels[label.Value]; exists {
		panic("repeated label: " + label.Value) // TODO: improve errors
	}
	c.labels[label.Value] = c.addressCounter.Pos()
}

func (c *CodeGenerator) processDirective(expression token.Expression) {
	directive := expression[0].(token.Directive)

	switch directive.Value {
	case string(token.Org):
		addr := expression[1].(token.NumericLiteral)
		c.addressCounter.SetPos(addr.Value)
	case string(token.Db):
		for i := 1; i < len(expression); i++ {
			literal := expression[i].(token.NumericLiteral)
			// set the byte on the current addres
			c.opcodes[c.addressCounter.Pos()] = byte(literal.Value)
			c.addressCounter.Advance()
		}
	}
}

func (c *CodeGenerator) processInstruction(expression token.Expression) {
	instr := expression[0].(token.Instruction)
	switch instr.Value {
	case string(token.CLS):
		c.appendOpcode(0x00, 0xE0)
	case string(token.RET):
		c.appendOpcode(0x00, 0xEE)
	case string(token.JP):
		c.processJPInstruction(expression)
	case string(token.CALL):
		c.processNNNInstruction(0x20, expression)
	case string(token.SE):
		c.processSEInstruction(0x30, expression)
	case string(token.SNE):
		c.processSEInstruction(0x40, expression)
	case string(token.LD):
		c.processLDInstruction(expression)
	case string(token.ADD):
		c.processADDInstruction(expression)
	case string(token.SUB):
		c.processSUBInstruction(expression)
	case string(token.SUBN):
		c.processSUBNInstruction(expression)
	case string(token.OR):
		c.processORInstruction(expression)
	case string(token.AND):
		c.processANDInstruction(expression)
	case string(token.XOR):
		c.processXORInstruction(expression)
	case string(token.SHR):
		c.processSHRInstruction(expression)
	case string(token.SHL):
		c.processSHLInstruction(expression)
	case string(token.RND):
		c.processRNDInstruction(expression)
	case string(token.DRW):
		c.processDRWInstruction(expression)
	case string(token.SKP):
		c.processSKPInstruction(expression)
	case string(token.SKNP):
		c.processSKNPInstruction(expression)
	}
}

func (c *CodeGenerator) Generate(expressions []token.Expression) []byte {
	for _, expression := range expressions {
		if _, ok := expression[0].(token.Label); ok {
			c.processLabel(expression)
		}

		if _, ok := expression[0].(token.Directive); ok {
			c.processDirective(expression)
		}

		if _, ok := expression[0].(token.Instruction); ok {
			c.processInstruction(expression)
		}

		panic("invalid expression")
	}

	return c.opcodes
}
