package compilador

import (
	"fmt"
	"strconv"
)

type Contexto struct {
	tmp           int
	lb            int
	Consola       string
	Errores       string
	Simbolos      string
	Funciones     string
	DisplaySwitch []string
}

func NewContexto() *Contexto {
	return &Contexto{
		tmp:           0,
		lb:            0,
		Consola:       "",
		Errores:       "",
		Simbolos:      "",
		Funciones:     "",
		DisplaySwitch: make([]string, 0),
	}
}

func (ctx *Contexto) Gen(out string) {
	ctx.Consola += out + "\n"
	fmt.Println(out)
}

func (ctx *Contexto) NewTemp() string {
	ctx.tmp++
	return "t" + strconv.Itoa(ctx.tmp)
}

func (ctx *Contexto) NewEtq() string {
	ctx.lb++
	return "L" + strconv.Itoa(ctx.lb)
}

func (ctx *Contexto) ImprimirEtq(etiquetas []string) {
	for _, e := range etiquetas {
		ctx.Gen(e + ":")
	}
}

func (ctx *Contexto) Unir(etq1 []string, etq2 []string) []string {
	etiquetas := append(etq1, etq2...)
	return etiquetas
}

//Errores

func (ctx *Contexto) AddError(err string) {
	ctx.Errores += err + "\n"
}

// DisplaySwitch
func (ctx *Contexto) PushDisplaySwitch() string {
	nuevaEtq := ctx.NewEtq()
	ctx.DisplaySwitch = append(ctx.DisplaySwitch, nuevaEtq)
	return nuevaEtq
}

func (ctx *Contexto) PeekDisplaySwitch() string {
	if len(ctx.DisplaySwitch) > 0 {
		return ctx.DisplaySwitch[len(ctx.DisplaySwitch)-1]
	} else {
		ctx.AddError("Error al hacer peek en el DisplaySwitch vacio")
		return ""
	}
}

func (ctx *Contexto) PopDisplaySwitch() string {
	etiqueta := ctx.PeekDisplaySwitch()
	if len(ctx.DisplaySwitch) > 0 {
		ctx.DisplaySwitch = ctx.DisplaySwitch[:len(ctx.DisplaySwitch)-1]
	} else {
		ctx.AddError("Error al hacer pop en el DisplaySwitch vacio")
	}
	return etiqueta
}
