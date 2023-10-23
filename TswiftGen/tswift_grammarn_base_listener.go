// Code generated from Tswift_GrammarN.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftGen // Tswift_GrammarN
import "github.com/antlr4-go/antlr/v4"

// BaseTswift_GrammarNListener is a complete listener for a parse tree produced by Tswift_GrammarNParser.
type BaseTswift_GrammarNListener struct{}

var _ Tswift_GrammarNListener = &BaseTswift_GrammarNListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTswift_GrammarNListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTswift_GrammarNListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTswift_GrammarNListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTswift_GrammarNListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSLSentencias is called when production SLSentencias is entered.
func (s *BaseTswift_GrammarNListener) EnterSLSentencias(ctx *SLSentenciasContext) {}

// ExitSLSentencias is called when production SLSentencias is exited.
func (s *BaseTswift_GrammarNListener) ExitSLSentencias(ctx *SLSentenciasContext) {}

// EnterL_Sentencia is called when production L_Sentencia is entered.
func (s *BaseTswift_GrammarNListener) EnterL_Sentencia(ctx *L_SentenciaContext) {}

// ExitL_Sentencia is called when production L_Sentencia is exited.
func (s *BaseTswift_GrammarNListener) ExitL_Sentencia(ctx *L_SentenciaContext) {}

// EnterS_Print is called when production S_Print is entered.
func (s *BaseTswift_GrammarNListener) EnterS_Print(ctx *S_PrintContext) {}

// ExitS_Print is called when production S_Print is exited.
func (s *BaseTswift_GrammarNListener) ExitS_Print(ctx *S_PrintContext) {}

// EnterS_If is called when production S_If is entered.
func (s *BaseTswift_GrammarNListener) EnterS_If(ctx *S_IfContext) {}

// ExitS_If is called when production S_If is exited.
func (s *BaseTswift_GrammarNListener) ExitS_If(ctx *S_IfContext) {}

// EnterS_Switch is called when production S_Switch is entered.
func (s *BaseTswift_GrammarNListener) EnterS_Switch(ctx *S_SwitchContext) {}

// ExitS_Switch is called when production S_Switch is exited.
func (s *BaseTswift_GrammarNListener) ExitS_Switch(ctx *S_SwitchContext) {}

// EnterS_Guard is called when production S_Guard is entered.
func (s *BaseTswift_GrammarNListener) EnterS_Guard(ctx *S_GuardContext) {}

// ExitS_Guard is called when production S_Guard is exited.
func (s *BaseTswift_GrammarNListener) ExitS_Guard(ctx *S_GuardContext) {}

// EnterS_Trans is called when production S_Trans is entered.
func (s *BaseTswift_GrammarNListener) EnterS_Trans(ctx *S_TransContext) {}

// ExitS_Trans is called when production S_Trans is exited.
func (s *BaseTswift_GrammarNListener) ExitS_Trans(ctx *S_TransContext) {}

// EnterS_While is called when production S_While is entered.
func (s *BaseTswift_GrammarNListener) EnterS_While(ctx *S_WhileContext) {}

// ExitS_While is called when production S_While is exited.
func (s *BaseTswift_GrammarNListener) ExitS_While(ctx *S_WhileContext) {}

// EnterS_For is called when production S_For is entered.
func (s *BaseTswift_GrammarNListener) EnterS_For(ctx *S_ForContext) {}

// ExitS_For is called when production S_For is exited.
func (s *BaseTswift_GrammarNListener) ExitS_For(ctx *S_ForContext) {}

// EnterBreak is called when production Break is entered.
func (s *BaseTswift_GrammarNListener) EnterBreak(ctx *BreakContext) {}

// ExitBreak is called when production Break is exited.
func (s *BaseTswift_GrammarNListener) ExitBreak(ctx *BreakContext) {}

// EnterContinue is called when production Continue is entered.
func (s *BaseTswift_GrammarNListener) EnterContinue(ctx *ContinueContext) {}

// ExitContinue is called when production Continue is exited.
func (s *BaseTswift_GrammarNListener) ExitContinue(ctx *ContinueContext) {}

// EnterReturn is called when production Return is entered.
func (s *BaseTswift_GrammarNListener) EnterReturn(ctx *ReturnContext) {}

// ExitReturn is called when production Return is exited.
func (s *BaseTswift_GrammarNListener) ExitReturn(ctx *ReturnContext) {}

