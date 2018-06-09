package main

import (
	"github.com/jamiemansfield/lmcemu/emu"
)

func main() {
	cpu := emu.CreateLmcCpu()
	memory := emu.CreateMemory(emu.Add_Subt)
	cpu.Execute(memory)
}
