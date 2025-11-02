package codegenerator

import (
	"fmt"

	"github.com/xoesae/chip8/assembler/token"
)

func (c *CodeGenerator) processLabel(expression token.Expression) error {
	label := mustAs[token.Label](expression[0])
	if _, exists := c.labels[label.Value]; exists {
		return fmt.Errorf("repeated label")
	}

	// Set current address for the label
	c.labels[label.Value] = c.addressCounter.pos

	return nil
}
