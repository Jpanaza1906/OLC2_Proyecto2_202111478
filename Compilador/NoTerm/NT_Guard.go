package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
)

type NT_Guard struct {
	Condicion  compilador.CAbstractExpr
	Sentencias compilador.CAbstractExpr
	Linea      int
	Columna    int
}

// Constructor ================================================================================

func NewNT_Guard(condicion compilador.CAbstractExpr, sentencias compilador.CAbstractExpr, linea int, columna int) *NT_Guard {
	return &NT_Guard{
		Condicion:  condicion,
		Sentencias: sentencias,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion =============================================================================

func (NtGuard *NT_Guard) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//Etiquetas
	Lsalida := ctx.NewEtq()

	ctx.Gen("// Guard >>>")
	ctx.Gen("// codigo Condicion ------------------------------------")
	condicion := NtGuard.Condicion.Compilar(ctx) //Se genera codigo de la condicion

	ctx.Gen("// EV Cond ------------------------------------")
	ctx.ImprimirEtq(condicion.EV) //Se imprime la etiqueta verdadera
	ctx.Gen("goto " + Lsalida)    //Se genera el goto para la salida

	ctx.Gen("// EF Cond ------------------------------------")
	ctx.ImprimirEtq(condicion.EF) //Se imprime la etiqueta falsa

	ctx.Gen("// codigo sentencias ------------------------------------")

	NtGuard.Sentencias.Compilar(ctx) //Se genera codigo de las sentencias

	ctx.Gen("goto " + Lsalida) //Se genera el goto para la salida

	ctx.Gen(Lsalida + ":") //Se imprime la etiqueta de salida

	return compilador.NewNill()

}
