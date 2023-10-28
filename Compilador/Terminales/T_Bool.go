package terminales

import compilador "OLC2_Proyecto2_202111478/Compilador"

type T_Bool struct {
	Valor string
}

// Constructor ================================================================================\
func NewT_Bool(valor string) *T_Bool {
	return &T_Bool{
		Valor: valor,
	}
}

// Implementacion =============================================================================

func (Tb *T_Bool) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	if Tb.Valor == "true" {
		return compilador.NewBool("1")
	} else if Tb.Valor == "false" {
		return compilador.NewBool("0")
	}
	return compilador.NewNill()
}
