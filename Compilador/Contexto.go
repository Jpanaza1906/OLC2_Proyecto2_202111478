package compilador

import (
	"fmt"
	"strconv"
)

type Contexto struct {
	tmp     int
	lb      int
	Errores []string
}

func NewContexto() *Contexto {
	return &Contexto{
		tmp:     0,
		lb:      0,
		Errores: make([]string, 0),
	}
}

func (ctx *Contexto) Gen(out string) {
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

func (ctx *Contexto) AddError(err string) {
	ctx.Errores = append(ctx.Errores, err)
}
