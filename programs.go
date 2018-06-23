package main

import (
	"github.com/jamiemansfield/lmcemu/asm"
)

var start    = asm.CreatePositionLabel()
var addition = asm.CreatePositionLabel()
var subtract = asm.CreatePositionLabel()
var end      = asm.CreatePositionLabel()

var mode     = asm.CreateDataLabel(1)
var value_o  = asm.CreateDataLabel(1)
var value_t  = asm.CreateDataLabel(2)

var Calculator = []*asm.Instruction {
	// START:
	start.Apply(asm.INP()),
	asm.STA(mode),

	asm.INP(),
	asm.STA(value_o),
	asm.INP(),
	asm.STA(value_t),

	asm.LDA(mode),
	asm.BRZ(addition),
	asm.BRA(subtract),

	// ADDITION:
	addition.Apply(asm.LDA(value_o)),
	asm.ADD(value_t),
	asm.BRA(end),

	// SUBTRACT:
	subtract.Apply(asm.LDA(value_o)),
	asm.SUB(value_t),
	asm.BRA(end),

	// END:
	end.Apply(asm.OUT()),
	asm.HLT(),

	// Data
	asm.DAT(mode),
	asm.DAT(value_o),
	asm.DAT(value_t),
}

var BuiltinRegistry = map[string][]*asm.Instruction{
	"calculator": Calculator,
}
