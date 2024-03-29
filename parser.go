package main

import (
	"fmt"
	"strconv"
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
	ASSIGN     = 30
	DECL       = 31
	NEGATION   = 32
)

func scan(s string) (string, int) {
	for {
		switch {
		case len(s) == 0:
			return s, EOS
		case s[0] == ' ':
			s = s[1:]
		case s[0] == '0':
			return s[1:], ZERO
		case len(s) > 4 && s[0:4] == "true" && !IsLetter(string(s[4])):
			return s[4:], TRUE
		case len(s) > 5 && string(s[0:5]) == "while" && !IsLetter(string(s[5])):
			return s[5:], WHILE
		case len(s) > 5 && string(s[0:5]) == "print" && !IsLetter(string(s[5])):
			return s[5:], PRINT
		case len(s) > 5 && string(s[0:5]) == "false" && !IsLetter(string(s[5])):
			return s[5:], FALSE
		case len(s) > 4 && string(s[0:4]) == "else" && !IsLetter(string(s[4])):
			return s[4:], ELSE
		case s[0] == '1':
			return s[1:], ONE
		case s[0] == '2':
			return s[1:], TWO
		case s[0] == '3':
			return s[1:], THREE
		case s[0] == '4':
			return s[1:], FOUR
		case s[0] == '5':
			return s[1:], FIVE
		case s[0] == '6':
			return s[1:], SIX
		case s[0] == '7':
			return s[1:], SEVEN
		case s[0] == '8':
			return s[1:], EIGHT
		case s[0] == '9':
			return s[1:], NINE
		case s[0] == '+':
			return s[1:], PLUS
		case s[0] == '-':
			return s[1:], NEG
		case s[0] == '*':
			return s[1:], MULT
		case s[0] == '(':
			return s[1:], OPEN_GRP
		case s[0] == ')':
			return s[1:], CLOSE_GRP
		case s[0] == '{':
			return s[1:], OPEN_STMT
		case s[0] == '}':
			return s[1:], CLOSE_STMT
		case string(s[0:2]) == "if" && !IsLetter(string(s[2])):
			// if boolean then exp
			return s[2:], IF
		case s[0] == ';':
			return s[1:], SEQ
		case string(s[0:2]) == "==":
			return s[2:], EQU
		case s[0] == '=':
			return s[1:], ASSIGN
		case s[0] == '<':
			return s[1:], LESSER
		case string(s[0:2]) == ":=":
			return s[2:], DECL
		case s[0] == '!':
			return s[1:], NEGATION
		case string(s[0:2]) == "&&":
			return s[2:], AND
		case string(s[0:2]) == "||":
			return s[2:], OR
		case IsLower(s[0:0]):
			// falseVar
			// x, xy, xyz
			return s[0:], VARS
		default: // simply skip everything else
			s = s[1:]
		}
	}
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
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

/*
check string for numbers, returns number as string and integer of the position
*/
func parseTillNotAnNumberAnymore(s string) (string, int) {
	value := ""
	for i, c := range s {
		if unicode.IsNumber(c) {
			value = value + string(c)
		} else {
			return value, i
		}
	}
	return "", 0
}

/*
	parse number
*/

func parseNumber(s *State) (bool, Num) {
	numberString, index := parseTillNotAnNumberAnymore(*s.s)
	number, _ := strconv.Atoi(strconv.Itoa(s.tok) + numberString)
	*s.s = (*s.s)[(index):]
	return true, Num(number)
}

/*
parse Operation inside a group expression
| exp "+" exp           -- Addition
| exp "==" exp          -- Equality test
| exp "<" exp           -- Lesser test
| exp "*" exp           -- Multiplication
| exp "||" exp          -- Disjunction
| exp "&&" exp          -- Conjunction
*/
func parseOperation(exp Exp, s *State) (bool, Exp) {
	switch s.tok {
	case PLUS:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid2, Plus{exp, secondExp}
	case MULT:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid2, Mult{exp, secondExp}
	case AND:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid2, And{exp, secondExp}
	case OR:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid2, Or{exp, secondExp}
	case EQU:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid2, Equ{exp, secondExp}
	case LESSER:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid2, Less{exp, secondExp}
	}
	return false, errorExp("error")
}

/*
skipping parsing whatever is inside the group
*/
func parseGrp(s *State) (bool, Exp) {
	next(s)
	return parseExp(s)
}

func parseExp(s *State) (bool, Exp) {
	valid := false
	exp := errorExp("error")
	switch s.tok {
	case TRUE:
		valid = true
		exp = boolean(true)
	case FALSE:
		valid = true
		exp = boolean(false)
	case NEGATION:
		next(s)
		valid, exp = parseExp(s)
		exp = Neg{exp}
	case ZERO:
		valid = true
		exp = Num(0)
	case NEG:
		numberString, index := parseTillNotAnNumberAnymore(*s.s)
		number, _ := strconv.Atoi(numberString)
		*s.s = (*s.s)[(index):]
		valid = true
		exp = Num(0 - number)
	case VARS:
		varName := parseVarsString(s)
		valid, exp = parseVars(varName)
	case OPEN_GRP:
		valid, exp = parseGrp(s)
		if s.tok != CLOSE_GRP && s.tok != CLOSE_STMT {
			valid, exp = parseOperation(exp, s)
			exp = Grp{exp}
			if s.tok != CLOSE_GRP && s.tok != CLOSE_STMT {
				return false, errorExp("not closing grouped expression")
			}
		} else {
			exp = Grp{exp}
		}
	}
	// check if token is a number and not zero
	if s.tok < 10 && s.tok != 0 {
		valid, exp = parseNumber(s)
	}
	// check if statement is over if it isnt then there will be operation (+, *, ==, <, ||, &&)
	if s.tok != CLOSE_STMT {
		next(s)
	} else {
		return valid, exp
	}
	switch s.tok {
	case PLUS:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid && valid2, Plus{exp, secondExp}
	case MULT:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid && valid2, Mult{exp, secondExp}
	case AND:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid && valid2, And{exp, secondExp}
	case OR:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid && valid2, Or{exp, secondExp}
	case EQU:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid && valid2, Equ{exp, secondExp}
	case LESSER:
		next(s)
		valid2, secondExp := parseExp(s)
		return valid && valid2, Less{exp, secondExp}
	}
	return valid, exp
}

/*
|  "print" exp                       -- Print
*/
func parsePrint(s *State) (bool, Stmt) {
	valid, exp := parseExp(s)
	return valid, Print{exp}
}

/*
|  "if" exp block "else" block       -- If-then-else
*/
func parseIf(s *State) (bool, Stmt) {
	valid, exp := parseExp(s)
	if !valid {
		return false, errorStmt("invalid expression for if:" + exp.pretty())
	}
	next(s)
	validIfStmt, ifStmt := parseStmt(s)
	if !validIfStmt {
		return false, errorStmt("invalid statement inside if block")
	}
	if s.tok != CLOSE_STMT {
		return false, errorStmt("program not ending with }")
	}
	next(s)
	if s.tok != ELSE {
		next(s)
		if s.tok != ELSE {
			return false, errorStmt("else needs to follow after if")
		}
	}
	next(s)
	if s.tok != OPEN_STMT {
		return false, errorStmt("else needs to be followed by {")
	}
	next(s)
	validElse, elseStmt := parseStmt(s)
	if s.tok != CLOSE_STMT {
		return false, errorStmt("program not ending with }")
	}
	if !validElse {
		return false, errorStmt("invalid statement inside if block")
	}
	ifThen := IfThenElse{exp, ifStmt, elseStmt}
	next(s)
	return true, ifThen
}

/*
"while" exp block                 -- While
*/
func parseWhile(s *State) (bool, Stmt) {
	valid, exp := parseExp(s)
	if !valid {
		return false, errorStmt("invalid expression for while:" + exp.pretty())
	}
	next(s)
	validWhileStmt, whileStmt := parseStmt(s)

	if !validWhileStmt {
		return false, errorStmt("invalid statement inside while block")
	}
	if s.tok != CLOSE_STMT {
		return false, errorStmt("program not ending with }")
	}
	next(s)
	return true, While{exp, whileStmt}
}

func parseVars(varName string) (bool, Exp) {
	return true, Var(varName)
}

func parseAssign(varName string, s *State) (bool, Stmt) {
	validVars, expVars := parseVars(varName)
	if !validVars {
		return false, errorStmt("error when assigning variable: " + varName + "error is: " + expVars.pretty())
	} else {
		validExp, exp := parseExp(s)
		if !validExp {
			return false, errorStmt("error when evaluting: " + exp.pretty())
		}
		return validVars, Assign{Var(varName), exp}
	}
}

func parseDecl(varName string, s *State) (bool, Stmt) {
	valid, exp := parseExp(s)
	if !valid {
		return false, errorStmt("error when decaring variable: " + varName + " error is: " + exp.pretty())
	}
	return true, Decl{varName, exp}
}

/*
STATEMENTS
	|  vars ":=" exp                     -- Variable declaration
	|  vars "=" exp                      -- Variable assignment
*/

func parseVarsString(s *State) string {
	varName := ""
	index := 0
	for i, c := range *s.s {
		if c == ' ' || !unicode.IsLetter(c) {
			index = i
			break
		}
	}
	if index == 1 {
		varName = string((*s.s)[0])
		*s.s = (*s.s)[1:]
	} else {
		varName = (*s.s)[0:index]
		*s.s = (*s.s)[(index + 1):]
	}
	return varName
}

/*
|  vars ":=" exp                     -- Variable declaration
|  vars "=" exp                      -- Variable assignment
*/
func parseDeclOrAssign(s *State) (bool, Stmt) {
	varName := parseVarsString(s)
	next(s)
	switch s.tok {
	case ASSIGN:
		next(s)
		return parseAssign(varName, s)
	case DECL:
		next(s)
		return parseDecl(varName, s)
	}
	return false, errorStmt("error when assign or declare")
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
	//statement ; statement
	stmt := errorStmt("ERROR1")
	valid := false
	switch s.tok {
	case PRINT:
		next(s)
		valid, stmt = parsePrint(s)
	case IF:
		next(s)
		valid, stmt = parseIf(s)
	case WHILE:
		next(s)
		valid, stmt = parseWhile(s)
	case VARS:
		valid, stmt = parseDeclOrAssign(s)
	default:
		return false, errorStmt("ERROR2")
	}
	if s.tok == SEQ {
		//statement
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
		return false, errorStmt("failing to evaulute: " + stmt.pretty())
	}
	if s.tok != CLOSE_STMT {
		return false, errorStmt("program not ending with }")
	}
	return true, stmt
}

func parse(s string) Stmt {
	st := State{&s, EOS}
	next(&st)
	_, e := parseBlock(&st)
	next(&st)
	if st.tok == EOS {
		return e
	}
	return errorStmt("")
}

func debug(s string) {
	fmt.Printf("%s", s)
}
