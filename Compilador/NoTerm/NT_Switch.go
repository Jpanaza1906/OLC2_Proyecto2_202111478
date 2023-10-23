package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_Switch struct {
	Expr    compilador.CAbstractExpr
	Cases   []compilador.CAbstractExpr
	Default compilador.CAbstractExpr
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_Switch(expr compilador.CAbstractExpr, cases []compilador.CAbstractExpr, def compilador.CAbstractExpr, linea int, columna int) *NT_Switch {
	return &NT_Switch{
		Expr:    expr,
		Cases:   cases,
		Default: def,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (NtSwitch *NT_Switch) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	expr := NtSwitch.Expr.Compilar(ctx)
	Lprueba := ctx.NewEtq()
	ctx.Gen("// Switch>>>")
	ctx.Gen("goto " + Lprueba + " //Lprueba")

	ctx.PushDisplaySwitch() //Lsalida

	listaExpr := make([]compilador.CAbstractExpr, 0)
	listaEtq := make([]string, 0)

	for _, c := range NtSwitch.Cases {
		//Se castea la expresion
		listaExpr = append(listaExpr, c.(*NT_Case).Expr)

		//Se generan las etiquetas
		etq := ctx.NewEtq()
		ctx.Gen("// case ------------------------------------")
		ctx.Gen(etq + ": //etq case")
		listaEtq = append(listaEtq, etq)
		c.Compilar(ctx)
	}

	var Ln string
	// Se verifica si hay default
	if NtSwitch.Default != nil {
		Ln = ctx.NewEtq()
		ctx.Gen("// default ------------------------------------")
		NtSwitch.Default.Compilar(ctx)
		ctx.Gen(Ln + ": //Ln")
		ctx.Gen("goto " + ctx.PeekDisplaySwitch() + " //Lsalida")
	}

	// Se ven las pruebas
	ctx.Gen(Lprueba + ": //Lprueba")
	for i, e := range listaExpr {
		//Se debe manejar los tipos
		atributos := e.Compilar(ctx)
		ctx.Gen("If " + expr.Dir + " == " + atributos.Dir + " goto " + listaEtq[i])

	}
	if NtSwitch.Default != nil {
		ctx.Gen("goto " + Ln)
	}
	//se agrega la etiqueta y se saca
	ctx.Gen(ctx.PopDisplaySwitch() + ": //Lsalida")
	return compilador.NewNill()
}
