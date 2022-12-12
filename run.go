package main

import (
	"fmt"
	"os"
)

func run(e Exp) {
	s := make(map[string]Val)
	t := make(map[string]Type)
	fmt.Printf("\n ******* ")
	fmt.Printf("\n %s", e.pretty())
	fmt.Printf("\n %s", showVal(e.eval(s)))
	fmt.Printf("\n %s", showType(e.infer(t)))
}

func runstmt(stmt Stmt) {
	s := make(map[string]Val)
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

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
