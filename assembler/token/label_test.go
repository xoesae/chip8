package token

import (
	"testing"
)

func TestLabelKind(t *testing.T) {
	l := Label{Value: "start"}
	got := l.Kind()
	want := "Label"
	if got != want {
		t.Errorf("Kind() = %s; want %s", got, want)
	}
}

func TestLabelFormat(t *testing.T) {
	l := Label{Value: "start"}
	got := l.Format()
	want := "Label(start)"
	if got != want {
		t.Errorf("Format() = %s; want %s", got, want)
	}
}
