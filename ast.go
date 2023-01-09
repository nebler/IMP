package main

// Helper functions to build ASTs by hand

func number(x int) Exp {
	return Num(x)
}

func boolean(x bool) Exp {
	return Bool(x)
}

func plus(x, y Exp) Exp {
	return (Plus)([2]Exp{x, y})

	// The type Plus is defined as the two element array consisting of Exp elements.
	// Plus and [2]Exp are isomorphic but different types.
	// We first build the AST value [2]Exp{x,y}.
	// Then cast this value (of type [2]Exp) into a value of type Plus.

}

func mult(x, y Exp) Exp {
	return (Mult)([2]Exp{x, y})
}

func and(x, y Exp) Exp {
	return (And)([2]Exp{x, y})
}

func or(x, y Exp) Exp {
	return (Or)([2]Exp{x, y})
}

func neg(x, y Exp) Exp {
	return (Neg)([1]Exp{x})
}

func less(x, y Exp) Exp {
	return (Less)([2]Exp{x, y})
}

func grp(x Exp) Exp {
	return (Grp)([1]Exp{x})
}

func equ(x, y Exp) Equ {
	return (Equ)([2]Exp{x, y})
}

func seq(x, y Stmt) Stmt {
	return (Seq)([2]Stmt{x, y})
}

func decl(lhs string, exp Exp) Stmt {
	return (Decl{lhs, exp})
}

func variable(variableName string) Exp {
	return (Var(variableName))
}

func errorStmt(errorMessage string) Stmt {
	return (ErorrStatement(errorMessage))
}
