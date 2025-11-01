package token

import (
	"testing"
)

func TestLabelOperandKind(t *testing.T) {
	l := LabelOperand{Value: "num1"}
	got := l.Kind()
	want := "LabelOperand"
	if got != want {
		t.Errorf("Kind() = %s; want %s", got, want)
	}
}

func TestLabelOperandFormat(t *testing.T) {
	l := LabelOperand{Value: "num1"}
	got := l.Format()
	want := "LabelOperand(num1)"
	if got != want {
		t.Errorf("Format() = %s; want %s", got, want)
	}
}

func TestIsLabelOperandTrue(t *testing.T) {
	cases := []string{"#num1", "#LOOP", "#a", "#_"}

	for _, input := range cases {
		if !IsLabelOperand(input) {
			t.Errorf("IsLabelOperand(%q) = false; want true", input)
		}
	}
}

func TestIsLabelOperandFalse(t *testing.T) {
	cases := []string{"num1", "", "#", "##dup", "123"}

	for _, input := range cases {
		if IsLabelOperand(input) {
			t.Errorf("IsLabelOperand(%q) = true; want false", input)
		}
	}
}

func TestParseLabelOperand(t *testing.T) {
	if got := ParseLabelOperand("#num1"); got != "num1" {
		t.Errorf(`ParseLabelOperand("#num1") = %q; want "num1"`, got)
	}
	if got := ParseLabelOperand("#LOOP"); got != "LOOP" {
		t.Errorf(`ParseLabelOperand("#LOOP") = %q; want "LOOP"`, got)
	}
}
