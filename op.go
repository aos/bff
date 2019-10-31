package main

type Kind uint8

const (
	Invalid Kind = iota
	IncPtr
	DecPtr
	IncData
	DecData
	ReadStdin
	WriteStdin
	JumpIfDataZero
	JumpIfDataNonZero
)

type Instruction struct {
	Kind
	Argument int
}
