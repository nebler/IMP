package main

import "strconv"

// pretty print

func (stmt Seq) pretty() string {
	return stmt[0].pretty() + "; " + stmt[1].pretty()
}

func (decl Decl) pretty() string {
	return decl.lhs + " := " + decl.rhs.pretty()
}

func (ifThenElse IfThenElse) pretty() string {
	return "if " + ifThenElse.cond.pretty() + ifThenElse.thenStmt.pretty() + "else " + ifThenElse.thenStmt.pretty()
}

func (while While) pretty() string {
	return "while " + while.cond.pretty() + while.stmt.pretty()
}

func (print Print) pretty() string {
	return "print " + print.printExp.pretty()
}

func (assignment Assign) pretty() string {
	return assignment.variable.pretty() + " =" + assignment.rhs.pretty()
}

func (x Var) pretty() string {
	return (string)(x)
}

func (x Bool) pretty() string {
	if x {
		return "true"
	} else {
		return "false"
	}
}

func (x Num) pretty() string {
	return strconv.Itoa(int(x))
}

func (e Mult) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "*"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Plus) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "+"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e And) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "&&"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Or) pretty() string {

	var x string
	x = "("
	x += e[0].pretty()
	x += "||"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Neg) pretty() string {
	var x string
	x = "!"
	x += e[0].pretty()

	return x
}

func (e Equ) pretty() string {
	var x string
	x = "("
	x += e[0].pretty()
	x += "=="
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Less) pretty() string {
	var x string
	x = "("
	x += e[0].pretty()
	x += "<"
	x += e[1].pretty()
	x += ")"

	return x
}

func (e Grp) pretty() string {
	var x string
	x = "("
	x += e[0].pretty()
	x += ")"

	return x
}

func (block Block) pretty() string {
	var x string
	x = "{"
	x += block.pretty()
	x += "}"

	return x
}

func (errorStmt ErorrStatement) pretty() string {
	return string(errorStmt)
}

func (errorExp ErrorExp) pretty() string {
	return string(errorExp)
}
