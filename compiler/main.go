package main

import (
	"fmt"
	"os"
	"types"
)

func main() {
	var run_case types.RunType = types.Compile
	if len(os.Args) > 1 && os.Args[1] == "interpret" {
		run_case = types.Interpret
	} else if len(os.Args) > 1 && os.Args[1] == "help" {
		fmt.Println("Usage: main.exe [interpret|compile] [-s filename|-i console_input]")
		return
	}

	if run_case == types.Interpret {
		interpret()
	} else {
		compile()
	}
}

func interpret() {

}

func compile() {

}
