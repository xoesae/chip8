package codegenerator

import "github.com/xoesae/chip8/assembler/token"

func (c *CodeGenerator) processDirective(expression token.Expression) {
	directive := expression[0].(token.Directive)

	switch directive.Value {
	case string(token.Org):
		addr := mustAs[token.NumericLiteral](expression[1])
		c.addressCounter.setPos(addr.Value)
	case string(token.Db):
		for i := 1; i < len(expression); i++ {
			literal := mustAs[token.NumericLiteral](expression[i])
			// set the byte on the current address
			c.opcodes[c.addressCounter.pos] = byte(literal.Value)
			c.addressCounter.advance()
		}
	}
}
