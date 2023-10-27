package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
)

// NT_Print =============================================================================

type NT_Print struct {
	Expresiones []compilador.CAbstractExpr
	Linea       int
	Columna     int
}

// Constructor ================================================================================

func NewNT_Print(expresiones []compilador.CAbstractExpr, linea int, columna int) *NT_Print {
	return &NT_Print{
		Expresiones: expresiones,
		Linea:       linea,
		Columna:     columna,
	}
}

// Implementacion =============================================================================

func (NtPrint *NT_Print) Compilar(ctx *compilador.Contexto) *compilador.Atributos {

	for _, exp := range NtPrint.Expresiones {
		result := exp.Compilar(ctx)

		switch result.Tipo {
		case compilador.Integer:
			//si trae el valor nulo de -202111478 se imprime nil
			l1 := ctx.NewEtq()
			l2 := ctx.NewEtq()
			ctx.Gen("if( (int)" + result.Dir + " == -202111478) goto " + l1)
			//imprimir nil caracter por caracter
			ctx.Gen("printf(\"%d\", (int) " + result.Dir + ")")
			ctx.Gen("goto " + l2)
			ctx.GenLabel(l1 + ":")
			ctx.Gen("printf(\"%c\", (char) 110)")
			ctx.Gen("printf(\"%c\", (char) 105)")
			ctx.Gen("printf(\"%c\", (char) 108)")
			ctx.GenLabel(l2 + ":")
			ctx.Gen("printf(\"%c\", (char) 32)")
			ctx.Gen("printf(\"%c\", (char) 10)")
		case compilador.Float:
			l1 := ctx.NewEtq()
			l2 := ctx.NewEtq()
			ctx.Gen("if( (float)" + result.Dir + " == -202111478.0) goto " + l1)
			ctx.Gen("printf(\"%f\", (float) " + result.Dir + ")")
			ctx.Gen("goto " + l2)
			ctx.GenLabel(l1 + ":")
			ctx.Gen("printf(\"%c\", (char) 110)")
			ctx.Gen("printf(\"%c\", (char) 105)")
			ctx.Gen("printf(\"%c\", (char) 108)")
			ctx.GenLabel(l2 + ":")
			ctx.Gen("printf(\"%c\", (char) 32)")
			ctx.Gen("printf(\"%c\", (char) 10)")
		case compilador.String:
			l1 := ctx.NewEtq()
			l2 := ctx.NewEtq()
			ctx.Gen("if( (int)" + result.Dir + " == -202111478) goto " + l1)
			ctx.Nat_PrintString(result.Dir)
			ctx.Gen("goto " + l2)
			ctx.GenLabel(l1 + ":")
			ctx.Gen("printf(\"%c\", (char) 110)")
			ctx.Gen("printf(\"%c\", (char) 105)")
			ctx.Gen("printf(\"%c\", (char) 108)")
			ctx.GenLabel(l2 + ":")
			ctx.Gen("printf(\"%c\", (char) 32)")
			ctx.Gen("printf(\"%c\", (char) 10)")
		case compilador.Bool:
			l1 := ctx.NewEtq()
			l2 := ctx.NewEtq()
			ctx.Gen("if( (int)" + result.Dir + " == -202111478) goto " + l1)
			ctx.Nat_PrintBool(result.Dir)
			ctx.Gen("goto " + l2)
			ctx.GenLabel(l1 + ":")
			ctx.Gen("printf(\"%c\", (char) 110)")
			ctx.Gen("printf(\"%c\", (char) 105)")
			ctx.Gen("printf(\"%c\", (char) 108)")
			ctx.GenLabel(l2 + ":")
			ctx.Gen("printf(\"%c\", (char) 32)")
			ctx.Gen("printf(\"%c\", (char) 10)")
		case compilador.Char:
			l1 := ctx.NewEtq()
			l2 := ctx.NewEtq()
			ctx.Gen("if( (int)" + result.Dir + " == -202111478) goto " + l1)
			ctx.Gen("printf(\"%c\", (char) " + result.Dir + ")")
			ctx.Gen("goto " + l2)
			ctx.GenLabel(l1 + ":")
			ctx.Gen("printf(\"%c\", (char) 110)")
			ctx.Gen("printf(\"%c\", (char) 105)")
			ctx.Gen("printf(\"%c\", (char) 108)")
			ctx.GenLabel(l2 + ":")
			ctx.Gen("printf(\"%c\", (char) 32)")
			ctx.Gen("printf(\"%c\", (char) 10)")
		case compilador.Nil:
			ctx.Gen("printf(\"%c\", (char) 110)")
			ctx.Gen("printf(\"%c\", (char) 105)")
			ctx.Gen("printf(\"%c\", (char) 108)")
			ctx.Gen("printf(\"%c\", (char) 32)")
			ctx.Gen("printf(\"%c\", (char) 10)")

		}
	}

	return compilador.NewNill()
}
