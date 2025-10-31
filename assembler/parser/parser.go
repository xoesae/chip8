package parser

import (
	"github.com/xoesae/chip8/assembler/lexer/token"
)

type Expression []token.Token

func isOperand(t token.Token) bool {
	switch t.(type) {
	case token.Register, token.NumericLiteral, token.LabelOperand:
		return true
	}
	return false
}

type Parser struct {
	tokens []token.Token
	size   int
	pos    int
}

func NewParser(tokens []token.Token) *Parser {
	return &Parser{tokens: tokens, size: len(tokens), pos: 0}
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.pos]
}

func (p *Parser) advance() {
	p.pos++
}

func (p *Parser) canConsume() bool {
	return p.pos < p.size
}

func (p *Parser) parseLabel() (Expression, bool) {
	if t, ok := p.peek().(token.Label); ok {
		p.advance()
		return Expression{t}, true
	}

	return nil, false
}

func (p *Parser) parseDirective() (Expression, bool) {
	if t, ok := p.peek().(token.Directive); ok {
		expr := Expression{t}
		p.advance()

		for p.canConsume() {
			if n, ok := p.peek().(token.NumericLiteral); ok {
				expr = append(expr, n)
				p.advance()
				continue
			}

			break
		}
		return expr, true
	}

	return nil, false
}

func (p *Parser) parseInstruction() (Expression, bool) {
	if t, ok := p.peek().(token.Instruction); ok {
		expr := Expression{t}
		p.advance()

		// first
		if p.canConsume() && isOperand(p.peek()) {
			expr = append(expr, p.peek())
			p.advance()
		}

		// second
		if p.canConsume() && isOperand(p.peek()) {
			expr = append(expr, p.peek())
			p.advance()
		}

		// third
		if p.canConsume() && isOperand(p.peek()) {
			expr = append(expr, p.peek())
			p.advance()
		}

		return expr, true
	}

	return nil, false
}

func (p *Parser) Parse() []Expression {
	var expressions []Expression

	for p.canConsume() {
		if expr, ok := p.parseDirective(); ok {
			expressions = append(expressions, expr)
			continue
		}
		if expr, ok := p.parseLabel(); ok {
			expressions = append(expressions, expr)
			continue
		}
		if expr, ok := p.parseInstruction(); ok {
			expressions = append(expressions, expr)
			continue
		}
		p.advance()
	}

	return expressions
}
