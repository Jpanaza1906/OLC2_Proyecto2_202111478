package compilador

type Tsimbolos struct {
	Id         string
	TipoId     string
	Tipo       TipoE
	Ambiente   int
	Size       int
	Valores    []Rangos
	Referencia bool
	Mutable    bool
}

type Rangos struct {
	Inicio int
	Fin    int
}
