package main

import (
	"os"

	"fyne.io/fyne/v2/app"
	"github.com/xoesae/chip8/emulator/cpu"
	"github.com/xoesae/chip8/emulator/event"
	"github.com/xoesae/chip8/emulator/io/display"
	"github.com/xoesae/chip8/emulator/memory"
	"github.com/xoesae/chip8/logger"
)

const MemorySize uint16 = 4096

func main() {

	logger.Init("debug") // debug - info

	romFile := os.Args[1]

	logger.Get().Debug("Loading ROM " + romFile)

	// memory
	mem := memory.NewMemory(MemorySize)
	mem.Setup(romFile)

	// EventBus
	eventChannel := make(chan event.Event)

	logger.Get().Debug("Starting emulator")
	c := cpu.NewCPU(mem, eventChannel)
	go c.Run(60)

	// output

	logger.Get().Info("Starting display")
	a := app.New()
	_display := display.NewDisplay(&a, eventChannel)
	_display.StartEventLoop()
	_display.Show()
	_display.Run()

	//fmt.Println("\nRegisters")
	//c.PrintRegisters()

	logger.Get().Info("Stopping emulator")

	//fmt.Println("\nRAM")
	//mem.Print()
}
