package main

import (
	"fmt"
	"github.com/jamiemansfield/lmcemu/asm"
	"log"
	"github.com/jamiemansfield/lmcemu/emu"
)

func main() {
	// Assemble the program.
	prog, err := asm.AssembleProgram(asm.Calculator)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Debug assembly
	if false {
		fmt.Printf("%v", prog)
	}

	// Run
	if true {
		cpu := emu.CreateLmcCpu()
		memory := emu.CreateMemory(prog)
		if cpu.Execute(memory) != nil {
			log.Fatal(err)
			return
		}
	}
}
