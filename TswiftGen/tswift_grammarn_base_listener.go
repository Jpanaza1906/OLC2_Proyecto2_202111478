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

// EnterS_Declaracion is called when production S_Declaracion is entered.
func (s *BaseTswift_GrammarNListener) EnterS_Declaracion(ctx *S_DeclaracionContext) {}

// ExitS_Declaracion is called when production S_Declaracion is exited.
func (s *BaseTswift_GrammarNListener) ExitS_Declaracion(ctx *S_DeclaracionContext) {}

// EnterS_Asignacion is called when production S_Asignacion is entered.
func (s *BaseTswift_GrammarNListener) EnterS_Asignacion(ctx *S_AsignacionContext) {}

// ExitS_Asignacion is called when production S_Asignacion is exited.
func (s *BaseTswift_GrammarNListener) ExitS_Asignacion(ctx *S_AsignacionContext) {}

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

// EnterS_Declaracion_Vector is called when production S_Declaracion_Vector is entered.
func (s *BaseTswift_GrammarNListener) EnterS_Declaracion_Vector(ctx *S_Declaracion_VectorContext) {}

// ExitS_Declaracion_Vector is called when production S_Declaracion_Vector is exited.
func (s *BaseTswift_GrammarNListener) ExitS_Declaracion_Vector(ctx *S_Declaracion_VectorContext) {}

// EnterS_Funcion_Vector is called when production S_Funcion_Vector is entered.
func (s *BaseTswift_GrammarNListener) EnterS_Funcion_Vector(ctx *S_Funcion_VectorContext) {}

// ExitS_Funcion_Vector is called when production S_Funcion_Vector is exited.
func (s *BaseTswift_GrammarNListener) ExitS_Funcion_Vector(ctx *S_Funcion_VectorContext) {}

// EnterS_Declaracion_Matriz is called when production S_Declaracion_Matriz is entered.
func (s *BaseTswift_GrammarNListener) EnterS_Declaracion_Matriz(ctx *S_Declaracion_MatrizContext) {}

// ExitS_Declaracion_Matriz is called when production S_Declaracion_Matriz is exited.
func (s *BaseTswift_GrammarNListener) ExitS_Declaracion_Matriz(ctx *S_Declaracion_MatrizContext) {}

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

// EnterDeclaracion_Tipo_Val is called when production Declaracion_Tipo_Val is entered.
func (s *BaseTswift_GrammarNListener) EnterDeclaracion_Tipo_Val(ctx *Declaracion_Tipo_ValContext) {}

// ExitDeclaracion_Tipo_Val is called when production Declaracion_Tipo_Val is exited.
func (s *BaseTswift_GrammarNListener) ExitDeclaracion_Tipo_Val(ctx *Declaracion_Tipo_ValContext) {}

// EnterDeclaracion_Val is called when production Declaracion_Val is entered.
func (s *BaseTswift_GrammarNListener) EnterDeclaracion_Val(ctx *Declaracion_ValContext) {}

// ExitDeclaracion_Val is called when production Declaracion_Val is exited.
func (s *BaseTswift_GrammarNListener) ExitDeclaracion_Val(ctx *Declaracion_ValContext) {}

// EnterDeclaracion_Tipo_noVal is called when production Declaracion_Tipo_noVal is entered.
func (s *BaseTswift_GrammarNListener) EnterDeclaracion_Tipo_noVal(ctx *Declaracion_Tipo_noValContext) {
}

// ExitDeclaracion_Tipo_noVal is called when production Declaracion_Tipo_noVal is exited.
func (s *BaseTswift_GrammarNListener) ExitDeclaracion_Tipo_noVal(ctx *Declaracion_Tipo_noValContext) {
}

