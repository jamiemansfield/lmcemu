package main

import (
	"github.com/jamiemansfield/lmcemu/emu"
	"strconv"
	"fmt"
	"os"
)

const (
	PC = "PC"
	IR = "IR"
	AR = "AR"
	ACC = "ACC"
)

func main() {
	cpu := CreateLmcCpu()
	memory := emu.Memory{
		Values: [100]int{
			901, // INP
			399, // STA 99
			901, // INP
			199, // ADD 99
			902, // OUT
			000, // HLT
		},
	}
	cpu.Execute(&memory)
}

func CreateLmcCpu() *emu.CPU {
	// Create registers
	var pc = emu.CreateRegister(PC, 0)
	var ir = emu.CreateRegister(IR, 0)
	var ar = emu.CreateRegister(AR, 0)
	var acc = emu.CreateRegister(ACC, 0)

	// Create CPU
	return &emu.CPU{
		// Registers
		ProgramCounter: pc,
		InstructionRegister: ir,
		AddressRegister: ar,
		Accumulator: acc,

		// Other
		Instructions: map[int]emu.Instruction{
			0: inst_hlt,
			1: inst_add,
			2: inst_sub,
			3: inst_sto,
			// There is no 4
			5: inst_lda,
			6: inst_bra,
			7: inst_brz,
			8: inst_brp,
			9: inst_inp_out,
		},
	}
}

// See https://peterhigginson.co.uk/LMC/help.html

func inst_hlt (cpu *emu.CPU, memory *emu.Memory) bool {
	return false
}

func inst_add (cpu *emu.CPU, memory *emu.Memory) bool {
	var accValue = cpu.Accumulator.GetValue()
	var value = memory.GetValue(cpu.AddressRegister.GetValue())
	cpu.Accumulator.SetValue(accValue + value)
	return true
}

func inst_sub (cpu *emu.CPU, memory *emu.Memory) bool {
	var accValue = cpu.Accumulator.GetValue()
	var value = memory.GetValue(cpu.AddressRegister.GetValue())
	cpu.Accumulator.SetValue(accValue - value)
	return true
}

func inst_sto (cpu *emu.CPU, memory *emu.Memory) bool {
	var accValue = cpu.Accumulator.GetValue()
	var address = cpu.AddressRegister.GetValue()
	memory.SetValue(address, accValue)
	return true
}

func inst_lda (cpu *emu.CPU, memory *emu.Memory) bool {
	cpu.Accumulator.SetValue(cpu.AddressRegister.GetValue())
	return true
}

func inst_bra (cpu *emu.CPU, memory *emu.Memory) bool {
	cpu.ProgramCounter.SetValue(cpu.AddressRegister.GetValue())
	return true
}

func inst_brz (cpu *emu.CPU, memory *emu.Memory) bool {
	if cpu.Accumulator.GetValue() == 0 {
		cpu.ProgramCounter.SetValue(cpu.AddressRegister.GetValue())
	}
	return true
}

func inst_brp (cpu *emu.CPU, memory *emu.Memory) bool {
	if cpu.Accumulator.GetValue() >= 0 {
		cpu.ProgramCounter.SetValue(cpu.AddressRegister.GetValue())
	}
	return true
}

func inst_inp_out (cpu *emu.CPU, memory *emu.Memory) bool {
	if cpu.AddressRegister.GetValue() == 1 {
		// INP
		println("inp:")
		var inpRaw = ""
		fmt.Scanln(&inpRaw)

		inp, err := strconv.Atoi(inpRaw)
		if err != nil {
			panic(err)
			os.Exit(-1)
		}
		cpu.Accumulator.SetValue(inp)
	} else
	if cpu.AddressRegister.GetValue() == 2 {
		// OUT
		println("out: " + strconv.Itoa(cpu.Accumulator.GetValue()))
	}
	return true
}
