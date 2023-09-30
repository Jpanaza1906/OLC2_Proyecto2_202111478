package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

// NT_OpAritmeticas ============================================================================

type NT_OpAritmeticas struct {
	ExprIzq compilador.CAbstractExpr
	ExprDer compilador.CAbstractExpr
	Op      string
}

// Constructor ================================================================================

func NewNT_OpAritmeticas(exprIzq compilador.CAbstractExpr, exprDer compilador.CAbstractExpr, op string) *NT_OpAritmeticas {
	return &NT_OpAritmeticas{
		ExprIzq: exprIzq,
		ExprDer: exprDer,
		Op:      op,
	}
}

// Implementacion =============================================================================

func (Nmd *NT_OpAritmeticas) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	exprIzq := Nmd.ExprIzq.Compilar(ctx)
	exprDer := Nmd.ExprDer.Compilar(ctx)

	tmp := ctx.NewTemp()

	ctx.Gen(tmp + " = " + exprIzq.Dir + " " + Nmd.Op + " " + exprDer.Dir)

	return compilador.NewDirAtributo(tmp)
}

// NT_Negativo ====================================================================================

type NT_Negativo struct {
	Expr compilador.CAbstractExpr
}

// Constructor ================================================================================

func NewNT_Negativo(expr compilador.CAbstractExpr) *NT_Negativo {
	return &NT_Negativo{
		Expr: expr,
	}
}

// Implementacion =============================================================================

func (Nmd *NT_Negativo) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	expr := Nmd.Expr.Compilar(ctx)

	tmp := ctx.NewTemp()

	ctx.Gen(tmp + " = -1")
	ctx.Gen(tmp + " = " + tmp + " * " + expr.Dir)

	return compilador.NewDirAtributo(tmp)
}

// NT_Mod ====================================================================================

type NT_Mod struct {
	ExprIzq compilador.CAbstractExpr
	ExprDer compilador.CAbstractExpr
}

// Constructor ================================================================================

func NewNT_Mod(exprIzq compilador.CAbstractExpr, exprDer compilador.CAbstractExpr) *NT_Mod {
	return &NT_Mod{
		ExprIzq: exprIzq,
		ExprDer: exprDer,
	}
}

// Implementacion =============================================================================

func (Nmd *NT_Mod) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	exprIzq := Nmd.ExprIzq.Compilar(ctx)
	exprDer := Nmd.ExprDer.Compilar(ctx)

	tmp := ctx.NewTemp()

	ctx.Gen(tmp + " = " + exprIzq.Dir + " % " + exprDer.Dir)

	return compilador.NewDirAtributo(tmp)
}