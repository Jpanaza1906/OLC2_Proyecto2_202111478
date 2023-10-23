package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_Break struct {
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_Break(linea int, columna int) *NT_Break {
	return &NT_Break{
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================
func (NtBreak *NT_Break) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	return compilador.NewNill()
}

// NT_Continue ====================================================================================

type NT_Continue struct {
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_Continue(linea int, columna int) *NT_Continue {
	return &NT_Continue{
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (NtContinue *NT_Continue) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	return compilador.NewNill()
}

// NT_Return ====================================================================================
type NT_Return struct {
	Expr    compilador.CAbstractExpr
	Linea   int
	Columna int
}

// Constructor ================================================================================
func NewNT_Return(expr compilador.CAbstractExpr, linea int, columna int) *NT_Return {
	return &NT_Return{
		Expr:    expr,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (NtReturn *NT_Return) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	return compilador.NewNill()
}
