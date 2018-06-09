package main

import (
	"github.com/jamiemansfield/lmcemu/asm"
	"github.com/jamiemansfield/lmcemu/emu"
)

func main() {
	cpu := emu.CreateLmcCpu()
	memory := emu.CreateMemory(asm.AssembleProgram(asm.AddSubtProgram))
	cpu.Execute(memory)
}
