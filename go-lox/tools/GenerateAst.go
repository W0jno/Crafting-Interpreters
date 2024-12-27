package main

import (
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"strings"
)

func main () {
 if len(os.Args ) != 2{
	fmt.Println("Usage: generate_ast <output directory>")
	os.Exit(1)
 }
 	outputDir, err := filepath.Abs(os.Args[1])
 	if err != nil {
		panic(err)
	}
	defineAst(outputDir, "Expr", []string{
		"Binary   : Expr left, Token operator, Expr right",
    	"Grouping : Expr expression",
    	"Literal  : Object value",
    	"Unary    : Token operator, Expr right",
	})

	
}

func defineAst(output string, base string, types []string){
	path := output + "/" + base + ".go"

	var src string

	src += fmt.Sprintln("")
	src += fmt.Sprintln("package lox")
	src += fmt.Sprintln("")

	for _, t := range types {
		className := strings.TrimRight(strings.Split(t, ":")[0], "\t")
		fields := strings.TrimRight(strings.Split(t, ":")[1], " ")
		src += defineType(base, className, fields)

	}
	if err := saveFile(path, src); err != nil {
		panic(err)
	}
}

func defineType(base, className, fields string) string{
	var src string

	src += fmt.Sprintln("")
	src += fmt.Sprintf("type %s struct {\n", className)
	src += fmt.Sprintln("")

	flds := strings.Split(fields, ", ")

	for _, f := range flds {
		src += fmt.Sprintln(f)
	}
	src += fmt.Sprintln("}")
	return src
}

func saveFile(path, src string) error {
	buf, err := format.Source([]byte(src))
	if err != nil {
		return err
	}

	os.WriteFile(path, buf, 0644)
	return nil
}



