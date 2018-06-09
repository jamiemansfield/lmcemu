package asm

// Some basic programs for testing purposes
// See https://peterhigginson.co.uk/LMC/

// Adds two numbers together.
var AddProgram = [100]Line{
	CreateLine(INP_OUT, 1),
	CreateLine(STA, 99),
	CreateLine(INP_OUT, 1),
	CreateLine(ADD, 99),
	CreateLine(INP_OUT, 2),
	CreateLine(HLT, 0),
}

// Adds the first two numbers together,
// subtracts the first from the third.
var AddSubtProgram = [100]Line{
	CreateLine(INP_OUT, 1),
	CreateLine(STA, 99),
	CreateLine(INP_OUT, 1),
	CreateLine(ADD, 99),
	CreateLine(INP_OUT, 2),
	CreateLine(INP_OUT, 1),
	CreateLine(SUB, 99),
	CreateLine(INP_OUT, 2),
	CreateLine(HLT, 0),
}
