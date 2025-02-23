package translators

import (
	types "compilerTypes"
	"fmt"
	"log"
	"os"
)

type Interpreter struct{}

const memorySize = 30000 // 30,000 memory cells is the agreed standard for the Brainf*ck language

func (i Interpreter) Translate(code string) {
	var memory [memorySize]int
	var memoryPointer int = 0
	var codePointer int = 0

	var startLoopPointers types.Stack

	for codePointer < len(code) {
		switch code[codePointer] {
		case '>':
			memoryPointer++
		case '<':
			memoryPointer--
		case '+':
			memory[memoryPointer]++
		case '-':
			memory[memoryPointer]--
		case '.':
			fmt.Print(string(memory[memoryPointer]))
		case ',':
			fmt.Scan(&memory[memoryPointer])
		case '[':
			// If the memory location is already 0, skip to the end of the loop
			if memory[memoryPointer] == 0 {
				for codePointer < len(code) && code[codePointer] != ']' {
					codePointer++
				}
			}
			startLoopPointers.Push(memoryPointer, codePointer)
		case ']':
			originalMemoryPointer, original_codePointer, is_item := startLoopPointers.Peek()
			if !is_item {
				log.Fatal("No loop pointer found")
				os.Exit(-1)
			}

			// Check if the loop should continue (does initial memory location == 0)
			if memory[originalMemoryPointer] == 0 {
				startLoopPointers.Pop()
			} else {
				codePointer = original_codePointer
			}
		}
		codePointer++
	}
}
