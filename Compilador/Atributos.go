package compilador

type TipoE uint

const (
	// Tipos de expresiones
	Nil     TipoE = 0
	Bool    TipoE = 1
	Integer TipoE = 2
	Float   TipoE = 3
	Char    TipoE = 4
	String  TipoE = 5
	Vector  TipoE = 6
	Matriz  TipoE = 7
	StructT TipoE = 8
)

func (t TipoE) String() string {
	switch t {
	case Nil:
		return "Nil"
	case Bool:
		return "Bool"
	case Integer:
		return "Int"
	case Float:
		return "Float"
	case String:
		return "String"
	case Char:
		return "Char"
	case Vector:
		return "Vector"
	case Matriz:
		return "Matriz"
	case StructT:
		return "Struct"
	default:
		return "Unknown"
	}
}

type Atributos struct {
	EV  []string
	EF  []string
	Dir string
	Op  string
	// (tipos)
	Tipo TipoE
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

func NewAtributo(EV []string, EF []string, dir string, op string, tipo TipoE) *Atributos {
	return &Atributos{
		EV:   EV,
		EF:   EF,
		Dir:  dir,
		Op:   op,
		Tipo: tipo,
	}
}

//Atributos con tipo

func NewNill() *Atributos {
	return &Atributos{
		EV:   []string{},
		EF:   []string{},
		Dir:  "-202111478",
		Op:   "",
		Tipo: Nil,
	}
}

func NewBool(valor string) *Atributos {
	return &Atributos{
		EV:   []string{},
		EF:   []string{},
		Dir:  valor,
		Op:   "",
		Tipo: Bool,
	}
}

func NewInt(valor string) *Atributos {
	return &Atributos{
		EV:   []string{},
		EF:   []string{},
		Dir:  valor,
		Op:   "",
		Tipo: Integer,
	}
}

func NewFloat(valor string) *Atributos {
	return &Atributos{
		EV:   []string{},
		EF:   []string{},
		Dir:  valor,
		Op:   "",
		Tipo: Float,
	}
}

func NewString(valor string) *Atributos {
	return &Atributos{
		EV:   []string{},
		EF:   []string{},
		Dir:  valor,
		Op:   "",
		Tipo: String,
	}
}

func NewChar(valor string) *Atributos {
	return &Atributos{
		EV:   []string{},
		EF:   []string{},
		Dir:  valor,
		Op:   "",
		Tipo: Char,
	}
}

func NewVector(dir string) *Atributos {
	return &Atributos{
		EV:   []string{},
		EF:   []string{},
		Dir:  dir,
		Op:   "",
		Tipo: Vector,
	}
}

func NewMatriz(dir string) *Atributos {
	return &Atributos{
		EV:   []string{},
		EF:   []string{},
		Dir:  dir,
		Op:   "",
		Tipo: Matriz,
	}
}

func NewStruct() *Atributos {
	return &Atributos{
		EV:   []string{},
		EF:   []string{},
		Dir:  "",
		Op:   "",
		Tipo: StructT,
	}
}
