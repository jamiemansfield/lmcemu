package emu

import (
	"github.com/jamiemansfield/lmcemu/asm"
	"strings"
	"fmt"
	"strconv"
	"os"
)

// A representation of RAM within an emulated system.
type Memory struct {
	values [100]int
}

// Gets the integer value of a given memory address.
func (m *Memory) GetValue(address int) int {
	return m.values[address]
}

// Gets the value of the given memory address, as a Line.
func (m *Memory) GetValueAsLine(address int) asm.Line {
	// Convert to a string so I can splice
	var strOc = strings.Replace(fmt.Sprintf("%3d", m.values[address]), " ", "0", -1)

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

	return asm.CreateLine(asm.Opcode(inst), addr)
}

// Sets the integer value of a given memory address.
func (m *Memory) SetValue(address int, value int) {
	m.values[address] = value
}

// Creates memory with the provided initial values
func CreateMemory(initialValues [100]int) *Memory {
	return &Memory{
		values: initialValues,
	}
}
