package cpu

import (
	"fmt"
	"time"

	"github.com/xoesae/chip8/emulator/io/display"
	"github.com/xoesae/chip8/emulator/memory"
	"github.com/xoesae/chip8/emulator/shared"
)

type CPU struct {
	// internal
	v          [16]byte   // V0-VF registers
	i          uint16     // address register
	pc         *PC        // program counter
	stack      [16]uint16 // stack
	sp         byte       // StackPointer
	delayTimer byte
	soundTimer byte

	cu     *ControlUnit
	memory *memory.Memory

	pixels  [shared.DisplayHeight][shared.DisplayWidth]bool
	display *display.Display

	hasKeyPressed bool
	keyPressed    uint8

	running bool
}

func NewCPU(mem *memory.Memory, d *display.Display) *CPU {
	i := uint16(0x200)

	pc := NewPC(i)

	return &CPU{
		v:          [16]byte{},
		i:          i,
		pc:         pc,
		stack:      [16]uint16{},
		delayTimer: 0,
		soundTimer: 0,
		memory:     mem,
		cu:         &ControlUnit{},
		display:    d,
		running:    false,
	}
}

func (c *CPU) step() {
	memorySize := c.memory.Size()

	if c.pc.Current() >= memorySize-1 {
		c.running = false
	}

	keyEvents, running := c.display.PollEvents()

	c.running = running
	for _, e := range keyEvents {
		if e.Pressed {
			c.hasKeyPressed = true
			c.keyPressed = e.Key
			break
		}
	}

	c.cu.ExecuteCycle(c) // TODO: remove
	c.display.Render(c.pixels)

	c.pc.Count()
	c.hasKeyPressed = false

	if c.pc.Current() >= memorySize {
		c.running = false
	}
}

func (c *CPU) updateTimers() {
	if c.delayTimer > 0 {
		c.delayTimer--
	}

	if c.soundTimer > 0 {
		c.soundTimer--
	}
}

func (c *CPU) Run(fps int) {
	c.running = true

	cpuTick := time.NewTicker(time.Second / time.Duration(fps)) // FPS = instructions per second
	timerTick := time.NewTicker(time.Second / 60)               // Timers 60Hz

	defer cpuTick.Stop()
	defer timerTick.Stop()

	for c.running {
		select {
		case <-cpuTick.C:
			c.step()
		case <-timerTick.C:
			c.updateTimers()
		}
	}
}

func (c *CPU) PrintRegisters() {
	for i := 0; i < 16; i++ {
		fmt.Printf("V%X = 0x%02X  ", i, c.v[i])
		if (i+1)%4 == 0 {
			fmt.Println()
		}
	}
}
