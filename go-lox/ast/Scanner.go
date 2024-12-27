package ast

import "fmt"

type Scanner struct {
	Text string
}

 var keywords = map[string]TokenType{
	"and": And,
	"class": Class,
	"else": Else,
	"false": False,
	"for": For,
	"fun": Fun,
	"if": If,
	"nil": Nil,
	"or": Or,
	"print": Print,
	"return": Return,
	"super": Super,
	"this": This,
	"true": True,
	"var": Var,
	"while": While,
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

	peek := func() rune {
		if isAtEnd() {
			return '\000'
		}
		return runes[current]
	}

	isDigit := func(c rune) bool {
		return c >= '0' && c <= '9'
	}

	peekNext := func() rune {
		if current+1 >= len(runes) {
			return '\000'
		}
		return runes[current+1]
	}

	isLetter := func(c rune) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
	}

	

	number := func(){
		for isDigit(peek()){
			advance()
		}
		if peek() == '.' && isDigit(peekNext()){
			advance()
			for isDigit(peek()){
				advance()
			}
		}
		number := string(runes[start:current])
		tokens = append(tokens, Token{Number, number, number, line})
	}

	scanToken := func() error{
		c := advance()
		
		switch c {
		case '(':
			addToken(LeftParenthesis)
			
		case ')':
			addToken(RightParenthesis)
			
		case '{':
			addToken(LeftBrace)
			
		case '}':
			addToken(RightBrace)
			
		case ',':
			addToken(Comma)
			
		case '.':
			addToken(Dot)
			
		case '-':
			addToken(Minus)
			
		case '+':
			addToken(Plus)
			
		case ';':
			addToken(Semicolon)
			
		case '*':
			addToken(Star)
			
		case '!':
			if(match('=')){
				addToken(NotEqual)
			} else {
				addToken(Not)
			}
			
		case '=':
			if(match('=')){
				addToken(EqualEqual)
			} else {
				addToken(Equal)
			}
			
		case '<':
			if(match('=')){
				addToken(LessEqual)
			} else {
				addToken(Less)
			}
			
		case '>':
			if(match('=')){
				addToken(GreaterEqual)
			} else {
				addToken(Greater)
			}
			
		case '/':
			if(match('/')){
				for peek() != '\n' && !isAtEnd() {
					advance()
				}
			} else if(match('*')){
				for peek() != '*' && peekNext() != '/' && !isAtEnd() {
					if peek() == '\n' {
						line++
					}
					advance()
				}
				
			} else {
				addToken(Slash)
			}
		case ' ':
		case '\r':
		case '\t':
			// Ignore whitespace.
		case '\n':
			line++
		case '"':
			for peek() != '"' && !isAtEnd() {
				if peek() == '\n' {
					line++
				}
				advance()
			}
			if isAtEnd(){
				return fmt.Errorf("unterminated string at line %d", line)
			}
			advance()
			lexeme := string(runes[start:current])
			literal := string(runes[start+1:current-1])

			tokens = append(tokens, Token{String, lexeme, literal, line})
		default:
			if(isDigit(c)){
				number()
			} else if(isLetter(c)){
				start = current - 1

				for isLetter(peek()) || isDigit(peek()) {
						advance()
					}
				text := string(runes[start:current])
				if t, ok := keywords[string(runes[start:current])]; ok {
					tokens = append(tokens, Token{t, text, "", line})
				} else {
					tokens = append(tokens, Token{Identifier, text, "", line})
				}
			} else {
				fmt.Printf("unknown character '%v' at line %d\n", string(runes), line)
			}

		}
		return nil
	}

	for !isAtEnd() {
		start = current
		if err := scanToken(); err != nil {
			return nil
		}
	}


	//tokens = append(tokens, Token{Eof, "", "", line}) // Eof - End of File
	return tokens
}
