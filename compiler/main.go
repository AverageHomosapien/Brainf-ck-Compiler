package main

import (
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
	fmt.Println("Interpreting code... " + code)
	run(code)
}

func compile(code string) {
	fmt.Println("Compiling code... " + code)
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
			if memory[memory_pointer] == 0 {
				for code[code_pointer] != ']' {
					code_pointer++
				}
			}
		case ']':
			if memory[memory_pointer] != 0 {
				for code[code_pointer] != '[' {
					code_pointer--
				}
			}
		}
		code_pointer++
	}
}
