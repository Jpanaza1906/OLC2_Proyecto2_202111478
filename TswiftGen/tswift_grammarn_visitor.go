// Code generated from Tswift_GrammarN.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftGen // Tswift_GrammarN
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by Tswift_GrammarNParser.
type Tswift_GrammarNVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by Tswift_GrammarNParser#SLSentencias.
	VisitSLSentencias(ctx *SLSentenciasContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#L_Sentencia.
	VisitL_Sentencia(ctx *L_SentenciaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#S_Print.
	VisitS_Print(ctx *S_PrintContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#S_If.
	VisitS_If(ctx *S_IfContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#S_Switch.
	VisitS_Switch(ctx *S_SwitchContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#S_Guard.
	VisitS_Guard(ctx *S_GuardContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#S_Trans.
	VisitS_Trans(ctx *S_TransContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#S_While.
	VisitS_While(ctx *S_WhileContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#S_For.
	VisitS_For(ctx *S_ForContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Break.
	VisitBreak(ctx *BreakContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Continue.
	VisitContinue(ctx *ContinueContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#If.
	VisitIf(ctx *IfContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Guard.
	VisitGuard(ctx *GuardContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Switch.
	VisitSwitch(ctx *SwitchContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Case.
	VisitCase(ctx *CaseContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Default.
	VisitDefault(ctx *DefaultContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#While.
	VisitWhile(ctx *WhileContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#For.
	VisitFor(ctx *ForContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Rango.
	VisitRango(ctx *RangoContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Par.
	VisitCond_Par(ctx *Cond_ParContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Neg.
	VisitCond_Neg(ctx *Cond_NegContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Rel.
	VisitCond_Rel(ctx *Cond_RelContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Booleano.
	VisitCond_Booleano(ctx *Cond_BooleanoContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Cond_Logica.
	VisitCond_Logica(ctx *Cond_LogicaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_SumRes.
	VisitExpr_SumRes(ctx *Expr_SumResContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Id.
	VisitExpr_Id(ctx *Expr_IdContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Mod.
	VisitExpr_Mod(ctx *Expr_ModContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Neg.
	VisitExpr_Neg(ctx *Expr_NegContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_MulDiv.
	VisitExpr_MulDiv(ctx *Expr_MulDivContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Par.
	VisitExpr_Par(ctx *Expr_ParContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Entero.
	VisitExpr_Entero(ctx *Expr_EnteroContext) interface{}
}
