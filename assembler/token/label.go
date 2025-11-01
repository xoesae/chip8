package token

import "fmt"

type Label struct {
	Value string
}

func (l Label) Kind() string {
	return "Label"
}

func (l Label) Format() string {
	return fmt.Sprintf("%s(%s)", l.Kind(), l.Value)
}
