package compilador

type Atributos struct {
	EV  []string
	EF  []string
	Dir string
	Op  string
	// (tipos)
}

func NewEvAtributo(EV []string) *Atributos {
	return &Atributos{
		EV:  EV,
		EF:  []string{},
		Dir: "",
		Op:  "",
	}
}

func NewEfAtributo(EF []string) *Atributos {
	return &Atributos{
		EV:  []string{},
		EF:  EF,
		Dir: "",
		Op:  "",
	}
}

func NewDirAtributo(dir string) *Atributos {
	return &Atributos{
		EV:  []string{},
		EF:  []string{},
		Dir: dir,
		Op:  "",
	}
}

func NewOpAtributo(op string) *Atributos {
	return &Atributos{
		EV:  []string{},
		EF:  []string{},
		Dir: "",
		Op:  op,
	}
}

func NewAtributo(EV []string, EF []string, dir string, op string) *Atributos {
	return &Atributos{
		EV:  EV,
		EF:  EF,
		Dir: dir,
		Op:  op,
	}
}
func NewNill() *Atributos {
	return &Atributos{
		EV:  []string{},
		EF:  []string{},
		Dir: "",
		Op:  "",
	}
}
