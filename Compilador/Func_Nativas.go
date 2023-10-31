package compilador

import (
	"strconv"
)

//Llamada funciones nativas======================================================

// Print String
func (ctx *Contexto) Nat_PrintString(temp string) {
	llamada := "_printString()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + temp)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))

	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["printString"]; !ok {
		ctx.PrintString3d()
	}
}

// Print Bool
func (ctx *Contexto) Nat_PrintBool(temp string) {
	llamada := "_printBool()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + temp)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))

	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["printBool"]; !ok {
		ctx.PrintBool3d()
	}
}

// Equal String
func (ctx *Contexto) Nat_EqualString(temp string, tempIzq string, tempDer string) {
	llamada := "_equalString()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + tempIzq)
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt+1))
	ctx.Gen("stack[(int) " + t1 + "] = " + tempDer)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	//llamada := fmt.Sprintf("%s = _equalString(%s, %s)", temp, tempIzq, tempDer)
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))
	ctx.Gen(temp + " = stack[(int) " + strconv.Itoa(ctx.PosSt) + "]")
	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["equalString"]; !ok {
		ctx.EqualString3d()
	}
}

// Concat String
func (ctx *Contexto) Nat_ConcatString(tempIzq string, tempDer string) {
	llamada := "_concatString()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + tempIzq)
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt+1))
	ctx.Gen("stack[(int) " + t1 + "] = " + tempDer)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	//llamada := fmt.Sprintf("%s = _concatString(%s, %s)", tempIzq, tempIzq, tempDer)
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))
	ctx.Gen(tempIzq + " = stack[(int) " + strconv.Itoa(ctx.PosSt) + "]")
	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["concatString"]; !ok {
		ctx.ConcatString3d()
	}
}

// append vector
func (ctx *Contexto) Nat_AppendVector(inicioVec string, dirItem string) {
	llamada := "_appendVector()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + inicioVec)
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt+1))
	ctx.Gen("stack[(int) " + t1 + "] = " + dirItem)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	//llamada := fmt.Sprintf("_appendVector(%s, %s)", inicioVec, dirItem)
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))
	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["appendVector"]; !ok {
		ctx.AppendVector3d()
	}
}

// removeLast vector
func (ctx *Contexto) Nat_RemoveLastVector(inicioVec string) {
	llamada := "_removeLastVector()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + inicioVec)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	//llamada := fmt.Sprintf("_removeLastVector(%s)", inicioVec)
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))
	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["removeLastVector"]; !ok {
		ctx.RemoveLastVector3d()
	}
}

// remove vector
func (ctx *Contexto) Nat_RemoveVector(inicioVec string, dirItem string) {
	llamada := "_removeVector()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + inicioVec)
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt+1))
	ctx.Gen("stack[(int) " + t1 + "] = " + dirItem)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	//llamada := fmt.Sprintf("_removeVector(%s, %s)", inicioVec, dirItem)
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))
	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["removeVector"]; !ok {
		ctx.RemoveVector3d()
	}
}

// vector pos
func (ctx *Contexto) Nat_VectorPos(temp string, inicioVec string, dirItem string) {
	llamada := "_vectorPos()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + inicioVec)
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt+1))
	ctx.Gen("stack[(int) " + t1 + "] = " + dirItem)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	//llamada := fmt.Sprintf("%s = _vectorPos(%s, %s)", temp, inicioVec, dirItem)
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))
	ctx.Gen(temp + " = stack[(int) " + strconv.Itoa(ctx.PosSt) + "]")

	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["vectorPos"]; !ok {
		ctx.VectorPos3d()
	}
}

// count vector
func (ctx *Contexto) Nat_CountVector(temp string, inicioVec string) {
	llamada := "_countVector()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + inicioVec)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	//llamada := fmt.Sprintf("%s = _countVector(%s)", temp, inicioVec)
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))
	ctx.Gen(temp + " = stack[(int) " + strconv.Itoa(ctx.PosSt) + "]")

	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["countVector"]; !ok {
		ctx.CountVector3d()
	}
}

