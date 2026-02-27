package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xoesae/chip8/emulator/chip"
	"github.com/xoesae/chip8/emulator/memory"
	"github.com/xoesae/chip8/emulator/platform"
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
	p, err := platform.NewPlatform()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()

	logger.Get().Debug("Starting emulator")
	c := chip.NewChip(mem, p)

	c.Run(60)

	//fmt.Println("\nRegisters")
	//c.PrintRegisters()

	logger.Get().Info("Stopping emulator")

	fmt.Println("\nRAM")
	mem.Print()
}