// EnterAsignacion is called when production Asignacion is entered.
func (s *BaseTswift_GrammarNListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production Asignacion is exited.
func (s *BaseTswift_GrammarNListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterAsignacion_SumaResta is called when production Asignacion_SumaResta is entered.
func (s *BaseTswift_GrammarNListener) EnterAsignacion_SumaResta(ctx *Asignacion_SumaRestaContext) {}

// ExitAsignacion_SumaResta is called when production Asignacion_SumaResta is exited.
func (s *BaseTswift_GrammarNListener) ExitAsignacion_SumaResta(ctx *Asignacion_SumaRestaContext) {}

// EnterTipo_Int is called when production Tipo_Int is entered.
func (s *BaseTswift_GrammarNListener) EnterTipo_Int(ctx *Tipo_IntContext) {}

// ExitTipo_Int is called when production Tipo_Int is exited.
func (s *BaseTswift_GrammarNListener) ExitTipo_Int(ctx *Tipo_IntContext) {}

// EnterTipo_Float is called when production Tipo_Float is entered.
func (s *BaseTswift_GrammarNListener) EnterTipo_Float(ctx *Tipo_FloatContext) {}

// ExitTipo_Float is called when production Tipo_Float is exited.
func (s *BaseTswift_GrammarNListener) ExitTipo_Float(ctx *Tipo_FloatContext) {}

// EnterTipo_String is called when production Tipo_String is entered.
func (s *BaseTswift_GrammarNListener) EnterTipo_String(ctx *Tipo_StringContext) {}

// ExitTipo_String is called when production Tipo_String is exited.
func (s *BaseTswift_GrammarNListener) ExitTipo_String(ctx *Tipo_StringContext) {}

// EnterTipo_Bool is called when production Tipo_Bool is entered.
func (s *BaseTswift_GrammarNListener) EnterTipo_Bool(ctx *Tipo_BoolContext) {}

// ExitTipo_Bool is called when production Tipo_Bool is exited.
func (s *BaseTswift_GrammarNListener) ExitTipo_Bool(ctx *Tipo_BoolContext) {}

// EnterTipo_Character is called when production Tipo_Character is entered.
func (s *BaseTswift_GrammarNListener) EnterTipo_Character(ctx *Tipo_CharacterContext) {}

// ExitTipo_Character is called when production Tipo_Character is exited.
func (s *BaseTswift_GrammarNListener) ExitTipo_Character(ctx *Tipo_CharacterContext) {}

// EnterTipo_Struct is called when production Tipo_Struct is entered.
func (s *BaseTswift_GrammarNListener) EnterTipo_Struct(ctx *Tipo_StructContext) {}

// ExitTipo_Struct is called when production Tipo_Struct is exited.
func (s *BaseTswift_GrammarNListener) ExitTipo_Struct(ctx *Tipo_StructContext) {}

// EnterTipo_Vector is called when production Tipo_Vector is entered.
func (s *BaseTswift_GrammarNListener) EnterTipo_Vector(ctx *Tipo_VectorContext) {}

// ExitTipo_Vector is called when production Tipo_Vector is exited.
func (s *BaseTswift_GrammarNListener) ExitTipo_Vector(ctx *Tipo_VectorContext) {}

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

// EnterForInt is called when production ForInt is entered.
func (s *BaseTswift_GrammarNListener) EnterForInt(ctx *ForIntContext) {}

// ExitForInt is called when production ForInt is exited.
func (s *BaseTswift_GrammarNListener) ExitForInt(ctx *ForIntContext) {}

// EnterForList is called when production ForList is entered.
func (s *BaseTswift_GrammarNListener) EnterForList(ctx *ForListContext) {}

// ExitForList is called when production ForList is exited.
func (s *BaseTswift_GrammarNListener) ExitForList(ctx *ForListContext) {}

// EnterDeclaracion_Vector is called when production Declaracion_Vector is entered.
func (s *BaseTswift_GrammarNListener) EnterDeclaracion_Vector(ctx *Declaracion_VectorContext) {}

// ExitDeclaracion_Vector is called when production Declaracion_Vector is exited.
func (s *BaseTswift_GrammarNListener) ExitDeclaracion_Vector(ctx *Declaracion_VectorContext) {}

// EnterDeclaracion_Alterna is called when production Declaracion_Alterna is entered.
func (s *BaseTswift_GrammarNListener) EnterDeclaracion_Alterna(ctx *Declaracion_AlternaContext) {}

// ExitDeclaracion_Alterna is called when production Declaracion_Alterna is exited.
func (s *BaseTswift_GrammarNListener) ExitDeclaracion_Alterna(ctx *Declaracion_AlternaContext) {}

// EnterDef_Vector is called when production Def_Vector is entered.
func (s *BaseTswift_GrammarNListener) EnterDef_Vector(ctx *Def_VectorContext) {}

// ExitDef_Vector is called when production Def_Vector is exited.
func (s *BaseTswift_GrammarNListener) ExitDef_Vector(ctx *Def_VectorContext) {}

// EnterDef_Vector_Vacio is called when production Def_Vector_Vacio is entered.
func (s *BaseTswift_GrammarNListener) EnterDef_Vector_Vacio(ctx *Def_Vector_VacioContext) {}

// ExitDef_Vector_Vacio is called when production Def_Vector_Vacio is exited.
func (s *BaseTswift_GrammarNListener) ExitDef_Vector_Vacio(ctx *Def_Vector_VacioContext) {}

// EnterDef_Vector_Id is called when production Def_Vector_Id is entered.
func (s *BaseTswift_GrammarNListener) EnterDef_Vector_Id(ctx *Def_Vector_IdContext) {}

// ExitDef_Vector_Id is called when production Def_Vector_Id is exited.
func (s *BaseTswift_GrammarNListener) ExitDef_Vector_Id(ctx *Def_Vector_IdContext) {}

// EnterFunc_Vector_Append is called when production Func_Vector_Append is entered.
func (s *BaseTswift_GrammarNListener) EnterFunc_Vector_Append(ctx *Func_Vector_AppendContext) {}

// ExitFunc_Vector_Append is called when production Func_Vector_Append is exited.
func (s *BaseTswift_GrammarNListener) ExitFunc_Vector_Append(ctx *Func_Vector_AppendContext) {}

// EnterFunc_Vector_RemoveLast is called when production Func_Vector_RemoveLast is entered.
func (s *BaseTswift_GrammarNListener) EnterFunc_Vector_RemoveLast(ctx *Func_Vector_RemoveLastContext) {
}

// ExitFunc_Vector_RemoveLast is called when production Func_Vector_RemoveLast is exited.
func (s *BaseTswift_GrammarNListener) ExitFunc_Vector_RemoveLast(ctx *Func_Vector_RemoveLastContext) {
}

// EnterFunc_Vector_Remove is called when production Func_Vector_Remove is entered.
func (s *BaseTswift_GrammarNListener) EnterFunc_Vector_Remove(ctx *Func_Vector_RemoveContext) {}

// ExitFunc_Vector_Remove is called when production Func_Vector_Remove is exited.
func (s *BaseTswift_GrammarNListener) ExitFunc_Vector_Remove(ctx *Func_Vector_RemoveContext) {}

// EnterDeclaracion_Matriz is called when production Declaracion_Matriz is entered.
func (s *BaseTswift_GrammarNListener) EnterDeclaracion_Matriz(ctx *Declaracion_MatrizContext) {}

// ExitDeclaracion_Matriz is called when production Declaracion_Matriz is exited.
func (s *BaseTswift_GrammarNListener) ExitDeclaracion_Matriz(ctx *Declaracion_MatrizContext) {}

// EnterDef_Matriz_Lista is called when production Def_Matriz_Lista is entered.
func (s *BaseTswift_GrammarNListener) EnterDef_Matriz_Lista(ctx *Def_Matriz_ListaContext) {}

// ExitDef_Matriz_Lista is called when production Def_Matriz_Lista is exited.
func (s *BaseTswift_GrammarNListener) ExitDef_Matriz_Lista(ctx *Def_Matriz_ListaContext) {}

// EnterDef_Matriz_Simple is called when production Def_Matriz_Simple is entered.
func (s *BaseTswift_GrammarNListener) EnterDef_Matriz_Simple(ctx *Def_Matriz_SimpleContext) {}

// ExitDef_Matriz_Simple is called when production Def_Matriz_Simple is exited.
func (s *BaseTswift_GrammarNListener) ExitDef_Matriz_Simple(ctx *Def_Matriz_SimpleContext) {}

// EnterDef_Matriz_Lista_Valores is called when production Def_Matriz_Lista_Valores is entered.
func (s *BaseTswift_GrammarNListener) EnterDef_Matriz_Lista_Valores(ctx *Def_Matriz_Lista_ValoresContext) {
}

// ExitDef_Matriz_Lista_Valores is called when production Def_Matriz_Lista_Valores is exited.
func (s *BaseTswift_GrammarNListener) ExitDef_Matriz_Lista_Valores(ctx *Def_Matriz_Lista_ValoresContext) {
}

// EnterDef_Matriz_Expresion is called when production Def_Matriz_Expresion is entered.
func (s *BaseTswift_GrammarNListener) EnterDef_Matriz_Expresion(ctx *Def_Matriz_ExpresionContext) {}

// ExitDef_Matriz_Expresion is called when production Def_Matriz_Expresion is exited.
func (s *BaseTswift_GrammarNListener) ExitDef_Matriz_Expresion(ctx *Def_Matriz_ExpresionContext) {}

// EnterDef_Matriz_Simple2 is called when production Def_Matriz_Simple2 is entered.
func (s *BaseTswift_GrammarNListener) EnterDef_Matriz_Simple2(ctx *Def_Matriz_Simple2Context) {}

// ExitDef_Matriz_Simple2 is called when production Def_Matriz_Simple2 is exited.
func (s *BaseTswift_GrammarNListener) ExitDef_Matriz_Simple2(ctx *Def_Matriz_Simple2Context) {}

// EnterDef_Matriz_Simple3 is called when production Def_Matriz_Simple3 is entered.
func (s *BaseTswift_GrammarNListener) EnterDef_Matriz_Simple3(ctx *Def_Matriz_Simple3Context) {}

// ExitDef_Matriz_Simple3 is called when production Def_Matriz_Simple3 is exited.
func (s *BaseTswift_GrammarNListener) ExitDef_Matriz_Simple3(ctx *Def_Matriz_Simple3Context) {}

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

// EnterExpr_Rel is called when production Expr_Rel is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Rel(ctx *Expr_RelContext) {}

// ExitExpr_Rel is called when production Expr_Rel is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Rel(ctx *Expr_RelContext) {}

// EnterExpr_Decimal is called when production Expr_Decimal is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Decimal(ctx *Expr_DecimalContext) {}

// ExitExpr_Decimal is called when production Expr_Decimal is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Decimal(ctx *Expr_DecimalContext) {}

// EnterExpr_Caracter is called when production Expr_Caracter is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Caracter(ctx *Expr_CaracterContext) {}

// ExitExpr_Caracter is called when production Expr_Caracter is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Caracter(ctx *Expr_CaracterContext) {}

// EnterExpr_SumRes is called when production Expr_SumRes is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_SumRes(ctx *Expr_SumResContext) {}

// ExitExpr_SumRes is called when production Expr_SumRes is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_SumRes(ctx *Expr_SumResContext) {}

// EnterExpr_Matriz is called when production Expr_Matriz is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Matriz(ctx *Expr_MatrizContext) {}

// ExitExpr_Matriz is called when production Expr_Matriz is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Matriz(ctx *Expr_MatrizContext) {}

// EnterExpr_Neg is called when production Expr_Neg is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Neg(ctx *Expr_NegContext) {}

// ExitExpr_Neg is called when production Expr_Neg is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Neg(ctx *Expr_NegContext) {}

// EnterExpr_MulDiv is called when production Expr_MulDiv is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_MulDiv(ctx *Expr_MulDivContext) {}

// ExitExpr_MulDiv is called when production Expr_MulDiv is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_MulDiv(ctx *Expr_MulDivContext) {}

// EnterExpr_Nil is called when production Expr_Nil is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Nil(ctx *Expr_NilContext) {}

// ExitExpr_Nil is called when production Expr_Nil is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Nil(ctx *Expr_NilContext) {}

// EnterExpr_Cadena is called when production Expr_Cadena is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Cadena(ctx *Expr_CadenaContext) {}

// ExitExpr_Cadena is called when production Expr_Cadena is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Cadena(ctx *Expr_CadenaContext) {}

// EnterExpr_Count is called when production Expr_Count is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Count(ctx *Expr_CountContext) {}

// ExitExpr_Count is called when production Expr_Count is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Count(ctx *Expr_CountContext) {}

// EnterExpr_Id is called when production Expr_Id is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Id(ctx *Expr_IdContext) {}

// ExitExpr_Id is called when production Expr_Id is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Id(ctx *Expr_IdContext) {}

// EnterExpr_Mod is called when production Expr_Mod is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Mod(ctx *Expr_ModContext) {}

// ExitExpr_Mod is called when production Expr_Mod is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Mod(ctx *Expr_ModContext) {}

// EnterExpr_Par is called when production Expr_Par is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Par(ctx *Expr_ParContext) {}

// ExitExpr_Par is called when production Expr_Par is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Par(ctx *Expr_ParContext) {}

// EnterExpr_Logica is called when production Expr_Logica is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Logica(ctx *Expr_LogicaContext) {}

// ExitExpr_Logica is called when production Expr_Logica is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Logica(ctx *Expr_LogicaContext) {}

// EnterExpr_IsEmpty is called when production Expr_IsEmpty is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_IsEmpty(ctx *Expr_IsEmptyContext) {}

// ExitExpr_IsEmpty is called when production Expr_IsEmpty is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_IsEmpty(ctx *Expr_IsEmptyContext) {}

// EnterExpr_Booleano is called when production Expr_Booleano is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Booleano(ctx *Expr_BooleanoContext) {}

// ExitExpr_Booleano is called when production Expr_Booleano is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Booleano(ctx *Expr_BooleanoContext) {}

// EnterExpr_Vector is called when production Expr_Vector is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Vector(ctx *Expr_VectorContext) {}

// ExitExpr_Vector is called when production Expr_Vector is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Vector(ctx *Expr_VectorContext) {}

// EnterExpr_Entero is called when production Expr_Entero is entered.
func (s *BaseTswift_GrammarNListener) EnterExpr_Entero(ctx *Expr_EnteroContext) {}

// ExitExpr_Entero is called when production Expr_Entero is exited.
func (s *BaseTswift_GrammarNListener) ExitExpr_Entero(ctx *Expr_EnteroContext) {}
