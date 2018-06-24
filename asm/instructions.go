package asm

import (
	"fmt"
)

// The type of an instruction, used for establishing whether the label
// of an instruction needs to be evaluated when assembling.
type InstructionType int

const (
	INST_NORMAL  InstructionType = iota
	INST_LABELED
)

// Represents un-evaluated instructions.
type Instruction struct {
	Type InstructionType `json:"type"`
	Opcode Opcode `json:"opcode"`
	AddressRef *AddressRef `json:"ref"`

	// Labeled
	Label *AddressRef `json:"label"`
}

// Evaluates an instruction.
// NOTE: The address reference should have /already/ been evaluated
func (i *Instruction) Evaluate() (*EvaluatedInstruction, error) {
	// Check for errors
	if !IsValidAddress(i.AddressRef.Address) {
		if i.AddressRef.Type == ADDR_LABEL_POSITION {
			return nil, fmt.Errorf("compiler: unevaluated reference to labeled position '%d' used as an argument", i.AddressRef.Address)
		}
		if i.AddressRef.Type == ADDR_LABEL_DATA {
			return nil, fmt.Errorf("compiler: unevaluated reference to labeled data '%d' used as an argument", i.AddressRef.Address)
		}
		return nil, fmt.Errorf("compiler: invalid address '%d' used", i.AddressRef.Address)
	}

	// Handle Data properly
	if i.Opcode == OP_DAT {
		if !IsValidAddress(i.AddressRef.Value) {
			return nil, fmt.Errorf("compiler: invalid data value '%d'", i.AddressRef.Address)
		}
		return DisassembleInstruction(i.AddressRef.Value)
	}

	return CreateEvaluatedInstruction(i.Opcode, i.AddressRef.Address), nil
}

func CreateInstruction(opcode Opcode, address *AddressRef) *Instruction {
	return &Instruction{
		Type:       INST_NORMAL,
		Opcode:     opcode,
		AddressRef: address,
	}
}

// Represents an evaluated instruction.
type EvaluatedInstruction struct {
	Opcode Opcode
	Address int
}

// Compiles the instruction.
func (i *EvaluatedInstruction) Compile() int {
	return int(i.Opcode) * 100 + i.Address
}

func CreateEvaluatedInstruction(opcode Opcode, address int) *EvaluatedInstruction {
	return &EvaluatedInstruction{
		Opcode: opcode,
		Address: address,
	}
}
