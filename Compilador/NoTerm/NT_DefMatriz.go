package noterm

import compilador "OLC2_Proyecto2_202111478/Compilador"

type NT_DefMatriz struct {
	ListaExp []compilador.CAbstractExpr
	Linea    int
	Columna  int
}

// Constructor ================================================================================

func NewNT_DefMatriz(listaexp []compilador.CAbstractExpr, linea int, columna int) *NT_DefMatriz {
	return &NT_DefMatriz{
		ListaExp: listaexp,
		Linea:    linea,
		Columna:  columna,
	}
}

// Implementacion =============================================================================

func (NtDefMatriz *NT_DefMatriz) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	vecdir := []string{}
	for _, exp := range NtDefMatriz.ListaExp {
		expresion := exp.Compilar(ctx)
		vecdir = append(vecdir, expresion.Dir)
	}

	for _, dir := range vecdir {
		//si el nil se omite
		if dir != "-202111478" {
			ctx.Gen("heap[(int)H] = " + dir)
			ctx.Gen("H = H + 1")
		}
	}

	return compilador.NewNill()
}
