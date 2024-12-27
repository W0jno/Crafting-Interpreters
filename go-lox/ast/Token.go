package ast

import "fmt"

type TokenType int

const (
	And TokenType = iota
	Class
	Comma
	Dot
	Else
	Eof
	Equal
	EqualEqual
	False
	For
	Fun
	Greater
	GreaterEqual
	Identifier
	If
	LeftParenthesis
	LeftBrace
	Less
	LessEqual
	Minus
	Nil
	Not
	NotEqual
	Number
	Or
	Plus
	Print
	Return
	RightParenthesis
	RightBrace
	Semicolon
	Slash
	Star
	String
	Super
	This
	True
	Var
	While
	
)

type Token struct {
	TokenType
	Lexeme  string
	Literal string
	Line    int
}

func (t Token) String() string {
	return fmt.Sprintf("Token type: %d Lexeme: %s Literal: %s", t.TokenType, t.Lexeme, t.Literal)
}