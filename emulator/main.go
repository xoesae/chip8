package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/xoesae/chip8/emulator/cpu"
	"github.com/xoesae/chip8/emulator/event"
	"github.com/xoesae/chip8/emulator/io/display"
	"github.com/xoesae/chip8/emulator/memory"
)

const MemorySize uint16 = 4096

func main() {
	romFile := os.Args[1]

	// memory
	mem := memory.NewMemory(MemorySize)
	mem.Setup(romFile)

	// EventBus
	eventChannel := make(chan event.Event)

	c := cpu.NewCPU(mem, eventChannel)
	go c.Run(60)

	// output
	a := app.New()
	_display := display.NewDisplay(&a, eventChannel)
	_display.StartEventLoop()
	_display.Show()
	_display.Run()

	//fmt.Println("\nRegisters")
	//c.PrintRegisters()

	fmt.Println("\nRAM")
	mem.Print()
}
