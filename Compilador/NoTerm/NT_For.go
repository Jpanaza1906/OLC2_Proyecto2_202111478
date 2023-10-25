package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

// NT_For Integer===============================================================

type NT_ForInt struct {
	Id         string
	ExpIzq     compilador.CAbstractExpr
	ExpDer     compilador.CAbstractExpr
	Sentencias compilador.CAbstractExpr
	Linea      int
	Columna    int
}

// Constructor ================================================================================

func NewNT_ForInt(id string, expIzq compilador.CAbstractExpr, expDer compilador.CAbstractExpr, sentencias compilador.CAbstractExpr, linea int, columna int) *NT_ForInt {
	return &NT_ForInt{
		Id:         id,
		ExpIzq:     expIzq,
		ExpDer:     expDer,
		Sentencias: sentencias,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion =============================================================================

func (NtForInt *NT_ForInt) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//
	return compilador.NewNill()
}
