package memory

import (
	"fmt"
	"log"
	"os"
)

type Memory struct {
	size   uint16
	memory []byte
}

func NewMemory(size uint16) *Memory {
	return &Memory{
		size:   size,
		memory: make([]byte, size),
	}
}

func (m *Memory) Read(address uint16) byte {
	return m.memory[address]
}

func (m *Memory) Write(address uint16, value byte) {
	m.memory[address] = value
}

func (m *Memory) LoadFontSet() {
	// Load font set in 0x050–0x09F
	initialAddress := 0x050

	f := []byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}

	copy(m.memory[initialAddress:], f)
}

func (m *Memory) LoadProgram(program []byte) {
	initialAddress := 0x200

	// write in memory from 0x200
	copy(m.memory[initialAddress:], program)
}

func (m *Memory) Setup(filename string) {
	m.LoadFontSet()

	// load instructions
	program, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error on read file: %v", err)
	}

	m.LoadProgram(program)
}

func (m *Memory) Size() uint16 {
	return m.size
}

func (m *Memory) Print() {
	for i := 0; i < len(m.memory); i += 16 {
		fmt.Printf("%04X: ", i)

		end := i + 16
		if end > len(m.memory) {
			end = len(m.memory)
		}

		for j := i; j < end; j++ {
			fmt.Printf("%02X ", m.memory[j])
		}

		fmt.Println()
	}
}