// empty vector
func (ctx *Contexto) Nat_EmptyVector(temp string, inicioVec string) {
	llamada := "_emptyVector()"
	t1 := ctx.NewTemp()
	ctx.Gen(t1 + " = P + " + strconv.Itoa(ctx.PosSt))
	ctx.Gen("stack[(int) " + t1 + "] = " + inicioVec)
	ctx.Gen("P = P + " + strconv.Itoa(ctx.PosSt))
	//llamada := fmt.Sprintf("%s = _emptyVector(%s)", temp, inicioVec)
	ctx.Gen(llamada)
	ctx.Gen("P = P - " + strconv.Itoa(ctx.PosSt))
	ctx.Gen(temp + " = stack[(int) " + strconv.Itoa(ctx.PosSt) + "]")

	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["emptyVector"]; !ok {
		ctx.EmptyVector3d()
	}
}

//Genera codigo 3d para las funciones nativas====================================

// Concat String
func (ctx *Contexto) ConcatString3d() {
	//en cada posicion de la lista viene la direccion inicial de una cadena
	//se recorre la lista para concatenar las cadenas
	//se genera el codigo 3d para concatenar las cadenas
	nativa := ""

	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()
	L3 := ctx.NewEtq()
	L4 := ctx.NewEtq()

	nativa = "\n// Funcion concat strings\n"
	nativa += "void _concatString(){\n"

	//en tempIzq viene la direccion de la cadena izquierda
	//en tempDer viene la direccion de la cadena derecha
	//se recorre la cadena izquierda para concatenarla
	nativa += "\t" + t1 + " = stack[(int) P];\n"         //inicio de la cadena izquierda
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n" //se obtiene el primer caracter de la cadena
	nativa += "\tstack[(int) P] = H;\n"                  //se guarda la direccion de inicio de la cadena
	nativa += "\t" + L1 + ": //etiqueta inicio\n"
	nativa += "\tif((int) " + t2 + " == 0) goto " + L2 + ";\n"
	//se agrega el caracter al heap
	nativa += "\theap[(int) H] = " + t2 + ";\n"  //se agrega el caracter al heap
	nativa += "\tH = H + 1;\n"                   //se aumenta el contador
	nativa += "\t" + t1 + " = " + t1 + " + 1;\n" //se aumenta la direccion de la cadena
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	//se recorre la cadena derecha para concatenarla
	nativa += "\t" + t1 + " = P + 1;\n"
	nativa += "\t" + t1 + " = stack[(int) " + t1 + "];\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n"
	//nativa += "\tt2 = t0;\n"
	nativa += "\t" + L3 + ":\n"
	nativa += "\tif((int) " + t2 + " == 0) goto " + L4 + ";\n"
	//se agrega el caracter al heap
	nativa += "\theap[(int) H] = " + t2 + ";\n"
	nativa += "\tH = H + 1;\n"
	nativa += "\t" + t1 + " = " + t1 + " + 1;\n"
	nativa += "\t" + t2 + " = heap[(int)" + t1 + "];\n"
	nativa += "\tgoto " + L3 + ";\n"
	nativa += "\t" + L4 + ":\n"
	//se agrega el caracter nulo
	nativa += "\theap[(int) H] = 0;\n"
	nativa += "\tH = H + 1;\n"
	nativa += "\treturn;\n}"

	ctx.AddNativas("concatString", nativa)

}

