package ast

import "fmt"

type Scanner struct {
	Text string
}

func (s *Scanner) Scan() []Token {
	tokens := make([]Token, 0)
	runes := []rune(s.Text)
	start := 0
	current := 0
	line := 1

	isAtEnd := func() bool {
		return current >= len(s.Text)
	}
	
	

	addToken := func(tokenType TokenType) {
		tokens = append(tokens, Token{tokenType, string(runes[start:current]), "", line})
	}

	match := func(expected rune) bool {
		if isAtEnd() {
			return false
		}
		if runes[current] != expected {
			return false
		}

		current++
		return true
	}

	advance := func() rune {
		current++
		return runes[current-1]
	}

	scanToken := func() {
		c := advance()
		
		switch c {
		case '(':
			addToken(LeftParenthesis)
			break
		case ')':
			addToken(RightParenthesis)
			break
		case '{':
			addToken(LeftBrace)
			break
		case '}':
			addToken(RightBrace)
			break
		case ',':
			addToken(Comma)
			break
		case '.':
			addToken(Dot)
			break
		case '-':
			addToken(Minus)
			break
		case '+':
			addToken(Plus)
			break
		case ';':
			addToken(Semicolon)
			break
		case '*':
			addToken(Star)
			break
		case '!':
			if(match('=')){
				addToken(NotEqual)
			} else {
				addToken(Not)
			}
			break
		case '=':
			if(match('=')){
				addToken(EqualEqual)
			} else {
				addToken(Equal)
			}
			break
		case '<':
			if(match('=')){
				addToken(LessEqual)
			} else {
				addToken(Less)
			}
			break
		case '>':
			if(match('=')){
				addToken(GreaterEqual)
			} else {
				addToken(Greater)
			}
			break
		default:
			fmt.Errorf("unknown character '%v' at line %d", string(runes), line)
			break
		}
	}

	for isAtEnd() {
		start = current
		scanToken()
	}


	tokens = append(tokens, Token{Eof, "", "", line}) // Eof - End of File
	return tokens

}
