package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
)

// NT_Asignacion =============================================================================

type NT_Asignacion struct {
	ID        string
	Expresion compilador.CAbstractExpr
	Linea     int
	Columna   int
}

// Constructor ================================================================================
func NewNT_Asignacion(id string, expresion compilador.CAbstractExpr, linea int, columna int) *NT_Asignacion {
	return &NT_Asignacion{
		ID:        id,
		Expresion: expresion,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion =============================================================================

func (NtAsignacion *NT_Asignacion) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	return compilador.NewNill()
}
