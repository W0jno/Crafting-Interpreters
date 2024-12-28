package ast

import "fmt"

type AstPrinter struct {
	indents int
}

func NewAstPrinter() *AstPrinter {
	return &AstPrinter{indents: 0}
}

func getIndents(num int) string {
	indents := ""

	for i := 0; i < num; i++ {
		indents += "	"
	}
	return indents
}

func (p *AstPrinter) Print(stmts []Stmt) {
	ast := ""

	for _, stmt := range stmts {
		val, _ := stmt.Accept(p).(string)
		ast += val
	}

	fmt.Println(ast)
}

func (p *AstPrinter) VisitUnaryExpr(expr *Unary) interface{} {
	return p.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (p *AstPrinter) parenthesize(name string, values ...interface{}) string {
	ast := "(" + name

	for _, obj := range values {
		ast += " "
		switch v := obj.(type) {
		case Expr:
			val, _ := v.Accept(p).(string)
			ast += val
		case Stmt:
			val, _ := v.Accept(p).(string)
			ast += val
		case *Token:
			ast += v.Lexeme
		default:
			val, _ := obj.(string)
			ast += val
		}
	}
	ast += ")"
	return ast
}