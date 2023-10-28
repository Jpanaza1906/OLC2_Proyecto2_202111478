package noterm

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"strconv"
)

// NT_Append ==================================================================================

type NT_Append struct {
	Id        string
	Expresion compilador.CAbstractExpr
	Linea     int
	Columna   int
}

// Constructor ================================================================================
func NewNT_Append(id string, expresion compilador.CAbstractExpr, linea int, columna int) *NT_Append {
	return &NT_Append{
		Id:        id,
		Expresion: expresion,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion =============================================================================
func (NtAppend *NT_Append) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//con el id se obtiene el simbolo

	idsimbolo := ctx.GetSimbolo(NtAppend.Id)

	//si el simbolo no existe
	if idsimbolo.Tipo == 0 {
		ctx.AddErrorLine("Semantico", "La variable "+NtAppend.Id+" no existe", NtAppend.Linea, NtAppend.Columna)
		return compilador.NewNill()
	}

	//si el simbolo no es un vector
	if idsimbolo.TipoId != compilador.Vector.String() {
		ctx.AddErrorLine("Semantico", "La variable "+NtAppend.Id+" no es un vector", NtAppend.Linea, NtAppend.Columna)
		return compilador.NewNill()
	}

	//si el simbolo no es mutable
	if !idsimbolo.Mutable {
		ctx.AddErrorLine("Semantico", "El vector "+NtAppend.Id+" no es mutable", NtAppend.Linea, NtAppend.Columna)
		return compilador.NewNill()
	}

	//se obtiene el tipo del vector y se compila la expresion
	exp := NtAppend.Expresion.Compilar(ctx)

	//se verifica que el tipo de la expresion sea el mismo que el del vector
	if exp.Tipo != idsimbolo.Tipo {
		ctx.AddErrorLine("Semantico", "El tipo de la expresion no es el mismo que el del vector", NtAppend.Linea, NtAppend.Columna)
		return compilador.NewNill()
	}

	//se genera el codigo 3d para asignar el valor a la variable
	ctx.GenComentario("Append")
	t1 := ctx.NewTemp()

	ctx.Gen(t1 + " = P + " + strconv.Itoa(idsimbolo.Size)) //la posicion del inicio del vector en el stack
	t2 := ctx.NewTemp()
	ctx.Gen(t2 + " = stack[(int) " + t1 + "]") //se obtiene la posicion del heap donde inicia el vector

	ctx.Nat_AppendVector(t2, exp.Dir) //se llama a la funcion nativa para agregar el valor al vector

	return compilador.NewNill()
}

// NT_REMOVELAST ==============================================================================
type NT_RemoveLast struct {
	Id      string
	Linea   int
	Columna int
}

// Constructor ================================================================================
func NewNT_RemoveLast(id string, linea int, columna int) *NT_RemoveLast {
	return &NT_RemoveLast{
		Id:      id,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================
func (NtRemoveLast *NT_RemoveLast) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	idsimbolo := ctx.GetSimbolo(NtRemoveLast.Id)

	//si el simbolo no existe
	if idsimbolo.Tipo == 0 {
		ctx.AddErrorLine("Semantico", "La variable "+NtRemoveLast.Id+" no existe", NtRemoveLast.Linea, NtRemoveLast.Columna)
		return compilador.NewNill()
	}
	//si el simbolo no es un vector
	if idsimbolo.TipoId != compilador.Vector.String() {
		ctx.AddErrorLine("Semantico", "La variable "+NtRemoveLast.Id+" no es un vector", NtRemoveLast.Linea, NtRemoveLast.Columna)
		return compilador.NewNill()
	}
	//si el simbolo no es mutable
	if !idsimbolo.Mutable {
		ctx.AddErrorLine("Semantico", "El vector "+NtRemoveLast.Id+" no es mutable", NtRemoveLast.Linea, NtRemoveLast.Columna)
		return compilador.NewNill()
	}
	//se genera el codigo 3d para quitar el ultimo elemento del vector

	ctx.GenComentario("RemoveLast")
	t1 := ctx.NewTemp()

	ctx.Gen(t1 + " = P + " + strconv.Itoa(idsimbolo.Size)) //la posicion del inicio del vector en el stack
	t2 := ctx.NewTemp()
	ctx.Gen(t2 + " = stack[(int) " + t1 + "]") //se obtiene la posicion del heap donde inicia el vector

	//se hace un count para saber el tamaño del vector
	t3 := ctx.NewTemp()
	ctx.Nat_CountVector(t3, t2)
	l1 := ctx.NewEtq()

	//si el tamaño del vector es 0
	ctx.Gen("if (" + t3 + " == 0) goto " + l1)
	ctx.Nat_RemoveLastVector(t2) //se llama a la funcion nativa para quitar el ultimo elemento del vector
	ctx.GenLabel(l1 + ":")

	return compilador.NewNill()
}

// NT_REMOVE =================================================================================
type NT_Remove struct {
	Id        string
	Expresion compilador.CAbstractExpr
	Linea     int
	Columna   int
}

// Constructor ================================================================================
func NewNT_Remove(id string, expresion compilador.CAbstractExpr, linea int, columna int) *NT_Remove {
	return &NT_Remove{
		Id:        id,
		Expresion: expresion,
		Linea:     linea,
		Columna:   columna,
	}
}

// Implementacion =============================================================================
func (NtRemove *NT_Remove) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	//DEBE VENIR UN ENTERO EN LA EXPRESION
	idsimbolo := ctx.GetSimbolo(NtRemove.Id)

	//si el simbolo no existe
	if idsimbolo.Tipo == 0 {
		ctx.AddErrorLine("Semantico", "La variable "+NtRemove.Id+" no existe", NtRemove.Linea, NtRemove.Columna)
		return compilador.NewNill()
	}

	//si el simbolo no es un vector
	if idsimbolo.TipoId != compilador.Vector.String() {
		ctx.AddErrorLine("Semantico", "La variable "+NtRemove.Id+" no es un vector", NtRemove.Linea, NtRemove.Columna)
		return compilador.NewNill()
	}

	//si el simbolo no es mutable
	if !idsimbolo.Mutable {
		ctx.AddErrorLine("Semantico", "El vector "+NtRemove.Id+" no es mutable", NtRemove.Linea, NtRemove.Columna)
		return compilador.NewNill()
	}

	//se compila la expresion
	exp := NtRemove.Expresion.Compilar(ctx)

	//se verifica que la expresion sea de tipo entero
	if exp.Tipo != compilador.Integer {
		ctx.AddErrorLine("Semantico", "La expresion debe ser de tipo entero", NtRemove.Linea, NtRemove.Columna)
		return compilador.NewNill()
	}

	//se genera el codigo 3d para quitar el elemento del vector
	ctx.GenComentario("Remove")
	t1 := ctx.NewTemp()

	ctx.Gen(t1 + " = P + " + strconv.Itoa(idsimbolo.Size)) //la posicion del inicio del vector en el stack
	t2 := ctx.NewTemp()
	ctx.Gen(t2 + " = stack[(int) " + t1 + "]") //se obtiene la posicion del heap donde inicia el vector

	//hacemos un count para saber el tamaño del vector
	t3 := ctx.NewTemp()
	ctx.Nat_CountVector(t3, t2)
	l1 := ctx.NewEtq()
	//si el valor de la expresion es mayor al tamaño del vector
	ctx.Gen("if (" + exp.Dir + " >= " + t3 + ") goto " + l1)
	ctx.Nat_RemoveVector(t2, exp.Dir) //se llama a la funcion nativa para quitar el elemento del vector
	ctx.GenLabel(l1 + ":")

	return compilador.NewNill()
}

// NT_COUNT ==================================================================================

type NT_Count struct {
	Id      string
	Linea   int
	Columna int
}

// Constructor ================================================================================

func NewNT_Count(id string, linea int, columna int) *NT_Count {
	return &NT_Count{
		Id:      id,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (NtCount *NT_Count) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	idsimbolo := ctx.GetSimbolo(NtCount.Id)

	//si el simbolo no existe
	if idsimbolo.Tipo == 0 {
		ctx.AddErrorLine("Semantico", "La variable "+NtCount.Id+" no existe", NtCount.Linea, NtCount.Columna)
		return compilador.NewNill()
	}

	//si el simbolo no es un vector
	if idsimbolo.TipoId != compilador.Vector.String() {
		ctx.AddErrorLine("Semantico", "La variable "+NtCount.Id+" no es un vector", NtCount.Linea, NtCount.Columna)
		return compilador.NewNill()
	}

	//se genera el codigo 3d para obtener el tamaño del vector
	ctx.GenComentario("Count")
	t1 := ctx.NewTemp()

	ctx.Gen(t1 + " = P + " + strconv.Itoa(idsimbolo.Size)) //la posicion del inicio del vector en el stack
	t2 := ctx.NewTemp()
	ctx.Gen(t2 + " = stack[(int) " + t1 + "]") //se obtiene la posicion del heap donde inicia el vector

	t3 := ctx.NewTemp()

	ctx.Nat_CountVector(t3, t2) //se llama a la funcion nativa para obtener el tamaño del vector

	return compilador.NewInt(t3)
}

// NT_ISEMPTY =================================================================================

type NT_IsEmpty struct {
	Id      string
	Linea   int
	Columna int
}

// Constructor ================================================================================
func NewNT_IsEmpty(id string, linea int, columna int) *NT_IsEmpty {
	return &NT_IsEmpty{
		Id:      id,
		Linea:   linea,
		Columna: columna,
	}
}

// Implementacion =============================================================================

func (NtIsEmpty *NT_IsEmpty) Compilar(ctx *compilador.Contexto) *compilador.Atributos {
	idsimbolo := ctx.GetSimbolo(NtIsEmpty.Id)

	//si el simbolo no existe
	if idsimbolo.Tipo == 0 {
		ctx.AddErrorLine("Semantico", "La variable "+NtIsEmpty.Id+" no existe", NtIsEmpty.Linea, NtIsEmpty.Columna)
		return compilador.NewNill()
	}

	//si el simbolo no es un vector
	if idsimbolo.TipoId != compilador.Vector.String() {
		ctx.AddErrorLine("Semantico", "La variable "+NtIsEmpty.Id+" no es un vector", NtIsEmpty.Linea, NtIsEmpty.Columna)
		return compilador.NewNill()
	}

	//se genera el codigo 3d para saber si el vector esta vacio
	ctx.GenComentario("IsEmpty")
	t1 := ctx.NewTemp()

	ctx.Gen(t1 + " = P + " + strconv.Itoa(idsimbolo.Size)) //la posicion del inicio del vector en el stack
	t2 := ctx.NewTemp()
	ctx.Gen(t2 + " = stack[(int) " + t1 + "]") //se obtiene la posicion del heap donde inicia el vector

	t3 := ctx.NewTemp()

	ctx.Nat_EmptyVector(t3, t2) //se llama a la funcion nativa para saber si el vector esta vacio

	return compilador.NewBool(t3)
}
