package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_S struct {
	Sentencias compilador.CAbstractExpr
}

// Constructor ================================================================================

func NewNT_S(sentencias compilador.CAbstractExpr) *NT_S {
	return &NT_S{
		Sentencias: sentencias,
	}
}

// Implementacion =============================================================================

func (NtS *NT_S) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	return NtS.Sentencias.Compilar(ctx)
}