// Equal String
func (ctx *Contexto) EqualString3d() {
	//parametro 1 -> direccion de la cadena izquierda
	//parametro 2 -> direccion de la cadena derecha

	nativa := ""

	nativa = "\n// Funcion equal strings\n"
	nativa += "void _equalString(){\n"

	//tempIzq -> direccion de la cadena izquierda
	//tempDer -> direccion de la cadena derecha

	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()
	t3 := ctx.NewTemp()
	t4 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()
	L3 := ctx.NewEtq()
	L4 := ctx.NewEtq()

	nativa += "\t" + t1 + " = stack[(int) P];\n" //inicio de la cadena izquierda
	nativa += "\t" + t2 + " = P + 1;\n"
	nativa += "\t" + t2 + " = stack[(int) " + t2 + "];\n" //inicio de la cadena derecha
	nativa += "\t" + L1 + ":\n"
	nativa += "\t" + t3 + " = heap[(int) " + t1 + "];\n"
	nativa += "\t" + t4 + " = heap[(int) " + t2 + "];\n"
	nativa += "\tif(" + t3 + " != " + t4 + ") goto " + L2 + ";\n"
	nativa += "\tif((int) " + t3 + " == 0) goto " + L3 + ";\n"
	nativa += "\t" + t1 + " = " + t1 + " + 1;\n"
	nativa += "\t" + t2 + " = " + t2 + " + 1;\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	nativa += "\tstack[(int) P] = 0;\n"
	nativa += "\tgoto " + L4 + ";\n"
	nativa += "\t" + L3 + ":\n"
	nativa += "\tstack[(int) P] = 1;\n"
	nativa += "\t" + L4 + ":\n"
	nativa += "\treturn;\n"
	nativa += "}"

	ctx.AddNativas("equalString", nativa)
}

// Print String
func (ctx *Contexto) PrintString3d() {
	//parametro 1 -> direccion de inicio de la cadena

	nativa := ""

	nativa = "\n// Funcion print string\n"
	nativa += "void _printString(){\n"

	//temp -> direccion de inicio de la cadena

	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = stack[(int) P];\n"
	nativa += "\t" + L1 + ":\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n"
	nativa += "\tif((int) " + t2 + " == 0) goto " + L2 + ";\n"
	nativa += "\tprintf(\"%c\", (int) " + t2 + ");\n"
	nativa += "\t" + t1 + " = " + t1 + " + 1;\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	nativa += "\treturn;\n"
	nativa += "}"

	ctx.AddNativas("printString", nativa)

}

// Print Bool
func (ctx *Contexto) PrintBool3d() {
	//parametro 1 -> numero 1 o 0

	nativa := ""

	nativa = "\n// Funcion print bool\n"
	nativa += "void _printBool(){\n"

	//temp -> numero 1 o 0

	t1 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = stack[(int) P];\n"
	nativa += "\tif((int) " + t1 + " == 1) goto " + L1 + ";\n"
	//imprimir false caracter por caracter
	nativa += "\tprintf(\"%c\", (int) 102);\n"
	nativa += "\tprintf(\"%c\", (int) 97);\n"
	nativa += "\tprintf(\"%c\", (int) 108);\n"
	nativa += "\tprintf(\"%c\", (int) 115);\n"
	nativa += "\tprintf(\"%c\", (int) 101);\n"
	nativa += "\tgoto " + L2 + ";\n"
	nativa += "\t" + L1 + ":\n"
	//imprimir true caracter por caracter
	nativa += "\tprintf(\"%c\", (int) 116);\n"
	nativa += "\tprintf(\"%c\", (int) 114);\n"
	nativa += "\tprintf(\"%c\", (int) 117);\n"
	nativa += "\tprintf(\"%c\", (int) 101);\n"
	nativa += "\t" + L2 + ":\n"
	nativa += "\treturn;\n"
	nativa += "}"

	ctx.AddNativas("printBool", nativa)
}

// Append Vector

