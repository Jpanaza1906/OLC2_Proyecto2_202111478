package compilador

type Tsimbolos struct {
	Id         string
	TipoId     string
	Tipo       TipoE
	Ambiente   int
	Size       int //me da posicion relativa en la pila
	Valores    []Rangos
	Referencia bool
	Mutable    bool
}

type Rangos struct {
	Inicio int
	Fin    int
}
