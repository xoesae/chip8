package cpu

type Opcode struct {
	Raw uint16

	Group uint16
	X     byte
	Y     byte
	N     byte
	NN    byte
	NNN   uint16
}

func NewOpcode(high, low byte) Opcode {
	opcode := (uint16(high) << 8) | uint16(low)

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
