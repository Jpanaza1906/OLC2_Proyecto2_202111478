package terminales

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

// T_VectorPos ==================================================================================

type T_VectorPos struct {
	Id        string
	Expresion compilador.CAbstractExpr
	Linea     int
	Columna   int
}

// Constructor ================================================================================

func NewT_VectorPos(id string, expresion compilador.CAbstractExpr, linea int, columna int) *T_VectorPos {
	return &T_VectorPos{
		Id:        id,
		Expresion: expresion,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion =============================================================================

func (TVectorPos *T_VectorPos) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//con el id se obtiene el simbolo

	idsimbolo := ctx.GetSimbolo(TVectorPos.Id)

	//si el simbolo no existe
	if idsimbolo.Tipo == 0 {
		ctx.AddErrorLine("Semantico", "La variable "+TVectorPos.Id+" no existe", TVectorPos.Linea, TVectorPos.Columna)
		return compilador.NewNill()
	}

	//si el simbolo no es un vector
	if idsimbolo.TipoId != compilador.Vector.String() {
		ctx.AddErrorLine("Semantico", "La variable "+TVectorPos.Id+" no es un vector", TVectorPos.Linea, TVectorPos.Columna)
		return compilador.NewNill()
	}

	//se obtiene el tipo del vector y se compila la expresion
	exp := TVectorPos.Expresion.Compilar(ctx)

	//se verifica que la expresion sea de tipo entero
	if exp.Tipo != compilador.Integer {
		ctx.AddErrorLine("Semantico", "La expresion no es de tipo entero", TVectorPos.Linea, TVectorPos.Columna)
		return compilador.NewNill()
	}

	//se genera el codigo 3d para obtener el item del vector
	ctx.GenComentario("VectorPos")
	t1 := ctx.NewTemp()

	ctx.Gen(t1 + " = P + " + strconv.Itoa(idsimbolo.Size)) //la posision de inicio del vector
	t2 := ctx.NewTemp()
	ctx.Gen(t2 + " = stack[(int) " + t1 + "]") //la posicion del vector en el heap

	//hacemos un count para saber si la posicion es valida
	t3 := ctx.NewTemp()
	ctx.Nat_CountVector(t3, t2)
	l1 := ctx.NewEtq()

	t4 := ctx.NewTemp()
	ctx.Gen(t4 + " = -202111478")
	//si el valor de la expresion es mayor al tamaño del vector
	ctx.Gen("if (" + exp.Dir + " >= " + t3 + ") goto " + l1)

	ctx.Nat_VectorPos(t4, t2, exp.Dir) //en t4 queda el valor del item del vector
	ctx.GenLabel(l1 + ":")

	return compilador.NewAtributo(nil, nil, t4, "", idsimbolo.Tipo)
}
