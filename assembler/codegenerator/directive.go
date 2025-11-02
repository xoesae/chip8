package codegenerator

import "github.com/xoesae/chip8/assembler/token"

func (c *CodeGenerator) processDirective(expression token.Expression) {
	directive := expression[0].(token.Directive)

	switch directive.Value {
	case string(token.Org):
		addr := expression[1].(token.NumericLiteral)
		c.addressCounter.setPos(addr.Value)
	case string(token.Db):
		for i := 1; i < len(expression); i++ {
			literal := expression[i].(token.NumericLiteral)
			// set the byte on the current addres
			c.opcodes[c.addressCounter.pos] = byte(literal.Value)
			c.addressCounter.advance()
		}
	}
}
