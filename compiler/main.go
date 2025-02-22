package main

import (
	types "compiler_types"
	"flag"
	"fmt"
	"log"
	"os"
)

const memory_size = 30000 // 30,000 memory cells is the agreed standard for the Brainf*ck language
const help_message = `Usage: main.exe -m [interpret|compile|help] [-s filepath|-i console_input]`

func main() {
	// Parse command line arguments
	var method string
	var cli_code string
	var filepath string

	flag.StringVar(&method, "m", "", "Method to use (interpret or compile)")
	flag.StringVar(&cli_code, "i", "", "Brainf*ck code to interpret")
	flag.StringVar(&filepath, "s", "", "Filepath to Brainf*ck code to interpret")
	flag.Parse()

	var code = extract_code_to_run(cli_code, filepath)
	if len(code) == 0 {
		fmt.Println(help_message)
		fmt.Println("No code to run. Exiting...")
		return
	}

	if method == "interpret" {
		interpret(code)
	} else if method == "compile" {
		compile(code)
	} else {
		fmt.Println(help_message)
		return
	}
}

func interpret(code string) {
	fmt.Println("Interpreting code... \n")
	run(code)
}

func compile(code string) {
	fmt.Println("Compiling code... \n")
}

func extract_code_to_run(cli_code string, filepath string) string {
	var code string
	if cli_code != "" {
		code = cli_code
	} else if filepath != "" {
		data, err := os.ReadFile(filepath)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		code = string(data)
	}
	return code
}

func run(code string) {
	var memory [memory_size]int
	var memory_pointer int = 0
	var code_pointer int = 0

	var start_loop_pointers types.Stack

	for code_pointer < len(code) {
		switch code[code_pointer] {
		case '>':
			memory_pointer++
		case '<':
			memory_pointer--
		case '+':
			memory[memory_pointer]++
		case '-':
			memory[memory_pointer]--
		case '.':
			fmt.Print(string(memory[memory_pointer]))
		case ',':
			fmt.Scan(&memory[memory_pointer])
		case '[':
			// If the memory location is already 0, skip to the end of the loop
			if memory[memory_pointer] == 0 {
				for code_pointer < len(code) && code[code_pointer] != ']' {
					code_pointer++
				}
			}
			start_loop_pointers.Push(memory_pointer, code_pointer)
		case ']':
			original_memory_pointer, original_code_pointer, is_item := start_loop_pointers.Peek()
			if !is_item {
				log.Fatal("No loop pointer found")
				os.Exit(-1)
			}

			// Check if the loop should continue (does initial memory location == 0)
			if memory[original_memory_pointer] == 0 {
				start_loop_pointers.Pop()
			} else {
				code_pointer = original_code_pointer
			}
		}
		code_pointer++
	}
}
