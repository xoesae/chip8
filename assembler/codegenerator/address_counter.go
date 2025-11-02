package codegenerator

type AddressCounter struct {
	pos uint32
}

func (a *AddressCounter) setPos(pos uint32) {
	a.pos = pos
}

func (a *AddressCounter) advance() {
	a.pos++
}