// EnterPrint is called when production Print is entered.
func (s *BaseTswift_GrammarNListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production Print is exited.
func (s *BaseTswift_GrammarNListener) ExitPrint(ctx *PrintContext) {}

// EnterIf is called when production If is entered.
func (s *BaseTswift_GrammarNListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production If is exited.
func (s *BaseTswift_GrammarNListener) ExitIf(ctx *IfContext) {}

// EnterGuard is called when production Guard is entered.
func (s *BaseTswift_GrammarNListener) EnterGuard(ctx *GuardContext) {}

// ExitGuard is called when production Guard is exited.
func (s *BaseTswift_GrammarNListener) ExitGuard(ctx *GuardContext) {}

// EnterSwitch is called when production Switch is entered.
func (s *BaseTswift_GrammarNListener) EnterSwitch(ctx *SwitchContext) {}

// ExitSwitch is called when production Switch is exited.
func (s *BaseTswift_GrammarNListener) ExitSwitch(ctx *SwitchContext) {}

// EnterCase is called when production Case is entered.
func (s *BaseTswift_GrammarNListener) EnterCase(ctx *CaseContext) {}

// ExitCase is called when production Case is exited.
func (s *BaseTswift_GrammarNListener) ExitCase(ctx *CaseContext) {}

// EnterDefault is called when production Default is entered.
func (s *BaseTswift_GrammarNListener) EnterDefault(ctx *DefaultContext) {}

// ExitDefault is called when production Default is exited.
func (s *BaseTswift_GrammarNListener) ExitDefault(ctx *DefaultContext) {}

// EnterWhile is called when production While is entered.
func (s *BaseTswift_GrammarNListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production While is exited.
func (s *BaseTswift_GrammarNListener) ExitWhile(ctx *WhileContext) {}

// EnterFor is called when production For is entered.
func (s *BaseTswift_GrammarNListener) EnterFor(ctx *ForContext) {}

// ExitFor is called when production For is exited.
func (s *BaseTswift_GrammarNListener) ExitFor(ctx *ForContext) {}

// EnterRango is called when production Rango is entered.
func (s *BaseTswift_GrammarNListener) EnterRango(ctx *RangoContext) {}

// ExitRango is called when production Rango is exited.
func (s *BaseTswift_GrammarNListener) ExitRango(ctx *RangoContext) {}

// EnterCond_Par is called when production Cond_Par is entered.
func (s *BaseTswift_GrammarNListener) EnterCond_Par(ctx *Cond_ParContext) {}

// ExitCond_Par is called when production Cond_Par is exited.
func (s *BaseTswift_GrammarNListener) ExitCond_Par(ctx *Cond_ParContext) {}

// EnterCond_Neg is called when production Cond_Neg is entered.
func (s *BaseTswift_GrammarNListener) EnterCond_Neg(ctx *Cond_NegContext) {}

// ExitCond_Neg is called when production Cond_Neg is exited.
func (s *BaseTswift_GrammarNListener) ExitCond_Neg(ctx *Cond_NegContext) {}

// EnterCond_Rel is called when production Cond_Rel is entered.
func (s *BaseTswift_GrammarNListener) EnterCond_Rel(ctx *Cond_RelContext) {}

// ExitCond_Rel is called when production Cond_Rel is exited.
func (s *BaseTswift_GrammarNListener) ExitCond_Rel(ctx *Cond_RelContext) {}

// EnterCond_Booleano is called when production Cond_Booleano is entered.
func (s *BaseTswift_GrammarNListener) EnterCond_Booleano(ctx *Cond_BooleanoContext) {}

// ExitCond_Booleano is called when production Cond_Booleano is exited.
func (s *BaseTswift_GrammarNListener) ExitCond_Booleano(ctx *Cond_BooleanoContext) {}

// EnterCond_Logica is called when production Cond_Logica is entered.
func (s *BaseTswift_GrammarNListener) EnterCond_Logica(ctx *Cond_LogicaContext) {}

// ExitCond_Logica is called when production Cond_Logica is exited.
func (s *BaseTswift_GrammarNListener) ExitCond_Logica(ctx *Cond_LogicaContext) {}

// EnterExpr_SumRes is called when production Expr_SumRes is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_SumRes(ctx *Expr_SumResContext) {}

// ExitExpr_SumRes is called when production Expr_SumRes is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_SumRes(ctx *Expr_SumResContext) {}

// EnterExpr_Id is called when production Expr_Id is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Id(ctx *Expr_IdContext) {}

// ExitExpr_Id is called when production Expr_Id is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Id(ctx *Expr_IdContext) {}

// EnterExpr_Mod is called when production Expr_Mod is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Mod(ctx *Expr_ModContext) {}

// ExitExpr_Mod is called when production Expr_Mod is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Mod(ctx *Expr_ModContext) {}

// EnterExpr_Neg is called when production Expr_Neg is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Neg(ctx *Expr_NegContext) {}

// ExitExpr_Neg is called when production Expr_Neg is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Neg(ctx *Expr_NegContext) {}

// EnterExpr_MulDiv is called when production Expr_MulDiv is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_MulDiv(ctx *Expr_MulDivContext) {}

// ExitExpr_MulDiv is called when production Expr_MulDiv is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_MulDiv(ctx *Expr_MulDivContext) {}

// EnterExpr_Par is called when production Expr_Par is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Par(ctx *Expr_ParContext) {}

// ExitExpr_Par is called when production Expr_Par is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Par(ctx *Expr_ParContext) {}

// EnterExpr_Entero is called when production Expr_Entero is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Entero(ctx *Expr_EnteroContext) {}

// ExitExpr_Entero is called when production Expr_Entero is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Entero(ctx *Expr_EnteroContext) {}
