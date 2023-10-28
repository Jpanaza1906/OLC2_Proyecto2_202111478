package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

// NT_Or =====================================================================================

type NT_Or struct {
	CondIzq compilador.CAbstractExpr
	CondDer compilador.CAbstractExpr
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_Or(condIzq compilador.CAbstractExpr, condDer compilador.CAbstractExpr, linea int, columna int) *NT_Or {
	return &NT_Or{
		CondIzq: condIzq,
		CondDer: condDer,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (NtOr *NT_Or) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	condIzq := NtOr.CondIzq.Compilar(ctx)
	//marcador
	ctx.ImprimirEtq(condIzq.EF)
	condDer := NtOr.CondDer.Compilar(ctx)

	return compilador.NewAtributo(ctx.Unir(condIzq.EV, condDer.EV), condDer.EF, "", "", compilador.Bool)

}

// NT_And ====================================================================================

type NT_And struct {
	CondIzq compilador.CAbstractExpr
	CondDer compilador.CAbstractExpr
	Linea   int
	Columna int
}

// Constructor ================================================================================
func NewNT_And(condIzq compilador.CAbstractExpr, condDer compilador.CAbstractExpr, linea int, columna int) *NT_And {
	return &NT_And{
		CondIzq: condIzq,
		CondDer: condDer,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (NtAnd *NT_And) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	condIzq := NtAnd.CondIzq.Compilar(ctx)
	//marcador
	ctx.ImprimirEtq(condIzq.EV)
	condDer := NtAnd.CondDer.Compilar(ctx)

	return compilador.NewAtributo(condDer.EV, ctx.Unir(condIzq.EF, condDer.EF), "", "", compilador.Bool)

}

// NT_Not ====================================================================================

type NT_Not struct {
	Cond    compilador.CAbstractExpr
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_Not(cond compilador.CAbstractExpr, linea int, columna int) *NT_Not {
	return &NT_Not{
		Cond:    cond,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (NtNot *NT_Not) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	cond := NtNot.Cond.Compilar(ctx)

	return compilador.NewAtributo(cond.EF, cond.EV, "", "", compilador.Bool)

}

// NT_NotExp =================================================================================

type NT_NotExp struct {
	Exp    compilador.CAbstractExpr
	Linea  int
	Columa int
}

// Constructor ================================================================================

func NewNT_NotExp(exp compilador.CAbstractExpr, linea int, columna int) *NT_NotExp {
	return &NT_NotExp{
		Exp:    exp,
		Linea:  linea,
		Columa: columna,
	}
}

// Implementacion =============================================================================

func (NtNotExp *NT_NotExp) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	exp := NtNotExp.Exp.Compilar(ctx)

	//verificar que sea booleano
	if exp.Tipo != compilador.Bool {
		ctx.AddErrorLine("Semantico", "La expresion no es booleana", NtNotExp.Linea, NtNotExp.Columa)
		return compilador.NewNill()
	}

	// si en dir viene 1, se cambia por 0
	// si en dir viene 0, se cambia por 1

	//se genera el codigo 3d para asignar el valor a la variable
	ctx.GenComentario("Not")
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = " + exp.Dir)
	//if que si viene 1, se cambia por 0
	l1 := ctx.NewEtq()
	l2 := ctx.NewEtq()

	ctx.Gen("if (" + t1 + " == 1)" + " goto " + l1)
	ctx.Gen(t1 + " = 1")
	ctx.Gen("goto " + l2)
	ctx.GenLabel(l1 + ":")
	ctx.Gen(t1 + " = 0")
	ctx.GenLabel(l2 + ":")

	return compilador.NewBool(t1)

}
