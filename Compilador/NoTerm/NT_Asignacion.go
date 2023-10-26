package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
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
	// Verificar si existe la variable
	simbolo := ctx.GetSimbolo(NtAsignacion.ID)
	if simbolo.Tipo == 0 { //si es nil, no existe
		ctx.AddErrorLine("Semantico", "La variable "+NtAsignacion.ID+" no existe", NtAsignacion.Linea, NtAsignacion.Columna)
		return compilador.NewNill()
	}

	// Verificar si la expresion es del mismo tipo
	expresion := NtAsignacion.Expresion.Compilar(ctx)
	if expresion.Tipo != simbolo.Tipo {
		ctx.AddErrorLine("Semantico", "La variable "+NtAsignacion.ID+" es de tipo "+simbolo.Tipo.String()+" y la expresion es de tipo "+expresion.Tipo.String(), NtAsignacion.Linea, NtAsignacion.Columna)
		return compilador.NewNill()
	}
	// Verificar si la expresion es de tipo nil
	if expresion.Tipo == compilador.Nil {
		ctx.AddErrorLine("Semantico", "La variable "+NtAsignacion.ID+" es de tipo "+simbolo.Tipo.String()+" y la expresion es de tipo "+expresion.Tipo.String(), NtAsignacion.Linea, NtAsignacion.Columna)
		return compilador.NewNill()
	}

	// verificar si es mutable
	if !simbolo.Mutable {
		ctx.AddErrorLine("Semantico", "La variable "+NtAsignacion.ID+" no es mutable", NtAsignacion.Linea, NtAsignacion.Columna)
		return compilador.NewNill()
	}

	// codigo de la asignacion
	ctx.GenComentario("Asignacion de la variable " + NtAsignacion.ID)
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(simbolo.Size))
	ctx.Gen("stack[(int) " + t1 + "] = " + expresion.Dir)

	return compilador.NewNill()
}
