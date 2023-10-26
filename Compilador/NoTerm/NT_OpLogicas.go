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
