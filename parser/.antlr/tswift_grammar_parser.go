// Code generated from /home/josep-ubu/Lab_Compiladores2/OLC2_Proyecto2_202111478/parser/Tswift_Grammar.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // Tswift_Grammar

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

//Es el codigo que incrusta al principio del archivo
//Importacionies
import "OLC2_Proyecto2_202111478/gen"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 79, 102,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 36, 10, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	7, 3, 50, 10, 3, 12, 3, 14, 3, 53, 11, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5,
	71, 10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	7, 5, 83, 10, 5, 12, 5, 14, 5, 86, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 100, 10, 6, 3, 6, 2, 4,
	4, 8, 7, 2, 4, 6, 8, 10, 2, 4, 4, 2, 19, 19, 23, 23, 3, 2, 21, 22, 2, 112,
	2, 12, 3, 2, 2, 2, 4, 35, 3, 2, 2, 2, 6, 54, 3, 2, 2, 2, 8, 70, 3, 2, 2,
	2, 10, 99, 3, 2, 2, 2, 12, 13, 5, 4, 3, 2, 13, 14, 7, 2, 2, 3, 14, 15,
	8, 2, 1, 2, 15, 3, 3, 2, 2, 2, 16, 17, 8, 3, 1, 2, 17, 18, 7, 32, 2, 2,
	18, 19, 5, 4, 3, 9, 19, 20, 8, 3, 1, 2, 20, 36, 3, 2, 2, 2, 21, 22, 5,
	8, 5, 2, 22, 23, 5, 10, 6, 2, 23, 24, 5, 8, 5, 2, 24, 25, 8, 3, 1, 2, 25,
	36, 3, 2, 2, 2, 26, 27, 7, 4, 2, 2, 27, 28, 5, 4, 3, 2, 28, 29, 7, 3, 2,
	2, 29, 30, 8, 3, 1, 2, 30, 36, 3, 2, 2, 2, 31, 32, 7, 36, 2, 2, 32, 36,
	8, 3, 1, 2, 33, 34, 7, 37, 2, 2, 34, 36, 8, 3, 1, 2, 35, 16, 3, 2, 2, 2,
	35, 21, 3, 2, 2, 2, 35, 26, 3, 2, 2, 2, 35, 31, 3, 2, 2, 2, 35, 33, 3,
	2, 2, 2, 36, 51, 3, 2, 2, 2, 37, 38, 12, 8, 2, 2, 38, 39, 7, 30, 2, 2,
	39, 40, 5, 6, 4, 2, 40, 41, 5, 4, 3, 9, 41, 42, 8, 3, 1, 2, 42, 50, 3,
	2, 2, 2, 43, 44, 12, 7, 2, 2, 44, 45, 7, 31, 2, 2, 45, 46, 5, 6, 4, 2,
	46, 47, 5, 4, 3, 8, 47, 48, 8, 3, 1, 2, 48, 50, 3, 2, 2, 2, 49, 37, 3,
	2, 2, 2, 49, 43, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 51,
	52, 3, 2, 2, 2, 52, 5, 3, 2, 2, 2, 53, 51, 3, 2, 2, 2, 54, 55, 8, 4, 1,
	2, 55, 7, 3, 2, 2, 2, 56, 57, 8, 5, 1, 2, 57, 58, 7, 21, 2, 2, 58, 59,
	5, 8, 5, 8, 59, 60, 8, 5, 1, 2, 60, 71, 3, 2, 2, 2, 61, 62, 7, 4, 2, 2,
	62, 63, 5, 8, 5, 2, 63, 64, 7, 3, 2, 2, 64, 65, 8, 5, 1, 2, 65, 71, 3,
	2, 2, 2, 66, 67, 7, 72, 2, 2, 67, 71, 8, 5, 1, 2, 68, 69, 7, 76, 2, 2,
	69, 71, 8, 5, 1, 2, 70, 56, 3, 2, 2, 2, 70, 61, 3, 2, 2, 2, 70, 66, 3,
	2, 2, 2, 70, 68, 3, 2, 2, 2, 71, 84, 3, 2, 2, 2, 72, 73, 12, 7, 2, 2, 73,
	74, 9, 2, 2, 2, 74, 75, 5, 8, 5, 8, 75, 76, 8, 5, 1, 2, 76, 83, 3, 2, 2,
	2, 77, 78, 12, 6, 2, 2, 78, 79, 9, 3, 2, 2, 79, 80, 5, 8, 5, 7, 80, 81,
	8, 5, 1, 2, 81, 83, 3, 2, 2, 2, 82, 72, 3, 2, 2, 2, 82, 77, 3, 2, 2, 2,
	83, 86, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2, 84, 85, 3, 2, 2, 2, 85, 9, 3, 2,
	2, 2, 86, 84, 3, 2, 2, 2, 87, 88, 7, 24, 2, 2, 88, 100, 8, 6, 1, 2, 89,
	90, 7, 25, 2, 2, 90, 100, 8, 6, 1, 2, 91, 92, 7, 28, 2, 2, 92, 100, 8,
	6, 1, 2, 93, 94, 7, 29, 2, 2, 94, 100, 8, 6, 1, 2, 95, 96, 7, 26, 2, 2,
	96, 100, 8, 6, 1, 2, 97, 98, 7, 27, 2, 2, 98, 100, 8, 6, 1, 2, 99, 87,
	3, 2, 2, 2, 99, 89, 3, 2, 2, 2, 99, 91, 3, 2, 2, 2, 99, 93, 3, 2, 2, 2,
	99, 95, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 11, 3, 2, 2, 2, 9, 35, 49,
	51, 70, 82, 84, 99,
}
var literalNames = []string{
	"", "')'", "'('", "'{'", "'}'", "'['", "']'", "':'", "';'", "'?'", "','",
	"'.'", "'_'", "'&'", "'+='", "'-='", "'='", "'/'", "'%'", "'-'", "'+'",
	"'*'", "'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "'&&'", "'||'", "'!'",
	"'print'", "'var'", "'let'", "'true'", "'false'", "'if'", "'else'", "'switch'",
	"'case'", "'default'", "'while'", "'for'", "'in'", "'...'", "'guard'",
	"'continue'", "'return'", "'break'", "'nil'", "'append'", "'removeLast'",
	"'remove'", "'at'", "'IsEmpty'", "'count'", "'func'", "'repeating'", "'struct'",
	"'mutating'", "'inout'", "'atoi'", "'iota'", "'self'", "'Int'", "'Float'",
	"'Bool'", "'Character'", "'String'",
}
var symbolicNames = []string{
	"", "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER",
	"DOSPT", "PTCOMA", "INTERROGACION", "COMA", "PUNTO", "GUIONBAJO", "DIR",
	"MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR",
	"IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR",
	"AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "IF", "ELSE",
	"SWITCH", "CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", "GUARD", "CONTINUE",
	"RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", "REMOVE", "AT", "ISEMPTY",
	"COUNT", "FUNC", "REPEATING", "STRUCT", "MUTATING", "INOUT", "ATOI", "IOTA",
	"SELF", "INT", "FLOAT", "BOOL", "CHARACTER", "STRING", "ENBLANCO", "ENTERO",
	"DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", "ML_C", "ERROR",
}

