package terminales

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

// T_MatrizPos ==================================================================================

type T_MatrizPos struct {
	Id        string
	Expresion []compilador.CAbstractExpr
	Linea     int
	Columna   int
}

// Constructor ================================================================================

func NewT_MatrizPos(id string, expresion []compilador.CAbstractExpr, linea int, columna int) *T_MatrizPos {
	return &T_MatrizPos{
		Id:        id,
		Expresion: expresion,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion =============================================================================

func (TMatrizPos *T_MatrizPos) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	idsimbolo := ctx.GetSimbolo(TMatrizPos.Id)

	//si el simbolo no existe
	if idsimbolo.Tipo == 0 {
		ctx.AddErrorLine("Semantico", "La variable "+TMatrizPos.Id+" no existe", TMatrizPos.Linea, TMatrizPos.Columna)
		return compilador.NewNill()
	}

	//si el simbolo no es una matriz
	if idsimbolo.TipoId != compilador.Matriz.String() {
		ctx.AddErrorLine("Semantico", "La variable "+TMatrizPos.Id+" no es una matriz", TMatrizPos.Linea, TMatrizPos.Columna)
		return compilador.NewNill()
	}

	//se obtiene el tipo de la matriz y se compilan las expresiones
	exp := []compilador.Atributos{}
	for _, e := range TMatrizPos.Expresion {
		exp = append(exp, *e.Compilar(ctx))
	}
	tsuma := ctx.NewTemp()
	ctx.Gen(tsuma + " = 0")
	idsimbolo.CorregirTam()

	for i, ex := range exp {
		//si es el ultimo
		temp := ctx.NewTemp()
		if i == len(exp)-1 {
			ctx.Gen(temp + " = " + ex.Dir)
		} else {
			ctx.Gen(temp + " = " + ex.Dir + " * " + strconv.Itoa(idsimbolo.Valores[i+1]))
		}
		ctx.Gen(tsuma + " = " + tsuma + " + " + temp)
	}

	//codigo tres direcciones
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(idsimbolo.Size))
	t2 := ctx.NewTemp()
	ctx.Gen(t2 + " = stack[(int) " + t1 + "]")
	ctx.Gen(tsuma + " = " + tsuma + " + " + t2)

	t3 := ctx.NewTemp()
	ctx.Gen(t3 + " = heap[(int) " + tsuma + "]")

	return compilador.NewInt(t3)

}
