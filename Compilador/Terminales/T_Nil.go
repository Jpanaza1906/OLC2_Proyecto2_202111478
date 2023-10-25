package terminales

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
)

type T_Nil struct {
}

// Constructor ================================================================================
func NewT_Nil() *T_Nil {
	return &T_Nil{}
}

// Implementacion =============================================================================
func (Tn *T_Nil) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	return compilador.NewNill()
}
