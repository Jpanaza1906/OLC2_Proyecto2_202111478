package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_OpRelacionales struct {
	ExprIzq compilador.CAbstractExpr
	ExprDer compilador.CAbstractExpr
	Op      string
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_OpRelacionales(exprIzq compilador.CAbstractExpr, exprDer compilador.CAbstractExpr, op string, linea int, columna int) *NT_OpRelacionales {
	return &NT_OpRelacionales{
		ExprIzq: exprIzq,
		ExprDer: exprDer,
		Op:      op,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (Nopr *NT_OpRelacionales) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	exprIzq := Nopr.ExprIzq.Compilar(ctx)
	exprDer := Nopr.ExprDer.Compilar(ctx)

	EV := ctx.NewEtq()
	EF := ctx.NewEtq()

	EVs := make([]string, 0)
	EFs := make([]string, 0)

	EVs = append(EVs, EV)
	EFs = append(EFs, EF)

	ctx.Gen("if " + exprIzq.Dir + " " + Nopr.Op + " " + exprDer.Dir + " then goto " + EV)
	ctx.Gen("goto " + EF)

	return compilador.NewAtributo(EVs, EFs, "", "")
}
