package chip

import (
	oc "github.com/xoesae/chip8/emulator/opcode"
	"github.com/xoesae/chip8/emulator/shared"
	"github.com/xoesae/chip8/logger"
)

func (c *Chip) handleGroupD(opcode oc.Opcode) {
	logger.Get().Debug("DRW")

	vx := int(c.v[opcode.X])
	vy := int(c.v[opcode.Y])

	c.v[0xF] = 0
	collision := false

	// 0..5
	for row := 0; row < int(opcode.N); row++ {
		_addr := c.i + uint16(row)
		spriteByte := c.memory.Read(_addr + 0x200)

		// 1 byte column
		for col := 0; col < 8; col++ {
			// 0x80 == 1000 0000
			// 0x80>>1 == 0100 0000
			// 0x80>>2 == 0010 0000
			// 0x80>>3 == 0001 0000
			// ...
			shiftedCol := byte(0x80 >> uint(col))

			// spriteByte == 0xF0 (1111 0000)
			// spriteByte&1000 0000 == 1000 0000
			// spriteByte&0100 0000 == 0100 0000
			// spriteByte&0010 0000 == 0010 0000
			// ...
			// spriteByte&0000 1000 == 0000 0000 -> no pixel ON
			if spriteByte&shiftedCol == 0 {
				continue
			}

			x := (vx + col) % shared.DisplayWidth
			y := (vy + row) % shared.DisplayHeight

			// [0, 0, 0, 1, 1]
			// [0, 1, 1, 0, 0]
			// [0, 1, 1, 0, 0]
			// [0, 1, 1, 0, 0]

			if c.pixels[y][x] {
				collision = true
			}

			c.pixels[y][x] = !c.pixels[y][x]
		}
	}

	if collision {
		c.v[0xF] = 1
	} else {
		c.v[0xF] = 0
	}
}
