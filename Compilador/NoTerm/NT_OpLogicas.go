package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

// NT_Or =====================================================================================

type NT_Or struct {
	CondIzq compilador.CAbstractExpr
	CondDer compilador.CAbstractExpr
}

// Constructor ================================================================================

func NewNT_Or(condIzq compilador.CAbstractExpr, condDer compilador.CAbstractExpr) *NT_Or {
	return &NT_Or{
		CondIzq: condIzq,
		CondDer: condDer,
	}
}

// Implementacion =============================================================================

func (NtOr *NT_Or) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	condIzq := NtOr.CondIzq.Compilar(ctx)
	//marcador
	ctx.ImprimirEtq(condIzq.EF)
	condDer := NtOr.CondDer.Compilar(ctx)

	return compilador.NewAtributo(ctx.Unir(condIzq.EV, condDer.EV), condDer.EF, "", "")

}

// NT_And ====================================================================================

type NT_And struct {
	CondIzq compilador.CAbstractExpr
	CondDer compilador.CAbstractExpr
}

// Constructor ================================================================================
func NewNT_And(condIzq compilador.CAbstractExpr, condDer compilador.CAbstractExpr) *NT_And {
	return &NT_And{
		CondIzq: condIzq,
		CondDer: condDer,
	}
}

// Implementacion =============================================================================

func (NtAnd *NT_And) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	condIzq := NtAnd.CondIzq.Compilar(ctx)
	//marcador
	ctx.ImprimirEtq(condIzq.EV)
	condDer := NtAnd.CondDer.Compilar(ctx)

	return compilador.NewAtributo(condDer.EV, ctx.Unir(condIzq.EF, condDer.EF), "", "")

}

// NT_Not ====================================================================================

type NT_Not struct {
	Cond compilador.CAbstractExpr
}

// Constructor ================================================================================

func NewNT_Not(cond compilador.CAbstractExpr) *NT_Not {
	return &NT_Not{
		Cond: cond,
	}
}

// Implementacion =============================================================================

func (NtNot *NT_Not) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	cond := NtNot.Cond.Compilar(ctx)

	return compilador.NewAtributo(cond.EF, cond.EV, "", "")

}
