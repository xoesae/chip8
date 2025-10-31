package token

import "fmt"

type RegisterType string

const (
	V0 RegisterType = "V0"
	V1 RegisterType = "V1"
	V2 RegisterType = "V2"
	V3 RegisterType = "V3"
	V4 RegisterType = "V4"
	V5 RegisterType = "V5"
	V6 RegisterType = "V6"
	V7 RegisterType = "V7"
	V8 RegisterType = "V8"
	V9 RegisterType = "V9"
	VA RegisterType = "VA"
	VB RegisterType = "VB"
	VC RegisterType = "VC"
	VD RegisterType = "VD"
	VE RegisterType = "VE"
	VF RegisterType = "VF"
	I  RegisterType = "I"
	DT RegisterType = "DT"
	ST RegisterType = "ST"
	F  RegisterType = "F"
	VI RegisterType = "[I]"
)

type Register struct {
	Value string
}

func (r Register) Kind() string {
	return "Register"
}

func (r Register) Format() string {
	return fmt.Sprintf("%s(%s)", r.Kind(), r.Value)
}

func IsRegister(t string) bool {
	switch t {
	case string(V0):
		return true
	case string(V1):
		return true
	case string(V2):
		return true
	case string(V3):
		return true
	case string(V4):
		return true
	case string(V5):
		return true
	case string(V6):
		return true
	case string(V7):
		return true
	case string(V8):
		return true
	case string(V9):
		return true
	case string(VA):
		return true
	case string(VB):
		return true
	case string(VC):
		return true
	case string(VD):
		return true
	case string(VE):
		return true
	case string(VF):
		return true
	case string(I):
		return true
	case string(DT):
		return true
	case string(ST):
		return true
	case string(F):
		return true
	case string(VI):
		return true
	default:
		return false
	}
}
