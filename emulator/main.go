package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xoesae/chip8/emulator/cpu"
	"github.com/xoesae/chip8/emulator/io/display"
	"github.com/xoesae/chip8/emulator/memory"
	"github.com/xoesae/chip8/logger"
)

func main() {
	logger.Init("debug") // debug - info

	romFile := os.Args[1]

	logger.Get().Debug("Loading ROM " + romFile)

	// memory
	mem := memory.NewMemory()
	mem.Setup(romFile)

	logger.Get().Info("Starting display")
	d, err := display.NewDisplay()
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()

	logger.Get().Debug("Starting emulator")
	c := cpu.NewCPU(mem, d)

	c.Run(60)

	//fmt.Println("\nRegisters")
	//c.PrintRegisters()

	logger.Get().Info("Stopping emulator")

	fmt.Println("\nRAM")
	mem.Print()
}