var ruleNames = []string{
	"s", "cond", "marcador", "expr", "oprel",
}

type Tswift_GrammarParser struct {
	*antlr.BaseParser
}

// NewTswift_GrammarParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Tswift_GrammarParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewTswift_GrammarParser(input antlr.TokenStream) *Tswift_GrammarParser {
	this := new(Tswift_GrammarParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Tswift_Grammar.g4"

	return this
}

//instancias de variables
//contador de temportales
var temp int = 1

// Tswift_GrammarParser tokens.
const (
	Tswift_GrammarParserEOF           = antlr.TokenEOF
	Tswift_GrammarParserPARDER        = 1
	Tswift_GrammarParserPARIZQ        = 2
	Tswift_GrammarParserLLAVEIZQ      = 3
	Tswift_GrammarParserLLAVEDER      = 4
	Tswift_GrammarParserCORCHETEIZQ   = 5
	Tswift_GrammarParserCORCHETEDER   = 6
	Tswift_GrammarParserDOSPT         = 7
	Tswift_GrammarParserPTCOMA        = 8
	Tswift_GrammarParserINTERROGACION = 9
	Tswift_GrammarParserCOMA          = 10
	Tswift_GrammarParserPUNTO         = 11
	Tswift_GrammarParserGUIONBAJO     = 12
	Tswift_GrammarParserDIR           = 13
	Tswift_GrammarParserMASIGUAL      = 14
	Tswift_GrammarParserMENOSIGUAL    = 15
	Tswift_GrammarParserIGUAL         = 16
	Tswift_GrammarParserDIV           = 17
	Tswift_GrammarParserMOD           = 18
	Tswift_GrammarParserMENOS         = 19
	Tswift_GrammarParserMAS           = 20
	Tswift_GrammarParserPOR           = 21
	Tswift_GrammarParserIGUALIGUAL    = 22
	Tswift_GrammarParserDIFERENTE     = 23
	Tswift_GrammarParserMAYORIGUAL    = 24
	Tswift_GrammarParserMENORIGUAL    = 25
	Tswift_GrammarParserMAYOR         = 26
	Tswift_GrammarParserMENOR         = 27
	Tswift_GrammarParserAND           = 28
	Tswift_GrammarParserOR            = 29
	Tswift_GrammarParserNOT           = 30
	Tswift_GrammarParserPRINT         = 31
	Tswift_GrammarParserVAR           = 32
	Tswift_GrammarParserLET           = 33
	Tswift_GrammarParserTRUE          = 34
	Tswift_GrammarParserFALSE         = 35
	Tswift_GrammarParserIF            = 36
	Tswift_GrammarParserELSE          = 37
	Tswift_GrammarParserSWITCH        = 38
	Tswift_GrammarParserCASE          = 39
	Tswift_GrammarParserDEFAULT       = 40
	Tswift_GrammarParserWHILE         = 41
	Tswift_GrammarParserFOR           = 42
	Tswift_GrammarParserIN            = 43
	Tswift_GrammarParserRANGO         = 44
	Tswift_GrammarParserGUARD         = 45
	Tswift_GrammarParserCONTINUE      = 46
	Tswift_GrammarParserRETURN        = 47
	Tswift_GrammarParserBREAK         = 48
	Tswift_GrammarParserNIL           = 49
	Tswift_GrammarParserAPPEND        = 50
	Tswift_GrammarParserREMOVELAST    = 51
	Tswift_GrammarParserREMOVE        = 52
	Tswift_GrammarParserAT            = 53
	Tswift_GrammarParserISEMPTY       = 54
	Tswift_GrammarParserCOUNT         = 55
	Tswift_GrammarParserFUNC          = 56
	Tswift_GrammarParserREPEATING     = 57
	Tswift_GrammarParserSTRUCT        = 58
	Tswift_GrammarParserMUTATING      = 59
	Tswift_GrammarParserINOUT         = 60
	Tswift_GrammarParserATOI          = 61
	Tswift_GrammarParserIOTA          = 62
	Tswift_GrammarParserSELF          = 63
	Tswift_GrammarParserINT           = 64
	Tswift_GrammarParserFLOAT         = 65
	Tswift_GrammarParserBOOL          = 66
	Tswift_GrammarParserCHARACTER     = 67
	Tswift_GrammarParserSTRING        = 68
	Tswift_GrammarParserENBLANCO      = 69
	Tswift_GrammarParserENTERO        = 70
	Tswift_GrammarParserDECIMAL       = 71
	Tswift_GrammarParserCARACTER      = 72
	Tswift_GrammarParserCADENA        = 73
	Tswift_GrammarParserID            = 74
	Tswift_GrammarParserUL_C          = 75
	Tswift_GrammarParserML_C          = 76
	Tswift_GrammarParserERROR         = 77
)

