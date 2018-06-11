package emu

import (
	"github.com/jamiemansfield/lmcemu/new-asm"
)

// This file has a number of functions that can be associated with the
// Control Unit of a CPU.

// Gets the value of the given memory address, as an EvaluatedLine.
func DecodeInstruction(memory *Memory, address int) (*new_asm.EvaluatedInstruction, error) {
	return new_asm.DisassembleInstruction(memory.GetValue(address))
}
