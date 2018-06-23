package asm

import (
	"strconv"
	"github.com/oleiade/lane"
	"os"
	"bufio"
	"strings"
	"errors"
)

type Parser struct {
	tokenQueue *lane.Queue
	addressRegistry AddressRegistry
}

func CreateParser() *Parser {
	return &Parser{
		tokenQueue: lane.NewQueue(),
		addressRegistry: map[string]*AddressRef{},
	}
}

func (p *Parser) ReadFromFile(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var str = scanner.Text()

		if len(str) > 0 {
			for _, v := range strings.Fields(str) {
				p.tokenQueue.Enqueue(GetToken(v))
			}
		}
	}
}

func (p *Parser) Assemble() ([]*Instruction, error) {
	// Create instructions
	var instructions []*Instruction
	for !p.tokenQueue.Empty() {
		inst, err := p.parseInstruction()
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, inst)
	}

	return instructions, nil
}

func (p *Parser) parseInstruction() (*Instruction, error) {
	token := p.tokenQueue.Pop().(*Token)

	// Labeled instructions
	if token.Type == TKN_LABEL {
		nextToken := p.tokenQueue.Head().(*Token)
		label := p.addressRegistry.GetMapping(nextToken.Name)

		// Data Label
		if nextToken.Type == TKN_DAT {
			p.tokenQueue.Pop()
			nextNextToken := p.tokenQueue.Pop().(*Token)
			label.Type = ADDR_LABEL_DATA

			if nextNextToken.Type == TKN_VALUE {
				// Parse the integer value - we can ignore the error as it should NEVER
				// fail as it needs to pass to be identified as TKN_VALUE.
				val, _ := strconv.Atoi(nextNextToken.Name)
				label.Value = val

				// Create the instruction
				return DAT(label), nil
			}
			return nil, errors.New("Invalid input supplied!")
		}

		// Position Label
		label.Type = ADDR_LABEL_POSITION
		inst, err := p.parseInstruction()
		if err != nil {
			return nil, err
		}
		return label.Apply(inst), nil
	}

	// Instructions that require no arguments
	switch token.Type {
	case TKN_HLT:
		return HLT(), nil
	case TKN_INP:
		return INP(), nil
	case TKN_OUT:
		return OUT(), nil
	}

	// Instructions that require arguments
	nextToken := p.tokenQueue.Pop().(*Token)

	// Parse the argument
	var ref *AddressRef
	if nextToken.Type == TKN_LABEL {
		ref = p.addressRegistry.GetMapping(nextToken.Name)
	} else {
		val, err := strconv.Atoi(nextToken.Name)
		if err != nil {
			return nil, err
		}
		ref = CreateAddressRef(val)
	}

	// Create the instruction
	switch token.Type {
	case TKN_ADD:
		return ADD(ref), nil
	case TKN_SUB:
		return SUB(ref), nil
	case TKN_STA:
		return STA(ref), nil
	case TKN_LDA:
		return LDA(ref), nil
	case TKN_BRA:
		return BRA(ref), nil
	case TKN_BRZ:
		return BRZ(ref), nil
	case TKN_BRP:
		return BRP(ref), nil
	case TKN_DAT:
		return DAT(ref), nil
	}

	// If we're still here error out
	return nil, errors.New("Invalid input supplied!")
}