// Tswift_GrammarParser rules.
const (
	Tswift_GrammarParserRULE_s        = 0
	Tswift_GrammarParserRULE_cond     = 1
	Tswift_GrammarParserRULE_marcador = 2
	Tswift_GrammarParserRULE_expr     = 3
	Tswift_GrammarParserRULE_oprel    = 4
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetC1 returns the c1 rule contexts.
	GetC1() ICondContext

	// SetC1 sets the c1 rule contexts.
	SetC1(ICondContext)

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	c1     ICondContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_s
	return p
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) GetC1() ICondContext { return s.c1 }

func (s *SContext) SetC1(v ICondContext) { s.c1 = v }

func (s *SContext) EOF() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserEOF, 0)
}

func (s *SContext) Cond() ICondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *Tswift_GrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Tswift_GrammarParserRULE_s)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)

		var _x = p.cond(0)

		localctx.(*SContext).c1 = _x
	}
	{
		p.SetState(11)
		p.Match(Tswift_GrammarParserEOF)
	}

	return localctx
}

// ICondContext is an interface to support dynamic dispatch.
type ICondContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// GetC1 returns the c1 rule contexts.
	GetC1() ICondContext

	// GetC returns the c rule contexts.
	GetC() ICondContext

	// GetE1 returns the e1 rule contexts.
	GetE1() IExprContext

	// GetOpr returns the opr rule contexts.
	GetOpr() IOprelContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExprContext

	// GetC2 returns the c2 rule contexts.
	GetC2() ICondContext

	// SetC1 sets the c1 rule contexts.
	SetC1(ICondContext)

	// SetC sets the c rule contexts.
	SetC(ICondContext)

	// SetE1 sets the e1 rule contexts.
	SetE1(IExprContext)

	// SetOpr sets the opr rule contexts.
	SetOpr(IOprelContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExprContext)

	// SetC2 sets the c2 rule contexts.
	SetC2(ICondContext)

	// GetEV returns the EV attribute.
	GetEV() []string

	// GetEF returns the EF attribute.
	GetEF() []string

	// SetEV sets the EV attribute.
	SetEV([]string)

	// SetEF sets the EF attribute.
	SetEF([]string)

	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	EV     []string
	EF     []string
	c1     ICondContext
	op     antlr.Token
	c      ICondContext
	e1     IExprContext
	opr    IOprelContext
	e2     IExprContext
	c2     ICondContext
}

