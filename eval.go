package main

import "fmt"

func (stmt Seq) eval(s ValState) {
	stmt[0].eval(s)
	stmt[1].eval(s)
}

func (ite IfThenElse) eval(s ValState) {
	v := ite.cond.eval(s)
	if v.flag == ValueBool {
		s2 := s
		switch {
		case v.valB:
			ite.thenStmt.eval(s2)
		case !v.valB:
			ite.elseStmt.eval(s)
		}
		update(s, s2)
	} else {
		fmt.Printf("if-then-else eval fail")
	}

}

// Maps are represented via points.
// Hence, maps are passed by "reference" and the update is visible for the caller as well.
func (decl Decl) eval(s ValState) {
	v := decl.rhs.eval(s)
	x := (string)(decl.lhs)
	s[x] = v
}

// Maps are represented via points.
// Hence, maps are passed by "reference" and the update is visible for the caller as well.
func (assign Assign) eval(s ValState) {
	v := assign.rhs.eval(s)
	val := assign.variable.eval(s)
	if val.flag == Undefined {
		fmt.Printf("ERROR: cannot assign %v to %v because it isnt declared yet.", v, assign.variable.pretty())
		return
	} else if v.flag == Undefined {
		fmt.Printf("ERROR: cannot eval expression for variable: %v", assign.variable.pretty())
		return
	} else if v.flag != val.flag {
		fmt.Printf("ERROR: cannot assign value because different types")
		return
	} else {
		s[assign.variable.pretty()] = v
	}
}

func update(s1 ValState, s2 ValState) {
	for k := range s2 {
		val, ok := s1[k]
		if ok {
			s1[k] = val
		}
	}
}

func (while While) eval(s ValState) {
	cond := while.cond.eval(s)
	if cond.flag == ValueBool {
		s2 := s
		for cond.valB {
			while.stmt.eval(s2)
			update(s, s2)
		}
	} else {
		fmt.Printf("ERROR: Condition isnt boolean cannot eval while-loop")
	}
}

func (print Print) eval(s ValState) {
	isBool := print.printExp.eval(s).flag == ValueBool
	if isBool {
		fmt.Printf("%v", print.printExp.eval(s).valB)
	} else {
		fmt.Printf("%v", print.printExp.eval(s).valI)
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

func (varName Var) eval(s ValState) Val {
	value, ok := s[varName.pretty()]
	if ok {
		if value.flag == ValueInt {
			return mkInt(value.valI)
		} else {
			return mkBool(value.valB)
		}
	} else {
		return mkUndefined()
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
	switch {
	case b1.flag == ValueBool && b1.valB == false:
		//false && 0 kommt hier rein soll das sein?
		return mkBool(false)
	case b1.flag == ValueBool && b2.flag == ValueBool:
		return mkBool(b1.valB && b2.valB)
	}
	return mkUndefined()
}

func (e Or) eval(s ValState) Val {
	b1 := e[0].eval(s)
	b2 := e[1].eval(s)
	switch {
	case b1.flag == ValueBool && b1.valB == true:
		return mkBool(true)
	case b1.flag == ValueBool && b2.flag == ValueBool:
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
