package cpu

import (
	"fmt"

	"github.com/xoesae/chip8/emulator/event"
	"github.com/xoesae/chip8/emulator/io/display"
)

type ControlUnit struct{}

// Fetch returns opcode from actual PC
func (c *ControlUnit) fetch(cpu *CPU) Opcode {
	high := uint16(cpu.memory.Read(cpu.pc.Current()))
	low := uint16(cpu.memory.Read(cpu.pc.Current() + 1))

	//fmt.Printf("high: 0x%04X | low: 0x%04X\n", high, low)

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

	fmt.Printf("opcode: 0x%04X\n", opcode.Raw)

	switch opcode.Group {
	case 0x0000:
		c.handle0Group(cpu, opcode)
		return
	case 0x1000:
		fmt.Println("JUMP NNN")
		cpu.pc.JumpTo(opcode.NNN)
		return
	case 0x2000:
		// CALL
		fmt.Println("0x2000")
		return
	case 0x3000:
		fmt.Println("0x3000")
		return
	case 0x4000:
		fmt.Println("0x4000")
		return
	case 0x5000:
		fmt.Println("0x5000")
		return
	case 0x6000:
		// v[x] := NN
		fmt.Printf("V%d := %d\n", opcode.X, opcode.NN)
		cpu.v[opcode.X] = opcode.NN
		return
	case 0x7000:
		// v[x] += NN
		fmt.Printf("V%d += %d\n", opcode.X, opcode.NN)
		cpu.v[opcode.X] += opcode.NN
		return
	case 0x8000:
		c.handle8Group(cpu, opcode)
		return
	case 0x9000:
		fmt.Println("0x9000")
		return
	case 0xA000:
		fmt.Println("0xA000")
		return
	case 0xB000:
		fmt.Println("0xB000")
		return
	case 0xC000:
		fmt.Println("0xC000")
		return
	case 0xD000:
		fmt.Println("DRW")

		vx := int(cpu.v[opcode.X]) % display.DisplayWidth
		vy := int(cpu.v[opcode.Y]) % display.DisplayHeight

		collision := false

		for row := 0; row < int(opcode.N); row++ {
			_addr := cpu.i + uint16(row)
			if _addr >= 0xFFF {
				continue
			}

			// get the byte from db instruction [F8 F8 F8 F8 F8] and add the 0x200 from chip8 offset program
			spriteByte := cpu.memory.Read(_addr + 0x200)

			// 1 byte column
			for col := 0; col < 8; col++ {
				bit := (spriteByte & (0x80 >> col)) != 0
				if !bit {
					continue
				}

				x := (vx + col) % display.DisplayWidth
				y := (vy + row) % display.DisplayHeight

				wasOn := cpu.display[y][x]
				cpu.display[y][x] = !wasOn

				if wasOn {
					collision = true
				}

				ev := event.PixelUpdatedEvent{
					X: x, Y: y,
					State: cpu.display[y][x],
				}

				cpu.emitEvent(ev)
			}

		}

		if collision {
			cpu.v[0xF] = 1
		} else {
			cpu.v[0xF] = 0
		}

		return
	case 0xE000:
		fmt.Println("0xE000")
		return
	case 0xF000:
		c.handleFGroup(cpu, opcode)
		return
	}
}

func (c *ControlUnit) handle0Group(cpu *CPU, opcode Opcode) {
	switch opcode.Raw & 0x00FF {
	case 0xE0:
		for y := 0; y < display.DisplayHeight; y++ {
			for x := 0; x < display.DisplayWidth; x++ {
				cpu.display[y][x] = false
			}
		}
		cpu.emitEvent(event.DisplayClearEvent{})
	case 0xEE:
		fmt.Println("Return")
	}
}

func (c *ControlUnit) handle8Group(cpu *CPU, opcode Opcode) {
	switch opcode.N {
	case 0x0: // 8XY0
		fmt.Printf("V%d := V%d\n", opcode.X, opcode.Y)
		cpu.v[opcode.X] = cpu.v[opcode.Y]
	case 0x1: // 8XY1
		fmt.Printf("V%d |= V%d\n", opcode.X, opcode.Y)
		cpu.v[opcode.X] |= cpu.v[opcode.Y]
	case 0x2: // 8XY2
		fmt.Printf("V%d &= V%d\n", opcode.X, opcode.Y)
		cpu.v[opcode.X] &= cpu.v[opcode.Y]
	case 0x3: // 8XY3
		fmt.Printf("V%d ^= V%d\n", opcode.X, opcode.Y)
		cpu.v[opcode.X] ^= cpu.v[opcode.Y]
	case 0x4: // 8XY4
		fmt.Printf("V%d += V%d\n", opcode.X, opcode.Y)
		sum := uint16(cpu.v[opcode.X]) + uint16(cpu.v[opcode.Y])
		if sum > 0xFF {
			cpu.v[0xF] = 1 // set carry
		} else {
			cpu.v[0xF] = 0
		}
		cpu.v[opcode.X] = byte(sum & 0xFF)
	case 0x5: // 8XY5
		fmt.Printf("V%d -= V%d\n", opcode.X, opcode.Y)
		if cpu.v[opcode.X] > cpu.v[opcode.Y] {
			cpu.v[0xF] = 1
		} else {
			cpu.v[0xF] = 0
		}
		cpu.v[opcode.X] -= cpu.v[opcode.Y]
	case 0x6: // 8XY6
		fmt.Printf("V%d >>= V%d\n", opcode.X, opcode.Y)
		cpu.v[0xF] = cpu.v[opcode.X] & 0x1
		cpu.v[opcode.X] >>= 1
	case 0x7: // 8XY7
		fmt.Printf("V%d =- V%d\n", opcode.X, opcode.Y)
		if cpu.v[opcode.Y] > cpu.v[opcode.X] {
			cpu.v[0xF] = 1
		} else {
			cpu.v[0xF] = 0
		}
		cpu.v[opcode.X] = cpu.v[opcode.Y] - cpu.v[opcode.X]
	case 0xE: // 8XYE
		fmt.Printf("V%d <<= V%d\n", opcode.X, opcode.Y)
		cpu.v[0xF] = (cpu.v[opcode.X] & 0x80) >> 7
		cpu.v[opcode.X] <<= 1
	}
}

func (c *ControlUnit) handleFGroup(cpu *CPU, opcode Opcode) {
	x := uint16(opcode.X)

	switch opcode.NN {
	case 0x55: // FX55
		fmt.Printf("SAVE V%d\n", opcode.X)
		for i := uint16(0); i <= x; i++ {
			cpu.memory.Write(cpu.i+i, cpu.v[i])
		}
	case 0x65: // FX65
		fmt.Printf("LOAD V%d\n", opcode.X)
		for i := uint16(0); i <= x; i++ {
			cpu.v[i] = cpu.memory.Read(cpu.i + i)
		}
	}
}
