package chip

import (
	"fmt"

	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroupF(opcode oc.Opcode) {
	x := uint16(opcode.X)

	switch opcode.NN {
	case 0x0A:
		logger.Get().Debug(fmt.Sprintf("LD V%d, K", opcode.X))

		// Wait for a keypress and store the result in register VX
		if c.hasKeyPressed {
			c.v[opcode.X] = c.keyPressed
			logger.Get().Debug("Key pressed", "key", c.keyPressed)
		} else {
			c.pc.Undo()
			logger.Get().Debug("Waiting for a keypress")
		}

		return
	case 0x29:
		logger.Get().Debug(fmt.Sprintf("LD F, V%d", opcode.X))

		c.i = uint16(c.v[opcode.X])

	case 0x55: // FX55
		logger.Get().Debug(fmt.Sprintf("SAVE V%d", opcode.X))
		for i := uint16(0); i <= x; i++ {
			c.memory.Write(c.i+i, c.v[i])
		}
	case 0x65: // FX65
		logger.Get().Debug(fmt.Sprintf("LOAD V%d", opcode.X))
		for i := uint16(0); i <= x; i++ {
			c.v[i] = c.memory.Read(c.i + i)
		}
	}
}
