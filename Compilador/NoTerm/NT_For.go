package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

// NT_For Integer===============================================================

type NT_ForInt struct {
	Id         string
	ExpIzq     compilador.CAbstractExpr
	ExpDer     compilador.CAbstractExpr
	Sentencias compilador.CAbstractExpr
	Linea      int
	Columna    int
}

// Constructor ================================================================================

func NewNT_ForInt(id string, expIzq compilador.CAbstractExpr, expDer compilador.CAbstractExpr, sentencias compilador.CAbstractExpr, linea int, columna int) *NT_ForInt {
	return &NT_ForInt{
		Id:         id,
		ExpIzq:     expIzq,
		ExpDer:     expDer,
		Sentencias: sentencias,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion =============================================================================

func (NtForInt *NT_ForInt) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//Etiqueta de inicio
	ctx.PushAmbito()
	ctx.GenComentario("For >>>")
	Linicio := ctx.NewEtq()
	Lsalida := ctx.NewEtq()

	//se agrega el display de transiciones para un break o continue
	ctx.PushDisplayTrans(Lsalida, Linicio, "for", NtForInt.Id)

	//Se genera codigo de la asignacion
	ctx.GenComentario("codigo Asignacion ------------------------------------")
	//se declara la variable en el ambito
	//se convierte a un nt_declaracion
	ntDeclaracion := NewNT_Declaracion(NtForInt.Id, compilador.Integer.String(), true, NtForInt.ExpIzq, NtForInt.Linea, NtForInt.Columna)
	ntDeclaracion.Compilar(ctx)

	//segunda expresion
	ctx.GenComentario("codigo segunda expresion ------------------------------------")
	segundaExp := NtForInt.ExpDer.Compilar(ctx)

	//se comienza a generar el codigo de la condicion
	ctx.GenComentario("codigo Condicion ------------------------------------")
	ctx.GenLabel(Linicio + ":") //Se imprime la etiqueta de inicio

	//se genera el codigo de la condicion

	//se obtiene el valor de la variable
	simbolo := ctx.GetAmbitoSimbolo(NtForInt.Id)
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(simbolo.Size))
	t2 := ctx.NewTemp()
	ctx.Gen(t2 + " = stack[(int) " + t1 + "]")
	//se genera el codigo para la condicion
	ctx.Gen("if (" + t2 + " > " + segundaExp.Dir + ") goto " + Lsalida)
	ctx.GenComentario("codigo sentencias ------------------------------------")
	NtForInt.Sentencias.Compilar(ctx)

	//se genera el codigo para aumentar la variable
	ctx.GenComentario("codigo aumento ------------------------------------")
	t3 := ctx.NewTemp()
	ctx.Gen(t3 + " = " + t2 + " + 1")
	ctx.Gen("stack[(int) " + t1 + "] = " + t3)

	//se regresa al inicio
	ctx.Gen("goto " + Linicio)
	//Se imprime la etiqueta falsa
	ctx.GenComentario("EF Cond ------------------------------------")
	ctx.GenLabel(Lsalida + ":")

	ctx.PopDisplayTrans()
	ctx.PopAmbito()

	return compilador.NewNill()
}

// NT_ForList================================================================

type NT_ForList struct {
	Id         string
	Expresion  compilador.CAbstractExpr
	Sentencias compilador.CAbstractExpr
	Linea      int
	Columna    int
}

// Constructor ================================================================================

func NewNT_ForList(id string, expresion compilador.CAbstractExpr, sentencias compilador.CAbstractExpr, linea int, columna int) *NT_ForList {
	return &NT_ForList{
		Id:         id,
		Expresion:  expresion,
		Sentencias: sentencias,
		Linea:      linea,
		Columna:    columna,
	}
}

// Implementacion =============================================================================

func (NtForList *NT_ForList) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	ctx.PushAmbito()
	ctx.GenComentario("For >>>")
	Linicio := ctx.NewEtq()
	Lsalida := ctx.NewEtq()

	//se agrega el display de transiciones para un break o continue
	ctx.PushDisplayTrans(Lsalida, Linicio, "for", NtForList.Id)

	//se compila la expresion
	ctx.GenComentario("codigo expresion ------------------------------------")
	expresion := NtForList.Expresion.Compilar(ctx)

	//si es de tipo string
	if expresion.Tipo == compilador.String {
		//H = 0
		t11 := ctx.NewTemp()
		ctx.Gen(t11 + " = H")
		ctx.Gen("H = H + 1") //t11 = 1
		ctx.Gen("heap[(int) H] = 0")
		ctx.Gen("H = H + 1") //t11 = 2
		ctx.GenLabel(Linicio + ":")

		t1 := ctx.NewTemp()
		ctx.Gen(t1 + " = heap[(int) " + expresion.Dir + "]")

		//se declara en el ambito
		//ctx.GenComentario("Declaracion variable")
		//se asigna la variable en el heap
		ctx.Gen("heap[(int) " + t11 + "] = " + t1)

		t2 := ctx.NewTemp()
		ctx.Gen(t2 + " = P + " + strconv.Itoa(ctx.PosSt)) //la tabla de simbolos me dira el desplazamiento
		ctx.Gen("stack[(int) " + t2 + "] = " + t11)

		//se agrega el simbolo a la tabla de simbolos
		ctx.AddSimbolo(NtForList.Id, "variable", compilador.String, ctx.PeekAmbito(), ctx.PosSt, nil, false, true)
		ctx.PosSt++

		ctx.GenComentario("codigo Condicion ------------------------------------")
		ctx.Gen("if (" + t1 + " == 0) goto " + Lsalida)

		//aumento de la variable
		ctx.GenComentario("codigo aumento ------------------------------------")
		ctx.Gen(expresion.Dir + " = " + expresion.Dir + " + 1")

		ctx.GenComentario("codigo sentencias ------------------------------------")
		NtForList.Sentencias.Compilar(ctx)

		ctx.Gen("goto " + Linicio)

		ctx.GenComentario("EF Cond ------------------------------------")
		ctx.GenLabel(Lsalida + ":")

	} else if expresion.Tipo == compilador.Vector {
		ctx.PosSt++
	}

	ctx.PopDisplayTrans()
	ctx.PopAmbito()

	return compilador.NewNill()
}
