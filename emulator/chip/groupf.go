package chip

import (
	"fmt"

	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroupF(opcode oc.Opcode) {
	x := uint16(opcode.X)

	switch opcode.NN {
	case 0x07: // FX07
		logger.Get().Debug(fmt.Sprintf("LD V%d, DT", opcode.X))
		c.v[opcode.X] = c.delayTimer
		return
	case 0x0A: // FX0A
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
	case 0x15: // FX15
		logger.Get().Debug(fmt.Sprintf("LD DT, V%d", opcode.X))
		c.delayTimer = c.v[opcode.X]
		return
	case 0x18: // FX18
		logger.Get().Debug(fmt.Sprintf("LD ST, V%d", opcode.X))
		c.soundTimer = c.v[opcode.X]
		return
	case 0x29: // FX29
		logger.Get().Debug(fmt.Sprintf("LD F, V%d", opcode.X))
		c.i = uint16(c.v[opcode.X])
		return
	case 0x1E: // FX1E
		logger.Get().Debug(fmt.Sprintf("ADD I, V%d", opcode.X))
		c.i += uint16(c.v[opcode.X])
		return
	case 0x33: // FX33
		logger.Get().Debug(fmt.Sprintf("LD B, V%d", opcode.X))

		value := c.v[opcode.X] // Ex: 123

		c.memory.Write(c.i, value/100)       // 1
		c.memory.Write(c.i+1, (value/10)%10) // 2
		c.memory.Write(c.i+2, value%10)      //3
		return
	case 0x55: // FX55
		logger.Get().Debug(fmt.Sprintf("SAVE V%d", opcode.X))
		for i := uint16(0); i <= x; i++ {
			c.memory.Write(c.i+i, c.v[i])
		}
		return
	case 0x65: // FX65
		logger.Get().Debug(fmt.Sprintf("LOAD V%d", opcode.X))
		for i := uint16(0); i <= x; i++ {
			c.v[i] = c.memory.Read(c.i + i)
		}
		return
	}
}
