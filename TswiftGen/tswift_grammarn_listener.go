// Code generated from Tswift_GrammarN.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftGen // Tswift_GrammarN
import "github.com/antlr4-go/antlr/v4"

// Tswift_GrammarNListener is a complete listener for a parse tree produced by Tswift_GrammarNParser.
type Tswift_GrammarNListener interface {
	antlr.ParseTreeListener

	// EnterSLSentencias is called when entering the SLSentencias production.
	EnterSLSentencias(c *SLSentenciasContext)

	// EnterL_Sentencia is called when entering the L_Sentencia production.
	EnterL_Sentencia(c *L_SentenciaContext)

	// EnterS_Print is called when entering the S_Print production.
	EnterS_Print(c *S_PrintContext)

	// EnterS_Condicion is called when entering the S_Condicion production.
	EnterS_Condicion(c *S_CondicionContext)

	// EnterS_If is called when entering the S_If production.
	EnterS_If(c *S_IfContext)

	// EnterS_Switch is called when entering the S_Switch production.
	EnterS_Switch(c *S_SwitchContext)

	// EnterPrint is called when entering the Print production.
	EnterPrint(c *PrintContext)

	// EnterIf is called when entering the If production.
	EnterIf(c *IfContext)

	// EnterSwitch is called when entering the Switch production.
	EnterSwitch(c *SwitchContext)

	// EnterCase is called when entering the Case production.
	EnterCase(c *CaseContext)

	// EnterDefault is called when entering the Default production.
	EnterDefault(c *DefaultContext)

	// EnterCond_Par is called when entering the Cond_Par production.
	EnterCond_Par(c *Cond_ParContext)

	// EnterCond_Neg is called when entering the Cond_Neg production.
	EnterCond_Neg(c *Cond_NegContext)

	// EnterCond_Rel is called when entering the Cond_Rel production.
	EnterCond_Rel(c *Cond_RelContext)

	// EnterCond_Booleano is called when entering the Cond_Booleano production.
	EnterCond_Booleano(c *Cond_BooleanoContext)

	// EnterCond_Logica is called when entering the Cond_Logica production.
	EnterCond_Logica(c *Cond_LogicaContext)

	// EnterExpr_SumRes is called when entering the Expr_SumRes production.
	EnterExpr_SumRes(c *Expr_SumResContext)

	// EnterExpr_Id is called when entering the Expr_Id production.
	EnterExpr_Id(c *Expr_IdContext)

	// EnterExpr_Mod is called when entering the Expr_Mod production.
	EnterExpr_Mod(c *Expr_ModContext)

	// EnterExpr_Neg is called when entering the Expr_Neg production.
	EnterExpr_Neg(c *Expr_NegContext)

	// EnterExpr_MulDiv is called when entering the Expr_MulDiv production.
	EnterExpr_MulDiv(c *Expr_MulDivContext)

	// EnterExpr_Par is called when entering the Expr_Par production.
	EnterExpr_Par(c *Expr_ParContext)

	// EnterExpr_Entero is called when entering the Expr_Entero production.
	EnterExpr_Entero(c *Expr_EnteroContext)

	// ExitSLSentencias is called when exiting the SLSentencias production.
	ExitSLSentencias(c *SLSentenciasContext)

	// ExitL_Sentencia is called when exiting the L_Sentencia production.
	ExitL_Sentencia(c *L_SentenciaContext)

	// ExitS_Print is called when exiting the S_Print production.
	ExitS_Print(c *S_PrintContext)

	// ExitS_Condicion is called when exiting the S_Condicion production.
	ExitS_Condicion(c *S_CondicionContext)

	// ExitS_If is called when exiting the S_If production.
	ExitS_If(c *S_IfContext)

	// ExitS_Switch is called when exiting the S_Switch production.
	ExitS_Switch(c *S_SwitchContext)

	// ExitPrint is called when exiting the Print production.
	ExitPrint(c *PrintContext)

	// ExitIf is called when exiting the If production.
	ExitIf(c *IfContext)

	// ExitSwitch is called when exiting the Switch production.
	ExitSwitch(c *SwitchContext)

	// ExitCase is called when exiting the Case production.
	ExitCase(c *CaseContext)

	// ExitDefault is called when exiting the Default production.
	ExitDefault(c *DefaultContext)

	// ExitCond_Par is called when exiting the Cond_Par production.
	ExitCond_Par(c *Cond_ParContext)

	// ExitCond_Neg is called when exiting the Cond_Neg production.
	ExitCond_Neg(c *Cond_NegContext)

	// ExitCond_Rel is called when exiting the Cond_Rel production.
	ExitCond_Rel(c *Cond_RelContext)

	// ExitCond_Booleano is called when exiting the Cond_Booleano production.
	ExitCond_Booleano(c *Cond_BooleanoContext)

	// ExitCond_Logica is called when exiting the Cond_Logica production.
	ExitCond_Logica(c *Cond_LogicaContext)

	// ExitExpr_SumRes is called when exiting the Expr_SumRes production.
	ExitExpr_SumRes(c *Expr_SumResContext)

	// ExitExpr_Id is called when exiting the Expr_Id production.
	ExitExpr_Id(c *Expr_IdContext)

	// ExitExpr_Mod is called when exiting the Expr_Mod production.
	ExitExpr_Mod(c *Expr_ModContext)

	// ExitExpr_Neg is called when exiting the Expr_Neg production.
	ExitExpr_Neg(c *Expr_NegContext)

	// ExitExpr_MulDiv is called when exiting the Expr_MulDiv production.
	ExitExpr_MulDiv(c *Expr_MulDivContext)

	// ExitExpr_Par is called when exiting the Expr_Par production.
	ExitExpr_Par(c *Expr_ParContext)

	// ExitExpr_Entero is called when exiting the Expr_Entero production.
	ExitExpr_Entero(c *Expr_EnteroContext)
}
