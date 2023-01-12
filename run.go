package main

import (
	"fmt"
)

func run(e Exp) {
	m := make(map[ValName]Val)
	s := ValState{"global", m}
	t := make(map[string]Type)
	fmt.Printf("\n ******* ")
	fmt.Printf("\n %s", e.pretty())
	fmt.Printf("\n %s", showVal(e.eval(s)))
	fmt.Printf("\n %s", showType(e.infer(t)))
}

func runstmt(stmt Stmt) {
	m := make(map[ValName]Val)
	s := ValState{"global", m}
	t := make(map[string]Type)
	fmt.Printf("\n ******* ")
	fmt.Printf("\n %s", stmt.pretty())
	stmt.eval(s)
	fmt.Printf("\n %v", s)
	fmt.Printf("\n %v", stmt.check(t))
}

func ex1() {
	ast := plus(mult(number(1), number(2)), number(0))

	run(ast)
}

func ex2() {
	ast := and(boolean(false), number(0))
	run(ast)
}

func ex3() {
	ast := or(boolean(false), number(0))
	run(ast)
}

func ex4() {
	ast := less(boolean(false), number(0))
	run(ast)
}

func ex5() {
	ast := less(number(0), number(1))
	run(ast)
}

func ex6() {
	ast := less(number(1), number(0))
	run(ast)
}

func ex7() {
	ast := seq(decl("x", number(1)), decl("y", plus(number(6), variable("x"))))
	runstmt(ast)
}

func ex8() {
	ast := seq(decl("x", number(1)), decl("y", plus(number(6), variable("x"))))
	runstmt(ast)
}

func ex9() {
	ast := seq(decl("x", number(1)), printStatement(variable("x")))
	runstmt(ast)
}

func ex10() {
	ast := ifThenElse(boolean(true), printStatement(number(1)), printStatement(number(2)))
	ast2 := ifThenElse(boolean(false), printStatement(number(1)), printStatement(number(2)))
	astWhile := seq(decl("x", number(1)), while(less(variable("x"), number(4)), seq(assign(variable("x"), plus(variable("x"), number(1))), printStatement(variable("x")))))
	runstmt(ast)
	runstmt(ast2)
	runstmt(astWhile)
}

func ex11() {
	astWhile := seq(decl("x", number(1)), seq(while(less(variable("x"), number(4)), seq(decl("x", plus(variable("x"), number(1))), printStatement(variable("x")))), printStatement(variable("x"))))
	runstmt(astWhile)
}

func ex12() {
	astIf := seq(decl("x", number(1)), seq(ifThenElse(less(variable("x"), number(4)), seq(decl("x", plus(variable("x"), number(1))), printStatement(variable("x"))), seq(decl("x", plus(variable("x"), number(1))), printStatement(variable("x")))), printStatement(variable("x"))))
	runstmt(astIf)
}

func main() {
	ex12()
}
