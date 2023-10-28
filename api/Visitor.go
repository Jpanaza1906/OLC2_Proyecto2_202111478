package api

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	noterm "OLC2_Proyecto2_202111478/Compilador/NoTerm"
	terminales "OLC2_Proyecto2_202111478/Compilador/Terminales"
	"OLC2_Proyecto2_202111478/TswiftGen"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	antlr.ParseTreeVisitor
	Raiz compilador.CAbstractExpr
}

func NewVisitor() TswiftGen.Tswift_GrammarNVisitor {
	return &Visitor{
		ParseTreeVisitor: &TswiftGen.BaseTswift_GrammarNVisitor{},
		Raiz:             nil,
	}
}

func (tV *Visitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		return noterm.NewNT_Error(val.GetText())
	default:
		nodoCompilador, ok := tree.Accept(tV).(compilador.CAbstractExpr)
		if ok {
			return nodoCompilador
		}
		return noterm.NewNT_Error("Nodo desconocido")
	}
}

func (tV *Visitor) VisitChildren(node antlr.RuleNode) interface{} {
	return ""
}

func (tV *Visitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return ""
}

func (tV *Visitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return noterm.NewNT_Error(node.GetText())
}

// Visit a parse tree produced by Tswift_GrammarNParser#SLSentencias.
func (tV *Visitor) VisitSLSentencias(ctx *TswiftGen.SLSentenciasContext) interface{} {
	lsentencias := ctx.L_sentencias().Accept(tV).(compilador.CAbstractExpr)
	nodoS := noterm.NewNT_S(lsentencias)
	tV.Raiz = nodoS
	return nodoS
}

// Visit a parse tree produced by Tswift_GrammarNParser#L_Sentencia.
func (tV *Visitor) VisitL_Sentencia(ctx *TswiftGen.L_SentenciaContext) interface{} {
	lsentencias := noterm.NewNT_LSentencias()
	sentenciasAntlr := ctx.AllSentencia()
	for _, sentenciaAntlr := range sentenciasAntlr {
		nodoSentencia, ok := sentenciaAntlr.Accept(tV).(compilador.CAbstractExpr)
		if ok {
			lsentencias.AddSentencia(nodoSentencia)
		}
	}
	return lsentencias
}

// SENTENCIAS ==========================================================

