package asm

// Some basic programs for testing purposes
// See https://peterhigginson.co.uk/LMC/

// Adds two numbers together.
var add_first = CreateLabelRef()
var AddProgram = []Line{
	INP(),
	STA(add_first),
	INP(),
	ADD(add_first),
	OUT(),
	HLT(),
}

// Adds the first two numbers together,
// subtracts the first from the third.
var add_subt_first = CreateLabelRef()
var AddSubtProgram = []Line{
	INP(),
	STA(add_subt_first),
	INP(),
	ADD(add_subt_first),
	OUT(),
	INP(),
	SUB(add_subt_first),
	OUT(),
	HLT(),
}
