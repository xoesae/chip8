package chip

import (
	"fmt"

	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroup8(opcode oc.Opcode) {
	switch opcode.N {
	case 0x0: // 8XY0
		logger.Get().Debug(fmt.Sprintf("V%d := V%d", opcode.X, opcode.Y))
		c.v[opcode.X] = c.v[opcode.Y]
	case 0x1: // 8XY1
		logger.Get().Debug(fmt.Sprintf("V%d |= V%d", opcode.X, opcode.Y))
		c.v[opcode.X] |= c.v[opcode.Y]
	case 0x2: // 8XY2
		logger.Get().Debug(fmt.Sprintf("V%d &= V%d", opcode.X, opcode.Y))
		c.v[opcode.X] &= c.v[opcode.Y]
	case 0x3: // 8XY3
		logger.Get().Debug(fmt.Sprintf("V%d ^= V%d", opcode.X, opcode.Y))
		c.v[opcode.X] ^= c.v[opcode.Y]
	case 0x4: // 8XY4
		logger.Get().Debug(fmt.Sprintf("V%d += V%d", opcode.X, opcode.Y))
		sum := uint16(c.v[opcode.X]) + uint16(c.v[opcode.Y])
		if sum > 0xFF {
			c.v[0xF] = 1 // set carry
		} else {
			c.v[0xF] = 0
		}
		c.v[opcode.X] = byte(sum & 0xFF)
	case 0x5: // 8XY5
		logger.Get().Debug(fmt.Sprintf("V%d -= V%d", opcode.X, opcode.Y))
		if c.v[opcode.X] > c.v[opcode.Y] {
			c.v[0xF] = 1
		} else {
			c.v[0xF] = 0
		}
		c.v[opcode.X] -= c.v[opcode.Y]
	case 0x6: // 8XY6
		logger.Get().Debug(fmt.Sprintf("V%d >>= V%d", opcode.X, opcode.Y))
		c.v[0xF] = c.v[opcode.X] & 0x1
		c.v[opcode.X] >>= 1
	case 0x7: // 8XY7
		logger.Get().Debug(fmt.Sprintf("V%d =- V%d", opcode.X, opcode.Y))
		if c.v[opcode.Y] > c.v[opcode.X] {
			c.v[0xF] = 1
		} else {
			c.v[0xF] = 0
		}
		c.v[opcode.X] = c.v[opcode.Y] - c.v[opcode.X]
	case 0xE: // 8XYE
		logger.Get().Debug(fmt.Sprintf("V%d <<= V%d", opcode.X, opcode.Y))
		c.v[0xF] = (c.v[opcode.X] & 0x80) >> 7
		c.v[opcode.X] <<= 1
	}
}
