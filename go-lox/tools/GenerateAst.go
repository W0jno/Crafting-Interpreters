package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: generate_ast <output directory>")
        os.Exit(64)
    }
    outputDir := os.Args[1]
    defineAst(outputDir, "Expr", []string{
        "Binary   : Expr left, Token operator, Expr right",
        "Grouping : Expr expression",
        "Literal  : Object value",
        "Unary    : Token operator, Expr right",
    })
}

func defineAst(outputDir, base string, types []string) {
    path := fmt.Sprintf("%s/%s.go", outputDir, strings.ToLower(base))
    var src string

    src += fmt.Sprintln("package ast")
    src += fmt.Sprintln("")
    src += defineVisitor(base, types)

    for _, t := range types {
        className := strings.TrimRight(strings.Split(t, ":")[0], " ")
        fields := strings.TrimRight(strings.Split(t, ":")[1], " ")
        src += defineType(base, className, fields)
    }

    if err := saveFile(path, src); err != nil {
        panic(err)
    }
}

func defineVisitor(base string, types []string) string {
    var src string
    src += fmt.Sprintf("type %sVisitor interface {\n", base)
    for _, t := range types {
        typeName := strings.TrimRight(strings.Split(t, ":")[0], " ")
        src += fmt.Sprintf("    Visit%s%s(%s *%s) interface{}\n", typeName, base, strings.ToLower(base), typeName)
    }
    src += "}\n"
    return src
}

func defineType(base, className, fieldList string) string {
    var src string
    src += fmt.Sprintf("type %s struct {\n", className)
    fields := strings.Split(fieldList, ", ")
    for _, field := range fields {
        fieldParts := strings.Split(field, " ")
        fieldName := fieldParts[1]
        fieldType := fieldParts[0]
        src += fmt.Sprintf("    %s %s\n", fieldName, fieldType)
    }
    src += "}\n"
    src += fmt.Sprintf("func (e *%s) Accept(visitor %sVisitor) interface{} {\n", className, base)
    src += fmt.Sprintf("    return visitor.Visit%s%s(e)\n", className, base)
    src += "}\n"
    return src
}

func saveFile(path, src string) error {
    return os.WriteFile(path, []byte(src), 0644)
}