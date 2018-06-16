package asm

import (
	"os"
	"bufio"
	"strconv"
	"errors"
	"strings"
)

type TokenType int

const (
	// Instructions
	TKN_HLT TokenType = iota
	TKN_ADD
	TKN_SUB
	TKN_STA
	TKN_LDA
	TKN_BRA
	TKN_BRZ
	TKN_BRP
	TKN_INP
	TKN_OUT
	TKN_DAT

	// Other
	TKN_VALUE
	TKN_LABEL
)

func (t TokenType) NeedsMore() bool {
	if t == TKN_HLT ||
		t == TKN_INP ||
		t == TKN_OUT ||
		t == TKN_VALUE {
		return false
	}
	return true
}

type AddressRegistry map[string]*AddressRef

func (r AddressRegistry) GetMapping(labelName string) *AddressRef {
	if r[labelName] == nil {
		panic("bad")
	}
	return r[labelName]
}

type Token struct {
	Type TokenType
	Name string
}

func CreateToken(ttype TokenType, name string) *Token {
	return &Token{
		Type: ttype,
		Name: name,
	}
}

func GetTokenType(str string) TokenType {
	switch str {
	case "HLT":
		return TKN_HLT
	case "ADD":
		return TKN_ADD
	case "SUB":
		return TKN_SUB
	case "STA":
		return TKN_STA
	case "LDA":
		return TKN_LDA
	case "BRA":
		return TKN_BRA
	case "BRZ":
		return TKN_BRZ
	case "BRP":
		return TKN_BRP
	case "INP":
		return TKN_INP
	case "OUT":
		return TKN_OUT
	case "DAT":
		return TKN_DAT
	default:
		_, err := strconv.Atoi(str)
		if err != nil {
			return TKN_LABEL
		}
		return TKN_VALUE
	}
}

func GetToken(str string) *Token {
	return CreateToken(GetTokenType(str), str)
}

type TokenisedFile [][]*Token

func (f TokenisedFile) Assemble() []*Instruction {
	// First pass (to make Address references)
	var refs = map[string]*AddressRef{}

	for _, line := range f {
		// Positions
		if line[0].Type == TKN_LABEL && line[1].Type != TKN_DAT && refs[line[0].Name] == nil {
			refs[line[0].Name] = CreatePositionLabel()
		}
		// Data
		if line[0].Type == TKN_LABEL && line[1].Type == TKN_DAT && refs[line[0].Name] == nil {
			val, err := strconv.Atoi(line[2].Name)
			if err != nil {
				panic(err) // TODO:
			}
			refs[line[0].Name] = CreateDataLabel(val)
		}
	}

	// Create instructions
	var instructions []*Instruction
	for _, line := range f {
		var instruction *Instruction

		if len(line) == 1 {
			instruction = getInstruction1(line[0])
		} else
		if len(line) == 2 {
			instruction = getInstruction2(line[0], line[1], refs)
		} else
		if len(line) == 3 {
			instruction = getInstruction3(line[0], line[1], line[2], refs)
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

func getInstruction1(token1 *Token) *Instruction {
	// HLT, INP, OUT
	switch token1.Type {
	case TKN_HLT:
		return HLT()
	case TKN_INP:
		return INP()
	case TKN_OUT:
		return OUT()
	}
	panic(errors.New("bad"))
}

func getInstruction2(token1 *Token, token2 *Token, refs AddressRegistry) *Instruction {
	// Position
	if token1.Type == TKN_LABEL {
		return refs.GetMapping(token1.Name).Apply(getInstruction1(token2))
	}

	// Requires address
	var ref *AddressRef
	if token2.Type == TKN_LABEL {
		ref = refs.GetMapping(token2.Name)
	} else {
		val, err := strconv.Atoi(token2.Name)
		if err != nil {
			panic(err) // TODO:
		}
		ref = CreateAddressRef(val)
	}

	switch token1.Type {
	case TKN_ADD:
		return ADD(ref)
	case TKN_SUB:
		return SUB(ref)
	case TKN_STA:
		return STA(ref)
	case TKN_LDA:
		return LDA(ref)
	case TKN_BRA:
		return BRA(ref)
	case TKN_BRZ:
		return BRZ(ref)
	case TKN_BRP:
		return BRP(ref)
	case TKN_DAT:
		return DAT(ref)
	}
	panic(errors.New("bad"))
}


func getInstruction3(token1 *Token, token2 *Token, token3 *Token, refs AddressRegistry) *Instruction {
	// Handle Data
	if token1.Type == TKN_LABEL && token2.Type == TKN_DAT {
		return DAT(refs.GetMapping(token1.Name))
	}

	return refs.GetMapping(token1.Name).Apply(getInstruction2(token2, token3, refs))
}

func TokeniseFile(file *os.File) TokenisedFile {
	var lines TokenisedFile

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var str = scanner.Text()

		if len(str) > 0 {
			var line []*Token

			for _, v := range strings.Fields(str) {
				line = append(line, GetToken(v))
			}

			lines = append(lines, line)
		}
	}

	return lines
}
