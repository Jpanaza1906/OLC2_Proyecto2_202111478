package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

// NT_While ...
type NT_While struct {
	Condicion  compilador.CAbstractExpr
	Sentencias compilador.CAbstractExpr
	Linea      int
	Columna    int
}

// Constructor ================================================================================

func NewNT_While(condicion compilador.CAbstractExpr, sentencias compilador.CAbstractExpr, linea int, columna int) *NT_While {
	return &NT_While{
		Condicion:  condicion,
		Sentencias: sentencias,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion =============================================================================

func (NtWhile *NT_While) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//Se agrega un ambito
	ctx.PushAmbito()

	ctx.GenComentario("While >>>")
	Linicio := ctx.NewEtq() //Se crea una nueva etiqueta para el inicio
	Lsalida := ctx.NewEtq() //Se crea una nueva etiqueta para la salida

	//se agrega el display de transiciones para un break o continue
	ctx.PushDisplayTrans(Lsalida, Linicio)

	ctx.GenLabel(Linicio + ":") //Se imprime la etiqueta de inicio

	//Se genera codigo de la condicion
	ctx.GenComentario("codigo Condicion ------------------------------------")
	condicion := NtWhile.Condicion.Compilar(ctx)

	//Se imprime la etiqueta verdadera
	ctx.GenComentario("EV Cond ------------------------------------")
	ctx.ImprimirEtq(condicion.EV)

	//Se genera codigo de las sentencias
	ctx.GenComentario("codigo sentencias ------------------------------------")
	NtWhile.Sentencias.Compilar(ctx)

	//se regresa al inicio
	ctx.Gen("goto " + Linicio)

	//Se imprime la etiqueta falsa
	ctx.GenComentario("EF Cond ------------------------------------")
	ctx.ImprimirEtq(condicion.EF)
	ctx.GenLabel(Lsalida + ":") //Se imprime la etiqueta de salida

	//se hace pop del display de transiciones
	ctx.PopDisplayTrans()

	//Se elimina el ambito
	ctx.PopAmbito()

	return compilador.NewNill()
}
