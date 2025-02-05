package main

import (
	"flag"
	"fmt"
)

const memory_size = 30000 // 30,000 memory cells is the agreed standard for the Brainf*ck language

func main() {
	var method = *flag.String("m", "", "Method to use (interpret or compile)")
	var code = *flag.String("i", "", "Brainf*ck code to interpret")
	var filepath = *flag.String("s", "", "Filepath to Brainf*ck code to interpret")
	flag.Parse()

	if method == "interpret" {
		interpret(code)
	} else if method == "compile" {
		compile(filepath)
	} else {
		fmt.Println("Usage: main.exe -m [interpret|compile|help] [-s filepath|-i console_input]")
		return
	}
}

func interpret(code string) {
	fmt.Println(code)

	for {
		if len(code) != 0 {
			break
		}
		fmt.Println("Enter your Brainf*ck code: ")
		fmt.Scan(&code)
	}
}

func compile(code string) {

}
