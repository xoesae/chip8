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
