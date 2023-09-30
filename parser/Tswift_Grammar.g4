grammar Tswift_Grammar;
import Tswift_Lexer;

//opciones de header y member
options {
    language = 'Go';
}
@parser::header{
    //Es el codigo que incrusta al principio del archivo
    //Importacionies
    import "OLC2_Proyecto2_202111478/gen"
}

@parser::members{
    //instancias de variables
    //contador de temportales
    var temp int = 1;
}

/* reglas */
s
    :c1 = cond EOF  {}
    ;

cond returns[[]string EV, []string EF]
    : <assoc=right> op=NOT c=cond   { 
                                        $EV = $c.EF; 
                                        $EF = $c.EV; 
                                    }
    | c1=cond op=AND marcador[$c1.EV] c2=cond  { 
                                                $EV = $c2.EV;
                                                $EF = gen.Unir($c1.EF, $c2.EF);
                                            }
    | c1=cond op=OR marcador[$c1.EF] c2=cond   { 
                                                $EV = gen.Unir($c1.EV, $c2.EV);
                                                $EF = $c2.EF;
                                            }
    | e1=expr opr=oprel e2=expr     { 
                                        $EV = append($EV, gen.NewEtq());
                                        $EF = append($EF, gen.NewEtq());
                                        var cad string;
                                        cad = $e1.dir + " " + $opr.op + " " + $e2.dir;
                                        gen.Gen("if " + cad + " then goto " + $EV[0]);
                                        gen.Gen("goto " + $EF[0]);

                                    }
    | PARIZQ c=cond PARDER          { 
                                        $EV = $c.EV; 
                                        $EF = $c.EF; 
                                    }
    | TRUE                         { 
                                        $EV = append($EV, gen.NewEtq());
                                        $EF = append($EF, gen.NewEtq());
                                        gen.Gen("goto " + $EV[0]);
                                    }
    | FALSE                        { 
                                        $EV = append($EV, gen.NewEtq());
                                        $EF = append($EF, gen.NewEtq());
                                        gen.Gen("goto " + $EF[0]);
                                    }       
    ;

marcador [[]string ES]:{
    gen.ImprimirEtq(ES);
};

expr returns[string dir]
    :<assoc=right> op=MENOS e1=expr {
                                        $dir = gen.NewTemp();
                                        gen.Gen($dir + " = -1");
                                        gen.Gen($dir + " = " + $dir + " * " + $e1.dir);
                                    } 
    |e1=expr op=('*'|'/') e2=expr  {
                                        $dir = gen.NewTemp();
                                        gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir);
                                    }
    | e1=expr op=('+'|'-') e2=expr  {
                                        $dir = gen.NewTemp();
                                        gen.Gen($dir + " = " + $e1.dir + " " + $op.text + " " + $e2.dir);
                                    }
    | PARIZQ e1=expr PARDER         {$dir = $e1.dir;}
    | n=ENTERO                      {$dir = $n.text;}
    | id=ID                         {$dir = $id.text;}
    ;

oprel returns[string op]
    : ope=IGUALIGUAL { $op = $ope.text ;}
    | ope=DIFERENTE { $op = $ope.text ;}
    | ope=MAYOR { $op = $ope.text; }
    | ope=MENOR { $op = $ope.text ;}
    | ope=MAYORIGUAL { $op = $ope.text ;}
    | ope=MENORIGUAL { $op = $ope.text ;}
    ;