package cpu

type PC struct {
	current, i uint16
}

func NewPC(i uint16) *PC {
	return &PC{
		current: i,
		i:       i,
	}
}

func (p *PC) Count() {
	p.current += 2
}

func (p *PC) JumpTo(address uint16) {
	p.current = address + p.i
}

func (p *PC) Current() uint16 {
	return p.current
}