func NewEmptyCondContext() *CondContext {
	var p = new(CondContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_cond
	return p
}

func (*CondContext) IsCondContext() {}

func NewCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondContext {
	var p = new(CondContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_cond

	return p
}

func (s *CondContext) GetParser() antlr.Parser { return s.parser }

func (s *CondContext) GetOp() antlr.Token { return s.op }

func (s *CondContext) SetOp(v antlr.Token) { s.op = v }

func (s *CondContext) GetC1() ICondContext { return s.c1 }

func (s *CondContext) GetC() ICondContext { return s.c }

func (s *CondContext) GetE1() IExprContext { return s.e1 }

func (s *CondContext) GetOpr() IOprelContext { return s.opr }

func (s *CondContext) GetE2() IExprContext { return s.e2 }

func (s *CondContext) GetC2() ICondContext { return s.c2 }

func (s *CondContext) SetC1(v ICondContext) { s.c1 = v }

func (s *CondContext) SetC(v ICondContext) { s.c = v }

func (s *CondContext) SetE1(v IExprContext) { s.e1 = v }

func (s *CondContext) SetOpr(v IOprelContext) { s.opr = v }

func (s *CondContext) SetE2(v IExprContext) { s.e2 = v }

func (s *CondContext) SetC2(v ICondContext) { s.c2 = v }

func (s *CondContext) GetEV() []string { return s.EV }

func (s *CondContext) GetEF() []string { return s.EF }

func (s *CondContext) SetEV(v []string) { s.EV = v }

func (s *CondContext) SetEF(v []string) { s.EF = v }

func (s *CondContext) NOT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserNOT, 0)
}

func (s *CondContext) AllCond() []ICondContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICondContext)(nil)).Elem())
	var tst = make([]ICondContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICondContext)
		}
	}

	return tst
}

func (s *CondContext) Cond(i int) ICondContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *CondContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CondContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CondContext) Oprel() IOprelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOprelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOprelContext)
}

func (s *CondContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *CondContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *CondContext) TRUE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserTRUE, 0)
}

func (s *CondContext) FALSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserFALSE, 0)
}

func (s *CondContext) Marcador() IMarcadorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarcadorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMarcadorContext)
}

func (s *CondContext) AND() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserAND, 0)
}

func (s *CondContext) OR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserOR, 0)
}

func (s *CondContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *Tswift_GrammarParser) Cond() (localctx ICondContext) {
	return p.cond(0)
}

