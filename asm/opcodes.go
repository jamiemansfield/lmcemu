package asm

// Represents the operation code of an instruction, specifying what
// operation to perform.
type Opcode int

const (
	// This is not actually an operation code, but rather a psuedo
	// Opcode for DAT instructions.
	OP_DAT     Opcode = -1

	// The HLT operation code will stop the program when run.
	OP_HLT     Opcode = 0

	// The ADD operation code will add the given argument to the
	// current value on the accumulator, the result will be stored
	// on the accumulator.
	OP_ADD     Opcode = 1

	// The SUB operation code will subtract the given argument from
	// the current value on the accumulator, the result will be stored
	// on the accumulator.
	OP_SUB     Opcode = 2

	// The STA operation code will store the value on the accumulator
	// to the address provided.
	OP_STA     Opcode = 3

	// The LDA operation code will load the value from the provided
	// address into the accumulator.
	OP_LDA     Opcode = 5

	// The BRA operation code will branch to the provided address.
	OP_BRA     Opcode = 6

	// The BRZ operation code will branch to the provided address,
	// given that the value on the accumulator is zero.
	OP_BRZ     Opcode = 7

	// The BRP operation code will branch to the provided address,
	// given that the value on the accumulator is positive.
	OP_BRP     Opcode = 8

	// The INP_OUT operation code performs two operations depending
	// on the provided address.
	// If the address is 1, the program will ask for input to be stored
	//   in the accumulator.
	// If the address is 2, the program will output the value stored in
	//   the accumulator.
	OP_INP_OUT Opcode = 9
)
