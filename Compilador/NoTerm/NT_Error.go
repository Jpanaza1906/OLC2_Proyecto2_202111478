package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_Error struct {
	ErrorS string
}

// Constructor ================================================================================

func NewNT_Error(errorS string) *NT_Error {
	return &NT_Error{
		ErrorS: errorS,
	}
}

// Implementacion =============================================================================

func (NtError *NT_Error) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	ctx.AddError(NtError.ErrorS)
	return compilador.NewNill()
}
