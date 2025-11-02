package codegenerator

import (
	"testing"
)

func TestAddressCounterSetPos(t *testing.T) {
	a := &AddressCounter{}
	a.setPos(0x200)
	if a.pos != 0x200 {
		t.Errorf("setPos(0x200): pos = 0x%X, expected 0x200", a.pos)
	}

	a.setPos(0xFFF)
	if a.pos != 0xFFF {
		t.Errorf("setPos(0xFFF): pos = 0x%X, expected 0xFFF", a.pos)
	}
}

func TestAddressCounterAdvance(t *testing.T) {
	a := &AddressCounter{pos: 0x300}
	a.advance()
	if a.pos != 0x301 {
		t.Errorf("advance: pos = 0x%X, expected 0x301", a.pos)
	}

	a.advance()
	if a.pos != 0x302 {
		t.Errorf("second advance: pos = 0x%X, expected 0x302", a.pos)
	}
}
