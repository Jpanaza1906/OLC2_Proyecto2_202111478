package terminales

import compilador "OLC2_Proyecto2_202111478/Compilador"

type T_BoolCond struct {
	Valor string
}

// Constructor ================================================================================

func NewT_BoolCond(valor string) *T_BoolCond {
	return &T_BoolCond{
		Valor: valor,
	}
}

// Implementacion =============================================================================

func (Tb *T_BoolCond) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	EV := ctx.NewEtq()
	EF := ctx.NewEtq()

	if Tb.Valor == "true" {
		ctx.Gen("goto " + EV)
	} else {
		ctx.Gen("goto " + EF)
	}

	//Se concatenand las etiquetas verdaders y falsas
	EVs := make([]string, 0)
	EVs = append(EVs, EV)

	EFs := make([]string, 0)
	EFs = append(EFs, EF)

	return compilador.NewAtributo(EVs, EFs, "", "", compilador.Bool)
}
