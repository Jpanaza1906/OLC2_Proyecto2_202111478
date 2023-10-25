package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_Case struct {
	Expr       compilador.CAbstractExpr
	Sentencias compilador.CAbstractExpr
}

// Constructor ================================================================================

func NewNT_Case(expr compilador.CAbstractExpr, sentencias compilador.CAbstractExpr) *NT_Case {
	return &NT_Case{
		Expr:       expr,
		Sentencias: sentencias,
	}
}

// Implementacion =============================================================================

func (NtCase *NT_Case) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	/*
		<cod sentencias 3d>
		<goto Lsalida>
	*/
	ctx.GenComentario("Case >>>")
	NtCase.Sentencias.Compilar(ctx)
	Lsalida := ctx.PeekDisplaySwitch()
	ctx.GenComentario("goto Lsalida case ")
	ctx.Gen("goto " + Lsalida)
	return compilador.NewNill()
}
