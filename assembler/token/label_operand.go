package token

import (
	"fmt"
	"strings"
)

type LabelOperand struct {
	Value string
}

func (l LabelOperand) Kind() string {
	return "LabelOperand"
}

func (l LabelOperand) Format() string {
	return fmt.Sprintf("%s(%s)", l.Kind(), l.Value)
}

func IsLabelOperand(word string) bool {
	return strings.HasPrefix(word, "#") && len(word) > 1 && !strings.HasPrefix(word[1:], "#")
}

func ParseLabelOperand(word string) string {
	return word[1:] // remove the first #
}
