// Generated from /home/josep/USAC/6to_Semestre/Lab_Compi/Proyecto2/OLC2_Proyecto2_202111478/parser/Tswift_GrammarN.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link Tswift_GrammarNParser}.
 */
public interface Tswift_GrammarNListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by the {@code SLSentencias}
	 * labeled alternative in {@link Tswift_GrammarNParser#s}.
	 * @param ctx the parse tree
	 */
	void enterSLSentencias(Tswift_GrammarNParser.SLSentenciasContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SLSentencias}
	 * labeled alternative in {@link Tswift_GrammarNParser#s}.
	 * @param ctx the parse tree
	 */
	void exitSLSentencias(Tswift_GrammarNParser.SLSentenciasContext ctx);
	/**
	 * Enter a parse tree produced by the {@code L_Sentencia}
	 * labeled alternative in {@link Tswift_GrammarNParser#l_sentencias}.
	 * @param ctx the parse tree
	 */
	void enterL_Sentencia(Tswift_GrammarNParser.L_SentenciaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code L_Sentencia}
	 * labeled alternative in {@link Tswift_GrammarNParser#l_sentencias}.
	 * @param ctx the parse tree
	 */
	void exitL_Sentencia(Tswift_GrammarNParser.L_SentenciaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_Print}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_Print(Tswift_GrammarNParser.S_PrintContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_Print}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_Print(Tswift_GrammarNParser.S_PrintContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_Declaracion}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_Declaracion(Tswift_GrammarNParser.S_DeclaracionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_Declaracion}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_Declaracion(Tswift_GrammarNParser.S_DeclaracionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_Asignacion}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_Asignacion(Tswift_GrammarNParser.S_AsignacionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_Asignacion}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_Asignacion(Tswift_GrammarNParser.S_AsignacionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_If}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_If(Tswift_GrammarNParser.S_IfContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_If}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_If(Tswift_GrammarNParser.S_IfContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_Switch}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_Switch(Tswift_GrammarNParser.S_SwitchContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_Switch}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_Switch(Tswift_GrammarNParser.S_SwitchContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_Guard}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_Guard(Tswift_GrammarNParser.S_GuardContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_Guard}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_Guard(Tswift_GrammarNParser.S_GuardContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_Trans}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_Trans(Tswift_GrammarNParser.S_TransContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_Trans}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_Trans(Tswift_GrammarNParser.S_TransContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_While}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_While(Tswift_GrammarNParser.S_WhileContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_While}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_While(Tswift_GrammarNParser.S_WhileContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_For}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_For(Tswift_GrammarNParser.S_ForContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_For}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_For(Tswift_GrammarNParser.S_ForContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_Declaracion_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_Declaracion_Vector(Tswift_GrammarNParser.S_Declaracion_VectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_Declaracion_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_Declaracion_Vector(Tswift_GrammarNParser.S_Declaracion_VectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_Funcion_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_Funcion_Vector(Tswift_GrammarNParser.S_Funcion_VectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_Funcion_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_Funcion_Vector(Tswift_GrammarNParser.S_Funcion_VectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code S_Declaracion_Matriz}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void enterS_Declaracion_Matriz(Tswift_GrammarNParser.S_Declaracion_MatrizContext ctx);
	/**
	 * Exit a parse tree produced by the {@code S_Declaracion_Matriz}
	 * labeled alternative in {@link Tswift_GrammarNParser#sentencia}.
	 * @param ctx the parse tree
	 */
	void exitS_Declaracion_Matriz(Tswift_GrammarNParser.S_Declaracion_MatrizContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Break}
	 * labeled alternative in {@link Tswift_GrammarNParser#trans_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterBreak(Tswift_GrammarNParser.BreakContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Break}
	 * labeled alternative in {@link Tswift_GrammarNParser#trans_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitBreak(Tswift_GrammarNParser.BreakContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Continue}
	 * labeled alternative in {@link Tswift_GrammarNParser#trans_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterContinue(Tswift_GrammarNParser.ContinueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Continue}
	 * labeled alternative in {@link Tswift_GrammarNParser#trans_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitContinue(Tswift_GrammarNParser.ContinueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Return}
	 * labeled alternative in {@link Tswift_GrammarNParser#trans_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterReturn(Tswift_GrammarNParser.ReturnContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Return}
	 * labeled alternative in {@link Tswift_GrammarNParser#trans_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitReturn(Tswift_GrammarNParser.ReturnContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Print}
	 * labeled alternative in {@link Tswift_GrammarNParser#print_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterPrint(Tswift_GrammarNParser.PrintContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Print}
	 * labeled alternative in {@link Tswift_GrammarNParser#print_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitPrint(Tswift_GrammarNParser.PrintContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Declaracion_Tipo_Val}
	 * labeled alternative in {@link Tswift_GrammarNParser#declaracion_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracion_Tipo_Val(Tswift_GrammarNParser.Declaracion_Tipo_ValContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Declaracion_Tipo_Val}
	 * labeled alternative in {@link Tswift_GrammarNParser#declaracion_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracion_Tipo_Val(Tswift_GrammarNParser.Declaracion_Tipo_ValContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Declaracion_Val}
	 * labeled alternative in {@link Tswift_GrammarNParser#declaracion_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracion_Val(Tswift_GrammarNParser.Declaracion_ValContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Declaracion_Val}
	 * labeled alternative in {@link Tswift_GrammarNParser#declaracion_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracion_Val(Tswift_GrammarNParser.Declaracion_ValContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Declaracion_Tipo_noVal}
	 * labeled alternative in {@link Tswift_GrammarNParser#declaracion_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracion_Tipo_noVal(Tswift_GrammarNParser.Declaracion_Tipo_noValContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Declaracion_Tipo_noVal}
	 * labeled alternative in {@link Tswift_GrammarNParser#declaracion_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracion_Tipo_noVal(Tswift_GrammarNParser.Declaracion_Tipo_noValContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion}
	 * labeled alternative in {@link Tswift_GrammarNParser#asignacion_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion(Tswift_GrammarNParser.AsignacionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion}
	 * labeled alternative in {@link Tswift_GrammarNParser#asignacion_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion(Tswift_GrammarNParser.AsignacionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Asignacion_SumaResta}
	 * labeled alternative in {@link Tswift_GrammarNParser#asignacion_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterAsignacion_SumaResta(Tswift_GrammarNParser.Asignacion_SumaRestaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Asignacion_SumaResta}
	 * labeled alternative in {@link Tswift_GrammarNParser#asignacion_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitAsignacion_SumaResta(Tswift_GrammarNParser.Asignacion_SumaRestaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Tipo_Int}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo_Int(Tswift_GrammarNParser.Tipo_IntContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Tipo_Int}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo_Int(Tswift_GrammarNParser.Tipo_IntContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Tipo_Float}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo_Float(Tswift_GrammarNParser.Tipo_FloatContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Tipo_Float}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo_Float(Tswift_GrammarNParser.Tipo_FloatContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Tipo_String}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo_String(Tswift_GrammarNParser.Tipo_StringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Tipo_String}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo_String(Tswift_GrammarNParser.Tipo_StringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Tipo_Bool}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo_Bool(Tswift_GrammarNParser.Tipo_BoolContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Tipo_Bool}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo_Bool(Tswift_GrammarNParser.Tipo_BoolContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Tipo_Character}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo_Character(Tswift_GrammarNParser.Tipo_CharacterContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Tipo_Character}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo_Character(Tswift_GrammarNParser.Tipo_CharacterContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Tipo_Struct}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo_Struct(Tswift_GrammarNParser.Tipo_StructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Tipo_Struct}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo_Struct(Tswift_GrammarNParser.Tipo_StructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Tipo_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void enterTipo_Vector(Tswift_GrammarNParser.Tipo_VectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Tipo_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#tipo}.
	 * @param ctx the parse tree
	 */
	void exitTipo_Vector(Tswift_GrammarNParser.Tipo_VectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code If}
	 * labeled alternative in {@link Tswift_GrammarNParser#if_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterIf(Tswift_GrammarNParser.IfContext ctx);
	/**
	 * Exit a parse tree produced by the {@code If}
	 * labeled alternative in {@link Tswift_GrammarNParser#if_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitIf(Tswift_GrammarNParser.IfContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Guard}
	 * labeled alternative in {@link Tswift_GrammarNParser#guard_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterGuard(Tswift_GrammarNParser.GuardContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Guard}
	 * labeled alternative in {@link Tswift_GrammarNParser#guard_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitGuard(Tswift_GrammarNParser.GuardContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Switch}
	 * labeled alternative in {@link Tswift_GrammarNParser#switch_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterSwitch(Tswift_GrammarNParser.SwitchContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Switch}
	 * labeled alternative in {@link Tswift_GrammarNParser#switch_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitSwitch(Tswift_GrammarNParser.SwitchContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Case}
	 * labeled alternative in {@link Tswift_GrammarNParser#lcasos}.
	 * @param ctx the parse tree
	 */
	void enterCase(Tswift_GrammarNParser.CaseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Case}
	 * labeled alternative in {@link Tswift_GrammarNParser#lcasos}.
	 * @param ctx the parse tree
	 */
	void exitCase(Tswift_GrammarNParser.CaseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Default}
	 * labeled alternative in {@link Tswift_GrammarNParser#cdefault}.
	 * @param ctx the parse tree
	 */
	void enterDefault(Tswift_GrammarNParser.DefaultContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Default}
	 * labeled alternative in {@link Tswift_GrammarNParser#cdefault}.
	 * @param ctx the parse tree
	 */
	void exitDefault(Tswift_GrammarNParser.DefaultContext ctx);
	/**
	 * Enter a parse tree produced by the {@code While}
	 * labeled alternative in {@link Tswift_GrammarNParser#while_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterWhile(Tswift_GrammarNParser.WhileContext ctx);
	/**
	 * Exit a parse tree produced by the {@code While}
	 * labeled alternative in {@link Tswift_GrammarNParser#while_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitWhile(Tswift_GrammarNParser.WhileContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForInt}
	 * labeled alternative in {@link Tswift_GrammarNParser#for_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterForInt(Tswift_GrammarNParser.ForIntContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForInt}
	 * labeled alternative in {@link Tswift_GrammarNParser#for_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitForInt(Tswift_GrammarNParser.ForIntContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForList}
	 * labeled alternative in {@link Tswift_GrammarNParser#for_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterForList(Tswift_GrammarNParser.ForListContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForList}
	 * labeled alternative in {@link Tswift_GrammarNParser#for_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitForList(Tswift_GrammarNParser.ForListContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Declaracion_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#dec_vector}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracion_Vector(Tswift_GrammarNParser.Declaracion_VectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Declaracion_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#dec_vector}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracion_Vector(Tswift_GrammarNParser.Declaracion_VectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Declaracion_Alterna}
	 * labeled alternative in {@link Tswift_GrammarNParser#dec_vector}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracion_Alterna(Tswift_GrammarNParser.Declaracion_AlternaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Declaracion_Alterna}
	 * labeled alternative in {@link Tswift_GrammarNParser#dec_vector}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracion_Alterna(Tswift_GrammarNParser.Declaracion_AlternaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_vector}.
	 * @param ctx the parse tree
	 */
	void enterDef_Vector(Tswift_GrammarNParser.Def_VectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_vector}.
	 * @param ctx the parse tree
	 */
	void exitDef_Vector(Tswift_GrammarNParser.Def_VectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Vector_Vacio}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_vector}.
	 * @param ctx the parse tree
	 */
	void enterDef_Vector_Vacio(Tswift_GrammarNParser.Def_Vector_VacioContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Vector_Vacio}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_vector}.
	 * @param ctx the parse tree
	 */
	void exitDef_Vector_Vacio(Tswift_GrammarNParser.Def_Vector_VacioContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Vector_Id}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_vector}.
	 * @param ctx the parse tree
	 */
	void enterDef_Vector_Id(Tswift_GrammarNParser.Def_Vector_IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Vector_Id}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_vector}.
	 * @param ctx the parse tree
	 */
	void exitDef_Vector_Id(Tswift_GrammarNParser.Def_Vector_IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Func_Vector_Append}
	 * labeled alternative in {@link Tswift_GrammarNParser#func_vector}.
	 * @param ctx the parse tree
	 */
	void enterFunc_Vector_Append(Tswift_GrammarNParser.Func_Vector_AppendContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Func_Vector_Append}
	 * labeled alternative in {@link Tswift_GrammarNParser#func_vector}.
	 * @param ctx the parse tree
	 */
	void exitFunc_Vector_Append(Tswift_GrammarNParser.Func_Vector_AppendContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Func_Vector_RemoveLast}
	 * labeled alternative in {@link Tswift_GrammarNParser#func_vector}.
	 * @param ctx the parse tree
	 */
	void enterFunc_Vector_RemoveLast(Tswift_GrammarNParser.Func_Vector_RemoveLastContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Func_Vector_RemoveLast}
	 * labeled alternative in {@link Tswift_GrammarNParser#func_vector}.
	 * @param ctx the parse tree
	 */
	void exitFunc_Vector_RemoveLast(Tswift_GrammarNParser.Func_Vector_RemoveLastContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Func_Vector_Remove}
	 * labeled alternative in {@link Tswift_GrammarNParser#func_vector}.
	 * @param ctx the parse tree
	 */
	void enterFunc_Vector_Remove(Tswift_GrammarNParser.Func_Vector_RemoveContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Func_Vector_Remove}
	 * labeled alternative in {@link Tswift_GrammarNParser#func_vector}.
	 * @param ctx the parse tree
	 */
	void exitFunc_Vector_Remove(Tswift_GrammarNParser.Func_Vector_RemoveContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Declaracion_Matriz}
	 * labeled alternative in {@link Tswift_GrammarNParser#dec_matriz}.
	 * @param ctx the parse tree
	 */
	void enterDeclaracion_Matriz(Tswift_GrammarNParser.Declaracion_MatrizContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Declaracion_Matriz}
	 * labeled alternative in {@link Tswift_GrammarNParser#dec_matriz}.
	 * @param ctx the parse tree
	 */
	void exitDeclaracion_Matriz(Tswift_GrammarNParser.Declaracion_MatrizContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Matriz}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_matriz}.
	 * @param ctx the parse tree
	 */
	void enterDef_Matriz(Tswift_GrammarNParser.Def_MatrizContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Matriz}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_matriz}.
	 * @param ctx the parse tree
	 */
	void exitDef_Matriz(Tswift_GrammarNParser.Def_MatrizContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Matriz_Simple}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_matriz}.
	 * @param ctx the parse tree
	 */
	void enterDef_Matriz_Simple(Tswift_GrammarNParser.Def_Matriz_SimpleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Matriz_Simple}
	 * labeled alternative in {@link Tswift_GrammarNParser#def_matriz}.
	 * @param ctx the parse tree
	 */
	void exitDef_Matriz_Simple(Tswift_GrammarNParser.Def_Matriz_SimpleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Matriz_Valor}
	 * labeled alternative in {@link Tswift_GrammarNParser#listavalores_matriz}.
	 * @param ctx the parse tree
	 */
	void enterDef_Matriz_Valor(Tswift_GrammarNParser.Def_Matriz_ValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Matriz_Valor}
	 * labeled alternative in {@link Tswift_GrammarNParser#listavalores_matriz}.
	 * @param ctx the parse tree
	 */
	void exitDef_Matriz_Valor(Tswift_GrammarNParser.Def_Matriz_ValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Matriz_Valor4}
	 * labeled alternative in {@link Tswift_GrammarNParser#listavalores_matriz2}.
	 * @param ctx the parse tree
	 */
	void enterDef_Matriz_Valor4(Tswift_GrammarNParser.Def_Matriz_Valor4Context ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Matriz_Valor4}
	 * labeled alternative in {@link Tswift_GrammarNParser#listavalores_matriz2}.
	 * @param ctx the parse tree
	 */
	void exitDef_Matriz_Valor4(Tswift_GrammarNParser.Def_Matriz_Valor4Context ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Matriz_Valor3}
	 * labeled alternative in {@link Tswift_GrammarNParser#listavalores_matriz2}.
	 * @param ctx the parse tree
	 */
	void enterDef_Matriz_Valor3(Tswift_GrammarNParser.Def_Matriz_Valor3Context ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Matriz_Valor3}
	 * labeled alternative in {@link Tswift_GrammarNParser#listavalores_matriz2}.
	 * @param ctx the parse tree
	 */
	void exitDef_Matriz_Valor3(Tswift_GrammarNParser.Def_Matriz_Valor3Context ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Matriz_Valor2}
	 * labeled alternative in {@link Tswift_GrammarNParser#listavalores_matriz2}.
	 * @param ctx the parse tree
	 */
	void enterDef_Matriz_Valor2(Tswift_GrammarNParser.Def_Matriz_Valor2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Matriz_Valor2}
	 * labeled alternative in {@link Tswift_GrammarNParser#listavalores_matriz2}.
	 * @param ctx the parse tree
	 */
	void exitDef_Matriz_Valor2(Tswift_GrammarNParser.Def_Matriz_Valor2Context ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Matriz_Simple2}
	 * labeled alternative in {@link Tswift_GrammarNParser#simple_matriz}.
	 * @param ctx the parse tree
	 */
	void enterDef_Matriz_Simple2(Tswift_GrammarNParser.Def_Matriz_Simple2Context ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Matriz_Simple2}
	 * labeled alternative in {@link Tswift_GrammarNParser#simple_matriz}.
	 * @param ctx the parse tree
	 */
	void exitDef_Matriz_Simple2(Tswift_GrammarNParser.Def_Matriz_Simple2Context ctx);
	/**
	 * Enter a parse tree produced by the {@code Def_Matriz_Simple3}
	 * labeled alternative in {@link Tswift_GrammarNParser#simple_matriz}.
	 * @param ctx the parse tree
	 */
	void enterDef_Matriz_Simple3(Tswift_GrammarNParser.Def_Matriz_Simple3Context ctx);
	/**
	 * Exit a parse tree produced by the {@code Def_Matriz_Simple3}
	 * labeled alternative in {@link Tswift_GrammarNParser#simple_matriz}.
	 * @param ctx the parse tree
	 */
	void exitDef_Matriz_Simple3(Tswift_GrammarNParser.Def_Matriz_Simple3Context ctx);
	/**
	 * Enter a parse tree produced by the {@code Cond_Par}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void enterCond_Par(Tswift_GrammarNParser.Cond_ParContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Cond_Par}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void exitCond_Par(Tswift_GrammarNParser.Cond_ParContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Cond_Neg}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void enterCond_Neg(Tswift_GrammarNParser.Cond_NegContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Cond_Neg}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void exitCond_Neg(Tswift_GrammarNParser.Cond_NegContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Cond_Rel}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void enterCond_Rel(Tswift_GrammarNParser.Cond_RelContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Cond_Rel}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void exitCond_Rel(Tswift_GrammarNParser.Cond_RelContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Cond_Booleano}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void enterCond_Booleano(Tswift_GrammarNParser.Cond_BooleanoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Cond_Booleano}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void exitCond_Booleano(Tswift_GrammarNParser.Cond_BooleanoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Cond_Logica}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void enterCond_Logica(Tswift_GrammarNParser.Cond_LogicaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Cond_Logica}
	 * labeled alternative in {@link Tswift_GrammarNParser#condicion}.
	 * @param ctx the parse tree
	 */
	void exitCond_Logica(Tswift_GrammarNParser.Cond_LogicaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Rel}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Rel(Tswift_GrammarNParser.Expr_RelContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Rel}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Rel(Tswift_GrammarNParser.Expr_RelContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Decimal}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Decimal(Tswift_GrammarNParser.Expr_DecimalContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Decimal}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Decimal(Tswift_GrammarNParser.Expr_DecimalContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Caracter}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Caracter(Tswift_GrammarNParser.Expr_CaracterContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Caracter}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Caracter(Tswift_GrammarNParser.Expr_CaracterContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_SumRes}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_SumRes(Tswift_GrammarNParser.Expr_SumResContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_SumRes}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_SumRes(Tswift_GrammarNParser.Expr_SumResContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Neg}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Neg(Tswift_GrammarNParser.Expr_NegContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Neg}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Neg(Tswift_GrammarNParser.Expr_NegContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_MulDiv}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_MulDiv(Tswift_GrammarNParser.Expr_MulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_MulDiv}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_MulDiv(Tswift_GrammarNParser.Expr_MulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Nil}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Nil(Tswift_GrammarNParser.Expr_NilContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Nil}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Nil(Tswift_GrammarNParser.Expr_NilContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Cadena}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Cadena(Tswift_GrammarNParser.Expr_CadenaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Cadena}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Cadena(Tswift_GrammarNParser.Expr_CadenaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Count}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Count(Tswift_GrammarNParser.Expr_CountContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Count}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Count(Tswift_GrammarNParser.Expr_CountContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Id}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Id(Tswift_GrammarNParser.Expr_IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Id}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Id(Tswift_GrammarNParser.Expr_IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Mod}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Mod(Tswift_GrammarNParser.Expr_ModContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Mod}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Mod(Tswift_GrammarNParser.Expr_ModContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Par}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Par(Tswift_GrammarNParser.Expr_ParContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Par}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Par(Tswift_GrammarNParser.Expr_ParContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Logica}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Logica(Tswift_GrammarNParser.Expr_LogicaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Logica}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Logica(Tswift_GrammarNParser.Expr_LogicaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_IsEmpty}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_IsEmpty(Tswift_GrammarNParser.Expr_IsEmptyContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_IsEmpty}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_IsEmpty(Tswift_GrammarNParser.Expr_IsEmptyContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Booleano}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Booleano(Tswift_GrammarNParser.Expr_BooleanoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Booleano}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Booleano(Tswift_GrammarNParser.Expr_BooleanoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Vector(Tswift_GrammarNParser.Expr_VectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Vector}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Vector(Tswift_GrammarNParser.Expr_VectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Expr_Entero}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void enterExpr_Entero(Tswift_GrammarNParser.Expr_EnteroContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Expr_Entero}
	 * labeled alternative in {@link Tswift_GrammarNParser#e}.
	 * @param ctx the parse tree
	 */
	void exitExpr_Entero(Tswift_GrammarNParser.Expr_EnteroContext ctx);
}