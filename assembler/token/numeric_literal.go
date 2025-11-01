package token

import (
	"fmt"
	"strconv"
	"strings"
)

type NumericLiteral struct {
	Value uint32
}

func (n NumericLiteral) Kind() string {
	return "NumericLiteral"
}

func (n NumericLiteral) Format() string {
	return fmt.Sprintf("%s(%d)", n.Kind(), n.Value)
}

func IsNumericLiteral(word string) bool {
	// Hex
	if strings.HasPrefix(word, "$") && len(word) > 1 {
		_, err := strconv.ParseUint(word[1:], 16, 32)
		return err == nil
	}

	// Decimal
	_, err := strconv.ParseUint(word, 10, 32)
	return err == nil
}

func ParseNumericLiteral(word string) uint32 {
	// Hex
	if strings.HasPrefix(word, "$") {
		val, _ := strconv.ParseUint(word[1:], 16, 32)
		return uint32(val)
	}

	// Decimal
	val, _ := strconv.ParseUint(word, 10, 32)
	return uint32(val)
}
