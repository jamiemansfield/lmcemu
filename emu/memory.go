package emu

// A representation of RAM within an emulated system.
type Memory struct {
	Values [100]int
}

// Gets the integer value of a given memory address.
func (m *Memory) GetValue(address int) int {
	return m.Values[address]
}

// Sets the integer value of a given memory address.
func (m *Memory) SetValue(address int, value int) {
	m.Values[address] = value
}
