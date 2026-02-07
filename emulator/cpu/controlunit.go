package cpu

import (
	"fmt"

	"github.com/xoesae/chip8/emulator/event"
	"github.com/xoesae/chip8/emulator/io/display"
	"github.com/xoesae/chip8/logger"
)

type ControlUnit struct{}

// Fetch returns opcode from actual PC
func (c *ControlUnit) fetch(cpu *CPU) Opcode {
	high := uint16(cpu.memory.Read(cpu.pc.Current()))
	low := uint16(cpu.memory.Read(cpu.pc.Current() + 1))

	opcode := (high << 8) | low

	return Opcode{
		Raw:   opcode,
		Group: opcode & 0xF000,
		X:     byte((opcode & 0x0F00) >> 8),
		Y:     byte((opcode & 0x00F0) >> 4),
		N:     byte(opcode & 0x000F),
		NN:    byte(opcode & 0x00FF),
		NNN:   opcode & 0x0FFF,
	}
}

func (c *ControlUnit) ExecuteCycle(cpu *CPU) {
	opcode := c.fetch(cpu)

	if opcode.Raw != 0x00 {
		logger.Get().Debug(fmt.Sprintf("Opcode: 0x%04X", opcode.Raw))
	}

	switch opcode.Group {
	case 0x0000:
		c.handle0Group(cpu, opcode)
		return
	case 0x1000:
		logger.Get().Debug("JUMP NNN")
		cpu.pc.JumpTo(opcode.NNN)
		return
	case 0x2000:
		// CALL
		logger.Get().Debug("0x2000")
		return
	case 0x3000:
		logger.Get().Debug("0x3000")
		return
	case 0x4000:
		logger.Get().Debug("0x4000")
		return
	case 0x5000:
		logger.Get().Debug(fmt.Sprintf("SE V%d, V%d", opcode.X, opcode.Y))

		if cpu.v[opcode.X] == cpu.v[opcode.Y] {
			// Ignore next instruction, add +2 on program counter
			cpu.pc.Count()

			logger.Get().Debug(fmt.Sprintf("V%d == V%d", opcode.X, opcode.Y))
		}

		return
	case 0x6000:
		// v[x] := NN
		logger.Get().Debug(fmt.Sprintf("V%d := %d", opcode.X, opcode.NN))
		cpu.v[opcode.X] = opcode.NN
		return
	case 0x7000:
		// v[x] += NN
		logger.Get().Debug(fmt.Sprintf("V%d += %d", opcode.X, opcode.NN))
		cpu.v[opcode.X] += opcode.NN
		logger.Get().Debug(fmt.Sprintf("V%d == %d", opcode.X, cpu.v[opcode.X]))
		return
	case 0x8000:
		c.handle8Group(cpu, opcode)
		return
	case 0x9000:
		logger.Get().Debug("0x9000") // SNE
		return
	case 0xA000:
		logger.Get().Debug(fmt.Sprintf("I = %d", opcode.NNN))

		cpu.i = opcode.NNN

		return
	case 0xB000:
		logger.Get().Debug("0xB000")
		return
	case 0xC000:
		logger.Get().Debug("0xC000")
		return
	case 0xD000:
		logger.Get().Debug("DRW")

		vx := int(cpu.v[opcode.X])
		vy := int(cpu.v[opcode.Y])

		cpu.v[0xF] = 0
		collision := false

		// 0..5
		for row := 0; row < int(opcode.N); row++ {
			_addr := cpu.i + uint16(row)
			spriteByte := cpu.memory.Read(_addr + 0x200)

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

				x := (vx + col) % display.DisplayWidth
				y := (vy + row) % display.DisplayHeight

				// [0, 0, 0, 1, 1]
				// [0, 1, 1, 0, 0]
				// [0, 1, 1, 0, 0]
				// [0, 1, 1, 0, 0]

				if cpu.display[y][x] {
					collision = true
				}

				cpu.display[y][x] = !cpu.display[y][x]
			}
		}

		if collision {
			cpu.v[0xF] = 1
		} else {
			cpu.v[0xF] = 0
		}

		cpu.emitEvent(event.DisplayUpdatedEvent{
			Pixels: cpu.display,
		})

		return
	case 0xE000:
		logger.Get().Debug("0xE000")
		return
	case 0xF000:
		c.handleFGroup(cpu, opcode)
		return
	}
}

