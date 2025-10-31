package lexer

import (
	"bufio"
	"io"
	"strings"
	"unicode"

	"github.com/xoesae/chip8/assembler/lexer/token"
)

type Lexer struct {
	pos    Position
	reader *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    Position{line: 1, column: 1},
		reader: bufio.NewReader(reader),
	}
}

func (l *Lexer) isEOF() bool {
	_, _, err := l.reader.ReadRune()
	l.pos.NextColumn()
	if err == io.EOF {
		return true
	}

	err = l.reader.UnreadRune()
	l.pos.PreviousColumn()
	if err != nil {
		panic(err)
	}

	return false
}

func (l *Lexer) readWord() string {
	var sb strings.Builder

	for {
		// read current rune (char)
		char, _, err := l.reader.ReadRune()
		l.pos.NextColumn()

		if err == io.EOF {
			return "EOF"
		}

		if err != nil {
			panic(err)
		}

		if char == ',' {
			if sb.Len() > 0 {
				break
			}

			sb.Reset()
			l.pos.NextColumn()
			break
		}

		if char == ';' {
			if sb.Len() > 0 {
				err = l.reader.UnreadRune()
				l.pos.PreviousColumn()
				if err != nil {
					panic(err)
				}

				break
			}

			_, _, err := l.reader.ReadLine()
			if err == io.EOF {
				return "EOF"
			}
			if err != nil {
				panic(err)
			}

			l.pos.NextLine()
			break
		}

		if char == '\n' {

			if sb.Len() > 0 {
				err = l.reader.UnreadRune()
				l.pos.PreviousColumn()
				if err != nil {
					panic(err)
				}

				break
			}

			l.pos.NextLine()
			break
		}

		if unicode.IsSpace(char) {
			// if cursor is end of line
			if sb.Len() > 0 {
				err = l.reader.UnreadRune()
				l.pos.PreviousColumn()
				if err != nil {
					panic(err)
				}
			}

			break
		}

		isLetter := unicode.IsLetter(char)
		isDigit := unicode.IsDigit(char)
		isSpecial := char == '_' || char == '$' || char == '#'

		if !isLetter && !isDigit && !isSpecial {
			err = l.reader.UnreadRune()
			l.pos.PreviousColumn()
			if err != nil {
				panic(err)
			}

			break
		}

		// append char to string (word)
		sb.WriteRune(char)

		// checks if the next character is the EOF and consume the current word
		_, err = l.reader.Peek(1)
		if err == io.EOF {
			break
		}
	}

	return sb.String()
}

func (l *Lexer) NextToken() token.Token {
	word := l.readWord()

	switch word {
	case "EOF":
		return token.EOF{}
	case string(token.Org), string(token.Db):
		return token.Directive{Value: word}
	default:
		if len(word) > 0 {
			if token.IsInstruction(word) {
				return token.Instruction{Value: word}
			}

			if token.IsRegister(word) {
				return token.Register{Value: word}
			}

			if token.IsNumericLiteral(word) {
				return token.NumericLiteral{Value: token.ParseNumericLiteral(word)}
			}

			if token.IsLabelOperand(word) {
				return token.LabelOperand{Value: word}
			}

			return token.Label{Value: word}
		}
	}

	return nil
}

// Lex scans the input for the next token. It returns the position of the token,
// the token's type, and the literal value.
func (l *Lexer) Lex() []token.Token {
	var tokens []token.Token

	for {
		tkn := l.NextToken()

		if tkn == nil {
			continue
		}

		if _, ok := tkn.(token.EOF); ok {
			tokens = append(tokens, tkn)
			break
		}

		tokens = append(tokens, tkn)
	}

	return tokens
}
