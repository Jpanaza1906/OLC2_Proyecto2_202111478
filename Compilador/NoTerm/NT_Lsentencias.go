package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_LSentencias struct {
	LSentencias []compilador.CAbstractExpr
}

// Constructor ================================================================================
func NewNT_LSentencias() *NT_LSentencias {
	return &NT_LSentencias{
		LSentencias: make([]compilador.CAbstractExpr, 0),
	}
}

func (NtLS *NT_LSentencias) AddSentencia(sentencia compilador.CAbstractExpr) {
	NtLS.LSentencias = append(NtLS.LSentencias, sentencia)
}

// Implementacion =============================================================================
func (NtLS *NT_LSentencias) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	for _, sentencia := range NtLS.LSentencias {
		sentencia.Compilar(ctx)
	}
	return compilador.NewNill()
}
