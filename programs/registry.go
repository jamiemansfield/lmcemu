package programs

import "github.com/jamiemansfield/lmcemu/asm"

// A registry of all the builtin programs.
// NOTE: All names should be in lower case, as all queries should be.
var BuiltinRegistry = map[string][]*asm.Instruction{
	"calculator": Calculator,
}