func (p *Tswift_GrammarParser) cond(_p int) (localctx ICondContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCondContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICondContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, Tswift_GrammarParserRULE_cond, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(15)

			var _m = p.Match(Tswift_GrammarParserNOT)

			localctx.(*CondContext).op = _m
		}
		{
			p.SetState(16)

			var _x = p.cond(7)

			localctx.(*CondContext).c = _x
		}

		localctx.(*CondContext).SetEV(localctx.(*CondContext).GetC().GetEF())
		localctx.(*CondContext).SetEF(localctx.(*CondContext).GetC().GetEV())

	case 2:
		{
			p.SetState(19)

			var _x = p.expr(0)

			localctx.(*CondContext).e1 = _x
		}
		{
			p.SetState(20)

			var _x = p.Oprel()

			localctx.(*CondContext).opr = _x
		}
		{
			p.SetState(21)

			var _x = p.expr(0)

			localctx.(*CondContext).e2 = _x
		}

		localctx.(*CondContext).SetEV(append(localctx.(*CondContext).EV, gen.NewEtq()))
		localctx.(*CondContext).SetEF(append(localctx.(*CondContext).EF, gen.NewEtq()))
		var cad string
		cad = localctx.(*CondContext).GetE1().GetDir() + " " + localctx.(*CondContext).GetOpr().GetOp() + " " + localctx.(*CondContext).GetE2().GetDir()
		gen.Gen("if " + cad + " then goto " + localctx.(*CondContext).EV[0])
		gen.Gen("goto " + localctx.(*CondContext).EF[0])

	case 3:
		{
			p.SetState(24)
			p.Match(Tswift_GrammarParserPARIZQ)
		}
		{
			p.SetState(25)

			var _x = p.cond(0)

			localctx.(*CondContext).c = _x
		}
		{
			p.SetState(26)
			p.Match(Tswift_GrammarParserPARDER)
		}

		localctx.(*CondContext).SetEV(localctx.(*CondContext).GetC().GetEV())
		localctx.(*CondContext).SetEF(localctx.(*CondContext).GetC().GetEF())

	case 4:
		{
			p.SetState(29)
			p.Match(Tswift_GrammarParserTRUE)
		}

		localctx.(*CondContext).SetEV(append(localctx.(*CondContext).EV, gen.NewEtq()))
		localctx.(*CondContext).SetEF(append(localctx.(*CondContext).EF, gen.NewEtq()))
		gen.Gen("goto " + localctx.(*CondContext).EV[0])

	case 5:
		{
			p.SetState(31)
			p.Match(Tswift_GrammarParserFALSE)
		}

		localctx.(*CondContext).SetEV(append(localctx.(*CondContext).EV, gen.NewEtq()))
		localctx.(*CondContext).SetEF(append(localctx.(*CondContext).EF, gen.NewEtq()))
		gen.Gen("goto " + localctx.(*CondContext).EF[0])

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCondContext(p, _parentctx, _parentState)
				localctx.(*CondContext).c1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_cond)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(36)

					var _m = p.Match(Tswift_GrammarParserAND)

					localctx.(*CondContext).op = _m
				}
				{
					p.SetState(37)
					p.Marcador(localctx.(*CondContext).GetC1().GetEV())
				}
				{
					p.SetState(38)

					var _x = p.cond(7)

					localctx.(*CondContext).c2 = _x
				}

				localctx.(*CondContext).SetEV(localctx.(*CondContext).GetC2().GetEV())
				localctx.(*CondContext).SetEF(gen.Unir(localctx.(*CondContext).GetC1().GetEF(), localctx.(*CondContext).GetC2().GetEF()))

			case 2:
				localctx = NewCondContext(p, _parentctx, _parentState)
				localctx.(*CondContext).c1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_cond)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(42)

					var _m = p.Match(Tswift_GrammarParserOR)

					localctx.(*CondContext).op = _m
				}
				{
					p.SetState(43)
					p.Marcador(localctx.(*CondContext).GetC1().GetEF())
				}
				{
					p.SetState(44)

					var _x = p.cond(6)

					localctx.(*CondContext).c2 = _x
				}

				localctx.(*CondContext).SetEV(gen.Unir(localctx.(*CondContext).GetC1().GetEV(), localctx.(*CondContext).GetC2().GetEV()))
				localctx.(*CondContext).SetEF(localctx.(*CondContext).GetC2().GetEF())

			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IMarcadorContext is an interface to support dynamic dispatch.
type IMarcadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetES returns the ES attribute.
	GetES() []string

	// SetES sets the ES attribute.
	SetES([]string)

	// IsMarcadorContext differentiates from other interfaces.
	IsMarcadorContext()
}

type MarcadorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ES     []string
}

func NewEmptyMarcadorContext() *MarcadorContext {
	var p = new(MarcadorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_marcador
	return p
}

func (*MarcadorContext) IsMarcadorContext() {}

func NewMarcadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, ES []string) *MarcadorContext {
	var p = new(MarcadorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_marcador

	p.ES = ES

	return p
}

func (s *MarcadorContext) GetParser() antlr.Parser { return s.parser }

func (s *MarcadorContext) GetES() []string { return s.ES }

func (s *MarcadorContext) SetES(v []string) { s.ES = v }

func (s *MarcadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarcadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *Tswift_GrammarParser) Marcador(ES []string) (localctx IMarcadorContext) {
	localctx = NewMarcadorContext(p, p.GetParserRuleContext(), p.GetState(), ES)
	p.EnterRule(localctx, 4, Tswift_GrammarParserRULE_marcador)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)

	gen.ImprimirEtq(ES)

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetN returns the n token.
	GetN() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExprContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExprContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExprContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExprContext)

	// GetDir returns the dir attribute.
	GetDir() string

	// SetDir sets the dir attribute.
	SetDir(string)

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    string
	e1     IExprContext
	op     antlr.Token
	n      antlr.Token
	id     antlr.Token
	e2     IExprContext
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) GetOp() antlr.Token { return s.op }

func (s *ExprContext) GetN() antlr.Token { return s.n }

func (s *ExprContext) GetId() antlr.Token { return s.id }

func (s *ExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprContext) SetN(v antlr.Token) { s.n = v }

func (s *ExprContext) SetId(v antlr.Token) { s.id = v }

func (s *ExprContext) GetE1() IExprContext { return s.e1 }

func (s *ExprContext) GetE2() IExprContext { return s.e2 }

func (s *ExprContext) SetE1(v IExprContext) { s.e1 = v }

func (s *ExprContext) SetE2(v IExprContext) { s.e2 = v }

func (s *ExprContext) GetDir() string { return s.dir }

func (s *ExprContext) SetDir(v string) { s.dir = v }

func (s *ExprContext) MENOS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOS, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARIZQ, 0)
}

func (s *ExprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPARDER, 0)
}

func (s *ExprContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserENTERO, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserID, 0)
}

func (s *ExprContext) POR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserPOR, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDIV, 0)
}

func (s *ExprContext) MAS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAS, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *Tswift_GrammarParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *Tswift_GrammarParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, Tswift_GrammarParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserMENOS:
		{
			p.SetState(55)

			var _m = p.Match(Tswift_GrammarParserMENOS)

			localctx.(*ExprContext).op = _m
		}
		{
			p.SetState(56)

			var _x = p.expr(6)

			localctx.(*ExprContext).e1 = _x
		}

		localctx.(*ExprContext).SetDir(gen.NewTemp())
		gen.Gen(localctx.(*ExprContext).dir + " = -1")
		gen.Gen(localctx.(*ExprContext).dir + " = " + localctx.(*ExprContext).dir + " * " + localctx.(*ExprContext).GetE1().GetDir())

	case Tswift_GrammarParserPARIZQ:
		{
			p.SetState(59)
			p.Match(Tswift_GrammarParserPARIZQ)
		}
		{
			p.SetState(60)

			var _x = p.expr(0)

			localctx.(*ExprContext).e1 = _x
		}
		{
			p.SetState(61)
			p.Match(Tswift_GrammarParserPARDER)
		}
		localctx.(*ExprContext).SetDir(localctx.(*ExprContext).GetE1().GetDir())

	case Tswift_GrammarParserENTERO:
		{
			p.SetState(64)

			var _m = p.Match(Tswift_GrammarParserENTERO)

			localctx.(*ExprContext).n = _m
		}
		localctx.(*ExprContext).SetDir((func() string {
			if localctx.(*ExprContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetN().GetText()
			}
		}()))

	case Tswift_GrammarParserID:
		{
			p.SetState(66)

			var _m = p.Match(Tswift_GrammarParserID)

			localctx.(*ExprContext).id = _m
		}
		localctx.(*ExprContext).SetDir((func() string {
			if localctx.(*ExprContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetId().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_expr)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(71)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarParserDIV || _la == Tswift_GrammarParserPOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(72)

					var _x = p.expr(6)

					localctx.(*ExprContext).e2 = _x
				}

				localctx.(*ExprContext).SetDir(gen.NewTemp())
				gen.Gen(localctx.(*ExprContext).dir + " = " + localctx.(*ExprContext).GetE1().GetDir() + " " + (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()) + " " + localctx.(*ExprContext).GetE2().GetDir())

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_expr)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(76)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarParserMENOS || _la == Tswift_GrammarParserMAS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(77)

					var _x = p.expr(5)

					localctx.(*ExprContext).e2 = _x
				}

				localctx.(*ExprContext).SetDir(gen.NewTemp())
				gen.Gen(localctx.(*ExprContext).dir + " = " + localctx.(*ExprContext).GetE1().GetDir() + " " + (func() string {
					if localctx.(*ExprContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExprContext).GetOp().GetText()
					}
				}()) + " " + localctx.(*ExprContext).GetE2().GetDir())

			}

		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IOprelContext is an interface to support dynamic dispatch.
type IOprelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpe returns the ope token.
	GetOpe() antlr.Token

	// SetOpe sets the ope token.
	SetOpe(antlr.Token)

	// GetOp returns the op attribute.
	GetOp() string

	// SetOp sets the op attribute.
	SetOp(string)

	// IsOprelContext differentiates from other interfaces.
	IsOprelContext()
}

type OprelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     string
	ope    antlr.Token
}

func NewEmptyOprelContext() *OprelContext {
	var p = new(OprelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_oprel
	return p
}

func (*OprelContext) IsOprelContext() {}

func NewOprelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OprelContext {
	var p = new(OprelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarParserRULE_oprel

	return p
}

func (s *OprelContext) GetParser() antlr.Parser { return s.parser }

func (s *OprelContext) GetOpe() antlr.Token { return s.ope }

func (s *OprelContext) SetOpe(v antlr.Token) { s.ope = v }

func (s *OprelContext) GetOp() string { return s.op }

func (s *OprelContext) SetOp(v string) { s.op = v }

func (s *OprelContext) IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserIGUALIGUAL, 0)
}

func (s *OprelContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserDIFERENTE, 0)
}

func (s *OprelContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAYOR, 0)
}

func (s *OprelContext) MENOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENOR, 0)
}

func (s *OprelContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMAYORIGUAL, 0)
}

func (s *OprelContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarParserMENORIGUAL, 0)
}

func (s *OprelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OprelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *Tswift_GrammarParser) Oprel() (localctx IOprelContext) {
	localctx = NewOprelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Tswift_GrammarParserRULE_oprel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(97)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserIGUALIGUAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)

			var _m = p.Match(Tswift_GrammarParserIGUALIGUAL)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).SetOp((func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}()))

	case Tswift_GrammarParserDIFERENTE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)

			var _m = p.Match(Tswift_GrammarParserDIFERENTE)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).SetOp((func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}()))

	case Tswift_GrammarParserMAYOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(89)

			var _m = p.Match(Tswift_GrammarParserMAYOR)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).SetOp((func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}()))

	case Tswift_GrammarParserMENOR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(91)

			var _m = p.Match(Tswift_GrammarParserMENOR)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).SetOp((func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}()))

	case Tswift_GrammarParserMAYORIGUAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(93)

			var _m = p.Match(Tswift_GrammarParserMAYORIGUAL)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).SetOp((func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}()))

	case Tswift_GrammarParserMENORIGUAL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(95)

			var _m = p.Match(Tswift_GrammarParserMENORIGUAL)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).SetOp((func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}()))

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Tswift_GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *CondContext = nil
		if localctx != nil {
			t = localctx.(*CondContext)
		}
		return p.Cond_Sempred(t, predIndex)

	case 3:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Tswift_GrammarParser) Cond_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Tswift_GrammarParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
