package chip

import (
	"fmt"

	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) executeCycle() {
	opcode := oc.NewOpcode(
		c.memory.Read(c.pc.Current()),
		c.memory.Read(c.pc.Current()+1),
	)

	if opcode.Raw != 0x00 {
		logger.Get().Debug(fmt.Sprintf("Opcode: 0x%04X", opcode.Raw))
	}

	switch opcode.Group {
	case 0x0000:
		c.handleGroup0(opcode)
		return
	case 0x1000:
		c.handleGroup1(opcode)
		return
	case 0x2000:
		// CALL
		logger.Get().Debug("0x2000")
		return
	case 0x3000:
		c.handleGroup3(opcode)
		return
	case 0x4000:
		c.handleGroup4(opcode)
		return
	case 0x5000:
		c.handleGroup5(opcode)
		return
	case 0x6000:
		c.handleGroup6(opcode)
		return
	case 0x7000:
		c.handleGroup7(opcode)
		return
	case 0x8000:
		c.handleGroup8(opcode)
		return
	case 0x9000:
		logger.Get().Debug("0x9000") // SNE
		return
	case 0xA000:
		c.handleGroupA(opcode)
		return
	case 0xB000:
		logger.Get().Debug("0xB000")
		return
	case 0xC000:
		logger.Get().Debug("0xC000")
		return
	case 0xD000:
		c.handleGroupD(opcode)
		return
	case 0xE000:
		logger.Get().Debug("0xE000")
		return
	case 0xF000:
		c.handleGroupF(opcode)
		return
	}
}
