package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
)

type NT_ExprLogica struct {
	ExprIzq  compilador.CAbstractExpr
	ExprDer  compilador.CAbstractExpr
	Operador string
	Linea    int
	Columna  int
}

// Constructor ================================================================================

func NewNT_ExprLogica(exprIzq compilador.CAbstractExpr, exprDer compilador.CAbstractExpr, operador string, linea int, columna int) *NT_ExprLogica {
	return &NT_ExprLogica{
		ExprIzq:  exprIzq,
		ExprDer:  exprDer,
		Operador: operador,
		Linea:    linea,
		Columna:  columna,
	}
}

// Implementacion =============================================================================

func (NtExprLogica *NT_ExprLogica) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//se compila la expresion izquierda
	exprIzq := NtExprLogica.ExprIzq.Compilar(ctx)
	//se compila la expresion derecha
	exprDer := NtExprLogica.ExprDer.Compilar(ctx)

	//se verifica que los tipos sean iguales
	if exprIzq.Tipo != exprDer.Tipo {
		ctx.AddErrorLine("Semantico", "Los tipos de las expresiones no son iguales", NtExprLogica.Linea, NtExprLogica.Columna)
		return compilador.NewNill()
	}

	//se verifica que los tipos sean booleanos
	if exprIzq.Tipo != compilador.Bool {
		ctx.AddErrorLine("Semantico", "Los tipos de las expresiones no son booleanos", NtExprLogica.Linea, NtExprLogica.Columna)
		return compilador.NewNill()
	}

	switch NtExprLogica.Operador {
	case "&&":
		//se genera el codigo 3d para la operacion
		ctx.GenComentario("&&")
		//se crea el temporal
		t1 := ctx.NewTemp()
		l1 := ctx.NewEtq()
		ctx.Gen("if (" + exprIzq.Dir + " == 0) goto " + l1)
		ctx.Gen("if (" + exprDer.Dir + " == 0) goto " + l1)
		ctx.Gen(t1 + " = 1")
		l2 := ctx.NewEtq()
		ctx.Gen("goto " + l2)
		ctx.GenLabel(l1 + ":")
		ctx.Gen(t1 + " = 0")
		ctx.GenLabel(l2 + ":")
		return compilador.NewBool(t1)
	case "||":
		//se genera el codigo 3d para la operacion
		ctx.GenComentario("||")

		//se crea el temporal
		t1 := ctx.NewTemp()
		l1 := ctx.NewEtq()
		ctx.Gen("if (" + exprIzq.Dir + " == 1) goto " + l1)
		ctx.Gen("if (" + exprDer.Dir + " == 1) goto " + l1)
		ctx.Gen(t1 + " = 0")
		l2 := ctx.NewEtq()
		ctx.Gen("goto " + l2)
		ctx.GenLabel(l1 + ":")
		ctx.Gen(t1 + " = 1")
		ctx.GenLabel(l2 + ":")
		return compilador.NewBool(t1)
	}
	return compilador.NewNill()
}

// NT_ExprRelacional ==========================================================================

type NT_ExprRelacional struct {
	ExprIzq  compilador.CAbstractExpr
	ExprDer  compilador.CAbstractExpr
	Operador string
	Linea    int
	Columna  int
}

// Constructor ================================================================================

func NewNT_ExprRelacional(exprIzq compilador.CAbstractExpr, exprDer compilador.CAbstractExpr, operador string, linea int, columna int) *NT_ExprRelacional {
	return &NT_ExprRelacional{
		ExprIzq:  exprIzq,
		ExprDer:  exprDer,
		Operador: operador,
		Linea:    linea,
		Columna:  columna,
	}
}

// Implementacion =============================================================================

