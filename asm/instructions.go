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
	if i.AddressRef.Address <= -1 {
		if i.AddressRef.Type == ADDR_LABEL_POSITION {
			return nil, errors.New("Undefined labeled position!")
		}
		if i.AddressRef.Type == ADDR_LABEL_DATA {
			return nil, errors.New("Undefined labeled data!")
		}
		return nil, errors.New("Invalid instruction address of '" + strconv.Itoa(i.AddressRef.Address) + "'!")
	}
	if i.Type == INST_LABELED && i.AddressRef.Value <= -1 {
		return nil, errors.New("Invalid instruction value of '" + strconv.Itoa(i.AddressRef.Value) + "'!")
	}

	return CreateEvaluatedInstruction(i.Type, i.Opcode, i.AddressRef.Address, i.AddressRef.Value), nil
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
	Type InstructionType // TODO: eliminate type on evaluated instructions
	Opcode Opcode
	Address int
	Value int
}

// Compiles the instruction.
func (i *EvaluatedInstruction) Compile() int {
	if i.Opcode == OP_DAT {
		return i.Value
	}
	return int(i.Opcode) * 100 + i.Address
}

func CreateEvaluatedInstruction(ttype InstructionType, opcode Opcode, address int, value int) *EvaluatedInstruction {
	return &EvaluatedInstruction{
		Type: ttype,
		Opcode: opcode,
		Address: address,
		Value: value,
	}
}
