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
    |dec_vector PTCOMA?#S_Declaracion_Vector
    |func_vector PTCOMA?#S_Funcion_Vector
    |dec_matriz PTCOMA?#S_Declaracion_Matriz
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
    FOR ID IN e RANGO e LLAVEIZQ l_sentencias LLAVEDER #ForInt
    | FOR ID IN e LLAVEIZQ l_sentencias LLAVEDER #ForList
    ;

//Vectores-----------------------------------------------------------------------
dec_vector:
    tipod=(VAR|LET) ID DOSPT CORCHETEIZQ tipo CORCHETEDER IGUAL def_vector #Declaracion_Vector
    |tipod=(VAR|LET) ID IGUAL CORCHETEIZQ tipo CORCHETEDER PARIZQ PARDER #Declaracion_Alterna
    ;

def_vector:
    CORCHETEIZQ e (',' e)* CORCHETEDER #Def_Vector
    |CORCHETEIZQ CORCHETEDER #Def_Vector_Vacio
    |ID #Def_Vector_Id
    ;

func_vector:
    ID PUNTO APPEND PARIZQ e PARDER #Func_Vector_Append
    |ID PUNTO REMOVELAST PARIZQ PARDER #Func_Vector_RemoveLast
    |ID PUNTO REMOVE PARIZQ AT DOSPT e PARDER #Func_Vector_Remove
    ;

//Matrices-----------------------------------------------------------------------
dec_matriz:
    tipod=(VAR|LET) ID (DOSPT  tipo )? IGUAL def_matriz #Declaracion_Matriz
    ;

def_matriz:
    listavalores_matriz #Def_Matriz_Lista
    |simple_matriz #Def_Matriz_Simple
    ;

listavalores_matriz:
    CORCHETEIZQ listavalores_matriz (COMA listavalores_matriz)* CORCHETEDER #Def_Matriz_Lista_Valores
    |CORCHETEIZQ e (COMA e)* CORCHETEDER #Def_Matriz_Expresion
    ;


simple_matriz:
    CORCHETEIZQ tipo CORCHETEDER PARIZQ REPEATING DOSPT simple_matriz COMA COUNT DOSPT e PARDER #Def_Matriz_Simple2
    | CORCHETEIZQ tipo CORCHETEDER PARIZQ REPEATING DOSPT e COMA COUNT DOSPT e PARDER #Def_Matriz_Simple3
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
    : <assoc=right> op=(MENOS|NOT) e1=e      # Expr_Neg
    | e1=e op=(POR | DIV) e2=e    # Expr_MulDiv
    | e1=e op=(MAS | MENOS) e2=e  # Expr_SumRes
    | e1=e op=MOD e2=e           # Expr_Mod
    | e op=(IGUALIGUAL | DIFERENTE | MAYORIGUAL | MAYOR | MENORIGUAL | MENOR) e     # Expr_Rel
    | e op=(AND | OR) e     # Expr_Logica
    | n=(TRUE | FALSE)    # Expr_Booleano
    | n=NIL              # Expr_Nil
    | ID CORCHETEIZQ e CORCHETEDER # Expr_Vector
    | ID CORCHETEIZQ e CORCHETEDER (CORCHETEIZQ e CORCHETEDER)+ # Expr_Matriz
    | ID PUNTO ISEMPTY  # Expr_IsEmpty
    | ID PUNTO COUNT    # Expr_Count
    | id=ID                # Expr_Id
    | n=ENTERO            # Expr_Entero
    | n=DECIMAL           # Expr_Decimal
    | n=CADENA            # Expr_Cadena
    | n=CARACTER          # Expr_Caracter
    | PARIZQ e1=e PARDER   # Expr_Par
    ;