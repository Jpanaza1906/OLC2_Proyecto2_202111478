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
    |if_sentencia #S_If
    |switch_sentencia #S_Switch
    |guard_sentencia #S_Guard
    |trans_sentencia PTCOMA? #S_Trans
    |while_sentencia #S_While
    |for_sentencia #S_For
    ;

//Sentencias de transferencias---------------------------------------------------
trans_sentencia:
    BREAK #Break
    |CONTINUE  #Continue
    |RETURN (e)? #Return
    ;

//Sentencia print---------------------------------------------------------------
print_sentencia:
    PRINT PARIZQ e (COMA e)* PARDER #Print
    ;

//Sentencia if-------------------------------------------------------------------
if_sentencia:
    IF condicion LLAVEIZQ l_sentencias LLAVEDER (ELSE (if_sentencia | LLAVEIZQ l_sentencias LLAVEDER))? #If
    ;

//Sentencia Guard----------------------------------------------------------------
guard_sentencia:
    GUARD condicion ELSE LLAVEIZQ l_sentencias trans_sentencia LLAVEDER #Guard
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

//Sentencia While----------------------------------------------------------------
while_sentencia:
    WHILE condicion LLAVEIZQ l_sentencias LLAVEDER #While
    ;

//Sentencia For------------------------------------------------------------------
for_sentencia:
    FOR id=ID IN (rango_p|e) LLAVEIZQ l_sentencias LLAVEDER #For
    ;
rango_p:
    e RANGO e #Rango
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