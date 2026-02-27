package chip

import (
	"fmt"

	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroup3(opcode oc.Opcode) {
	logger.Get().Debug(fmt.Sprintf("SE V%d, %d", opcode.X, opcode.NN))

	if c.v[opcode.X] == opcode.NN {
		// Ignore next instruction, add +2 on program counter
		c.pc.Count()

		logger.Get().Debug(fmt.Sprintf("V%d == %d", opcode.X, opcode.NN))
	}
}