func (ctx *Contexto) AppendVector3d() {
	//parametro 1 -> inicio del vector en el heap
	//parametro 2 -> direccion del item a agregar

	nativa := ""

	nativa = "\n// Funcion append vector\n"

	nativa += "void _appendVector(){\n"

	//tempIzq -> inicio del vector en el heap

	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()
	t3 := ctx.NewTemp()
	t4 := ctx.NewTemp()
	t11 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = stack[(int) P];\n"
	nativa += "\t" + L1 + ":\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n" //el item
	nativa += "\t" + t3 + " = " + t1 + "+ 1;\n"          //se obtiene la direccion del siguiente item
	nativa += "\t" + t4 + " = heap[(int) " + t3 + "];\n" //direccion del heap del siguiente item
	nativa += "\tif((int) " + t4 + " == -1) goto " + L2 + ";\n"
	nativa += "\t" + t1 + " = " + t4 + ";\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	nativa += "\t" + t11 + " = P + 1;\n"
	nativa += "\t" + t11 + " = stack[(int) " + t11 + "];\n"
	nativa += "\theap[(int) " + t1 + " ] = " + t11 + ";\n"
	nativa += "\theap[(int) " + t3 + " ] = H;\n"
	nativa += "\theap[(int) H] = 0;\n"
	nativa += "\tH = H + 1;\n"
	nativa += "\theap[(int) H] = -1;\n"
	nativa += "\tH = H + 1;\n"
	nativa += "\treturn;\n"

	nativa += "}"

	ctx.AddNativas("appendVector", nativa)

}

// RemoveLast Vector

func (ctx *Contexto) RemoveLastVector3d() {
	//parametro 1 -> inicio del vector en el heap

	nativa := ""

	nativa = "\n// Funcion removeLast vector\n"

	nativa += "void _removeLastVector(){\n"

	//temp -> inicio del vector en el heap

	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()
	t3 := ctx.NewTemp()
	t4 := ctx.NewTemp()
	tant := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = stack[(int) P];\n"
	nativa += "\t" + L1 + ":\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n" //el item
	nativa += "\t" + t3 + " = " + t1 + "+ 1;\n"          //se obtiene la direccion del siguiente item
	nativa += "\t" + t4 + " = heap[(int) " + t3 + "];\n" //direccion del heap del siguiente item
	nativa += "\tif((int) " + t4 + " == -1) goto " + L2 + ";\n"
	nativa += "\t" + tant + " = " + t1 + ";\n" //en tant queda la direccion de la ultima posicion
	nativa += "\t" + t1 + " = " + t4 + ";\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	nativa += "\theap[(int) " + tant + " ] = 0;\n"
	nativa += "\t" + t3 + " = " + tant + " + 1;\n"
	nativa += "\theap[(int) " + t3 + " ] = -1;\n"
	nativa += "\treturn;\n"

	nativa += "}"

	ctx.AddNativas("removeLastVector", nativa)
}

// Remove Vector

func (ctx *Contexto) RemoveVector3d() {
	//parametro 1 -> inicio del vector en el heap
	//parametro 2 -> direccion del item a eliminar

	nativa := ""

	nativa = "\n// Funcion remove vector\n"

	nativa += "void _removeVector(){\n"

	//tempIzq -> inicio del vector en el heap

	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()
	t3 := ctx.NewTemp()
	t4 := ctx.NewTemp()
	t5 := ctx.NewTemp()
	cont := ctx.NewTemp()
	tant := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = stack[(int) P];\n"
	nativa += "\t" + t5 + " = P + 1;\n"
	nativa += "\t" + t5 + " = stack[(int) " + t5 + "];\n"
	nativa += "\t" + cont + " = 0;\n"
	nativa += "\t" + L1 + ":\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n" //el item
	nativa += "\t" + t3 + " = " + t1 + "+ 1;\n"          //se obtiene la direccion del siguiente item
	nativa += "\t" + t4 + " = heap[(int) " + t3 + "];\n" //direccion del heap del siguiente item
	nativa += "\tif((int) " + cont + " == " + t5 + ") goto " + L2 + ";\n"
	nativa += "\t" + tant + " = " + t3 + ";\n" //en tant queda la direccion de la ultima posicion
	nativa += "\t" + t1 + " = " + t4 + ";\n"
	nativa += "\t" + cont + " = " + cont + " + 1;\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	//apunta a la direccion del siguiente item
	nativa += "\theap[(int) " + tant + " ] = " + t4 + ";\n"

	nativa += "\treturn;\n"

	nativa += "}"

	ctx.AddNativas("removeVector", nativa)
}

