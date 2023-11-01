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

	// EnterS_Asignacion is called when entering the S_Asignacion production.
	EnterS_Asignacion(c *S_AsignacionContext)

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

	// EnterS_Declaracion_Vector is called when entering the S_Declaracion_Vector production.
	EnterS_Declaracion_Vector(c *S_Declaracion_VectorContext)

	// EnterS_Funcion_Vector is called when entering the S_Funcion_Vector production.
	EnterS_Funcion_Vector(c *S_Funcion_VectorContext)

	// EnterS_Declaracion_Matriz is called when entering the S_Declaracion_Matriz production.
	EnterS_Declaracion_Matriz(c *S_Declaracion_MatrizContext)

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

	// EnterAsignacion is called when entering the Asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterAsignacion_SumaResta is called when entering the Asignacion_SumaResta production.
	EnterAsignacion_SumaResta(c *Asignacion_SumaRestaContext)

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

	// EnterDeclaracion_Vector is called when entering the Declaracion_Vector production.
	EnterDeclaracion_Vector(c *Declaracion_VectorContext)

	// EnterDeclaracion_Alterna is called when entering the Declaracion_Alterna production.
	EnterDeclaracion_Alterna(c *Declaracion_AlternaContext)

	// EnterDef_Vector is called when entering the Def_Vector production.
	EnterDef_Vector(c *Def_VectorContext)

	// EnterDef_Vector_Vacio is called when entering the Def_Vector_Vacio production.
	EnterDef_Vector_Vacio(c *Def_Vector_VacioContext)

	// EnterDef_Vector_Id is called when entering the Def_Vector_Id production.
	EnterDef_Vector_Id(c *Def_Vector_IdContext)

	// EnterFunc_Vector_Append is called when entering the Func_Vector_Append production.
	EnterFunc_Vector_Append(c *Func_Vector_AppendContext)

	// EnterFunc_Vector_RemoveLast is called when entering the Func_Vector_RemoveLast production.
	EnterFunc_Vector_RemoveLast(c *Func_Vector_RemoveLastContext)

	// EnterFunc_Vector_Remove is called when entering the Func_Vector_Remove production.
	EnterFunc_Vector_Remove(c *Func_Vector_RemoveContext)

	// EnterDeclaracion_Matriz is called when entering the Declaracion_Matriz production.
	EnterDeclaracion_Matriz(c *Declaracion_MatrizContext)

	// EnterDef_Matriz_Lista is called when entering the Def_Matriz_Lista production.
	EnterDef_Matriz_Lista(c *Def_Matriz_ListaContext)

	// EnterDef_Matriz_Simple is called when entering the Def_Matriz_Simple production.
	EnterDef_Matriz_Simple(c *Def_Matriz_SimpleContext)

	// EnterDef_Matriz_Lista_Valores is called when entering the Def_Matriz_Lista_Valores production.
	EnterDef_Matriz_Lista_Valores(c *Def_Matriz_Lista_ValoresContext)

	// EnterDef_Matriz_Expresion is called when entering the Def_Matriz_Expresion production.
	EnterDef_Matriz_Expresion(c *Def_Matriz_ExpresionContext)

	// EnterDef_Matriz_Simple2 is called when entering the Def_Matriz_Simple2 production.
	EnterDef_Matriz_Simple2(c *Def_Matriz_Simple2Context)

	// EnterDef_Matriz_Simple3 is called when entering the Def_Matriz_Simple3 production.
	EnterDef_Matriz_Simple3(c *Def_Matriz_Simple3Context)

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

	// EnterExpr_Rel is called when entering the Expr_Rel production.
	EnterExpr_Rel(c *Expr_RelContext)

	// EnterExpr_Decimal is called when entering the Expr_Decimal production.
	EnterExpr_Decimal(c *Expr_DecimalContext)

	// EnterExpr_Caracter is called when entering the Expr_Caracter production.
	EnterExpr_Caracter(c *Expr_CaracterContext)

	// EnterExpr_SumRes is called when entering the Expr_SumRes production.
	EnterExpr_SumRes(c *Expr_SumResContext)

	// EnterExpr_Matriz is called when entering the Expr_Matriz production.
	EnterExpr_Matriz(c *Expr_MatrizContext)

	// EnterExpr_Neg is called when entering the Expr_Neg production.
	EnterExpr_Neg(c *Expr_NegContext)

	// EnterExpr_MulDiv is called when entering the Expr_MulDiv production.
	EnterExpr_MulDiv(c *Expr_MulDivContext)

	// EnterExpr_Nil is called when entering the Expr_Nil production.
	EnterExpr_Nil(c *Expr_NilContext)

	// EnterExpr_Cadena is called when entering the Expr_Cadena production.
	EnterExpr_Cadena(c *Expr_CadenaContext)

	// EnterExpr_Count is called when entering the Expr_Count production.
	EnterExpr_Count(c *Expr_CountContext)

	// EnterExpr_Id is called when entering the Expr_Id production.
	EnterExpr_Id(c *Expr_IdContext)

	// EnterExpr_Mod is called when entering the Expr_Mod production.
	EnterExpr_Mod(c *Expr_ModContext)

	// EnterExpr_Par is called when entering the Expr_Par production.
	EnterExpr_Par(c *Expr_ParContext)

	// EnterExpr_Logica is called when entering the Expr_Logica production.
	EnterExpr_Logica(c *Expr_LogicaContext)

	// EnterExpr_IsEmpty is called when entering the Expr_IsEmpty production.
	EnterExpr_IsEmpty(c *Expr_IsEmptyContext)

	// EnterExpr_Booleano is called when entering the Expr_Booleano production.
	EnterExpr_Booleano(c *Expr_BooleanoContext)

	// EnterExpr_Vector is called when entering the Expr_Vector production.
	EnterExpr_Vector(c *Expr_VectorContext)

	// EnterExpr_Entero is called when entering the Expr_Entero production.
	EnterExpr_Entero(c *Expr_EnteroContext)

	// ExitSLSentencias is called when exiting the SLSentencias production.
	ExitSLSentencias(c *SLSentenciasContext)

	// ExitL_Sentencia is called when exiting the L_Sentencia production.
	ExitL_Sentencia(c *L_SentenciaContext)

	// ExitS_Print is called when exiting the S_Print production.
	ExitS_Print(c *S_PrintContext)

	// ExitS_Declaracion is called when exiting the S_Declaracion production.
	ExitS_Declaracion(c *S_DeclaracionContext)

	// ExitS_Asignacion is called when exiting the S_Asignacion production.
	ExitS_Asignacion(c *S_AsignacionContext)

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

	// ExitS_Declaracion_Vector is called when exiting the S_Declaracion_Vector production.
	ExitS_Declaracion_Vector(c *S_Declaracion_VectorContext)

	// ExitS_Funcion_Vector is called when exiting the S_Funcion_Vector production.
	ExitS_Funcion_Vector(c *S_Funcion_VectorContext)

	// ExitS_Declaracion_Matriz is called when exiting the S_Declaracion_Matriz production.
	ExitS_Declaracion_Matriz(c *S_Declaracion_MatrizContext)

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

	// ExitAsignacion is called when exiting the Asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitAsignacion_SumaResta is called when exiting the Asignacion_SumaResta production.
	ExitAsignacion_SumaResta(c *Asignacion_SumaRestaContext)

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

	// ExitDeclaracion_Vector is called when exiting the Declaracion_Vector production.
	ExitDeclaracion_Vector(c *Declaracion_VectorContext)

	// ExitDeclaracion_Alterna is called when exiting the Declaracion_Alterna production.
	ExitDeclaracion_Alterna(c *Declaracion_AlternaContext)

	// ExitDef_Vector is called when exiting the Def_Vector production.
	ExitDef_Vector(c *Def_VectorContext)

	// ExitDef_Vector_Vacio is called when exiting the Def_Vector_Vacio production.
	ExitDef_Vector_Vacio(c *Def_Vector_VacioContext)

	// ExitDef_Vector_Id is called when exiting the Def_Vector_Id production.
	ExitDef_Vector_Id(c *Def_Vector_IdContext)

	// ExitFunc_Vector_Append is called when exiting the Func_Vector_Append production.
	ExitFunc_Vector_Append(c *Func_Vector_AppendContext)

	// ExitFunc_Vector_RemoveLast is called when exiting the Func_Vector_RemoveLast production.
	ExitFunc_Vector_RemoveLast(c *Func_Vector_RemoveLastContext)

	// ExitFunc_Vector_Remove is called when exiting the Func_Vector_Remove production.
	ExitFunc_Vector_Remove(c *Func_Vector_RemoveContext)

	// ExitDeclaracion_Matriz is called when exiting the Declaracion_Matriz production.
	ExitDeclaracion_Matriz(c *Declaracion_MatrizContext)

	// ExitDef_Matriz_Lista is called when exiting the Def_Matriz_Lista production.
	ExitDef_Matriz_Lista(c *Def_Matriz_ListaContext)

	// ExitDef_Matriz_Simple is called when exiting the Def_Matriz_Simple production.
	ExitDef_Matriz_Simple(c *Def_Matriz_SimpleContext)

	// ExitDef_Matriz_Lista_Valores is called when exiting the Def_Matriz_Lista_Valores production.
	ExitDef_Matriz_Lista_Valores(c *Def_Matriz_Lista_ValoresContext)

	// ExitDef_Matriz_Expresion is called when exiting the Def_Matriz_Expresion production.
	ExitDef_Matriz_Expresion(c *Def_Matriz_ExpresionContext)

	// ExitDef_Matriz_Simple2 is called when exiting the Def_Matriz_Simple2 production.
	ExitDef_Matriz_Simple2(c *Def_Matriz_Simple2Context)

	// ExitDef_Matriz_Simple3 is called when exiting the Def_Matriz_Simple3 production.
	ExitDef_Matriz_Simple3(c *Def_Matriz_Simple3Context)

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

	// ExitExpr_Rel is called when exiting the Expr_Rel production.
	ExitExpr_Rel(c *Expr_RelContext)

	// ExitExpr_Decimal is called when exiting the Expr_Decimal production.
	ExitExpr_Decimal(c *Expr_DecimalContext)

	// ExitExpr_Caracter is called when exiting the Expr_Caracter production.
	ExitExpr_Caracter(c *Expr_CaracterContext)

	// ExitExpr_SumRes is called when exiting the Expr_SumRes production.
	ExitExpr_SumRes(c *Expr_SumResContext)

	// ExitExpr_Matriz is called when exiting the Expr_Matriz production.
	ExitExpr_Matriz(c *Expr_MatrizContext)

	// ExitExpr_Neg is called when exiting the Expr_Neg production.
	ExitExpr_Neg(c *Expr_NegContext)

	// ExitExpr_MulDiv is called when exiting the Expr_MulDiv production.
	ExitExpr_MulDiv(c *Expr_MulDivContext)

	// ExitExpr_Nil is called when exiting the Expr_Nil production.
	ExitExpr_Nil(c *Expr_NilContext)

	// ExitExpr_Cadena is called when exiting the Expr_Cadena production.
	ExitExpr_Cadena(c *Expr_CadenaContext)

	// ExitExpr_Count is called when exiting the Expr_Count production.
	ExitExpr_Count(c *Expr_CountContext)

	// ExitExpr_Id is called when exiting the Expr_Id production.
	ExitExpr_Id(c *Expr_IdContext)

	// ExitExpr_Mod is called when exiting the Expr_Mod production.
	ExitExpr_Mod(c *Expr_ModContext)

	// ExitExpr_Par is called when exiting the Expr_Par production.
	ExitExpr_Par(c *Expr_ParContext)

	// ExitExpr_Logica is called when exiting the Expr_Logica production.
	ExitExpr_Logica(c *Expr_LogicaContext)

	// ExitExpr_IsEmpty is called when exiting the Expr_IsEmpty production.
	ExitExpr_IsEmpty(c *Expr_IsEmptyContext)

	// ExitExpr_Booleano is called when exiting the Expr_Booleano production.
	ExitExpr_Booleano(c *Expr_BooleanoContext)

	// ExitExpr_Vector is called when exiting the Expr_Vector production.
	ExitExpr_Vector(c *Expr_VectorContext)

	// ExitExpr_Entero is called when exiting the Expr_Entero production.
	ExitExpr_Entero(c *Expr_EnteroContext)
}
