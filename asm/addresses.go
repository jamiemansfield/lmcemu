package asm

// The type of an address reference, used for checking whether additional
// processing needs to happen later on.
type AddressType int

const (
	ADDR_NORMAL AddressType = iota
	ADDR_LABEL_POSITION
	ADDR_LABEL_DATA
)

// Establishes whether the address type is a label type, position or data.
func (t AddressType) IsLabel() bool {
	return t == ADDR_LABEL_DATA || t == ADDR_LABEL_POSITION
}

// An address reference.
type AddressRef struct {
	Type AddressType `json:"type"`
	Address int `json:"address"`
	Value int `json:"value"`
}

// Applies the label to an instruction to create a
// labeled instruction.
func (r *AddressRef) Apply(i *Instruction) *Instruction {
	i.Type = INST_LABELED
	i.Label = r
	return i
}

// Creates an address reference, given a concrete
// address.
func CreateAddressRef(address int) *AddressRef {
	return &AddressRef{
		Type: ADDR_NORMAL,
		Address: address,
	}
}

// Creates an address reference for the use of
// position labels.
func CreatePositionLabel() *AddressRef {
	return &AddressRef{
		Type: ADDR_LABEL_POSITION,
		Address: -1,
	}
}

// Creates an address reference for the use of
// data labels.
func CreateDataLabel(value int) *AddressRef {
	return &AddressRef{
		Type: ADDR_LABEL_DATA,
		Address: -1,
		Value: value,
	}
}

// Establishes whether the given address is valid, noting that LMC address are
// only 3 digits.
func IsValidAddress(address int) bool {
	return address >= 0 && address <= 999
}
