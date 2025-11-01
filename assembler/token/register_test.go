package token

import (
	"testing"
)

func TestRegisterKind(t *testing.T) {
	r := Register{Value: "V0"}
	got := r.Kind()
	want := "Register"
	if got != want {
		t.Errorf("Kind() = %s; want %s", got, want)
	}
}

func TestRegisterFormat(t *testing.T) {
	r := Register{Value: "V0"}
	got := r.Format()
	want := "Register(V0)"
	if got != want {
		t.Errorf("Format() = %s; want %s", got, want)
	}
}

func TestIsRegisterTrue(t *testing.T) {
	registers := []string{
		"V0", "V1", "V2", "V3", "V4", "V5", "V6", "V7",
		"V8", "V9", "VA", "VB", "VC", "VD", "VE", "VF",
		"I", "DT", "ST", "F", "[I]", "K", "B",
	}
	for _, reg := range registers {
		if !IsRegister(reg) {
			t.Errorf("IsRegister(%q) = false; want true", reg)
		}
	}
}

func TestIsRegisterFalse(t *testing.T) {
	inputs := []string{
		"XX", "v0", "X", "V10", "", "123", "ABC", "VZ", "[J]", "i", "dT", "st", "Va", "vb",
	}
	for _, input := range inputs {
		if IsRegister(input) {
			t.Errorf("IsRegister(%q) = true; want false", input)
		}
	}
}
