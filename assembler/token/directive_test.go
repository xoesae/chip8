package token

import (
	"testing"
)

func TestDirectiveKind(t *testing.T) {
	d := Directive{Value: "org"}
	got := d.Kind()
	want := "Directive"
	if got != want {
		t.Errorf("Kind() = %s; want %s", got, want)
	}
}

func TestDirectiveFormat(t *testing.T) {
	d := Directive{Value: "db"}
	got := d.Format()
	want := "Directive(db)"
	if got != want {
		t.Errorf("Format() = %s; want %s", got, want)
	}
}
