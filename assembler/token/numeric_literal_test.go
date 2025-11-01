package token

import (
	"testing"
)

func TestNumericLiteralKind(t *testing.T) {
	n := NumericLiteral{Value: 10}
	got := n.Kind()
	want := "NumericLiteral"
	if got != want {
		t.Errorf("Kind() = %s; want %s", got, want)
	}
}

func TestNumericLiteralFormat(t *testing.T) {
	n := NumericLiteral{Value: 42}
	got := n.Format()
	want := "NumericLiteral(42)"
	if got != want {
		t.Errorf("Format() = %s; want %s", got, want)
	}
}

func TestIsNumericLiteralTrue(t *testing.T) {
	inputs := []string{"123", "$1A", "$200", "0"}
	for _, input := range inputs {
		if !IsNumericLiteral(input) {
			t.Errorf("IsNumericLiteral(%q) = false; want true", input)
		}
	}
}

func TestIsNumericLiteralFalse(t *testing.T) {
	inputs := []string{"abc", "$", "$ZZ", "", "-10"}
	for _, input := range inputs {
		if IsNumericLiteral(input) {
			t.Errorf("IsNumericLiteral(%q) = true; want false", input)
		}
	}
}

func TestParseNumericLiteralHex(t *testing.T) {
	input := "$200"
	got := ParseNumericLiteral(input)
	want := uint32(512)
	if got != want {
		t.Errorf("ParseNumericLiteral(%q) = %d; want %d", input, got, want)
	}
}

func TestParseNumericLiteralDecimal(t *testing.T) {
	input := "42"
	got := ParseNumericLiteral(input)
	want := uint32(42)
	if got != want {
		t.Errorf("ParseNumericLiteral(%q) = %d; want %d", input, got, want)
	}
}
