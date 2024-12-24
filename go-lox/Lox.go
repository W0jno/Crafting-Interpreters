package main

import (
	"fmt"
	"os"

	"github.com/w0jno/Crafting-Interpreters/go-lox/ast/ast"
)

func main() {
	
	if len(os.Args ) > 2{
		fmt.Println("Usage: Lox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
		
	} else {
		runPrompt()
	}
}

func runFile(path string){
	bytes, error := os.ReadFile(path)
	if error == nil {
		run(string(bytes))
	} else {
		report(0, "File", "Error reading file")
	}
}

func runPrompt(){
	for {
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)
		run(input)
		
	}
}

func run(source string){
	scanner := ast.Scanner{source}
	tokens := scanner.Scan()

	for _, token := range tokens {
		fmt.Println(token)
	}
}



func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
}