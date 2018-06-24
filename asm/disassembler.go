package asm

import (
	"strings"
	"fmt"
	"strconv"
)

// Disassemble machine code into the form of an /evaluated/ instruction.
func DisassembleInstruction(instruction int) (*EvaluatedInstruction, error) {
	// Convert to a string so I can splice
	var strOc = strings.Replace(fmt.Sprintf("%3d", instruction), " ", "0", -1)

	// Get the instruction code
	inst, err := strconv.Atoi(strOc[:1])
	if err != nil {
		return nil, fmt.Errorf("disassembler: failed to disassemble '%d'", instruction)
	}

	// Get the memory address
	addr, err := strconv.Atoi(strOc[1:])
	if err != nil {
		return nil, fmt.Errorf("disassembler: failed to disassemble '%d'", instruction)
	}

	return CreateEvaluatedInstruction(Opcode(inst), addr), nil
}
