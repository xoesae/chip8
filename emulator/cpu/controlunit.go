package cpu

import (
	"fmt"

	"github.com/xoesae/chip8/emulator/shared"
	"github.com/xoesae/chip8/logger"
)

func (c *CPU) executeCycle() {
	opcode := NewOpcode(
		c.memory.Read(c.pc.Current()),
		c.memory.Read(c.pc.Current()+1),
	)

	if opcode.Raw != 0x00 {
		logger.Get().Debug(fmt.Sprintf("Opcode: 0x%04X", opcode.Raw))
	}

	switch opcode.Group {
	case 0x0000:
		c.handle0Group(opcode)
		return
	case 0x1000:
		logger.Get().Debug("JUMP NNN")
		c.pc.JumpTo(opcode.NNN)
		return
	case 0x2000:
		// CALL
		logger.Get().Debug("0x2000")
		return
	case 0x3000:
		logger.Get().Debug(fmt.Sprintf("SE V%d, %d", opcode.X, opcode.NN))

		if c.v[opcode.X] == opcode.NN {
			// Ignore next instruction, add +2 on program counter
			c.pc.Count()

			logger.Get().Debug(fmt.Sprintf("V%d == %d", opcode.X, opcode.NN))
		}

		return
	case 0x4000:
		logger.Get().Debug(fmt.Sprintf("SNE V%d, %d", opcode.X, opcode.NN))

		if c.v[opcode.X] != opcode.NN {
			// Ignore next instruction, add +2 on program counter
			c.pc.Count()

			logger.Get().Debug(fmt.Sprintf("V%d != %d", opcode.X, opcode.NN))
		}

		return
	case 0x5000:
		logger.Get().Debug(fmt.Sprintf("SE V%d, V%d", opcode.X, opcode.Y))

		if c.v[opcode.X] == c.v[opcode.Y] {
			// Ignore next instruction, add +2 on program counter
			c.pc.Count()

			logger.Get().Debug(fmt.Sprintf("V%d == V%d", opcode.X, opcode.Y))
		}

		return
	case 0x6000:
		// v[x] := NN
		logger.Get().Debug(fmt.Sprintf("V%d := %d", opcode.X, opcode.NN))
		c.v[opcode.X] = opcode.NN
		return
	case 0x7000:
		// v[x] += NN
		logger.Get().Debug(fmt.Sprintf("V%d += %d", opcode.X, opcode.NN))
		c.v[opcode.X] += opcode.NN
		logger.Get().Debug(fmt.Sprintf("V%d == %d", opcode.X, c.v[opcode.X]))
		return
	case 0x8000:
		c.handle8Group(opcode)
		return
	case 0x9000:
		logger.Get().Debug("0x9000") // SNE
		return
	case 0xA000:
		logger.Get().Debug(fmt.Sprintf("I = %d", opcode.NNN))

		c.i = opcode.NNN

		return
	case 0xB000:
		logger.Get().Debug("0xB000")
		return
	case 0xC000:
		logger.Get().Debug("0xC000")
		return
	case 0xD000:
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

		return
	case 0xE000:
		logger.Get().Debug("0xE000")
		return
	case 0xF000:
		c.handleFGroup(opcode)
		return
	}
}

func (c *CPU) handle0Group(opcode Opcode) {
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

func (c *CPU) handle8Group(opcode Opcode) {
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

func (c *CPU) handleFGroup(opcode Opcode) {
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