func (NtExprRelacional *NT_ExprRelacional) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//se compila la expresion izquierda
	exprIzq := NtExprRelacional.ExprIzq.Compilar(ctx)
	//se compila la expresion derecha
	exprDer := NtExprRelacional.ExprDer.Compilar(ctx)

	//se verifica que los tipos sean iguales
	if exprIzq.Tipo != exprDer.Tipo {
		ctx.AddErrorLine("Semantico", "Los tipos de las expresiones no son iguales", NtExprRelacional.Linea, NtExprRelacional.Columna)
		return compilador.NewNill()
	}

	switch NtExprRelacional.Operador {
	case "==":
		//se genera el codigo 3d para la operacion
		ctx.GenComentario("==")

		//se crea el temporal
		t1 := ctx.NewTemp()
		l1 := ctx.NewEtq()

		ctx.Gen("if (" + exprIzq.Dir + " == " + exprDer.Dir + ") goto " + l1)
		ctx.Gen(t1 + " = 0")
		l2 := ctx.NewEtq()
		ctx.Gen("goto " + l2)
		ctx.GenLabel(l1 + ":")
		ctx.Gen(t1 + " = 1")
		ctx.GenLabel(l2 + ":")
		return compilador.NewBool(t1)
	case "!=":
		//se genera el codigo 3d para la operacion
		ctx.GenComentario("!=")

		//se crea el temporal
		t1 := ctx.NewTemp()
		l1 := ctx.NewEtq()

		ctx.Gen("if (" + exprIzq.Dir + " != " + exprDer.Dir + ") goto " + l1)
		ctx.Gen(t1 + " = 0")
		l2 := ctx.NewEtq()
		ctx.Gen("goto " + l2)
		ctx.GenLabel(l1 + ":")
		ctx.Gen(t1 + " = 1")
		ctx.GenLabel(l2 + ":")
		return compilador.NewBool(t1)
	case ">":
		//se genera el codigo 3d para la operacion
		ctx.GenComentario(">")
		//se crea el temporal
		t1 := ctx.NewTemp()
		l1 := ctx.NewEtq()

		ctx.Gen("if (" + exprIzq.Dir + " > " + exprDer.Dir + ") goto " + l1)
		ctx.Gen(t1 + " = 0")
		l2 := ctx.NewEtq()
		ctx.Gen("goto " + l2)
		ctx.GenLabel(l1 + ":")
		ctx.Gen(t1 + " = 1")
		ctx.GenLabel(l2 + ":")
		return compilador.NewBool(t1)
	case "<":
		//se genera el codigo 3d para la operacion
		ctx.GenComentario("<")
		//se crea el temporal
		t1 := ctx.NewTemp()
		l1 := ctx.NewEtq()

		ctx.Gen("if (" + exprIzq.Dir + " < " + exprDer.Dir + ") goto " + l1)
		ctx.Gen(t1 + " = 0")
		l2 := ctx.NewEtq()
		ctx.Gen("goto " + l2)
		ctx.GenLabel(l1 + ":")
		ctx.Gen(t1 + " = 1")
		ctx.GenLabel(l2 + ":")
		return compilador.NewBool(t1)
	case ">=":
		//se genera el codigo 3d para la operacion
		ctx.GenComentario(">=")
		//se crea el temporal
		t1 := ctx.NewTemp()
		l1 := ctx.NewEtq()

		ctx.Gen("if (" + exprIzq.Dir + " >= " + exprDer.Dir + ") goto " + l1)
		ctx.Gen(t1 + " = 0")
		l2 := ctx.NewEtq()
		ctx.Gen("goto " + l2)
		ctx.GenLabel(l1 + ":")
		ctx.Gen(t1 + " = 1")
		ctx.GenLabel(l2 + ":")
		return compilador.NewBool(t1)
	case "<=":
		//se genera el codigo 3d para la operacion
		ctx.GenComentario("<=")
		//se crea el temporal
		t1 := ctx.NewTemp()
		l1 := ctx.NewEtq()

		ctx.Gen("if (" + exprIzq.Dir + " <= " + exprDer.Dir + ") goto " + l1)
		ctx.Gen(t1 + " = 0")
		l2 := ctx.NewEtq()
		ctx.Gen("goto " + l2)
		ctx.GenLabel(l1 + ":")
		ctx.Gen(t1 + " = 1")
		ctx.GenLabel(l2 + ":")
		return compilador.NewBool(t1)
	}
	return compilador.NewNill()
}
