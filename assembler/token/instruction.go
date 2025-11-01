package token

import "fmt"

type InstructionType string

const (
	CLS  InstructionType = "CLS"
	RET  InstructionType = "RET"
	JP   InstructionType = "JP"
	CALL InstructionType = "CALL"
	SE   InstructionType = "SE"
	SNE  InstructionType = "SNE"
	LD   InstructionType = "LD"
	ADD  InstructionType = "ADD"
	SUB  InstructionType = "SUB"
	SUBN InstructionType = "SUBN"
	OR   InstructionType = "OR"
	AND  InstructionType = "AND"
	XOR  InstructionType = "XOR"
	SHR  InstructionType = "SHR"
	SHL  InstructionType = "SHL"
	RND  InstructionType = "RND"
	DRW  InstructionType = "DRW"
	SKP  InstructionType = "SKP"
	SKNP InstructionType = "SKNP"
)

type Instruction struct {
	Value string
}

func (i Instruction) Kind() string {
	return "Instruction"
}

func (i Instruction) Format() string {
	return fmt.Sprintf("%s(%s)", i.Kind(), i.Value)
}

func IsInstruction(t string) bool {
	switch t {
	case string(CLS):
		return true
	case string(RET):
		return true
	case string(JP):
		return true
	case string(CALL):
		return true
	case string(SE):
		return true
	case string(SNE):
		return true
	case string(LD):
		return true
	case string(ADD):
		return true
	case string(SUB):
		return true
	case string(SUBN):
		return true
	case string(OR):
		return true
	case string(AND):
		return true
	case string(XOR):
		return true
	case string(SHR):
		return true
	case string(SHL):
		return true
	case string(RND):
		return true
	case string(DRW):
		return true
	case string(SKP):
		return true
	case string(SKNP):
		return true
	default:
		return false
	}
}
