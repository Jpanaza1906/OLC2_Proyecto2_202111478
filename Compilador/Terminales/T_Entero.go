package terminales

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

type T_Num struct {
	Num string
}

// Constructor ================================================================================

func NewT_Num(num string) *T_Num {
	return &T_Num{
		Num: num,
	}
}

// Implementacion =============================================================================

func (Tn *T_Num) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	_, err := strconv.Atoi(Tn.Num)
	if err != nil {
		ctx.AddError("Error Semantico: No se pudo convertir el numero")
		return compilador.NewNill()
	}
	return compilador.NewInt(Tn.Num)
}
