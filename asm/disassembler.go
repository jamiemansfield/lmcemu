package asm

import (
	"strings"
	"fmt"
	"strconv"
)

func DisassembleInstruction(instruction int) (*EvaluatedInstruction, error) {
	// Convert to a string so I can splice
	var strOc = strings.Replace(fmt.Sprintf("%3d", instruction), " ", "0", -1)

	// Get the instruction code
	inst, err := strconv.Atoi(strOc[:1])
	if err != nil {
		return nil, err
	}

	// Get the memory address
	addr, err := strconv.Atoi(strOc[1:])
	if err != nil {
		return nil, err
	}

	return CreateEvaluatedInstruction(INST_NORMAL, Opcode(inst), addr, -1), nil
}
