package new_asm

type InstructionType int

const (
	UN_EVALUATED InstructionType = iota
	LABELED
	EVALUATED
)

// Represents instructions
type Instruction struct {
	Type InstructionType
	Opcode Opcode

	// Un-evaluated
	AddressRef *AddressRef

	// Labeled
	Label *AddressRef

	// Evaluated
	Address int
	Value int
}

// Evaluates an /un-evaluated/ instruction.
// NOTE: The address reference should have /already/ been evaluated
func (i *Instruction) Evaluate() *Instruction {
	if i.Type == UN_EVALUATED || i.Type == LABELED {
		i.Type = EVALUATED
		i.Address = i.AddressRef.Address
	}
	return i
}

// Compiles an /evaluated/ instruction.
func (i *Instruction) Compile() int {
	if i.Opcode == OP_DAT {
		return i.Value
	}
	return int(i.Opcode) * 100 + i.Address
}

func CreateInstruction(opcode Opcode, address *AddressRef) *Instruction {
	return &Instruction{
		Type: UN_EVALUATED,
		Opcode: opcode,
		AddressRef: address,
		Value: address.Address,
	}
}
