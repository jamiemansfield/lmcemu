package emu

import (
	"github.com/jamiemansfield/lmcemu/asm"
)

type Instruction func(cpu *CPU, memory *Memory) bool

type CPU struct {
	// Registers
	ProgramCounter *Register
	InstructionRegister *Register
	AddressRegister *Register
	Accumulator *Register

	// Other
	Instructions map[asm.Opcode]Instruction
}

func (c *CPU) Execute(memory *Memory) {
	// https://en.wikipedia.org/wiki/Little_man_computer#Execution_cycle

	var running = true
	for ; running ; {
		// 1. Decode instruction
		var line = memory.GetValueAsLine(c.ProgramCounter.GetValue())
		c.InstructionRegister.SetValue(int(line.Opcode))
		c.AddressRegister.SetValue(line.Address)

		// 2. Increment PC
		c.ProgramCounter.Increment()

		// 3. Prep for execute
		var inst = c.Instructions[asm.Opcode(c.InstructionRegister.GetValue())]

		// 4. Execute
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
		Instructions: map[asm.Opcode]Instruction{
			asm.HLT: inst_hlt,
			asm.ADD: inst_add,
			asm.SUB: inst_sub,
			asm.STA: inst_sta,
			// There is no 4
			asm.LDA: inst_lda,
			asm.BRA: inst_bra,
			asm.BRZ: inst_brz,
			asm.BRP: inst_brp,
			asm.INP_OUT: inst_inp_out,
		},
	}
}
