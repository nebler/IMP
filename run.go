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

func working() {
	testBrackets := "{z :=((-1))}"
	testAddition := "{z :=2+(-1)}"
	input3 := "{z :=(-1)}"
	testNegativeNumbers := "{x := -12345678; y := 3}"
	testMultipleOverwrite := "{x := -12345678; x = 1; y := 12; test := -12; zzz := 4; x = 2}"
	testOverwrite := "{x := -12345678; x = 1}"
	testMultiplication := "{x := -10; y := x * 2}"
	testMultiplication2 := "{x := -1; y := 4; z := x * y}"
	testMultiplicationWithBrackets := "{x := -1; y := 4; z := y*(x + 4)}"
	testBooleanTrue := "{z := true }"
	testBooleanFalse := "{z := false }"
	testNotFalse := "{z := false; x := !z }"
	testNotTrue := "{z := true; x := !z }"
	testAnd := "{z := true; x := z && false }"
	testOr := "{z := true; x := z || false }"
	testCompareTrue := "{z := 1; x := z == 1 }"
	testCompareFalse := "{z := 1; x := z == 2 }"
	testTrueAndTrue := "{z := true; x := true && z }"
	testCompareTrueIsZ := "{z := true; x := true == z }"
	testCompareZIsTrue := "{z := true; x := z == true }"
	input21 := "{x :=0; x = x+1}"
	input22 := "{z :=((-1))}"
	input23 := "{z :=(-1)+2}"
	testAndLesser := "{z := 3; z := true && (2 < z)}"
	testLesserAnd := "{z := 3; z := ((2 < z) && true)}"
	testIfElse := "{x := 0;y:=0; if x < 4 {x = x+1; x = x*10; y:= 100; y = y * 2} else {y:=1}}"
	testIfElseElse := "{x := 0;y:=0; if x < 4 {if x == 0 {x = x+1; x = x*10; y:= 100; y = y * 2} else {y:=100}}} else {y:=1}}"
	input28 := "{x := 0;y:=0; if 4 < x {x = x+1; x = x*10; y:= 100; y = y * 2} else {y=1}}"
	testPrintWithBrackets := "{x :=15; y:=100; print (x+y)}"
	testPrintWithBrackets2 := "{x :=0; while x<5{x = x+1}; print x}"
	testPrintWithBrackets3 := "{x :=0; while x<5{x = x+1; print x}}"
	testMultipleFunctions := "{x :=0; if x<5 {x = x+2; print x} else {x = 1}}"
	testMultipleFunctions2 := "{x :=0; if x<5 {x = x+2} else {x = 1}; print x}"
	runstmt(parse(testBrackets))
	runstmt(parse(testAddition))
	runstmt(parse(input3))
	runstmt(parse(testNegativeNumbers))
	runstmt(parse(testOverwrite))
	runstmt(parse(testMultipleOverwrite))
	runstmt(parse(testMultiplication))
	runstmt(parse(testMultiplication2))
	runstmt(parse(testMultiplicationWithBrackets))
	runstmt(parse(testBooleanTrue))
	runstmt(parse(testBooleanFalse))
	runstmt(parse(testNotFalse))
	runstmt(parse(testNotTrue))
	runstmt(parse(testAnd))
	runstmt(parse(testOr))
	runstmt(parse(testCompareTrue))
	runstmt(parse(testCompareFalse))
	runstmt(parse(testTrueAndTrue))
	runstmt(parse(testCompareTrueIsZ))
	runstmt(parse(testCompareZIsTrue))
	runstmt(parse(input21))
	runstmt(parse(input22))
	runstmt(parse(input23))
	runstmt(parse(testAndLesser))
	runstmt(parse(testLesserAnd))
	runstmt(parse(testIfElse))
	runstmt(parse(testIfElseElse))
	runstmt(parse(input28))
	runstmt(parse(testPrintWithBrackets))
	runstmt(parse(testPrintWithBrackets2))
	runstmt(parse(testPrintWithBrackets3))
	runstmt(parse(testMultipleFunctions))
	runstmt(parse(testMultipleFunctions2))
}

func ex12() {
	astIf := seq(decl("x", number(1)), seq(ifThenElse(less(variable("x"), number(4)), seq(assign("x", plus(variable("x"), number(1))), printStatement(variable("x"))), seq(decl("x", plus(variable("x"), number(1))), printStatement(variable("x")))), printStatement(variable("x"))))
	runstmt(astIf)

}

func experiment() {
	working()
}

func main() {
	experiment()
}
