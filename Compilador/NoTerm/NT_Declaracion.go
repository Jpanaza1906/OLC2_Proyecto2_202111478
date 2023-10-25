package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

// NT_Declaracion =============================================================================

type NT_Declaracion struct {
	Id        string
	Tipo      string
	Mutable   bool
	Expresion compilador.CAbstractExpr
	Linea     int
	Columna   int
}

// Constructor ================================================================================
func NewNT_Declaracion(id string, tipo string, mutable bool, expresion compilador.CAbstractExpr, linea int, columna int) *NT_Declaracion {
	return &NT_Declaracion{
		Id:        id,
		Tipo:      tipo,
		Mutable:   mutable,
		Expresion: expresion,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion =============================================================================

func (NtDeclaracion *NT_Declaracion) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	idsimbolo := ctx.GetAmbitoSimbolo(NtDeclaracion.Id)

	//si el simbolo ya existe
	if idsimbolo.Id != "" {
		ctx.AddErrorLine("Semantico", "La variable "+NtDeclaracion.Id+" ya existe en el ambito "+strconv.Itoa(ctx.PeekAmbito()), NtDeclaracion.Linea, NtDeclaracion.Columna)

		return compilador.NewNill()
	}
	ctx.GenComentario("Declaracion variable")
	//si la expresion no es nula
	if NtDeclaracion.Expresion != nil {
		//se compila la expresion
		expresion := NtDeclaracion.Expresion.Compilar(ctx)
		if expresion.Tipo.String() == NtDeclaracion.Tipo || NtDeclaracion.Tipo == "" {
			//se genera el codigo 3d para asignar el valor a la variable
			t1 := ctx.NewTemp()
			ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt)) //la tabla de simbolos me dira el desplazamiento
			ctx.Gen("stack[(int) " + t1 + "] = " + expresion.Dir)

			//se agrega el simbolo a la tabla de simbolos
			ctx.AddSimbolo(NtDeclaracion.Id, "variable", expresion.Tipo, ctx.PeekAmbito(), ctx.PosSt, nil, false, NtDeclaracion.Mutable)
			ctx.PosSt++
			return compilador.NewNill()
		}
		ctx.AddErrorLine("Semantico", "El tipo de la expresion no coincide con el tipo de la variable", NtDeclaracion.Linea, NtDeclaracion.Columna)
		return compilador.NewNill()
	}

	//se genera el codigo 3d para asignar el valor a la variable
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt)) //la tabla de simbolos me dira el desplazamiento
	ctx.Gen("stack[(int)" + t1 + "] = -202111478")

	//si la expresion es nula el simbolo
	tipoe := ctx.GetTipoE(NtDeclaracion.Tipo)
	ctx.AddSimbolo(NtDeclaracion.Id, "variable", tipoe, ctx.PeekAmbito(), ctx.PosSt, nil, false, NtDeclaracion.Mutable)
	ctx.PosSt++

	return compilador.NewNill()
}
