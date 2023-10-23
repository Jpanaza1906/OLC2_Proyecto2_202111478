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
	 * Enter a parse tree produced by the {@code For}
	 * labeled alternative in {@link Tswift_GrammarNParser#for_sentencia}.
	 * @param ctx the parse tree
	 */
	void enterFor(Tswift_GrammarNParser.ForContext ctx);
	/**
	 * Exit a parse tree produced by the {@code For}
	 * labeled alternative in {@link Tswift_GrammarNParser#for_sentencia}.
	 * @param ctx the parse tree
	 */
	void exitFor(Tswift_GrammarNParser.ForContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Rango}
	 * labeled alternative in {@link Tswift_GrammarNParser#rango_p}.
	 * @param ctx the parse tree
	 */
	void enterRango(Tswift_GrammarNParser.RangoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Rango}
	 * labeled alternative in {@link Tswift_GrammarNParser#rango_p}.
	 * @param ctx the parse tree
	 */
	void exitRango(Tswift_GrammarNParser.RangoContext ctx);
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