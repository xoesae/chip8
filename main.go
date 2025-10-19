package main

import (
	"fmt"
	"os"

	"github.com/xoesae/chip8/cpu"
	"github.com/xoesae/chip8/memory"
)

// MemorySize 12 bits
const MemorySize uint16 = 4096

func main() {
	romFile := os.Args[1]

	mem := memory.NewMemory(MemorySize)
	mem.Setup(romFile)

	c := cpu.NewCPU(mem)
	c.Run(100_000_000)

	fmt.Println("\nRegisters")
	c.PrintRegisters()

	fmt.Println("\nRAM")
	mem.Print()
}
