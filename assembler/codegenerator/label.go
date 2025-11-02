package codegenerator

import "github.com/xoesae/chip8/assembler/token"

func (c *CodeGenerator) processLabel(expression token.Expression) {
	label := expression[0].(token.Label)
	if _, exists := c.labels[label.Value]; exists {
		panic("repeated label: " + label.Value) // TODO: improve errors
	}

	// Set current address for the label
	c.labels[label.Value] = c.addressCounter.pos

	// TODO: maybe advance address?
}
