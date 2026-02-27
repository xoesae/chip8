package chip

import (
	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroup1(opcode oc.Opcode) {
	logger.Get().Debug("JUMP NNN")
	c.pc.JumpTo(opcode.NNN)
}
