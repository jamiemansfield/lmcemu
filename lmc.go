package main

import (
	"fmt"
	"github.com/jamiemansfield/lmcemu/asm"
	"log"
	"github.com/jamiemansfield/lmcemu/emu"
	"os"
)

func main() {
	// Read asm from file
	file, err := os.Open("calculator.asm")
	defer file.Close()

	// Assemble the program.
	prog, err := asm.AssembleProgram(asm.TokeniseFile(file).Assemble())
	if err != nil {
		log.Fatal(err)
		return
	}

	debug := false

	// Debug assembly
	if debug {
		fmt.Printf("%v", prog)
	}

	// Run
	if !debug {
		cpu := emu.CreateLmcCpu()
		memory := emu.CreateMemory(prog)
		if cpu.Execute(memory) != nil {
			log.Fatal(err)
			return
		}
	}
}
