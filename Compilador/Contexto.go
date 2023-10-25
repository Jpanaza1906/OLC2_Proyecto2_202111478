package compilador

import (
	"fmt"
	"strconv"
)

type Contexto struct {
	//temporales y labels
	tmp int
	lb  int
	//reportes y errores
	Consola   string
	Errores   string
	Simbolos  string
	Funciones string
	//displays
	DisplaySwitch []string
	DisplayTrans  []DisplayTrans
	PtrTrans      int
	//tabla de simbolos
	TablaSimbolos []Tsimbolos
	PosSt         int
	//ambito
	Ambito  int
	tmpList []string
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
		DisplayTrans:  make([]DisplayTrans, 15),
		PtrTrans:      0,
		TablaSimbolos: make([]Tsimbolos, 0),
		PosSt:         0,
		Ambito:        0,
	}
}

// Simbolos--------------------------------------------------------------------------------------

func (ctx *Contexto) Gen(out string) {
	ctx.Consola += "\t" + out + ";\n"
	fmt.Println(out)
}
func (ctx *Contexto) GenLabel(out string) {
	ctx.Consola += "\t" + out + "\n"
	fmt.Println(out)
}
func (ctx *Contexto) GenComentario(out string) {
	ctx.Consola += "\t//" + out + "\n"
	fmt.Println(out)
}

func (ctx *Contexto) NewTemp() string {
	ctx.tmp++
	//agregar el temporal a la lista de temporales
	ctx.tmpList = append(ctx.tmpList, "t"+strconv.Itoa(ctx.tmp))
	return "t" + strconv.Itoa(ctx.tmp)
}

func (ctx *Contexto) GetTemplist() []string {
	return ctx.tmpList
}

func (ctx *Contexto) NewEtq() string {
	ctx.lb++
	return "L" + strconv.Itoa(ctx.lb)
}

func (ctx *Contexto) ImprimirEtq(etiquetas []string) {
	for _, e := range etiquetas {
		ctx.GenLabel(e + ":")
	}
}

func (ctx *Contexto) Unir(etq1 []string, etq2 []string) []string {
	etiquetas := append(etq1, etq2...)
	return etiquetas
}

//Errores---------------------------------------------------------------------------------------

func (ctx *Contexto) AddErrorLine(tipo string, err string, linea int, columna int) {
	errorString := fmt.Sprintf("Error %s (%d,%d) : %s", tipo, linea, columna, err)
	ctx.Errores += errorString + "\n"
}

func (ctx *Contexto) AddError(err string) {
	ctx.Errores += err + "\n"
}

// DisplaySwitch---------------------------------------------------------------------------------
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
		return ""
	}
	return etiqueta
}

// DisplayTrans----------------------------------------------------------------------------------
func (ctx *Contexto) PushDisplayTrans(lbreak string, lwhile string) {
	ctx.DisplayTrans[ctx.PtrTrans] = DisplayTrans{
		Break:    lbreak,
		Continue: lwhile,
	}
	ctx.PtrTrans++
}

func (ctx *Contexto) PeekContinue() string {
	if ctx.PtrTrans > 0 {
		labelcontinue := ctx.DisplayTrans[ctx.PtrTrans-1].Continue
		return labelcontinue
	} else {
		ctx.AddError("No se puede utilizar continue fuera de un ciclo")
		return ""
	}
}

func (ctx *Contexto) PeekBreak() string {
	if ctx.PtrTrans > 0 {
		labelbreak := ctx.DisplayTrans[ctx.PtrTrans-1].Break
		return labelbreak
	} else {
		ctx.AddError("No se puede utilizar break fuera de un ciclo")
		return ""
	}
}

func (ctx *Contexto) PopDisplayTrans() {
	if ctx.PtrTrans > 0 {
		ctx.PtrTrans--
	} else {
		ctx.AddError("Error al hacer pop en el DisplayTrans vacio")
	}
}

// TablaSimbolos---------------------------------------------------------------------------------

func (ctx *Contexto) AddSimbolo(id string, tipoId string, tipo TipoE, ambiente int, size int, valores []Rangos, referencia bool, mutable bool) {
	simbolo := Tsimbolos{
		Id:         id,
		TipoId:     tipoId,
		Tipo:       tipo,
		Ambiente:   ambiente,
		Size:       size,
		Valores:    valores,
		Referencia: referencia,
		Mutable:    mutable,
	}
	ctx.TablaSimbolos = append(ctx.TablaSimbolos, simbolo)
}

func (ctx *Contexto) GetAmbitoSimbolo(id string) Tsimbolos {
	for i := len(ctx.TablaSimbolos) - 1; i >= 0; i-- {
		if ctx.TablaSimbolos[i].Id == id {
			if ctx.TablaSimbolos[i].Ambiente == ctx.Ambito {
				return ctx.TablaSimbolos[i]
			}
		}
	}

	return Tsimbolos{}
}

func (ctx *Contexto) GetSimbolo(id string) Tsimbolos {
	//se empieza por el ultimo ambito
	for j := ctx.Ambito; j >= 0; j-- {
		for i := len(ctx.TablaSimbolos) - 1; i >= 0; i-- {
			if ctx.TablaSimbolos[i].Id == id {
				//si el ambiente es el mismo
				if ctx.TablaSimbolos[i].Ambiente == j {
					return ctx.TablaSimbolos[i]
				}
			}
		}
	}
	return Tsimbolos{}
}

func (ctx *Contexto) RemoveSimbolosAmbito(ambito int) {
	//quitar cada ambito que tenga el mismo numero de ambito
	for i := len(ctx.TablaSimbolos) - 1; i >= 0; i-- {
		if ctx.TablaSimbolos[i].Ambiente == ambito {
			ctx.TablaSimbolos = append(ctx.TablaSimbolos[:i], ctx.TablaSimbolos[i+1:]...)
		}
	}
}

func (ctx *Contexto) GetSimbolosCadena() {
	ctx.Simbolos = "--------------------------TABLA DE SIMBOLOS-------------------------------\n"
	ctx.Simbolos += "Id | TipoId | Tipo | Ambiente | Size | Referencia | Mutable\n"
	ctx.Simbolos += "------------------------------------------------------------------------------------------\n"
	for _, s := range ctx.TablaSimbolos {
		ctx.Simbolos += fmt.Sprintf("%s | %s | %s | %d | %d | %t | %t\n", s.Id, s.TipoId, s.Tipo, s.Ambiente, s.Size, s.Referencia, s.Mutable)
	}
}

func (ctx *Contexto) GetTipoE(tipo string) TipoE {
	switch tipo {
	case "Int":
		return Integer
	case "Float":
		return Float
	case "Char":
		return Char
	case "String":
		return String
	case "Bool":
		return Bool
	}
	return Nil
}

// Ambito----------------------------------------------------------------------------------------

func (ctx *Contexto) PushAmbito() {
	ctx.Ambito++
}

func (ctx *Contexto) PeekAmbito() int {
	if ctx.Ambito > 0 {
		return ctx.Ambito
	} else {
		ctx.AddError("Error al hacer peek en el Ambito vacio")
		return 0
	}
}

func (ctx *Contexto) PopAmbito() {
	if ctx.Ambito > 0 {
		//se eliminan los simbolos del ambito
		ctx.RemoveSimbolosAmbito(ctx.Ambito)
		ctx.Ambito--
	} else {
		ctx.AddError("Error al hacer pop en el Ambito vacio")
	}
}
