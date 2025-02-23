package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"translators"
)

const helpMessage = `Usage: main.exe -m [interpret|compile|help] [-s filepath|-i console_input]`

func main() {
	// Parse command line arguments
	method, cliCode, filepath := parseCommands()

	var code = extractCodeToRun(cliCode, filepath)
	if len(code) == 0 {
		fmt.Println(helpMessage)
		fmt.Println("No code to run. Exiting...")
		return
	}

	if method == "interpret" {
		interpret(code)
	} else if method == "compile" {
		compile(code)
	} else {
		fmt.Println(helpMessage)
		return
	}
}

func parseCommands() (string, string, string) {
	var method string
	var cliCode string
	var filepath string

	flag.StringVar(&method, "m", "", "Method to use (interpret or compile)")
	flag.StringVar(&cliCode, "i", "", "Brainf*ck code to interpret")
	flag.StringVar(&filepath, "s", "", "Filepath to Brainf*ck code to interpret")
	flag.Parse()
	return method, cliCode, filepath
}

func interpret(code string) {
	fmt.Print("Interpreting code... \n\n")

	interpreter := translators.Interpreter{}
	interpreter.Translate(code)
}

func compile(code string) {
	fmt.Print("Compiling code... \n\n")
}

func extractCodeToRun(cliCode string, filepath string) string {
	var code string
	if cliCode != "" {
		code = cliCode
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
