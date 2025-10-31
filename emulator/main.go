package main

import (
	"os"

	"github.com/xoesae/chip8/emulator/cpu"
	"github.com/xoesae/chip8/emulator/io/display"
	"github.com/xoesae/chip8/emulator/memory"
)

// MemorySize 12 bits
const MemorySize uint16 = 4096

func main() {
	romFile := os.Args[1]

	// memory
	mem := memory.NewMemory(MemorySize)
	mem.Setup(romFile)

	// display
	displ := display.NewDisplay()
	displ.Clear()
	sprite := []byte{0xF0, 0x90, 0x90, 0x90, 0xF0} // "0"
	displ.DrawSprite(10, 5, sprite)
	//displ.Window.ShowAndRun()

	c := cpu.NewCPU(mem, displ)
	go c.Run(100_000_000)

	displ.Run()

	//fmt.Println("\nRegisters")
	//c.PrintRegisters()
	//
	//fmt.Println("\nRAM")
	//mem.Print()
}
