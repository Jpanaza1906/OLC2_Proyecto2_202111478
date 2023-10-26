package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_OpRelacionales struct {
	ExprIzq compilador.CAbstractExpr
	ExprDer compilador.CAbstractExpr
	Op      string
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_OpRelacionales(exprIzq compilador.CAbstractExpr, exprDer compilador.CAbstractExpr, op string, linea int, columna int) *NT_OpRelacionales {
	return &NT_OpRelacionales{
		ExprIzq: exprIzq,
		ExprDer: exprDer,
		Op:      op,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (Nopr *NT_OpRelacionales) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	exprIzq := Nopr.ExprIzq.Compilar(ctx)
	exprDer := Nopr.ExprDer.Compilar(ctx)

	//se verifica que los tipos sean iguales
	if exprIzq.Tipo != exprDer.Tipo {
		ctx.AddErrorLine("Semantico", "No se puede operar logicamente con tipos diferentes", Nopr.Linea, Nopr.Columna)
		return compilador.NewNill()
	}

	//si es de tipo string se compara caracter por caracter en el heap
	if exprIzq.Tipo == compilador.String {
		//se obtiene el valor de la cadena
		//exprIzq.Dir -> direccion del heap donde inicia la cadena
		//exprDer.Dir -> direccion del heap donde inicia la cadena

		//Si el operador es ==
		if Nopr.Op == "==" || Nopr.Op == "!=" {
			//funcion si son iguales devuelve 1, si no 0
			//se genera el codigo 3d para comparar las cadenas
			ctx.GenComentario("Funcion devuelve 1 si son iguales, 0 si no")
			t1 := ctx.NewTemp()
			ctx.Nat_EqualString(t1, exprIzq.Dir, exprDer.Dir) //se llama a la funcion nativa para comparar
			EV := ctx.NewEtq()
			EF := ctx.NewEtq()

			EVs := make([]string, 0)
			EFs := make([]string, 0)

			EVs = append(EVs, EV)
			EFs = append(EFs, EF)

			ctx.Gen("if ((int) " + t1 + " " + Nopr.Op + " 1)" + " goto " + EV)
			ctx.Gen("goto " + EF)

			return compilador.NewAtributo(EVs, EFs, "", "", compilador.Bool)

		}
		return compilador.NewNill()
	}

	EV := ctx.NewEtq()
	EF := ctx.NewEtq()

	EVs := make([]string, 0)
	EFs := make([]string, 0)

	EVs = append(EVs, EV)
	EFs = append(EFs, EF)

	ctx.Gen("if (" + exprIzq.Dir + " " + Nopr.Op + " " + exprDer.Dir + ") goto " + EV)
	ctx.Gen("goto " + EF)

	return compilador.NewAtributo(EVs, EFs, "", "", compilador.Bool)
}
