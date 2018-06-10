package emu

import (
	"github.com/jamiemansfield/lmcemu/asm"
	"strings"
	"fmt"
	"strconv"
	"os"
)

// This file has a number of functions that can be associated with the
// Control Unit of a CPU.

// Gets the value of the given memory address, as an EvaluatedLine.
func DecodeInstruction(memory *Memory, address int) asm.EvaluatedLine {
	// Convert to a string so I can splice
	var strOc = strings.Replace(fmt.Sprintf("%3d", memory.GetValue(address)), " ", "0", -1)

	// Get the instruction code
	inst, err := strconv.Atoi(strOc[:1])
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	// Get the memory address
	addr, err := strconv.Atoi(strOc[1:])
	if err != nil {
		panic(err)
		os.Exit(-1)
	}

	return asm.CreateEvaluatedLine(asm.Opcode(inst), addr)
}