// Visit a parse tree produced by Tswift_GrammarNParser#S_Print.
func (tV *Visitor) VisitS_Print(ctx *TswiftGen.S_PrintContext) interface{} {
	return ctx.Print_sentencia().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_Declaracion.
func (tV *Visitor) VisitS_Declaracion(ctx *TswiftGen.S_DeclaracionContext) interface{} {
	return ctx.Declaracion_sentencia().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_Asignacion.
func (tV *Visitor) VisitS_Asignacion(ctx *TswiftGen.S_AsignacionContext) interface{} {
	return ctx.Asignacion_sentencia().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_If.
func (tV *Visitor) VisitS_If(ctx *TswiftGen.S_IfContext) interface{} {
	return ctx.If_sentencia().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_Guard.
func (tV *Visitor) VisitS_Guard(ctx *TswiftGen.S_GuardContext) interface{} {
	return ctx.Guard_sentencia().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_Switch.
func (tV *Visitor) VisitS_Switch(ctx *TswiftGen.S_SwitchContext) interface{} {
	return ctx.Switch_sentencia().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_Trans.
func (tV *Visitor) VisitS_Trans(ctx *TswiftGen.S_TransContext) interface{} {
	return ctx.Trans_sentencia().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_While.
func (tV *Visitor) VisitS_While(ctx *TswiftGen.S_WhileContext) interface{} {
	return ctx.While_sentencia().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_For.
func (tV *Visitor) VisitS_For(ctx *TswiftGen.S_ForContext) interface{} {
	return ctx.For_sentencia().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_Declaracion_Vector.
func (tV *Visitor) VisitS_Declaracion_Vector(ctx *TswiftGen.S_Declaracion_VectorContext) interface{} {
	return ctx.Dec_vector().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#S_Funcion_Vector.
func (tV *Visitor) VisitS_Funcion_Vector(ctx *TswiftGen.S_Funcion_VectorContext) interface{} {
	return ctx.Func_vector().Accept(tV).(compilador.CAbstractExpr)
}

// SENTENCIAS DE TRANSFERENCIA =========================================

// Visit a parse tree produced by Tswift_GrammarNParser#Break.
func (tV *Visitor) VisitBreak(ctx *TswiftGen.BreakContext) interface{} {
	return noterm.NewNT_Break(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Continue.
func (tV *Visitor) VisitContinue(ctx *TswiftGen.ContinueContext) interface{} {
	return noterm.NewNT_Continue(ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Return.
func (tV *Visitor) VisitReturn(ctx *TswiftGen.ReturnContext) interface{} {
	//Se verifica si viene una expresion
	if ctx.E() != nil {
		return noterm.NewNT_Return(ctx.E().Accept(tV).(compilador.CAbstractExpr), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	return noterm.NewNT_Return(nil, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// PRINT ===============================================================

// Visit a parse tree produced by Tswift_GrammarNParser#Print.
func (tV *Visitor) VisitPrint(ctx *TswiftGen.PrintContext) interface{} {
	lExpresiones := ctx.AllE()

	//se crea un arreglo de expresiones
	expresiones := make([]compilador.CAbstractExpr, 0)

	//se recorre la lista de expresiones
	for _, e := range lExpresiones {
		expresiones = append(expresiones, e.Accept(tV).(compilador.CAbstractExpr))
	}

	return noterm.NewNT_Print(expresiones, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// DECLARACION VARIABLES Y CONSTANTES ===================================

// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Tipo_Val.
func (tV *Visitor) VisitDeclaracion_Tipo_Val(ctx *TswiftGen.Declaracion_Tipo_ValContext) interface{} {
	// get tipvar
	tipodec := ctx.GetTip().GetText()

	//si es let, mutable es false
	mutable := false

	if tipodec == "let" {
		mutable = false
	} else {
		mutable = true
	}

	//tipo de variable
	tipo := ctx.Tipo().Accept(tV).(string)

	//id
	id := ctx.ID().GetText()

	//expresion
	expresion := ctx.E().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Declaracion(id, tipo, mutable, expresion, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

}

// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Val.
func (tV *Visitor) VisitDeclaracion_Val(ctx *TswiftGen.Declaracion_ValContext) interface{} {
	// get tipvar
	tipodec := ctx.GetTip().GetText()

	//si es let, mutable es false
	mutable := false

	if tipodec == "let" {
		mutable = false
	} else {
		mutable = true
	}

	//id
	id := ctx.ID().GetText()

	//expresion
	expresion := ctx.E().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Declaracion(id, "", mutable, expresion, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Tipo_noVal.
func (tV *Visitor) VisitDeclaracion_Tipo_noVal(ctx *TswiftGen.Declaracion_Tipo_noValContext) interface{} {
	// get tipvar
	tipodec := ctx.GetTip().GetText()

	//si es let, mutable es false
	mutable := false

	if tipodec == "let" {
		mutable = false
	} else {
		mutable = true
	}

	//tipo de variable
	tipo := ctx.Tipo().Accept(tV).(string)

	//id
	id := ctx.ID().GetText()

	return noterm.NewNT_Declaracion(id, tipo, mutable, nil, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// ASIGNACION ==========================================================

// Visit a parse tree produced by Tswift_GrammarNParser#Asignacion.
func (tV *Visitor) VisitAsignacion(ctx *TswiftGen.AsignacionContext) interface{} {
	id := ctx.ID().GetText()
	expresion := ctx.E().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Asignacion(id, expresion, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Asignacion_SumaResta.
func (tV *Visitor) VisitAsignacion_SumaResta(ctx *TswiftGen.Asignacion_SumaRestaContext) interface{} {
	id := ctx.ID().GetText()
	oper := ctx.GetOp().GetText()
	//declarar una expresion
	var expresion compilador.CAbstractExpr

	if oper == "+=" {
		oper = "+"
		expresion = noterm.NewNT_OpAritmeticas(terminales.NewT_Identificador(id, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn()), ctx.E().Accept(tV).(compilador.CAbstractExpr), oper, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

	} else {
		oper = "-"
		expresion = noterm.NewNT_OpAritmeticas(terminales.NewT_Identificador(id, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn()), ctx.E().Accept(tV).(compilador.CAbstractExpr), oper, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	return noterm.NewNT_Asignacion(id, expresion, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

}

// IF ==================================================================

// Visit a parse tree produced by Tswift_GrammarNParser#If.
func (tV *Visitor) VisitIf(ctx *TswiftGen.IfContext) interface{} {
	var sentencias compilador.CAbstractExpr
	var sentenciasElse compilador.CAbstractExpr

	condicion := ctx.Condicion().Accept(tV).(compilador.CAbstractExpr)

	listaSentencias := ctx.AllL_sentencias()
	if listaSentencias != nil {
		if len(listaSentencias) == 2 {
			sentencias = listaSentencias[0].Accept(tV).(compilador.CAbstractExpr)
			sentenciasElse = listaSentencias[1].Accept(tV).(compilador.CAbstractExpr)
		} else {
			sentencias = listaSentencias[0].Accept(tV).(compilador.CAbstractExpr)
		}
	}

	//Si viene otro if
	if ctx.If_sentencia() != nil {
		sentenciasElse = ctx.If_sentencia().Accept(tV).(compilador.CAbstractExpr)
	}

	return noterm.NewNT_If(condicion, sentencias, sentenciasElse, ctx.Condicion().GetStart().GetLine(), ctx.Condicion().GetStart().GetColumn())
}

// GUARD ===============================================================

// Visit a parse tree produced by Tswift_GrammarNParser#Guard.
func (tV *Visitor) VisitGuard(ctx *TswiftGen.GuardContext) interface{} {
	//se obtiene la condicion
	condicion := ctx.Condicion().Accept(tV).(compilador.CAbstractExpr)

	//se obtiene la lista de sentencias
	sentencias := ctx.L_sentencias().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Guard(condicion, sentencias, ctx.Condicion().GetStart().GetLine(), ctx.Condicion().GetStart().GetColumn())
}

// SWITCH ==============================================================

// Visit a parse tree produced by Tswift_GrammarNParser#Switch.
func (tV *Visitor) VisitSwitch(ctx *TswiftGen.SwitchContext) interface{} {
	//Se obtiene la expresion
	expr := ctx.E().Accept(tV).(compilador.CAbstractExpr)

	//se crea un default
	var cdef compilador.CAbstractExpr = nil

	//Se obtiene la lista de casos
	cases := ctx.AllLcasos()

	if ctx.Cdefault() != nil {
		cdef = ctx.Cdefault().Accept(tV).(compilador.CAbstractExpr)
	}

	//Se crea un arreglo de casos
	listaCasos := make([]compilador.CAbstractExpr, 0)

	//Se recorre la lista de casos
	for _, c := range cases {
		listaCasos = append(listaCasos, c.Accept(tV).(compilador.CAbstractExpr))
	}

	return noterm.NewNT_Switch(expr, listaCasos, cdef, ctx.E().GetStart().GetLine(), ctx.E().GetStart().GetColumn())

}

// Visit a parse tree produced by Tswift_GrammarNParser#Case.
func (tV *Visitor) VisitCase(ctx *TswiftGen.CaseContext) interface{} {
	expr := ctx.E().Accept(tV).(compilador.CAbstractExpr)
	sentencias := ctx.L_sentencias().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Case(expr, sentencias)
}

// Visit a parse tree produced by Tswift_GrammarNParser#Default.
func (tV *Visitor) VisitDefault(ctx *TswiftGen.DefaultContext) interface{} {
	return ctx.L_sentencias().Accept(tV).(compilador.CAbstractExpr)
}

// WHILE ===============================================================

// Visit a parse tree produced by Tswift_GrammarNParser#While.
func (tV *Visitor) VisitWhile(ctx *TswiftGen.WhileContext) interface{} {
	//Se obtiene la condicion
	condicion := ctx.Condicion().Accept(tV).(compilador.CAbstractExpr)

	//Se obtiene la lista de sentencias
	sentencias := ctx.L_sentencias().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_While(condicion, sentencias, ctx.Condicion().GetStart().GetLine(), ctx.Condicion().GetStart().GetColumn())
}

// FOR =================================================================

// Visit a parse tree produced by Tswift_GrammarNParser#ForInt.
func (tV *Visitor) VisitForInt(ctx *TswiftGen.ForIntContext) interface{} {
	id := ctx.ID().GetText()
	expIzq := ctx.E(0).Accept(tV).(compilador.CAbstractExpr)
	expDer := ctx.E(1).Accept(tV).(compilador.CAbstractExpr)

	//Se obtiene la lista de sentencias
	sentencias := ctx.L_sentencias().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_ForInt(id, expIzq, expDer, sentencias, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

}

// Visit a parse tree produced by Tswift_GrammarNParser#ForList.
func (tV *Visitor) VisitForList(ctx *TswiftGen.ForListContext) interface{} {
	id := ctx.ID().GetText()
	exp := ctx.E().Accept(tV).(compilador.CAbstractExpr)

	//Se obtiene la lista de sentencias
	sentencias := ctx.L_sentencias().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_ForList(id, exp, sentencias, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// VECTORES =============================================================
// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Vector.
func (tV *Visitor) VisitDeclaracion_Vector(ctx *TswiftGen.Declaracion_VectorContext) interface{} {
	tipodec := ctx.GetTipod().GetText()

	//si es let, mutable es false
	mutable := false

	if tipodec == "let" {
		mutable = false
	} else {
		mutable = true
	}

	id := ctx.ID().GetText()

	tipo := ctx.Tipo().Accept(tV).(string)

	defvector := ctx.Def_vector().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_DecVector(id, tipo, mutable, defvector, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

}

// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Alterna.
func (tV *Visitor) VisitDeclaracion_Alterna(ctx *TswiftGen.Declaracion_AlternaContext) interface{} {
	tipodec := ctx.GetTipod().GetText()

	//si es let, mutable es false
	mutable := false

	if tipodec == "let" {
		mutable = false
	} else {
		mutable = true
	}

	id := ctx.ID().GetText()

	tipo := ctx.Tipo().Accept(tV).(string)

	return noterm.NewNT_DecVector(id, tipo, mutable, nil, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Def_Vector.
func (tV *Visitor) VisitDef_Vector(ctx *TswiftGen.Def_VectorContext) interface{} {
	lexpresiones := ctx.AllE()

	vector := []compilador.CAbstractExpr{}

	for _, exp := range lexpresiones {
		vector = append(vector, exp.Accept(tV).(compilador.CAbstractExpr))
	}

	return noterm.NewNT_DefVector("", vector, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Def_Vector_Vacio.
func (tV *Visitor) VisitDef_Vector_Vacio(ctx *TswiftGen.Def_Vector_VacioContext) interface{} {
	vector := []compilador.CAbstractExpr{}

	return noterm.NewNT_DefVector("", vector, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Def_Vector_Id.
func (tV *Visitor) VisitDef_Vector_Id(ctx *TswiftGen.Def_Vector_IdContext) interface{} {
	id := ctx.ID().GetText()

	return noterm.NewNT_DefVector(id, nil, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Func_Vector_Append.
func (tV *Visitor) VisitFunc_Vector_Append(ctx *TswiftGen.Func_Vector_AppendContext) interface{} {
	id := ctx.ID().GetText()
	exp := ctx.E().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Append(id, exp, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Func_Vector_RemoveLast.
func (tV *Visitor) VisitFunc_Vector_RemoveLast(ctx *TswiftGen.Func_Vector_RemoveLastContext) interface{} {
	id := ctx.ID().GetText()

	return noterm.NewNT_RemoveLast(id, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Func_Vector_Remove.
func (tV *Visitor) VisitFunc_Vector_Remove(ctx *TswiftGen.Func_Vector_RemoveContext) interface{} {
	id := ctx.ID().GetText()
	exp := ctx.E().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Remove(id, exp, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Count.
func (tV *Visitor) VisitExpr_Count(ctx *TswiftGen.Expr_CountContext) interface{} {
	id := ctx.ID().GetText()

	return noterm.NewNT_Count(id, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_IsEmpty.
func (tV *Visitor) VisitExpr_IsEmpty(ctx *TswiftGen.Expr_IsEmptyContext) interface{} {
	id := ctx.ID().GetText()

	return noterm.NewNT_IsEmpty(id, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// CONDICIONES ========================================================

// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Par.
func (tV *Visitor) VisitCond_Par(ctx *TswiftGen.Cond_ParContext) interface{} {
	return ctx.GetC().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Neg.
func (tV *Visitor) VisitCond_Neg(ctx *TswiftGen.Cond_NegContext) interface{} {
	cond := ctx.GetC().Accept(tV).(compilador.CAbstractExpr)
	return noterm.NewNT_Not(cond, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Rel.
func (tV *Visitor) VisitCond_Rel(ctx *TswiftGen.Cond_RelContext) interface{} {
	condIzq := ctx.GetE1().Accept(tV).(compilador.CAbstractExpr)
	condDer := ctx.GetE2().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_OpRelacionales(condIzq, condDer, ctx.GetOp().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Booleano.
func (tV *Visitor) VisitCond_Booleano(ctx *TswiftGen.Cond_BooleanoContext) interface{} {
	return terminales.NewT_BoolCond(ctx.GetOp().GetText())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Logica.
func (tV *Visitor) VisitCond_Logica(ctx *TswiftGen.Cond_LogicaContext) interface{} {
	condIzq := ctx.GetC1().Accept(tV).(compilador.CAbstractExpr)
	condDer := ctx.GetC2().Accept(tV).(compilador.CAbstractExpr)

	// se verifica que sean AND o OR
	if ctx.GetOp().GetText() == "&&" {
		return noterm.NewNT_And(condIzq, condDer, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	return noterm.NewNT_Or(condIzq, condDer, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// TIPOS ===============================================================

// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Int.
func (tV *Visitor) VisitTipo_Int(ctx *TswiftGen.Tipo_IntContext) interface{} {
	return ctx.INT().GetText()
}

// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Float.
func (tV *Visitor) VisitTipo_Float(ctx *TswiftGen.Tipo_FloatContext) interface{} {
	return ctx.FLOAT().GetText()
}

// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_String.
func (tV *Visitor) VisitTipo_String(ctx *TswiftGen.Tipo_StringContext) interface{} {
	return ctx.STRING().GetText()
}

// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Bool.
func (tV *Visitor) VisitTipo_Bool(ctx *TswiftGen.Tipo_BoolContext) interface{} {
	return ctx.BOOL().GetText()
}

// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Character.
func (tV *Visitor) VisitTipo_Character(ctx *TswiftGen.Tipo_CharacterContext) interface{} {
	return ctx.CHAR().GetText()
}

// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Struct.
func (tV *Visitor) VisitTipo_Struct(ctx *TswiftGen.Tipo_StructContext) interface{} {
	return "Struct"
}

// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Vector.
func (tV *Visitor) VisitTipo_Vector(ctx *TswiftGen.Tipo_VectorContext) interface{} {
	return "Vector"
}

// EXPRESIONES ========================================================

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_SumRes.
func (tV *Visitor) VisitExpr_SumRes(ctx *TswiftGen.Expr_SumResContext) interface{} {
	expIzq := ctx.GetE1().Accept(tV).(compilador.CAbstractExpr)
	expDer := ctx.GetE2().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_OpAritmeticas(expIzq, expDer, ctx.GetOp().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Id.
func (tV *Visitor) VisitExpr_Id(ctx *TswiftGen.Expr_IdContext) interface{} {
	return terminales.NewT_Identificador(ctx.ID().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Mod.
func (tV *Visitor) VisitExpr_Mod(ctx *TswiftGen.Expr_ModContext) interface{} {
	expIzq := ctx.GetE1().Accept(tV).(compilador.CAbstractExpr)
	expDer := ctx.GetE2().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Mod(expIzq, expDer, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Neg.
func (tV *Visitor) VisitExpr_Neg(ctx *TswiftGen.Expr_NegContext) interface{} {
	// expresion
	expr := ctx.GetE1().Accept(tV).(compilador.CAbstractExpr)

	if ctx.GetOp().GetText() == "!" {
		return noterm.NewNT_NotExp(expr, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
	}
	return noterm.NewNT_Negativo(expr, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_MulDiv.
func (tV *Visitor) VisitExpr_MulDiv(ctx *TswiftGen.Expr_MulDivContext) interface{} {
	expIzq := ctx.GetE1().Accept(tV).(compilador.CAbstractExpr)
	expDer := ctx.GetE2().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_OpAritmeticas(expIzq, expDer, ctx.GetOp().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Par.
func (tV *Visitor) VisitExpr_Par(ctx *TswiftGen.Expr_ParContext) interface{} {
	return ctx.GetE1().Accept(tV).(compilador.CAbstractExpr)
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Entero.
func (tV *Visitor) VisitExpr_Entero(ctx *TswiftGen.Expr_EnteroContext) interface{} {
	return terminales.NewT_Num(ctx.GetN().GetText())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Decimal.
func (tV *Visitor) VisitExpr_Decimal(ctx *TswiftGen.Expr_DecimalContext) interface{} {
	return terminales.NewT_Float(ctx.GetN().GetText())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Caracter.
func (tV *Visitor) VisitExpr_Caracter(ctx *TswiftGen.Expr_CaracterContext) interface{} {
	cadena := ctx.GetN().GetText()
	//eliminar las comillas
	cadena = cadena[1 : len(cadena)-1]
	return terminales.NewT_Char(cadena)
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Nil.
func (tV *Visitor) VisitExpr_Nil(ctx *TswiftGen.Expr_NilContext) interface{} {
	return terminales.NewT_Nil()
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Booleano.
func (tV *Visitor) VisitExpr_Booleano(ctx *TswiftGen.Expr_BooleanoContext) interface{} {
	return terminales.NewT_Bool(ctx.GetN().GetText())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Cadena.
func (tV *Visitor) VisitExpr_Cadena(ctx *TswiftGen.Expr_CadenaContext) interface{} {
	//Se obtiene el texto de la cadena
	cadena := ctx.GetN().GetText()

	//eliminar las comillas
	cadena = cadena[1 : len(cadena)-1]
	return terminales.NewT_Cad(cadena)
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Rel.
func (tV *Visitor) VisitExpr_Rel(ctx *TswiftGen.Expr_RelContext) interface{} {
	expizq := ctx.E(0).Accept(tV).(compilador.CAbstractExpr)
	expder := ctx.E(1).Accept(tV).(compilador.CAbstractExpr)

	operador := ctx.GetOp().GetText()

	return noterm.NewNT_ExprRelacional(expizq, expder, operador, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())

}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Logica.
func (tV *Visitor) VisitExpr_Logica(ctx *TswiftGen.Expr_LogicaContext) interface{} {
	expizq := ctx.E(0).Accept(tV).(compilador.CAbstractExpr)
	expder := ctx.E(1).Accept(tV).(compilador.CAbstractExpr)

	operador := ctx.GetOp().GetText()

	return noterm.NewNT_ExprLogica(expizq, expder, operador, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Vector.
func (tV *Visitor) VisitExpr_Vector(ctx *TswiftGen.Expr_VectorContext) interface{} {
	id := ctx.ID().GetText()
	exp := ctx.E().Accept(tV).(compilador.CAbstractExpr)

	return terminales.NewT_VectorPos(id, exp, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}
