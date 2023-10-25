package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_If struct {
	Condicion      compilador.CAbstractExpr
	Sentencias     compilador.CAbstractExpr
	SentenciasElse compilador.CAbstractExpr
	Linea          int
	Columna        int
}

// Constructor ================================================================================

func NewNT_If(condicion compilador.CAbstractExpr, sentencias compilador.CAbstractExpr, sentenciasElse compilador.CAbstractExpr, linea int, columna int) *NT_If {
	return &NT_If{
		Condicion:      condicion,
		Sentencias:     sentencias,
		SentenciasElse: sentenciasElse,
		Linea:          linea,
		Columna:        columna,
	}
}

// Implementacion =============================================================================

func (NtIf *NT_If) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	ctx.GenComentario("If >>>")
	ctx.GenComentario("codigo Condicion ------------------------------------")
	atributos := NtIf.Condicion.Compilar(ctx) //Se genera codigo de la condicion

	Lsalida := ctx.NewEtq() //Se crea una nueva etiqueta para la salida
	ctx.GenComentario("EV Cond ------------------------------------")
	ctx.ImprimirEtq(atributos.EV) //Se imprime la etiqueta verdadera

	//Se agrega un ambito
	ctx.PushAmbito()

	ctx.GenComentario("codigo sentencias ------------------------------------")
	NtIf.Sentencias.Compilar(ctx) //Se genera codigo de las sentencias

	//Se elimina el ambito
	ctx.PopAmbito()

	ctx.Gen("goto " + Lsalida) //Se genera el goto para la salida

	ctx.GenComentario("EF Cond ------------------------------------")
	ctx.ImprimirEtq(atributos.EF) //Se imprime la etiqueta falsa

	if NtIf.SentenciasElse != nil {
		//Se agrega un ambito
		ctx.PushAmbito()

		ctx.GenComentario("codigo sentencias else ------------------------------------")
		NtIf.SentenciasElse.Compilar(ctx) //Se genera codigo de las sentencias else

		//Se elimina el ambito
		ctx.PopAmbito()
	}
	ctx.GenLabel(Lsalida + ":") //Se imprime la etiqueta de salida

	//Se elimina el ambito
	ctx.PopAmbito()

	return compilador.NewNill()

}
