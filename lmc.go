package main

import (
	"fmt"
	"github.com/jamiemansfield/lmcemu/new-asm"
	"log"
	"github.com/jamiemansfield/lmcemu/emu"
)

func main() {
	// Assemble the program.
	prog, err := new_asm.AssembleProgram(new_asm.Calculator)
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
