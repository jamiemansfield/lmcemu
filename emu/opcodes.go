package emu

type Opcode int

const (
	HLT     Opcode = 0
	ADD     Opcode = 1
	SUB     Opcode = 2
	STA     Opcode = 3
	LDA     Opcode = 5
	BRA     Opcode = 6
	BRZ     Opcode = 7
	BRP     Opcode = 8
	INP_OUT Opcode = 9
)
