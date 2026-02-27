package chip

import (
	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/emulator/shared"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroup0(opcode oc.Opcode) {
	switch opcode.Raw & 0x00FF {
	case 0xE0:
		logger.Get().Debug("CLEAR")
		for y := 0; y < shared.DisplayHeight; y++ {
			for x := 0; x < shared.DisplayWidth; x++ {
				c.pixels[y][x] = false
			}
		}
	case 0xEE:
		logger.Get().Debug("Return")
	}
}
