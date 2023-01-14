package main

import (
	"fmt"
	"os"
	"strings"
)

func (stmt Seq) eval(s ValState) {
	stmt[0].eval(s)
	stmt[1].eval(s)
}

func (ite IfThenElse) eval(s ValState) {
	// x => 1
	//cond x < 1

	/*
		x := 1;
		if x < 1 {
		x := 2
		} else {
		x := 3
		};
		print x
		ausgabe 1

		x := 1;
		if x < 1 {
		x := 2
		} else {
		x = 3
		};
		print x
		ausgabe 3
	*/
	v := ite.cond.eval(s)
	if v.flag == ValueBool {
		s2 := makeNewScope(s, "-ifThenElse")
		switch {
		case v.valB:
			ite.thenStmt.eval(s2)
		case !v.valB:
			ite.elseStmt.eval(s2)
			// x => 3
		}
		// TODO: if you declare a new variable it will leak into the global scope
		// IDEA: attach name to  inside scope so we know when the variable has been declared or not: type ValName [2]string has name of variable and scope it is defined in
		update(s, s2)
	} else {
		fmt.Printf("if-then-else eval fail")
	}
}

// Maps are represented via points.
// Hence, maps are passed by "reference" and the update is visible for the caller as well.
// x := 1 + 1
// x := 2
// [x, scopeName] => 2
func (decl Decl) eval(s ValState) {
	v := decl.rhs.eval(s)
	x := (string)(decl.lhs)
	s.vals[ValName{x, s.name}] = v
}

// Maps are represented via points.
// Hence, maps are passed by "reference" and the update is visible for the caller as well.
func (assign Assign) eval(s ValState) {
	v := assign.rhs.eval(s)
	val := assign.variable.eval(s)
	if val.flag == Undefined {
		fmt.Printf("ERROR: cannot assign %v to %v because it isnt declared yet.", v, assign.variable.pretty())
		os.Exit(3)
		return
	} else if v.flag == Undefined {
		fmt.Printf("ERROR: cannot eval expression for variable: %v", assign.variable.pretty())
		return
	} else if v.flag != val.flag {
		fmt.Printf("ERROR: cannot assign value because different types")
		return
	} else {
		scope := findScopeOfVariable(s, assign.variable.pretty())
		s.vals[ValName{assign.variable.pretty(), scope}] = v
	}
}

/*
// s = global
x := 1;
// s {[x, global]: 1}
y := 2
// s {[x, global]: 1, [y, global]: 2}
//createNewScopeFrom()
// s2: gobal-else {[x, global]: 1, [y, global]: 2}
if x < 1 {
x := 2
} else {
// s2: gobal-else {[x, global]: 1, [y, global]: 2}
x := 3
// s2: gobal-else {[x, global-else]: 3, [y, global]: 2}
x = 5
x = 6
y = 1
// s2: gobal-else {[x, global-else]: 3, [y, global]: 1}
};

print x
print y
*/
func update(s1 ValState, s2 ValState) {
	for k := range s2.vals {
		_, ok := s1.vals[k]
		if ok {
			s1.vals[k] = s2.vals[k]
		}
	}
}

/*
global: if-else-then
global-if-else-then
global-if-else-then-if-else-then
global-if-else-then-if-else-then-while-while
*/
func makeNewScope(s ValState, prefix string) ValState {
	m := make(map[ValName]Val)
	s2 := ValState{s.name + prefix, m}
	for k, v := range s.vals {

		s2.vals[k] = v
	}
	return s2
}

func (while While) eval(s ValState) {
	cond := while.cond.eval(s)
	if cond.flag == ValueBool {
		s2 := makeNewScope(s, "-while")
		for cond.valB {
			while.stmt.eval(s2)
			cond = while.cond.eval(s2)
		}
		update(s, s2)
	} else {
		fmt.Printf("ERROR: Condition isnt boolean cannot eval while-loop")
	}
}

func (print Print) eval(s ValState) {
	isBool := print.printExp.eval(s).flag == ValueBool
	if isBool {
		fmt.Println()
		fmt.Printf("Output: %v", print.printExp.eval(s).valB)
	} else {
		fmt.Println()
		fmt.Printf("Output: %v", print.printExp.eval(s).valI)
	}
}

func (x Bool) eval(s ValState) Val {
	return mkBool((bool)(x))
}

func (x Num) eval(s ValState) Val {
	return mkInt((int)(x))
}

func (e Mult) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI * n2.valI)
	}
	return mkUndefined()
}

func findScopeOfVariable(s ValState, varName string) string {
	scope := "global"
	for k := range s.vals {
		if k[0] == varName {
			if strings.Count(scope, "-") < strings.Count(k[1], "-") {
				scope = k[1]
			}
		}
	}
	return scope
}

func printValState(s ValState, prefix string) {
	for k, v := range s.vals {
		println(prefix + " " + k[0] + k[1])
		println(v.valI)
	}
}

func (varName Var) eval(s ValState) Val {
	//printValState(s, "")
	value, ok := s.vals[ValName{varName.pretty(), s.name}]
	if ok {
		if value.flag == ValueInt {
			return mkInt(value.valI)
		} else {
			return mkBool(value.valB)
		}
	} else {
		scope := findScopeOfVariable(s, varName.pretty())
		valueScope, okScope := s.vals[ValName{varName.pretty(), scope}]
		if okScope {
			if valueScope.flag == ValueInt {
				return mkInt(valueScope.valI)
			} else {
				return mkBool(valueScope.valB)
			}
		} else {
			return mkUndefined()
		}
	}
}

func (e Plus) eval(s ValState) Val {
	n1 := e[0].eval(s)
	n2 := e[1].eval(s)
	if n1.flag == ValueInt && n2.flag == ValueInt {
		return mkInt(n1.valI + n2.valI)
	}
	return mkUndefined()
}

func (e And) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	if b1.flag == ValueBool && b2.flag == ValueBool {
		return mkBool(b1.valB && b2.valB)
	}
	return mkUndefined()
}

func (e Or) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	if b1.flag == ValueBool && b2.flag == ValueBool {
		return mkBool(b1.valB || b2.valB)
	}
	return mkUndefined()
}

func (e Neg) eval(s ValState) Val {
	b1 := e[0].eval(s)
	switch {
	case b1.flag == ValueBool:
		return mkBool(!b1.valB)
	}
	return mkUndefined()
}

func (e Equ) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB == b2.valB)
	case b1.flag == ValueInt && b2.flag == ValueInt:
		return mkBool(b1.valI == b2.valI)
	}
	return mkUndefined()
}

func (e Less) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueInt && b2.flag == ValueInt:
		return mkBool(b1.valI < b2.valI)
	}
	return mkUndefined()
}

/*
type Grp [2]Exp eval?
*/
func (e Grp) eval(s ValState) Val {
	b1 := e[0].eval(s)
	switch {
	case b1.flag == ValueBool:
		return mkBool(b1.valB)
	case b1.flag == ValueInt:
		return mkInt(b1.valI)
	}
	return mkUndefined()
}

func (errorStmt ErorrStatement) eval(s ValState) {
	print(errorStmt)
}

func (errorExp ErrorExp) eval(s ValState) Val {
	print(errorExp)
	return mkUndefined()
}
