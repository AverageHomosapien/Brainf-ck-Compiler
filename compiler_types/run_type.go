package compiler_types

type RunType int

const (
	Compile RunType = iota
	Interpret
)