func (c *ControlUnit) handle0Group(cpu *CPU, opcode Opcode) {
	switch opcode.Raw & 0x00FF {
	case 0xE0:
		logger.Get().Debug("CLEAR")
		for y := 0; y < display.DisplayHeight; y++ {
			for x := 0; x < display.DisplayWidth; x++ {
				cpu.display[y][x] = false
			}
		}
		cpu.emitEvent(event.DisplayClearEvent{})
	case 0xEE:
		logger.Get().Debug("Return")
	}
}

func (c *ControlUnit) handle8Group(cpu *CPU, opcode Opcode) {
	switch opcode.N {
	case 0x0: // 8XY0
		logger.Get().Debug(fmt.Sprintf("V%d := V%d\n", opcode.X, opcode.Y))
		cpu.v[opcode.X] = cpu.v[opcode.Y]
	case 0x1: // 8XY1
		logger.Get().Debug(fmt.Sprintf("V%d |= V%d\n", opcode.X, opcode.Y))
		cpu.v[opcode.X] |= cpu.v[opcode.Y]
	case 0x2: // 8XY2
		logger.Get().Debug(fmt.Sprintf("V%d &= V%d\n", opcode.X, opcode.Y))
		cpu.v[opcode.X] &= cpu.v[opcode.Y]
	case 0x3: // 8XY3
		logger.Get().Debug(fmt.Sprintf("V%d ^= V%d\n", opcode.X, opcode.Y))
		cpu.v[opcode.X] ^= cpu.v[opcode.Y]
	case 0x4: // 8XY4
		logger.Get().Debug(fmt.Sprintf("V%d += V%d\n", opcode.X, opcode.Y))
		sum := uint16(cpu.v[opcode.X]) + uint16(cpu.v[opcode.Y])
		if sum > 0xFF {
			cpu.v[0xF] = 1 // set carry
		} else {
			cpu.v[0xF] = 0
		}
		cpu.v[opcode.X] = byte(sum & 0xFF)
	case 0x5: // 8XY5
		logger.Get().Debug(fmt.Sprintf("V%d -= V%d\n", opcode.X, opcode.Y))
		if cpu.v[opcode.X] > cpu.v[opcode.Y] {
			cpu.v[0xF] = 1
		} else {
			cpu.v[0xF] = 0
		}
		cpu.v[opcode.X] -= cpu.v[opcode.Y]
	case 0x6: // 8XY6
		logger.Get().Debug(fmt.Sprintf("V%d >>= V%d\n", opcode.X, opcode.Y))
		cpu.v[0xF] = cpu.v[opcode.X] & 0x1
		cpu.v[opcode.X] >>= 1
	case 0x7: // 8XY7
		logger.Get().Debug(fmt.Sprintf("V%d =- V%d\n", opcode.X, opcode.Y))
		if cpu.v[opcode.Y] > cpu.v[opcode.X] {
			cpu.v[0xF] = 1
		} else {
			cpu.v[0xF] = 0
		}
		cpu.v[opcode.X] = cpu.v[opcode.Y] - cpu.v[opcode.X]
	case 0xE: // 8XYE
		logger.Get().Debug(fmt.Sprintf("V%d <<= V%d\n", opcode.X, opcode.Y))
		cpu.v[0xF] = (cpu.v[opcode.X] & 0x80) >> 7
		cpu.v[opcode.X] <<= 1
	}
}

func (c *ControlUnit) handleFGroup(cpu *CPU, opcode Opcode) {
	x := uint16(opcode.X)

	switch opcode.NN {
	case 0x55: // FX55
		logger.Get().Debug(fmt.Sprintf("SAVE V%d\n", opcode.X))
		for i := uint16(0); i <= x; i++ {
			cpu.memory.Write(cpu.i+i, cpu.v[i])
		}
	case 0x65: // FX65
		logger.Get().Debug(fmt.Sprintf("LOAD V%d\n", opcode.X))
		for i := uint16(0); i <= x; i++ {
			cpu.v[i] = cpu.memory.Read(cpu.i + i)
		}
	}
}
