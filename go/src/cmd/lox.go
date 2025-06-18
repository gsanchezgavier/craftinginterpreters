package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gsanchezgavier/craftinginterpreters/src/scanner"
)

func main() {
	args := os.Args[1:]

	switch {
	case len(args) > 1:
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	case len(args) == 1:
		runFile(args[0])
	default:
		runPrompt()
	}
}

func runPrompt() {
	for {
		fmt.Print("> ")
		var input string

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading input:", err)
			continue
		}
		run(input)
	}
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("reading file: %s", path)
		os.Exit(1)
	}
	run(string(bytes))
}

func run(source string) {
	s := scanner.NewScanner(source)

	tokens := s.ScanTokens()

	for _, t := range tokens {
		fmt.Printf("%v\n", t)
	}
}
