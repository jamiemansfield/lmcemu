package asm

import (
	"strconv"
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
