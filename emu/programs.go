package emu

// Some basic programs for testing purposes
// See https://peterhigginson.co.uk/LMC/

// Adds two numbers together.
var Add = [100]int{
	901, // INP
	399, // STA 99
	901, // INP
	199, // ADD 99
	902, // OUT
	000, // HLT
}

// Adds the first two numbers together,
// subtracts the first from the third.
var Add_Subt = [100]int{
	901, //       INP
	309, //       STA FIRST
	901, //       INP
	109, //       ADD FIRST
	902, //       OUT
	901, //       INP
	209, //       SUB FIRST
	902, //       OUT
	000, //       HLT
	000, // FIRST DAT
}
