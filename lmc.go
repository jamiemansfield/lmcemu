package main

import (
	"github.com/jamiemansfield/lmcemu/new-asm"
	"fmt"
)

func main() {
	//cpu := emu.CreateLmcCpu()
	//memory := emu.CreateMemory(asm.AssembleProgram(asm.AddSubtProgram))
	//cpu.Execute(memory)

	fmt.Printf("%v", new_asm.AssembleProgram(new_asm.Calculator))
}
