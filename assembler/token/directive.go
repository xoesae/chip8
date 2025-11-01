package token

import "fmt"

type DirectiveType string

const (
	Org DirectiveType = "org"
	Db  DirectiveType = "db"
)

type Directive struct {
	Value string
}

func (d Directive) Kind() string {
	return "Directive"
}

func (d Directive) Format() string {
	return fmt.Sprintf("%s(%s)", d.Kind(), d.Value)
}
