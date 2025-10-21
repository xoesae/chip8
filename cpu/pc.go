package cpu

type PC struct {
	current uint16
}

func NewPC() *PC {
	return &PC{current: 0x200}
}

func (p *PC) Count() {
	p.current += 2
}

func (p *PC) JumpTo(address uint16) {
	p.current = address
}

func (p *PC) Current() uint16 {
	return p.current
}
