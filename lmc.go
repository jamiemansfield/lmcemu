package main

import (
	"github.com/jamiemansfield/lmcemu/emu"
	"github.com/jamiemansfield/lmcemu/asm"
)

func main() {
	cpu := emu.CreateLmcCpu()
	memory := emu.CreateMemory(asm.AssembleProgram(asm.AddProgram))
	cpu.Execute(memory)
}
