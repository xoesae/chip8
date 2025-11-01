package codegenerator

import (
	"github.com/xoesae/chip8/assembler/parser"
	token2 "github.com/xoesae/chip8/assembler/token"
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

func (c *CodeGenerator) processLabel(expression parser.Expression) {
	label := expression[0].(token2.Label)
	if _, exists := c.labels[label.Value]; exists {
		panic("repeated label: " + label.Value) // TODO: improve errors
	}
	c.labels[label.Value] = c.addressCounter.Pos()
}

func (c *CodeGenerator) processDirective(expression parser.Expression) {
	directive := expression[0].(token2.Directive)

	switch directive.Value {
	case string(token2.Org):
		addr := expression[1].(token2.NumericLiteral)
		c.addressCounter.SetPos(addr.Value)
	case string(token2.Db):
		for i := 1; i < len(expression); i++ {
			literal := expression[i].(token2.NumericLiteral)
			// set the byte on the current addres
			c.opcodes[c.addressCounter.Pos()] = byte(literal.Value)
			c.addressCounter.Advance()
		}
	}
}

func (c *CodeGenerator) processInstruction(expression parser.Expression) {
	instr := expression[0].(token2.Instruction)
	switch instr.Value {
	case string(token2.CLS):
		c.appendOpcode(0x00, 0xE0)
	case string(token2.RET):
		c.appendOpcode(0x00, 0xEE)
	case string(token2.JP):
		c.processJPInstruction(expression)
	case string(token2.CALL):
		c.processNNNInstruction(0x20, expression)
	case string(token2.SE):
		c.processSEInstruction(0x30, expression)
	case string(token2.SNE):
		c.processSEInstruction(0x40, expression)
	case string(token2.LD):
		c.processLDInstruction(expression)
	case string(token2.ADD):
		c.processADDInstruction(expression)
	case string(token2.SUB):
		c.processSUBInstruction(expression)
	case string(token2.SUBN):
		c.processSUBNInstruction(expression)
	case string(token2.OR):
		c.processORInstruction(expression)
	case string(token2.AND):
		c.processANDInstruction(expression)
	case string(token2.XOR):
		c.processXORInstruction(expression)
	case string(token2.SHR):
		c.processSHRInstruction(expression)
	case string(token2.SHL):
		c.processSHLInstruction(expression)
	case string(token2.RND):
		c.processRNDInstruction(expression)
	case string(token2.DRW):
		c.processDRWInstruction(expression)
	case string(token2.SKP):
		c.processSKPInstruction(expression)
	case string(token2.SKNP):
		c.processSKNPInstruction(expression)
	}
}

func (c *CodeGenerator) Generate(expressions []parser.Expression) []byte {
	for _, expression := range expressions {
		if _, ok := expression[0].(token2.Label); ok {
			c.processLabel(expression)
		}

		if _, ok := expression[0].(token2.Directive); ok {
			c.processDirective(expression)
		}

		if _, ok := expression[0].(token2.Instruction); ok {
			c.processInstruction(expression)
		}

		panic("invalid expression")
	}

	return c.opcodes
}
