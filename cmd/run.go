package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Morzatt/glox/scanner"
)

type GLox struct {
	HadError bool 
}

// Read reads the standard input and runs the interpreter based on the prompt command line or based on a source file.
func (g *GLox) Read() {
	args := os.Args
	if len(args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(args) == 2 {
		g.RunFile(args[1])
	} else {
		g.RunPrompt()
	}
}

// RunFile reads the path to a file provided and executes it.
func (g *GLox) RunFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(64)	
	}
	g.Run(string(content))

	if g.HadError { os.Exit(65) }
}

// RunPrompt runs the interpreter interactively on the prompt command line interface. 
func (g *GLox) RunPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Print("> ")
		if scanner.Bytes() == nil { break }; 
		g.Run(scanner.Text())
		g.HadError = false
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		os.Exit(64)
	}
}

// Run runs the Scanner and Lexer on the provided source code.
func (g *GLox) Run(source string) {
	scanner := scanner.Scanner{
		Source: source,
		Tokens: make([]scanner.Token, 1),
		Start : 0, 
		Current: 0,
		Line: 1,
	}
	tokens := scanner.ScanTokens()

	for _, t := range tokens {
		fmt.Printf("Tokens: %+v\n", t)
	}
}