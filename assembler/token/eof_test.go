package token

import (
	"testing"
)

func TestEOFKind(t *testing.T) {
	e := EOF{}
	got := e.Kind()
	want := "EOF"
	if got != want {
		t.Errorf("Kind() = %s; want %s", got, want)
	}
}

func TestEOFFormat(t *testing.T) {
	e := EOF{}
	got := e.Format()
	want := "EOF"
	if got != want {
		t.Errorf("Format() = %s; want %s", got, want)
	}
}
