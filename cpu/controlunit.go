package cpu

import (
	"fmt"
)

type ControlUnit struct{}

// Fetch returns opcode from actual PC
func (c *ControlUnit) fetch(cpu *CPU) Opcode {
	high := uint16(cpu.memory.Read(cpu.pc))
	low := uint16(cpu.memory.Read(cpu.pc + 1))

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

	//fmt.Printf("opcode: 0x%04X\n", opcode.Raw)

	switch opcode.Group {
	case 0x0000:
		c.handle0Group(cpu, opcode)
	case 0x1000:
		fmt.Println("JUMP NNN")
		cpu.pc = opcode.NNN
	case 0x2000:
		fmt.Println("0x2000")
	case 0x3000:
		fmt.Println("0x3000")
	case 0x4000:
		fmt.Println("0x4000")
	case 0x5000:
		fmt.Println("0x5000")
	case 0x6000:
		// v[x] := NN
		fmt.Printf("V%d := %d\n", opcode.X, opcode.NN)
		cpu.v[opcode.X] = opcode.NN
	case 0x7000:
		// v[x] += NN
		fmt.Printf("V%d += %d\n", opcode.X, opcode.NN)
		cpu.v[opcode.X] += opcode.NN
	case 0x8000:
		c.handle8Group(cpu, opcode)
	case 0x9000:
		fmt.Println("0x9000")
	case 0xA000:
		fmt.Println("0xA000")
	case 0xB000:
		fmt.Println("0xB000")
	case 0xC000:
		fmt.Println("0xC000")
	case 0xD000:
		fmt.Println("0xD000")
	case 0xE000:
		fmt.Println("0xE000")
	case 0xF000:
		c.handleFGroup(cpu, opcode)
	}
}

func (c *ControlUnit) handle0Group(cpu *CPU, opcode Opcode) {
	switch opcode.Raw & 0x00FF {
	case 0xE0:
		fmt.Println("Clear")
	case 0xEE:
		fmt.Println("Return")
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
