package new_asm

import (
	"errors"
	"strconv"
)

type InstructionType int

const (
	NORMAL  InstructionType = iota
	LABELED
)

// Represents un-evaluated instructions.
type Instruction struct {
	Type InstructionType
	Opcode Opcode

	// Un-evaluated
	AddressRef *AddressRef

	// Labeled
	Label *AddressRef

	// Evaluated
	Address int
	Value int
}

// Evaluates an instruction.
// NOTE: The address reference should have /already/ been evaluated
func (i *Instruction) Evaluate() (*EvaluatedInstruction, error) {
	// Check for errors
	if i.Address <= -1 {
		return nil, errors.New("Invalid instruction address of '" + strconv.Itoa(i.Address) + "'!")
	}
	if i.Type == LABELED && i.Value <= -1 {
		return nil, errors.New("Invalid instruction value of '" + strconv.Itoa(i.Value) + "'!")
	}

	return CreateEvaluatedInstruction(i.Type, i.Opcode, i.AddressRef.Address, i.Value), nil
}

func CreateInstruction(opcode Opcode, address *AddressRef) *Instruction {
	return &Instruction{
		Type:       NORMAL,
		Opcode:     opcode,
		AddressRef: address,
		Value:      address.Address,
	}
}

// Represents an evaluated instruction.
type EvaluatedInstruction struct {
	Type InstructionType
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
