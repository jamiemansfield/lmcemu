package asm

type Opcode int

const (
	OP_DAT     Opcode = -1
	OP_HLT     Opcode = 0
	OP_ADD     Opcode = 1
	OP_SUB     Opcode = 2
	OP_STA     Opcode = 3
	OP_LDA     Opcode = 5
	OP_BRA     Opcode = 6
	OP_BRZ     Opcode = 7
	OP_BRP     Opcode = 8
	OP_INP_OUT Opcode = 9
)
