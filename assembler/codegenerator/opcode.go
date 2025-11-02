package codegenerator

type OpCode struct {
	// msb, lsb
	Bytes [2]byte
}

func NewOpCode(bytes [2]byte) *OpCode {
	return &OpCode{
		Bytes: bytes,
	}
}

// 1NNN, 2NNN, ANNN, BNNN
func NewOpCodePNNN(prefix byte, nnn uint16) *OpCode {
	msb := prefix<<4 | byte((nnn&0xF00)>>8)
	lsb := byte(nnn & 0xFF)
	return &OpCode{Bytes: [2]byte{msb, lsb}}
}

// 3XNN, 4XNN, 6XNN, 7XNN, CXNN
func NewOpCodePXNN(prefix, x, nn byte) *OpCode {
	msb := prefix<<4 | x
	lsb := nn
	return &OpCode{Bytes: [2]byte{msb, lsb}}
}

// 5XY0, 8XY0, 8XY1, ... 8XYE, 9XY0
func NewOpCodePXYS(prefix, x, y, suffix byte) *OpCode {
	msb := prefix<<4 | x
	lsb := (y << 4) | suffix
	return &OpCode{Bytes: [2]byte{msb, lsb}}
}

// EX9E, EXA1, FX07, FX0A, FX15, FX18, FX1E, FX29, FX33, FX55, FX65
func NewOpCodePXSS(prefix, x, suffix byte) *OpCode {
	msb := prefix<<4 | x
	lsb := suffix
	return &OpCode{Bytes: [2]byte{msb, lsb}}
}
