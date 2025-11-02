package codegenerator

import (
	"fmt"

	"github.com/xoesae/chip8/assembler/token"
)

type LabelMap map[string]uint32

type CodeGenerator struct {
	addressCounter AddressCounter
	labels         LabelMap
	opcodes        map[uint32]byte
}

func NewCodeGenerator() *CodeGenerator {
	return &CodeGenerator{
		addressCounter: AddressCounter{
			pos: 0,
		},
		labels:  make(LabelMap),
		opcodes: make(map[uint32]byte),
	}
}

func (c *CodeGenerator) appendOpcode(opcode *OpCode) {
	bytes := opcode.Bytes

	c.opcodes[c.addressCounter.pos] = bytes[0] // set MSB
	c.addressCounter.advance()

	c.opcodes[c.addressCounter.pos] = bytes[1] // set LSB
	c.addressCounter.advance()
}

func mustAs[T any](val any) T {
	result, ok := val.(T)
	if !ok {
		panic(fmt.Sprintf("invalid type: expected %T", result))
	}
	return result
}

func (c *CodeGenerator) Generate(expressions []token.Expression) map[uint32]byte {
	for _, expression := range expressions {
		if _, ok := expression[0].(token.Label); ok {
			err := c.processLabel(expression)
			if err != nil {
				panic(err)
			}

			continue
		}

		if _, ok := expression[0].(token.Directive); ok {
			c.processDirective(expression)
			continue
		}

		if _, ok := expression[0].(token.Instruction); ok {
			g := c.getInstructionGenerator(expression)
			opcode := g.generate()
			c.appendOpcode(opcode)
			continue
		}

		panic("invalid expression")
	}

	return c.opcodes
}
