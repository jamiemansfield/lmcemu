package emu

import (
	"fmt"
	"strconv"
	"os"
)

// See https://peterhigginson.co.uk/LMC/help.html

func inst_hlt (cpu *CPU, memory *Memory) bool {
	return false
}

func inst_add (cpu *CPU, memory *Memory) bool {
	var accValue = cpu.Accumulator.GetValue()
	var value = memory.GetValue(cpu.AddressRegister.GetValue())
	cpu.Accumulator.SetValue(accValue + value)
	return true
}

func inst_sub (cpu *CPU, memory *Memory) bool {
	var accValue = cpu.Accumulator.GetValue()
	var value = memory.GetValue(cpu.AddressRegister.GetValue())
	cpu.Accumulator.SetValue(accValue - value)
	return true
}

func inst_sto (cpu *CPU, memory *Memory) bool {
	var accValue = cpu.Accumulator.GetValue()
	var address = cpu.AddressRegister.GetValue()
	memory.SetValue(address, accValue)
	return true
}

func inst_lda (cpu *CPU, memory *Memory) bool {
	cpu.Accumulator.SetValue(cpu.AddressRegister.GetValue())
	return true
}

func inst_bra (cpu *CPU, memory *Memory) bool {
	cpu.ProgramCounter.SetValue(cpu.AddressRegister.GetValue())
	return true
}

func inst_brz (cpu *CPU, memory *Memory) bool {
	if cpu.Accumulator.GetValue() == 0 {
		cpu.ProgramCounter.SetValue(cpu.AddressRegister.GetValue())
	}
	return true
}

func inst_brp (cpu *CPU, memory *Memory) bool {
	if cpu.Accumulator.GetValue() >= 0 {
		cpu.ProgramCounter.SetValue(cpu.AddressRegister.GetValue())
	}
	return true
}

func inst_inp_out (cpu *CPU, memory *Memory) bool {
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
