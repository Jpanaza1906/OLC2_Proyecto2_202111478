package terminales

import compilador "OLC2_Proyecto2_202111478/Compilador"

type T_Bool struct {
	Valor string
}

// Constructor ================================================================================

func NewT_Bool(valor string) *T_Bool {
	return &T_Bool{
		Valor: valor,
	}
}

// Implementacion =============================================================================

func (Tb *T_Bool) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
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

	return compilador.NewAtributo(EVs, EFs, "", "")
}
