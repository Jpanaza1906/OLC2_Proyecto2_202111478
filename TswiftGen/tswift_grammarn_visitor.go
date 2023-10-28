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

	// Visit a parse tree produced by Tswift_GrammarNParser#S_Declaracion.
	VisitS_Declaracion(ctx *S_DeclaracionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#S_Asignacion.
	VisitS_Asignacion(ctx *S_AsignacionContext) interface{}

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

	// Visit a parse tree produced by Tswift_GrammarNParser#S_Declaracion_Vector.
	VisitS_Declaracion_Vector(ctx *S_Declaracion_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#S_Funcion_Vector.
	VisitS_Funcion_Vector(ctx *S_Funcion_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Break.
	VisitBreak(ctx *BreakContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Continue.
	VisitContinue(ctx *ContinueContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Return.
	VisitReturn(ctx *ReturnContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Print.
	VisitPrint(ctx *PrintContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Tipo_Val.
	VisitDeclaracion_Tipo_Val(ctx *Declaracion_Tipo_ValContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Val.
	VisitDeclaracion_Val(ctx *Declaracion_ValContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Tipo_noVal.
	VisitDeclaracion_Tipo_noVal(ctx *Declaracion_Tipo_noValContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Asignacion.
	VisitAsignacion(ctx *AsignacionContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Asignacion_SumaResta.
	VisitAsignacion_SumaResta(ctx *Asignacion_SumaRestaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Int.
	VisitTipo_Int(ctx *Tipo_IntContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Float.
	VisitTipo_Float(ctx *Tipo_FloatContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_String.
	VisitTipo_String(ctx *Tipo_StringContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Bool.
	VisitTipo_Bool(ctx *Tipo_BoolContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Character.
	VisitTipo_Character(ctx *Tipo_CharacterContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Struct.
	VisitTipo_Struct(ctx *Tipo_StructContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Tipo_Vector.
	VisitTipo_Vector(ctx *Tipo_VectorContext) interface{}

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

	// Visit a parse tree produced by Tswift_GrammarNParser#ForInt.
	VisitForInt(ctx *ForIntContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#ForList.
	VisitForList(ctx *ForListContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Vector.
	VisitDeclaracion_Vector(ctx *Declaracion_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Declaracion_Alterna.
	VisitDeclaracion_Alterna(ctx *Declaracion_AlternaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Def_Vector.
	VisitDef_Vector(ctx *Def_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Def_Vector_Vacio.
	VisitDef_Vector_Vacio(ctx *Def_Vector_VacioContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Def_Vector_Id.
	VisitDef_Vector_Id(ctx *Def_Vector_IdContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Func_Vector_Append.
	VisitFunc_Vector_Append(ctx *Func_Vector_AppendContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Func_Vector_RemoveLast.
	VisitFunc_Vector_RemoveLast(ctx *Func_Vector_RemoveLastContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Func_Vector_Remove.
	VisitFunc_Vector_Remove(ctx *Func_Vector_RemoveContext) interface{}

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

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Rel.
	VisitExpr_Rel(ctx *Expr_RelContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Decimal.
	VisitExpr_Decimal(ctx *Expr_DecimalContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Caracter.
	VisitExpr_Caracter(ctx *Expr_CaracterContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_SumRes.
	VisitExpr_SumRes(ctx *Expr_SumResContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Neg.
	VisitExpr_Neg(ctx *Expr_NegContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_MulDiv.
	VisitExpr_MulDiv(ctx *Expr_MulDivContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Nil.
	VisitExpr_Nil(ctx *Expr_NilContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Cadena.
	VisitExpr_Cadena(ctx *Expr_CadenaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Count.
	VisitExpr_Count(ctx *Expr_CountContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Id.
	VisitExpr_Id(ctx *Expr_IdContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Mod.
	VisitExpr_Mod(ctx *Expr_ModContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Par.
	VisitExpr_Par(ctx *Expr_ParContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Logica.
	VisitExpr_Logica(ctx *Expr_LogicaContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_IsEmpty.
	VisitExpr_IsEmpty(ctx *Expr_IsEmptyContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Booleano.
	VisitExpr_Booleano(ctx *Expr_BooleanoContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Vector.
	VisitExpr_Vector(ctx *Expr_VectorContext) interface{}

	// Visit a parse tree produced by Tswift_GrammarNParser#Expr_Entero.
	VisitExpr_Entero(ctx *Expr_EnteroContext) interface{}
}
