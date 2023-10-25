package terminales

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

type T_Cad struct {
	Cadena string
}

// Constructor ================================================================================

func NewT_Cad(cadena string) *T_Cad {
	return &T_Cad{
		Cadena: cadena,
	}
}

// Implementacion =============================================================================

func (Tc *T_Cad) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//recorrer la cadena para obtener el ascii de cada caracter
	//y agregarlo al heap

	//obtener el valor de la cadena
	cadena := Tc.Cadena

	//si viene vacia, se genera un string vacio
	if cadena == "" {
		return compilador.NewString("")
	}

	//se genera el codigo 3d para agregar el valor al heap
	t1 := ctx.NewTemp()

	ctx.Gen(t1 + " = H")
	//obtener el tamaño de la cadena
	for i := 0; i < len(cadena); i++ {
		//obtener el ascii del caracter
		ascii := int(cadena[i])
		ctx.Gen("heap[(int) H] = " + strconv.Itoa(ascii))
		ctx.Gen("H = H + 1")
	}
	//agregar el caracter nulo
	ctx.Gen("heap[(int) H] = 0")
	ctx.Gen("H = H + 1")

	return compilador.NewString(t1)
}
