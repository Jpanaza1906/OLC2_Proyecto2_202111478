package terminales

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

type T_Char struct {
	Char string
}

// Constructor ================================================================================
func NewT_Char(char string) *T_Char {
	return &T_Char{
		Char: char,
	}
}

// Implementacion =============================================================================

func (Tc *T_Char) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	///se obtiene el ascii del caracter
	ascii := int(Tc.Char[0])
	return compilador.NewChar(strconv.Itoa(ascii))
}
