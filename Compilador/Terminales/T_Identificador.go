package terminales

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

type T_Identificador struct {
	ID      string
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewT_Identificador(id string, linea int, columna int) *T_Identificador {
	return &T_Identificador{
		ID:      id,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (Tid *T_Identificador) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	simbolo := ctx.GetSimbolo(Tid.ID)

	if simbolo.Tipo == 0 { //si es nil, no existe
		ctx.AddErrorLine("Semantico", "La variable "+Tid.ID+" no existe", Tid.Linea, Tid.Columna)
		return compilador.NewNill()
	}

	//codigo tres direcciones
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(simbolo.Size))
	t2 := ctx.NewTemp()
	ctx.Gen(t2 + " = stack[(int) " + t1 + "]")

	return compilador.NewAtributo(nil, nil, t2, "", simbolo.Tipo)

}
