package compilador

type Tsimbolos struct {
	Id         string
	TipoId     string
	Tipo       TipoE
	Ambiente   int
	Size       int //me da posicion relativa en la pila
	Valores    []int
	Referencia bool
	Mutable    bool
}

func (ts *Tsimbolos) CorregirTam() {
	for i := 1; i < len(ts.Valores); i++ {
		for j := i + 1; j < len(ts.Valores); j++ {
			ts.Valores[i] = ts.Valores[i] * ts.Valores[j]
		}
	}
}
