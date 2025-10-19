package cpu

import (
	"fmt"
	"time"

	"github.com/xoesae/chip8/memory"
)

type CPU struct {
	v          [16]byte // V0-VF registers
	i          uint16   // address register
	pc         uint16   // program counter
	stack      [16]uint16
	sp         byte // StackPointer
	delayTimer byte
	soundTimer byte
	memory     *memory.Memory
	cu         *ControlUnit

	running bool
}

func NewCPU(mem *memory.Memory) *CPU {
	return &CPU{
		v:          [16]byte{},
		i:          0x300,
		pc:         0x200,
		stack:      [16]uint16{},
		delayTimer: 0,
		soundTimer: 0,
		memory:     mem,
		running:    false,
	}
}

func (c *CPU) step() {
	memorySize := c.memory.Size()

	if c.pc >= memorySize-1 {
		c.running = false
	}

	c.cu.ExecuteCycle(c)

	if c.pc+2 >= memorySize {
		c.running = false
		return
	}

	c.pc += 2
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
