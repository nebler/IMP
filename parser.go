package main

import (
	"fmt"
	"unicode"
)

// AST

// Simple imperative language

/*
vars       Variable names, start with lower-case letter

prog      ::= block
block     ::= "{" statement "}"
statement ::=  statement ";" statement           -- Command sequence
            |  vars ":=" exp                     -- Variable declaration
            |  vars "=" exp                      -- Variable assignment
            |  "while" exp block                 -- While
            |  "if" exp block "else" block       -- If-then-else
            |  "print" exp                       -- Print

exp ::= 0 | 1 | -1 | ...     -- Integers
     | "true" | "false"      -- Booleans
     | exp "+" exp           -- Addition
     | exp "*" exp           -- Multiplication
     | exp "||" exp          -- Disjunction
     | exp "&&" exp          -- Conjunction
     | "!" exp               -- Negation
     | exp "==" exp          -- Equality test
     | exp "<" exp           -- Lesser test
     | "(" exp ")"           -- Grouping of expressions
     | vars                  -- Variables

	 x := true
	 y := (x == false)
*/

// Tokens
const (
	EOS        = 99 // End of string
	ZERO       = 0
	ONE        = 1
	TWO        = 2
	THREE      = 3
	FOUR       = 4
	FIVE       = 5
	SIX        = 6
	SEVEN      = 7
	EIGHT      = 8
	NINE       = 9
	OPEN_STMT  = 10
	CLOSE_STMT = 11
	IF         = 12
	ELSE       = 13
	WHILE      = 14
	PRINT      = 15
	TRUE       = 16
	FALSE      = 17
	PLUS       = 18
	MINUS      = 19
	MULT       = 20
	AND        = 21
	OR         = 22
	EQU        = 23
	OPEN_GRP   = 24
	CLOSE_GRP  = 25
	LESSER     = 26
	NEG        = 27
	SEQ        = 28
	VARS       = 29
)

func scan(s string) (string, int) {
	for {
		switch {
		case len(s) == 0:
			return s, EOS
		case s[0] == '0':
			return s[1:], ZERO
		case s[0] == '1':
			return s[1:len(s)], ONE
		case s[0] == '2':
			return s[1:len(s)], TWO
		case s[0] == '3':
			return s[1:len(s)], THREE
		case s[0] == '4':
			return s[1:len(s)], FOUR
		case s[0] == '5':
			return s[1:len(s)], FIVE
		case s[0] == '6':
			return s[1:len(s)], SIX
		case s[0] == '7':
			return s[1:len(s)], SEVEN
		case s[0] == '8':
			return s[1:len(s)], EIGHT
		case s[0] == '9':
			return s[1:len(s)], NINE
		case s[0] == '+':
			return s[1:len(s)], PLUS
		case s[0] == '-':
			return s[1:len(s)], MINUS
		case s[0] == '*':
			return s[1:len(s)], MULT
		case s[0] == '(':
			return s[1:len(s)], OPEN_GRP
		case s[0] == ')':
			return s[1:len(s)], CLOSE_GRP
		case s[0] == '{':
			return s[1:len(s)], OPEN_STMT
		case s[0] == '}':
			return s[1:len(s)], CLOSE_STMT
		case s[0:2] == "if":
			return s[1:len(s)], IF
		case s[0] == ";":
			return s[1:len(s)], SEQ
		case s[0] == '<':
			return s[1:len(s)], LESSER
		case s[0:4] == "while":
			return s[4:len(s)], WHILE
		case s[0:2] == "&&":
			return s[2:len(s)], AND
		case s[0:2] == "||":
			return s[2:len(s)], OR
		case s[0:2] == "==":
			return s[2:len(s)], OR
		case s[0:4] == "true":
			return s[2:len(s)], TRUE
		case s[0:5] == "false":
			return s[5:len(s)], FALSE
		case unicode.IsLower(s[0]):

			s = s[1:len(s)]
			//logic missing here!!!
		case s[0] == ' ':
			s = s[1:len(s)]
		default: // simply skip everything else
			s = s[1:len(s)]
		}
	}
}

type State struct {
	s   *string
	tok int
}

func next(s *State) {
	s2, tok := scan(*s.s)
	s.s = &s2
	s.tok = tok
}

func parseExp(s *State) (bool, Exp) {

}

/*
statement ::=  statement ";" statement           -- Command sequence

	|  vars ":=" exp                     -- Variable declaration
	|  vars "=" exp                      -- Variable assignment
	|  "while" exp block                 -- While
	|  "if" exp block "else" block       -- If-then-else
	|  "print" exp                       -- Print
*/
func parseStmt(s *State) (bool, Stmt) {
	stmt := errorStmt("ERROR")
	valid := false
	switch s.tok {

	default:
		return false, errorStmt("ERROR")
	}

	if s.tok == SEQ {
		next(s)
		valid2, stmt2 := parseStmt(s)
		stmt = (Seq)([2]Stmt{stmt, stmt2})
		valid = valid && valid2
	}
	return valid, stmt
}

// block     ::= "{" statement "}"
func parseBlock(s *State) (bool, Stmt) {
	if s.tok != OPEN_STMT {
		return false, errorStmt("ERROR")
	}
	next(s)
	b, stmt := parseStmt(s)
	if !b {
		return false, errorStmt("ERROR")
	}
	return true, stmt
}

func parse(s string) Stmt {
	st := State{&s, EOS}
	next(&st)
	_, e := parseBlock(&st)
	if st.tok == EOS {
		return e
	}
	return errorStmt("ERROR")
}

func debug(s string) {
	fmt.Printf("%s", s)
}
