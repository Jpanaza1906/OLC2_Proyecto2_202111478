package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

// NT_OpAritmeticas ============================================================================

type NT_OpAritmeticas struct {
	ExprIzq compilador.CAbstractExpr
	ExprDer compilador.CAbstractExpr
	Op      string
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_OpAritmeticas(exprIzq compilador.CAbstractExpr, exprDer compilador.CAbstractExpr, op string, linea int, columna int) *NT_OpAritmeticas {
	return &NT_OpAritmeticas{
		ExprIzq: exprIzq,
		ExprDer: exprDer,
		Op:      op,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (Nmd *NT_OpAritmeticas) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	exprIzq := Nmd.ExprIzq.Compilar(ctx)
	exprDer := Nmd.ExprDer.Compilar(ctx)

	if exprIzq.Tipo != exprDer.Tipo {
		if exprIzq.Tipo > exprDer.Tipo {
			exprDer = ctx.Conversor.Ampliar(exprDer, exprIzq.Tipo)
		} else {
			exprIzq = ctx.Conversor.Ampliar(exprIzq, exprDer.Tipo)
		}
	}

	if exprIzq.Tipo == compilador.Nil {
		compilador.NewContexto().AddErrorLine("Semantico", "No se puede operar aritmeticamente con tipos nil", Nmd.Linea, Nmd.Columna)
		return compilador.NewNill()
	}

	conversionIzq := ""
	conversionDer := ""
	tmp := ""

	switch exprIzq.Tipo {
	case compilador.Bool:
		compilador.NewContexto().AddErrorLine("Semantico", "No se puede operar aritmeticamente con tipos bool", Nmd.Linea, Nmd.Columna)
		return compilador.NewNill()
	case compilador.Integer:
		conversionIzq = "(float) "
		conversionDer = "(float) "
		tmp = ctx.NewTemp()
		ctx.Gen(tmp + " = " + conversionIzq + exprIzq.Dir + " " + Nmd.Op + " " + conversionDer + exprDer.Dir)
	case compilador.Float:
		conversionIzq = ""
		conversionDer = ""
		tmp = ctx.NewTemp()
		ctx.Gen(tmp + " = " + conversionIzq + exprIzq.Dir + " " + Nmd.Op + " " + conversionDer + exprDer.Dir)
	case compilador.String:
		//se buscan en el heap los valores de las cadenas
		//si es suma se concatenan
		if Nmd.Op == "+" {
			//hacer una lista de temporales para concatenar
			tmp = exprIzq.Dir
			ctx.Nat_ConcatString(exprIzq.Dir, exprDer.Dir) //se llama a la funcion nativa para concatenar
		} else {
			compilador.NewContexto().AddErrorLine("Semantico", "No se puede operar aritmeticamente con tipos string", Nmd.Linea, Nmd.Columna)
			return compilador.NewNill()
		}
	}
	return compilador.NewAtributo(nil, nil, tmp, "", exprIzq.Tipo)
}

// NT_Negativo ====================================================================================

type NT_Negativo struct {
	Expr    compilador.CAbstractExpr
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_Negativo(expr compilador.CAbstractExpr, linea int, columna int) *NT_Negativo {
	return &NT_Negativo{
		Expr:    expr,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (Nmd *NT_Negativo) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	expr := Nmd.Expr.Compilar(ctx)

	tmp := ctx.NewTemp()

	ctx.Gen(tmp + " = -1")
	ctx.Gen(tmp + " = (int)" + tmp + " * (int)" + expr.Dir)

	return compilador.NewInt(tmp)
}

// NT_Mod ====================================================================================

type NT_Mod struct {
	ExprIzq compilador.CAbstractExpr
	ExprDer compilador.CAbstractExpr
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_Mod(exprIzq compilador.CAbstractExpr, exprDer compilador.CAbstractExpr, linea int, columna int) *NT_Mod {
	return &NT_Mod{
		ExprIzq: exprIzq,
		ExprDer: exprDer,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (Nmd *NT_Mod) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	exprIzq := Nmd.ExprIzq.Compilar(ctx)
	exprDer := Nmd.ExprDer.Compilar(ctx)

	tmp := ctx.NewTemp()

	ctx.Gen(tmp + " = (int)" + exprIzq.Dir + " % (int)" + exprDer.Dir)

	return compilador.NewInt(tmp)
}
