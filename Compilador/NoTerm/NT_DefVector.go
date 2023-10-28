package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	terminales "OLC2_Proyecto2_202111478/Compilador/Terminales"
)

type NT_DefVector struct {
	Id       string
	ListaExp []compilador.CAbstractExpr
	Linea    int
	Columna  int
}

// Constructor ================================================================================

func NewNT_DefVector(id string, listaexp []compilador.CAbstractExpr, linea int, columna int) *NT_DefVector {
	return &NT_DefVector{
		Id:       id,
		ListaExp: listaexp,
		Linea:    linea,
		Columna:  columna,
	}
}

// Implementacion =============================================================================

func (NtDefVector *NT_DefVector) Compilar(ctx *compilador.Contexto) *compilador.Atributos {

	if NtDefVector.Id != "" {
		//se verifica que la variable exista
		simbolo := ctx.GetSimbolo(NtDefVector.Id)
		if simbolo.Tipo == 0 {
			ctx.AddErrorLine("Semantico", "La variable "+NtDefVector.Id+" no existe", NtDefVector.Linea, NtDefVector.Columna)
			return compilador.NewNill()
		}

		//se verifica que sea un vector
		if simbolo.Tipo != compilador.Vector {
			ctx.AddErrorLine("Semantico", "La variable "+NtDefVector.Id+" no es un vector", NtDefVector.Linea, NtDefVector.Columna)
			return compilador.NewNill()
		}

		//se verifica que el vector sea mutable
		if !simbolo.Mutable {
			ctx.AddErrorLine("Semantico", "El vector "+NtDefVector.Id+" no es mutable", NtDefVector.Linea, NtDefVector.Columna)
			return compilador.NewNill()
		}

		//se devuelve el id
		idcom := terminales.NewT_Identificador(NtDefVector.Id, NtDefVector.Linea, NtDefVector.Columna)
		return idcom.Compilar(ctx)
	} else {
		//recorremos la lista de expresiones
		tipoant := ""
		vecdir := []string{}
		for _, exp := range NtDefVector.ListaExp {
			expresion := exp.Compilar(ctx)

			//si el tipo de la expresion es diferente al tipo anterior
			if tipoant != "" && tipoant != expresion.Tipo.String() {
				ctx.AddErrorLine("Semantico", "Los tipos de las expresiones no coinciden", NtDefVector.Linea, NtDefVector.Columna)
				return compilador.NewNill()
			}
			tipoant = expresion.Tipo.String()

			//se agrega al vector de direcciones
			vecdir = append(vecdir, expresion.Dir)
		}

		t1 := ctx.NewTemp()
		t2 := ""
		//se guarda la posicion donde inicia el vector
		ctx.Gen(t1 + " = H")
		//variable contador
		cont := 0
		//se recorre el vector de direcciones
		for _, dir := range vecdir {
			if cont == 0 {
				t2 = ctx.NewTemp()
			}
			//se genera el codigo 3d para asignar los valores en el heap
			ctx.Gen("heap[(int) H] = " + dir)
			ctx.Gen("H = H + 1")
			ctx.Gen(t2 + " = H + 1")
			ctx.Gen("heap[(int) H] = " + t2)
			ctx.Gen("H = H + 1")
			ctx.GenNewLine()
			cont++
		}

		//se agrega el valor nulo al final del vector
		ctx.Gen("heap[(int) H] = 0")
		ctx.Gen("H = H + 1")
		ctx.Gen("heap[(int) H] = -1")
		ctx.Gen("H = H + 1")

		return compilador.NewVector(t1)

	}
}
