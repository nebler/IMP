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
	println("\n________")
	stmt.eval(s)
	//TODO: printing schöner
	println("\nstate:")
	fmt.Printf("\n %v ", s)
	println("\ntyping:")
	fmt.Printf("\n %v ", stmt.check(t))
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

func working() {
	//input := "{x := 2; y := 3; z := x + y}"
	input := "{z :=((-1))}"
	input2 := "{z :=2+(-1)}"
	input3 := "{z :=(-1)}"
	input4 := "{x := -12345678; y := 3}"
	input5 := "{x := -12345678; x = 1; y := 12; test := -12; zzz := 4; x = 2}"
	input6 := "{x := -12345678; x = 1}"
	input7 := "{x := -10; y := x * 2}"
	input8 := "{x := -1; y := 4; z := x * y}"
	input9 := "{x := -1; y := 4; z := y*(x + 4)}"
	input10 := "{z := true }"
	input11 := "{z := false }"
	input12 := "{z := false; x := !z }"
	input13 := "{z := true; x := !z }"
	input14 := "{z := true; x := z && false }"
	input15 := "{z := true; x := z || false }"
	input16 := "{z := 1; x := z == 1 }"
	input17 := "{z := 1; x := z == 2 }"
	input18 := "{z := true; x := true && z }"
	input19 := "{z := true; x := true == z }"
	input20 := "{z := true; x := z == true }"
	input21 := "{x :=0; x = x+1}"
	runstmt(parse(input))
	runstmt(parse(input2))
	runstmt(parse(input3))
	runstmt(parse(input4))
	runstmt(parse(input5))
	runstmt(parse(input6))
	runstmt(parse(input7))
	runstmt(parse(input8))
	runstmt(parse(input9))
	runstmt(parse(input10))
	runstmt(parse(input11))
	runstmt(parse(input12))
	runstmt(parse(input13))
	runstmt(parse(input14))
	runstmt(parse(input15))
	runstmt(parse(input16))
	runstmt(parse(input17))
	runstmt(parse(input18))
	runstmt(parse(input19))
	runstmt(parse(input20))
	runstmt(parse(input21))
}

func experiment() {

	input := "{z :=((-1))}"
	input2 := "{z :=(-1)+2}"
	input20 := "{z := 3; z := true && (2 < z)}"
	input19 := "{z := 3; z := ((2 < z) && true)}"
	runstmt(parse(input))
	runstmt(parse(input2))
	runstmt(parse(input20))
	runstmt(parse(input19))
	//input := "{x := 0; if x < 4 {x = x+1} else {y:=1}}"
}

func main() {
	experiment()
	/*

		input2 := "{x := -12345678; y := 3}"
		input3 := "{x := 1;if x < 1 {x := 2} else { x := 3};print x}"
		input4 := "{x := 1; while x < 4 {x:=x+1; print x}; print x}"
		input5 := "{x := true; y:= x == false}"
		input6 := "{x := 1; x = 2}"
		input7 := "{x := 1;if x < 1 {x = 1} else { x = 3};print x}"
		input8 := "{x := true; y:=!x}"
		input9 := "{x := true; y:=x && true}"
		input10 := "{x := true; y:=x || true}"
		input11 := "{x := 1; y:=x * 10}"
		runstmt(parse(input))
	*/
}
