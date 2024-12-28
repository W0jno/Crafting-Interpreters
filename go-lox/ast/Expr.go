package ast

type ExprVisitor interface {
    VisitBinaryExpr(expr *Binary) interface{}
    VisitGroupingExpr(expr *Grouping) interface{}
    VisitLiteralExpr(expr *Literal) interface{}
    VisitUnaryExpr(expr *Unary) interface{}
}
type Binary struct {
    left Expr 
    operator *Token
    right Expr
}

type Expr interface {
	Accept(v ExprVisitor) interface{}
}

func (e *Binary) Accept(visitor ExprVisitor) interface{} {
    return visitor.VisitBinaryExpr(e)
}
type Grouping struct {
    Expression Expr 
}
func (e *Grouping) Accept(visitor ExprVisitor) interface{} {
    return visitor.VisitGroupingExpr(e)
}
type Literal struct {
    Value interface{}
}
func (e *Literal) Accept(visitor ExprVisitor) interface{} {
    return visitor.VisitLiteralExpr(e)
}
type Unary struct {
    Token 
    right Expr
}
func (e *Unary) Accept(visitor ExprVisitor) interface{} {
    return visitor.VisitUnaryExpr(e)
}