// Count Vector

func (ctx *Contexto) CountVector3d() {
	//parametro 1 -> inicio del vector en el heap

	nativa := ""

	nativa = "\n// Funcion count vector\n"

	nativa += "void _countVector(){\n"

	//temp -> inicio del vector en el heap

	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()
	t3 := ctx.NewTemp()
	t4 := ctx.NewTemp()
	cont := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = stack[(int) P];\n"
	nativa += "\t" + cont + " = 0;\n"
	nativa += "\t" + L1 + ":\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n" //el item
	nativa += "\t" + t3 + " = " + t1 + "+ 1;\n"          //se obtiene la direccion del siguiente item
	nativa += "\t" + t4 + " = heap[(int) " + t3 + "];\n" //direccion del heap del siguiente item
	nativa += "\tif((int) " + t4 + " == -1) goto " + L2 + ";\n"
	nativa += "\t" + t1 + " = " + t4 + ";\n"
	nativa += "\t" + cont + " = " + cont + " + 1;\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	nativa += "\tstack[(int) P] = " + cont + ";\n"
	nativa += "\treturn;\n"

	nativa += "}"

	ctx.AddNativas("countVector", nativa)
}

// Empty Vector

func (ctx *Contexto) EmptyVector3d() {

	//parametro 1 -> inicio del vector en el heap

	nativa := ""

	nativa = "\n// Funcion empty vector\n"

	nativa += "void _emptyVector(){\n"

	//temp -> inicio del vector en el heap
	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()

	l1 := ctx.NewEtq()
	l2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = stack[(int) P];\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n"
	nativa += "\t" + t1 + " = " + t1 + " + 1;\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n"
	nativa += "\tif((int) " + t2 + " != -1) goto " + l1 + ";\n"
	nativa += "\tstack[(int) P] = 1;\n"
	nativa += "\tgoto " + l2 + ";\n"
	nativa += "\t" + l1 + ":\n"
	nativa += "\tstack[(int) P] = 0;\n"
	nativa += "\t" + l2 + ":\n"
	nativa += "\treturn;\n"

	nativa += "}"

	ctx.AddNativas("emptyVector", nativa)
}

// Vector Pos

func (ctx *Contexto) VectorPos3d() {
	//parametro 1 -> inicio del vector en el heap
	//parametro 2 -> direccion del item a obtener

	nativa := ""

	nativa = "\n// Funcion vector pos\n"

	nativa += "void _vectorPos(){\n"

	//tempIzq -> inicio del vector en el heap

	t1 := ctx.NewTemp()
	t5 := ctx.NewTemp()
	cont := ctx.NewTemp()
	t2 := ctx.NewTemp()
	t3 := ctx.NewTemp()
	t4 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = stack[(int) P];\n"
	nativa += "\t" + t5 + " = P + 1;\n"
	nativa += "\t" + t5 + " = stack[(int) " + t5 + "];\n"
	nativa += "\t" + cont + " = 0;\n"
	nativa += "\t" + L1 + ":\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n" //el item
	nativa += "\t" + t3 + " = " + t1 + "+ 1;\n"          //se obtiene la direccion del siguiente item
	nativa += "\t" + t4 + " = heap[(int) " + t3 + "];\n" //direccion del heap del siguiente item
	nativa += "\tif((int) " + cont + " == " + t5 + ") goto " + L2 + ";\n"
	nativa += "\t" + t1 + " = " + t4 + ";\n"
	nativa += "\t" + cont + " = " + cont + " + 1;\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	nativa += "\tstack[(int) P] = " + t2 + ";\n"
	nativa += "\treturn;\n"

	nativa += "}"

	ctx.AddNativas("vectorPos", nativa)
}
