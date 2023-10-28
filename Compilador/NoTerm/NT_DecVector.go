package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

type NT_DecVector struct {
	Id         string
	Tipo       string
	Mutable    bool
	Def_Vector compilador.CAbstractExpr
	Linea      int
	Columna    int
}

// Constructor ================================================================================
func NewNT_DecVector(id string, tipo string, mutable bool, def_vector compilador.CAbstractExpr, linea int, columna int) *NT_DecVector {
	return &NT_DecVector{
		Id:         id,
		Tipo:       tipo,
		Mutable:    mutable,
		Def_Vector: def_vector,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion =============================================================================
func (NtDecVector *NT_DecVector) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//id simbolo
	idsimbolo := ctx.GetAmbitoSimbolo(NtDecVector.Id)

	//si el simbolo ya existe
	if idsimbolo.Tipo != 0 {
		ctx.AddErrorLine("Semantico", "La variable "+NtDecVector.Id+" ya existe en el ambito "+strconv.Itoa(ctx.PeekAmbito()), NtDecVector.Linea, NtDecVector.Columna)
		return compilador.NewNill()
	}

	//declaracion de la variable
	ctx.GenComentario("Declaracion vector")

	//si def_vector es nulo
	if NtDecVector.Def_Vector == nil {
		//se crea un vector vacio
		vectorv := []compilador.CAbstractExpr{}
		vecvacio := NewNT_DefVector("", vectorv, NtDecVector.Linea, NtDecVector.Columna)
		t1 := vecvacio.Compilar(ctx)

		t2 := ctx.NewTemp()
		ctx.Gen(t2 + " = P + " + strconv.Itoa(ctx.PosSt)) //la tabla de simbolos me dira el desplazamiento
		ctx.Gen("stack[(int) " + t2 + "] = " + t1.Dir)

		//se agrega el simbolo a la tabla de simbolos
		ctx.AddSimbolo(NtDecVector.Id, "Vector", ctx.GetTipoE(NtDecVector.Tipo), ctx.PeekAmbito(), ctx.PosSt, nil, false, NtDecVector.Mutable)
		ctx.PosSt++
		return compilador.NewNill()
	}
	//si def_vector no es nulo
	def_vector := NtDecVector.Def_Vector.Compilar(ctx)

	if def_vector.Tipo == compilador.Vector {
		//se genera el codigo 3d para asignar el valor a la variable
		t1 := ctx.NewTemp()
		ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt)) //la tabla de simbolos me dira el desplazamiento
		ctx.Gen("stack[(int) " + t1 + "] = " + def_vector.Dir)

		//se agrega el simbolo a la tabla de simbolos
		ctx.AddSimbolo(NtDecVector.Id, "Vector", ctx.GetTipoE(NtDecVector.Tipo), ctx.PeekAmbito(), ctx.PosSt, nil, false, NtDecVector.Mutable)
		ctx.PosSt++
		return compilador.NewNill()
	}

	//si el tipo de la expresion no coincide con el tipo de la variable
	ctx.AddErrorLine("Semantico", "El tipo de la expresion no coincide con el tipo de la variable", NtDecVector.Linea, NtDecVector.Columna)
	return compilador.NewNill()
}
