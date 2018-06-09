package emu

import (
	"strconv"
	"os"
	"fmt"
	"strings"
)

type Instruction func(cpu *CPU, memory *Memory) bool

type CPU struct {
	// Registers
	ProgramCounter *Register
	InstructionRegister *Register
	AddressRegister *Register
	Accumulator *Register

	// Other
	Instructions map[Opcode]Instruction
}

func (c *CPU) Execute(memory *Memory) {
	// https://en.wikipedia.org/wiki/Little_man_computer#Execution_cycle

	var running = true
	for ; running ; {
		// 1. Get address of next opcode
		var ocAddr = c.ProgramCounter.GetValue()

		// 2. Get next opcode
		var oc = memory.GetValue(ocAddr)

		// 3. Increment PC
		c.ProgramCounter.Increment()

		// 4. Decode instruction
		{
			// Convert to a string so I can splice
			var strOc = strings.Replace(fmt.Sprintf("%3d", oc), " ", "0", -1)
			// Get the instruction code
			{
				inst, err := strconv.Atoi(strOc[:1])
				if err != nil {
					panic(err)
					os.Exit(-1)
				}
				c.InstructionRegister.SetValue(inst)
			}
			// Get the memory address
			{
				addr, err := strconv.Atoi(strOc[1:])
				if err != nil {
					panic(err)
					os.Exit(-1)
				}
				c.AddressRegister.SetValue(addr)
			}
		}

		// 5. Prep for execute
		var inst = c.Instructions[Opcode(c.InstructionRegister.GetValue())]

		// 6. Execute
		running = inst(c, memory);
	}
}

func CreateLmcCpu() *CPU {
	return &CPU{
		// Registers
		ProgramCounter: CreateRegister(0),
		InstructionRegister: CreateRegister(0),
		AddressRegister: CreateRegister(0),
		Accumulator: CreateRegister(0),

		// Other
		Instructions: map[Opcode]Instruction{
			HLT: inst_hlt,
			ADD: inst_add,
			SUB: inst_sub,
			STA: inst_sta,
			// There is no 4
			LDA: inst_lda,
			BRA: inst_bra,
			BRZ: inst_brz,
			BRP: inst_brp,
			INP_OUT: inst_inp_out,
		},
	}
}
