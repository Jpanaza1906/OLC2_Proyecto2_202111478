grammar Tswift_GrammarN;
import Tswift_Lexer;

//Produccion inicial de donde salen las sentencias--------------------------------

s: l_sentencias #SLSentencias
    ;

//Lista de sentencias-------------------------------------------------------------
l_sentencias:
    sentencia* #L_Sentencia
    ;

//Sentencias---------------------------------------------------------------------
sentencia
    :print_sentencia PTCOMA? #S_Print
    |condicion PTCOMA? #S_Condicion
    |if_sentencia #S_If
    |switch_sentencia #S_Switch
    ;

//Sentencia print---------------------------------------------------------------
print_sentencia:
    PRINT PARIZQ e (COMA e)* PARDER #Print
    ;

//Sentencia if-------------------------------------------------------------------
if_sentencia:
    IF condicion LLAVEIZQ l_sentencias LLAVEDER (ELSE (if_sentencia | LLAVEIZQ l_sentencias LLAVEDER))? #If
    ;

//Sentencia Switch---------------------------------------------------------------
switch_sentencia:
    SWITCH e LLAVEIZQ lcasos* cdefault? LLAVEDER #Switch
    ;

lcasos:
    CASE e DOSPT l_sentencias #Case
    ;

cdefault:
    DEFAULT DOSPT l_sentencias #Default
    ;


//Condiciones--------------------------------------------------------------------

condicion
    : <assoc=right> op=NOT c=condicion      # Cond_Neg
    | e1=e op=(IGUALIGUAL | DIFERENTE | MAYORIGUAL | MAYOR | MENORIGUAL | MENOR) e2=e     # Cond_Rel
    | c1=condicion op=(AND | OR) c2=condicion     # Cond_Logica
    | op=(TRUE | FALSE)          # Cond_Booleano
    | PARIZQ c=condicion PARDER     # Cond_Par
    ;

//Expresiones--------------------------------------------------------------------
e
    : <assoc=right> op=MENOS e1=e      # Expr_Neg
    | e1=e op=(POR | DIV) e2=e    # Expr_MulDiv
    | e1=e op=(MAS | MENOS) e2=e  # Expr_SumRes
    | e1=e op=MOD e2=e           # Expr_Mod
    | id=ID                # Expr_Id
    | n=ENTERO            # Expr_Entero
    | PARIZQ e1=e PARDER   # Expr_Par
    ;