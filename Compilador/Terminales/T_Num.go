package terminales

import compilador "OLC2_Proyecto2_202111478/Compilador"

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
	return &compilador.Atributos{
		EV:  make([]string, 0),
		EF:  make([]string, 0),
		Dir: Tn.Num,
	}
}
