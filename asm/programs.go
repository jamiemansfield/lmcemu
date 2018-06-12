package asm

var start    = CreatePositionLabel()
var addition = CreatePositionLabel()
var subtract = CreatePositionLabel()
var end      = CreatePositionLabel()

var mode     = CreateDataLabel(1)
var value_o  = CreateDataLabel(1)
var value_t  = CreateDataLabel(2)

var Calculator = []*Instruction {
	// START:
	start.Apply(INP()),
	STA(mode),

	INP(),
	STA(value_o),
	INP(),
	STA(value_t),

	LDA(mode),
	BRZ(addition),
	BRA(subtract),

	// ADDITION:
	addition.Apply(LDA(value_o)),
	ADD(value_t),
	BRA(end),

	// SUBTRACT:
	subtract.Apply(LDA(value_o)),
	SUB(value_t),
	BRA(end),

	// END:
	end.Apply(OUT()),
	HLT(),

	// Data
	DAT(mode),
	DAT(value_o),
	DAT(value_t),
}
