package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
	//"strconv"
)

type NT_DecMatriz struct {
	Id         string
	Tipo       string
	Mutable    bool
	Def_Matriz compilador.CAbstractExpr
	Dim        int
	Linea      int
	Columna    int
}

// Constructor ================================================================================
func NewNT_DecMatriz(id string, tipo string, mutable bool, def_matriz compilador.CAbstractExpr, dim int, linea int, columna int) *NT_DecMatriz {
	return &NT_DecMatriz{
		Id:         id,
		Tipo:       tipo,
		Mutable:    mutable,
		Def_Matriz: def_matriz,
		Dim:        dim,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion =============================================================================

func (NtDecMatriz *NT_DecMatriz) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//id simbolo
	idsimbolo := ctx.GetAmbitoSimbolo(NtDecMatriz.Id)

	//si el simbolo ya existe
	if idsimbolo.Tipo != 0 {
		ctx.AddErrorLine("Semantico", "La variable "+NtDecMatriz.Id+" ya existe en el ambito "+strconv.Itoa(ctx.PeekAmbito()), NtDecMatriz.Linea, NtDecMatriz.Columna)
		return compilador.NewNill()
	}

	//declaracion de la variable
	ctx.GenComentario("Declaracion matriz")

	//si def_matriz no puede ser nulo
	if NtDecMatriz.Def_Matriz != nil {
		//vector de int
		vectorv := []int{}
		//castear def_matriz a NT_DefMatriz
		NtDefMatriz := NtDecMatriz.Def_Matriz.(*NT_DefMatriz)
		for i := 0; i < NtDecMatriz.Dim-1; i++ {
			vectorv = append(vectorv, len(NtDefMatriz.ListaExp))
			//se castea la primera lista de ntdefmatriz
			//si en la lista de expresion hay un numero
			NtDefMatriz = NtDefMatriz.ListaExp[0].(*NT_DefMatriz)
		}
		vectorv = append(vectorv, len(NtDefMatriz.ListaExp))

		ctx.AddSimbolo(NtDecMatriz.Id, "Matriz", compilador.Integer, ctx.PeekAmbito(), ctx.PosSt, vectorv, false, NtDecMatriz.Mutable)
		ctx.PosSt++

		t0 := ctx.NewTemp()
		ctx.Gen(t0 + " = H")

		NtDecMatriz.Def_Matriz.Compilar(ctx)

		t1 := ctx.NewTemp()
		ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt-1)) //la tabla de simbolos me dira el desplazamiento
		ctx.Gen("stack[(int) " + t1 + "] = " + t0)
	}

	return compilador.NewNill()
}
