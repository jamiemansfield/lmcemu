package asm

import (
	"errors"
	"strconv"
)

type InstructionType int

const (
	INST_NORMAL  InstructionType = iota
	INST_LABELED
)

// Represents un-evaluated instructions.
type Instruction struct {
	Type InstructionType
	Opcode Opcode
	AddressRef *AddressRef

	// Labeled
	Label *AddressRef
}

// Evaluates an instruction.
// NOTE: The address reference should have /already/ been evaluated
func (i *Instruction) Evaluate() (*EvaluatedInstruction, error) {
	// Check for errors
	if !IsValidAddress(i.AddressRef.Address) {
		if i.AddressRef.Type == ADDR_LABEL_POSITION {
			return nil, errors.New("Undefined labeled position!")
		}
		if i.AddressRef.Type == ADDR_LABEL_DATA {
			return nil, errors.New("Undefined labeled data!")
		}
		return nil, errors.New("Invalid instruction address of '" + strconv.Itoa(i.AddressRef.Address) + "'!")
	}
	if i.Type == INST_LABELED && !IsValidAddress(i.AddressRef.Value) {
		return nil, errors.New("Invalid instruction value of '" + strconv.Itoa(i.AddressRef.Value) + "'!")
	}

	// Handle Data properly
	if i.Opcode == OP_DAT {
		if !IsValidAddress(i.AddressRef.Value) {
			return nil, errors.New("Invalid default data value of '" + strconv.Itoa(i.AddressRef.Value) + "'!")
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
