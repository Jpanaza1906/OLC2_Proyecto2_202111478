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

	// EnterS_Declaracion is called when entering the S_Declaracion production.
	EnterS_Declaracion(c *S_DeclaracionContext)

	// EnterS_If is called when entering the S_If production.
	EnterS_If(c *S_IfContext)

	// EnterS_Switch is called when entering the S_Switch production.
	EnterS_Switch(c *S_SwitchContext)

	// EnterS_Guard is called when entering the S_Guard production.
	EnterS_Guard(c *S_GuardContext)

	// EnterS_Trans is called when entering the S_Trans production.
	EnterS_Trans(c *S_TransContext)

	// EnterS_While is called when entering the S_While production.
	EnterS_While(c *S_WhileContext)

	// EnterS_For is called when entering the S_For production.
	EnterS_For(c *S_ForContext)

	// EnterBreak is called when entering the Break production.
	EnterBreak(c *BreakContext)

	// EnterContinue is called when entering the Continue production.
	EnterContinue(c *ContinueContext)

	// EnterReturn is called when entering the Return production.
	EnterReturn(c *ReturnContext)

	// EnterPrint is called when entering the Print production.
	EnterPrint(c *PrintContext)

	// EnterDeclaracion_Tipo_Val is called when entering the Declaracion_Tipo_Val production.
	EnterDeclaracion_Tipo_Val(c *Declaracion_Tipo_ValContext)

	// EnterDeclaracion_Val is called when entering the Declaracion_Val production.
	EnterDeclaracion_Val(c *Declaracion_ValContext)

	// EnterDeclaracion_Tipo_noVal is called when entering the Declaracion_Tipo_noVal production.
	EnterDeclaracion_Tipo_noVal(c *Declaracion_Tipo_noValContext)

	// EnterTipo_Int is called when entering the Tipo_Int production.
	EnterTipo_Int(c *Tipo_IntContext)

	// EnterTipo_Float is called when entering the Tipo_Float production.
	EnterTipo_Float(c *Tipo_FloatContext)

	// EnterTipo_String is called when entering the Tipo_String production.
	EnterTipo_String(c *Tipo_StringContext)

	// EnterTipo_Bool is called when entering the Tipo_Bool production.
	EnterTipo_Bool(c *Tipo_BoolContext)

	// EnterTipo_Character is called when entering the Tipo_Character production.
	EnterTipo_Character(c *Tipo_CharacterContext)

	// EnterTipo_Struct is called when entering the Tipo_Struct production.
	EnterTipo_Struct(c *Tipo_StructContext)

	// EnterTipo_Vector is called when entering the Tipo_Vector production.
	EnterTipo_Vector(c *Tipo_VectorContext)

	// EnterIf is called when entering the If production.
	EnterIf(c *IfContext)

	// EnterGuard is called when entering the Guard production.
	EnterGuard(c *GuardContext)

	// EnterSwitch is called when entering the Switch production.
	EnterSwitch(c *SwitchContext)

	// EnterCase is called when entering the Case production.
	EnterCase(c *CaseContext)

	// EnterDefault is called when entering the Default production.
	EnterDefault(c *DefaultContext)

	// EnterWhile is called when entering the While production.
	EnterWhile(c *WhileContext)

	// EnterForInt is called when entering the ForInt production.
	EnterForInt(c *ForIntContext)

	// EnterForList is called when entering the ForList production.
	EnterForList(c *ForListContext)

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

	// EnterExpr_Decimal is called when entering the Expr_Decimal production.
	EnterExpr_Decimal(c *Expr_DecimalContext)

	// EnterExpr_Caracter is called when entering the Expr_Caracter production.
	EnterExpr_Caracter(c *Expr_CaracterContext)

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

	// EnterExpr_Nil is called when entering the Expr_Nil production.
	EnterExpr_Nil(c *Expr_NilContext)

	// EnterExpr_Par is called when entering the Expr_Par production.
	EnterExpr_Par(c *Expr_ParContext)

	// EnterExpr_Booleano is called when entering the Expr_Booleano production.
	EnterExpr_Booleano(c *Expr_BooleanoContext)

	// EnterExpr_Entero is called when entering the Expr_Entero production.
	EnterExpr_Entero(c *Expr_EnteroContext)

	// EnterExpr_Cadena is called when entering the Expr_Cadena production.
	EnterExpr_Cadena(c *Expr_CadenaContext)

	// ExitSLSentencias is called when exiting the SLSentencias production.
	ExitSLSentencias(c *SLSentenciasContext)

	// ExitL_Sentencia is called when exiting the L_Sentencia production.
	ExitL_Sentencia(c *L_SentenciaContext)

	// ExitS_Print is called when exiting the S_Print production.
	ExitS_Print(c *S_PrintContext)

	// ExitS_Declaracion is called when exiting the S_Declaracion production.
	ExitS_Declaracion(c *S_DeclaracionContext)

	// ExitS_If is called when exiting the S_If production.
	ExitS_If(c *S_IfContext)

	// ExitS_Switch is called when exiting the S_Switch production.
	ExitS_Switch(c *S_SwitchContext)

	// ExitS_Guard is called when exiting the S_Guard production.
	ExitS_Guard(c *S_GuardContext)

	// ExitS_Trans is called when exiting the S_Trans production.
	ExitS_Trans(c *S_TransContext)

	// ExitS_While is called when exiting the S_While production.
	ExitS_While(c *S_WhileContext)

	// ExitS_For is called when exiting the S_For production.
	ExitS_For(c *S_ForContext)

	// ExitBreak is called when exiting the Break production.
	ExitBreak(c *BreakContext)

	// ExitContinue is called when exiting the Continue production.
	ExitContinue(c *ContinueContext)

	// ExitReturn is called when exiting the Return production.
	ExitReturn(c *ReturnContext)

	// ExitPrint is called when exiting the Print production.
	ExitPrint(c *PrintContext)

	// ExitDeclaracion_Tipo_Val is called when exiting the Declaracion_Tipo_Val production.
	ExitDeclaracion_Tipo_Val(c *Declaracion_Tipo_ValContext)

	// ExitDeclaracion_Val is called when exiting the Declaracion_Val production.
	ExitDeclaracion_Val(c *Declaracion_ValContext)

	// ExitDeclaracion_Tipo_noVal is called when exiting the Declaracion_Tipo_noVal production.
	ExitDeclaracion_Tipo_noVal(c *Declaracion_Tipo_noValContext)

	// ExitTipo_Int is called when exiting the Tipo_Int production.
	ExitTipo_Int(c *Tipo_IntContext)

	// ExitTipo_Float is called when exiting the Tipo_Float production.
	ExitTipo_Float(c *Tipo_FloatContext)

	// ExitTipo_String is called when exiting the Tipo_String production.
	ExitTipo_String(c *Tipo_StringContext)

	// ExitTipo_Bool is called when exiting the Tipo_Bool production.
	ExitTipo_Bool(c *Tipo_BoolContext)

	// ExitTipo_Character is called when exiting the Tipo_Character production.
	ExitTipo_Character(c *Tipo_CharacterContext)

	// ExitTipo_Struct is called when exiting the Tipo_Struct production.
	ExitTipo_Struct(c *Tipo_StructContext)

	// ExitTipo_Vector is called when exiting the Tipo_Vector production.
	ExitTipo_Vector(c *Tipo_VectorContext)

	// ExitIf is called when exiting the If production.
	ExitIf(c *IfContext)

	// ExitGuard is called when exiting the Guard production.
	ExitGuard(c *GuardContext)

	// ExitSwitch is called when exiting the Switch production.
	ExitSwitch(c *SwitchContext)

	// ExitCase is called when exiting the Case production.
	ExitCase(c *CaseContext)

	// ExitDefault is called when exiting the Default production.
	ExitDefault(c *DefaultContext)

	// ExitWhile is called when exiting the While production.
	ExitWhile(c *WhileContext)

	// ExitForInt is called when exiting the ForInt production.
	ExitForInt(c *ForIntContext)

	// ExitForList is called when exiting the ForList production.
	ExitForList(c *ForListContext)

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

	// ExitExpr_Decimal is called when exiting the Expr_Decimal production.
	ExitExpr_Decimal(c *Expr_DecimalContext)

	// ExitExpr_Caracter is called when exiting the Expr_Caracter production.
	ExitExpr_Caracter(c *Expr_CaracterContext)

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

	// ExitExpr_Nil is called when exiting the Expr_Nil production.
	ExitExpr_Nil(c *Expr_NilContext)

	// ExitExpr_Par is called when exiting the Expr_Par production.
	ExitExpr_Par(c *Expr_ParContext)

	// ExitExpr_Booleano is called when exiting the Expr_Booleano production.
	ExitExpr_Booleano(c *Expr_BooleanoContext)

	// ExitExpr_Entero is called when exiting the Expr_Entero production.
	ExitExpr_Entero(c *Expr_EnteroContext)

	// ExitExpr_Cadena is called when exiting the Expr_Cadena production.
	ExitExpr_Cadena(c *Expr_CadenaContext)
}
