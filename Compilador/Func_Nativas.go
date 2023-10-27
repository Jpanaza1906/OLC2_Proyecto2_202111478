package compilador

import "fmt"

//Llamada funciones nativas======================================================

// Print String
func (ctx *Contexto) Nat_PrintString(temp string) {
	llamada := fmt.Sprintf("_printString(%s)", temp)
	ctx.Gen(llamada)

	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["printString"]; !ok {
		ctx.PrintString3d()
	}
}

// Print Bool
func (ctx *Contexto) Nat_PrintBool(temp string) {
	llamada := fmt.Sprintf("_printBool(%s)", temp)
	ctx.Gen(llamada)

	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["printBool"]; !ok {
		ctx.PrintBool3d()
	}
}

// Equal String
func (ctx *Contexto) Nat_EqualString(temp string, tempIzq string, tempDer string) {
	llamada := fmt.Sprintf("%s = _equalString(%s, %s)", temp, tempIzq, tempDer)
	ctx.Gen(llamada)
	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["equalString"]; !ok {
		ctx.EqualString3d()
	}
}

// Concat String
func (ctx *Contexto) Nat_ConcatString(tempIzq string, tempDer string) {
	llamada := fmt.Sprintf("%s = _concatString(%s, %s)", tempIzq, tempIzq, tempDer)
	ctx.Gen(llamada)
	//se verifica si ya existe una funcion en el contexto
	//si no existe, se agrega
	if _, ok := ctx.Fnativas["concatString"]; !ok {
		ctx.ConcatString3d()
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
	t3 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()
	L3 := ctx.NewEtq()
	L4 := ctx.NewEtq()

	nativa = "\n// Funcion concat strings\n"
	nativa += "int _concatString(int tempIzq, int tempDer){\n"

	//en tempIzq viene la direccion de la cadena izquierda
	//en tempDer viene la direccion de la cadena derecha
	//se recorre la cadena izquierda para concatenarla
	nativa += "\t" + t1 + " = tempIzq;\n"                           //inicio de la cadena izquierda
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n"            //se obtiene el primer caracter de la cadena
	nativa += "\t" + t3 + " = H; //se guarda inicio nueva cadena\n" //contador para recorrer la cadena
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
	nativa += "\t" + t1 + " = tempDer;\n"
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
	nativa += "\treturn " + t3 + ";\n}"

	ctx.AddNativas("concatString", nativa)

}

// Equal String
func (ctx *Contexto) EqualString3d() {
	//parametro 1 -> direccion de la cadena izquierda
	//parametro 2 -> direccion de la cadena derecha

	nativa := ""

	nativa = "\n// Funcion equal strings\n"
	nativa += "int _equalString(int tempIzq, int tempDer){\n"

	//tempIzq -> direccion de la cadena izquierda
	//tempDer -> direccion de la cadena derecha

	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()
	t3 := ctx.NewTemp()
	t4 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()
	L3 := ctx.NewEtq()

	nativa += "\t" + t1 + " = tempIzq;\n" //inicio de la cadena izquierda
	nativa += "\t" + t2 + " = tempDer;\n" //inicio de la cadena derecha
	nativa += "\t" + L1 + ":\n"
	nativa += "\t" + t3 + " = heap[(int) " + t1 + "];\n"
	nativa += "\t" + t4 + " = heap[(int) " + t2 + "];\n"
	nativa += "\tif(" + t3 + " != " + t4 + ") goto " + L2 + ";\n"
	nativa += "\tif((int) " + t3 + " == 0) goto " + L3 + ";\n"
	nativa += "\t" + t1 + " = " + t1 + " + 1;\n"
	nativa += "\t" + t2 + " = " + t2 + " + 1;\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	nativa += "\treturn 0;\n"
	nativa += "\t" + L3 + ":\n"
	nativa += "\treturn 1;\n"
	nativa += "}"

	ctx.AddNativas("equalString", nativa)
}

// Print String
func (ctx *Contexto) PrintString3d() {
	//parametro 1 -> direccion de inicio de la cadena

	nativa := ""

	nativa = "\n// Funcion print string\n"
	nativa += "void _printString(int temp){\n"

	//temp -> direccion de inicio de la cadena

	t1 := ctx.NewTemp()
	t2 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = temp;\n"
	nativa += "\t" + L1 + ":\n"
	nativa += "\t" + t2 + " = heap[(int) " + t1 + "];\n"
	nativa += "\tif((int) " + t2 + " == 0) goto " + L2 + ";\n"
	nativa += "\tprintf(\"%c\", (int) " + t2 + ");\n"
	nativa += "\t" + t1 + " = " + t1 + " + 1;\n"
	nativa += "\tgoto " + L1 + ";\n"
	nativa += "\t" + L2 + ":\n"
	nativa += "\t return;\n"
	nativa += "}"

	ctx.AddNativas("printString", nativa)

}

// Print Bool
func (ctx *Contexto) PrintBool3d() {
	//parametro 1 -> numero 1 o 0

	nativa := ""

	nativa = "\n// Funcion print bool\n"
	nativa += "void _printBool(int temp){\n"

	//temp -> numero 1 o 0

	t1 := ctx.NewTemp()

	L1 := ctx.NewEtq()
	L2 := ctx.NewEtq()

	nativa += "\t" + t1 + " = temp;\n"
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
	nativa += "\t return;\n"
	nativa += "}"

	ctx.AddNativas("printBool", nativa)
}
