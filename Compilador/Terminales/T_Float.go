package terminales

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

type T_Float struct {
	Num string
}

// Constructor ================================================================================
func NewT_Float(num string) *T_Float {
	return &T_Float{
		Num: num,
	}
}

// Implementacion =============================================================================
func (Tn *T_Float) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	num, err := strconv.ParseFloat(Tn.Num, 64)
	if err != nil {
		ctx.AddError("Error Semantico: No se pudo convertir el numero")
		return compilador.NewNill()
	}
	return compilador.NewFloat(num)
}
