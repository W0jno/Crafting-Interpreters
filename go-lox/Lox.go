package main

import (
	"fmt"
	"os"
)

func main() {
	hadError := false
	if len(os.Args ) > 2{
		fmt.Println("Usage: Lox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
		if hadError {
			os.Exit(65)
		}
	} else {
		runPrompt()
	}
}

func runFile(path string){
	bytes, error := os.ReadFile(path)
	if error == nil {
		run(string(bytes))
	}
}

func runPrompt(){
	for {
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)
		run(input)
		hadError = false
	}
}

func run(source string){
	scanner := ast.Scanner(source)
	tokens := scanner.scanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func error(int line, string message){
	report(line, "", message)
}

func report(int line, string where, string message){
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
}