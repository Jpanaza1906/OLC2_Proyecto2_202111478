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
	panic("not implemented") // TODO: Implement
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
	panic("not implemented") // TODO: Implement
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

	//se obtiene la transicion
	transicion := ctx.Trans_sentencia().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Guard(condicion, sentencias, transicion, ctx.Condicion().GetStart().GetLine(), ctx.Condicion().GetStart().GetColumn())
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
	panic("not implemented") // TODO: Implement
}

// FOR =================================================================

// Visit a parse tree produced by Tswift_GrammarNParser#For.
func (tV *Visitor) VisitFor(ctx *TswiftGen.ForContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Tswift_GrammarNParser#Rango.
func (tV *Visitor) VisitRango(ctx *TswiftGen.RangoContext) interface{} {
	panic("not implemented") // TODO: Implement
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
	return terminales.NewT_Bool(ctx.GetOp().GetText())
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

// EXPRESIONES ========================================================

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_SumRes.
func (tV *Visitor) VisitExpr_SumRes(ctx *TswiftGen.Expr_SumResContext) interface{} {
	expIzq := ctx.GetE1().Accept(tV).(compilador.CAbstractExpr)
	expDer := ctx.GetE2().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_OpAritmeticas(expIzq, expDer, ctx.GetOp().GetText(), ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Id.
func (tV *Visitor) VisitExpr_Id(ctx *TswiftGen.Expr_IdContext) interface{} {
	panic("not implemented") // TODO: Implement
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Mod.
func (tV *Visitor) VisitExpr_Mod(ctx *TswiftGen.Expr_ModContext) interface{} {
	expIzq := ctx.GetE1().Accept(tV).(compilador.CAbstractExpr)
	expDer := ctx.GetE2().Accept(tV).(compilador.CAbstractExpr)

	return noterm.NewNT_Mod(expIzq, expDer, ctx.GetStart().GetLine(), ctx.GetStart().GetColumn())
}

// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Neg.
func (tV *Visitor) VisitExpr_Neg(ctx *TswiftGen.Expr_NegContext) interface{} {
	expr := ctx.GetE1().Accept(tV).(compilador.CAbstractExpr)
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
