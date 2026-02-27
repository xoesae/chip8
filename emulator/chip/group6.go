package chip

import (
	"fmt"

	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroup6(opcode oc.Opcode) {
	// v[x] := NN
	logger.Get().Debug(fmt.Sprintf("V%d := %d", opcode.X, opcode.NN))
	c.v[opcode.X] = opcode.NN
}
