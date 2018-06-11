package emu

import (
	"github.com/jamiemansfield/lmcemu/asm"
)

// This file has a number of functions that can be associated with the
// Control Unit of a CPU.

// Gets the value of the given memory address, as an EvaluatedLine.
func DecodeInstruction(memory *Memory, address int) (*asm.EvaluatedInstruction, error) {
	return asm.DisassembleInstruction(memory.GetValue(address))
}
