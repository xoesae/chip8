package chip

import (
	"fmt"

	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroupA(opcode oc.Opcode) {
	logger.Get().Debug(fmt.Sprintf("I = %d", opcode.NNN))

	c.i = opcode.NNN
}
