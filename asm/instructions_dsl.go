package asm

// Creates a DAT instruction.
func DAT(address *AddressRef) *Instruction {
	return CreateInstruction(OP_DAT, address)
}

// Creates a HLT instruction.
func HLT() *Instruction {
	return CreateInstruction(OP_HLT, CreateAddressRef(0))
}

// Creates a ADD instruction.
func ADD(address *AddressRef) *Instruction {
	return CreateInstruction(OP_ADD, address)
}

// Creates a SUB instruction.
func SUB(address *AddressRef) *Instruction {
	return CreateInstruction(OP_SUB, address)
}

// Creates a STA instruction.
func STA(address *AddressRef) *Instruction {
	return CreateInstruction(OP_STA, address)
}

// Creates a LDA instruction.
func LDA(address *AddressRef) *Instruction {
	return CreateInstruction(OP_LDA, address)
}

// Creates a BRA instruction.
func BRA(address *AddressRef) *Instruction {
	return CreateInstruction(OP_BRA, address)
}

// Creates a BRZ instruction.
func BRZ(address *AddressRef) *Instruction {
	return CreateInstruction(OP_BRZ, address)
}

// Creates a BRP instruction.
func BRP(address *AddressRef) *Instruction {
	return CreateInstruction(OP_BRP, address)
}

// Creates a INP instruction.
func INP() *Instruction {
	return CreateInstruction(OP_INP_OUT, CreateAddressRef(1))
}

// Creates a OUT instruction.
func OUT() *Instruction {
	return CreateInstruction(OP_INP_OUT, CreateAddressRef(2))
}

