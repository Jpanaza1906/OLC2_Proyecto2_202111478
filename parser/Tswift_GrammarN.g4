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
    |declaracion_sentencia PTCOMA? #S_Declaracion
    |asignacion_sentencia PTCOMA? #S_Asignacion
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

//Declaración de variables-------------------------------------------------------
declaracion_sentencia:
    tip=(VAR | LET) ID DOSPT tipo IGUAL e #Declaracion_Tipo_Val
    | tip=(VAR | LET) ID IGUAL e #Declaracion_Val
    | tip=(VAR | LET) ID DOSPT tipo INTERROGACION #Declaracion_Tipo_noVal
    ;

//Asignación de variables--------------------------------------------------------
asignacion_sentencia:
    ID IGUAL e #Asignacion
    | ID op=(MASIGUAL|MENOSIGUAL) e #Asignacion_SumaResta
    ;

//Tipos de variables-------------------------------------------------------------
tipo:
    INT #Tipo_Int
    |FLOAT #Tipo_Float
    |STRING #Tipo_String
    |BOOL #Tipo_Bool
    |CHAR #Tipo_Character
    | ID #Tipo_Struct
    |CORCHETEIZQ tipo CORCHETEDER #Tipo_Vector
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
    FOR ID IN left=e RANGO right=e LLAVEIZQ l_sentencias LLAVEDER #ForInt
    | FOR ID IN e LLAVEIZQ l_sentencias LLAVEDER #ForList
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
    | n=(TRUE | FALSE)    # Expr_Booleano
    | n=NIL              # Expr_Nil
    | id=ID                # Expr_Id
    | n=ENTERO            # Expr_Entero
    | n=DECIMAL           # Expr_Decimal
    | n=CADENA            # Expr_Cadena
    | n=CARACTER          # Expr_Caracter
    | PARIZQ e1=e PARDER   # Expr_Par
    ;