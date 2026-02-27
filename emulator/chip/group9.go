package chip

import (
	"fmt"

	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroup9(opcode oc.Opcode) {
	logger.Get().Debug(fmt.Sprintf("SNE V%d, V%d", opcode.X, opcode.Y))

	if c.v[opcode.X] != c.v[opcode.Y] {
		// Ignore next instruction, add +2 on program counter
		c.pc.Count()

		logger.Get().Debug(fmt.Sprintf("V%d != V%d", opcode.X, opcode.Y))
	}
}
