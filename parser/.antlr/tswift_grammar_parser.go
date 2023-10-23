// Code generated from /home/josep/USAC/6to_Semestre/Lab_Compi/Proyecto2/OLC2_Proyecto2_202111478/parser/Tswift_Grammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tswift_Grammar

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

//Es el codigo que incrusta al principio del archivo
//Importacionies
import "OLC2_Proyecto2_202111478/gen"

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Tswift_GrammarParser struct {
	*antlr.BaseParser
}

var Tswift_GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tswift_grammarParserInit() {
	staticData := &Tswift_GrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "')'", "'('", "'{'", "'}'", "'['", "']'", "':'", "';'", "'?'", "','",
		"'.'", "'_'", "'&'", "'+='", "'-='", "'='", "'/'", "'%'", "'-'", "'+'",
		"'*'", "'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "'&&'", "'||'",
		"'!'", "'print'", "'var'", "'let'", "'true'", "'false'", "'if'", "'else'",
		"'switch'", "'case'", "'default'", "'while'", "'for'", "'in'", "'...'",
		"'guard'", "'continue'", "'return'", "'break'", "'nil'", "'append'",
		"'removeLast'", "'remove'", "'at'", "'IsEmpty'", "'count'", "'func'",
		"'repeating'", "'struct'", "'mutating'", "'inout'", "'atoi'", "'iota'",
		"'self'", "'Int'", "'Float'", "'Bool'", "'Character'", "'String'",
	}
	staticData.SymbolicNames = []string{
		"", "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER",
		"DOSPT", "PTCOMA", "INTERROGACION", "COMA", "PUNTO", "GUIONBAJO", "DIR",
		"MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR",
		"IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR",
		"AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "IF", "ELSE",
		"SWITCH", "CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", "GUARD",
		"CONTINUE", "RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", "REMOVE",
		"AT", "ISEMPTY", "COUNT", "FUNC", "REPEATING", "STRUCT", "MUTATING",
		"INOUT", "ATOI", "IOTA", "SELF", "INT", "FLOAT", "BOOL", "CHARACTER",
		"STRING", "ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID",
		"UL_C", "ML_C", "ERROR",
	}
	staticData.RuleNames = []string{
		"s", "cond", "marcador", "expr", "oprel",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 77, 100, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3,
		1, 34, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 5, 1, 48, 8, 1, 10, 1, 12, 1, 51, 9, 1, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		3, 3, 3, 69, 8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 3, 5, 3, 81, 8, 3, 10, 3, 12, 3, 84, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 98, 8, 4, 1, 4, 0, 2,
		2, 6, 5, 0, 2, 4, 6, 8, 0, 2, 2, 0, 17, 17, 21, 21, 1, 0, 19, 20, 110,
		0, 10, 1, 0, 0, 0, 2, 33, 1, 0, 0, 0, 4, 52, 1, 0, 0, 0, 6, 68, 1, 0, 0,
		0, 8, 97, 1, 0, 0, 0, 10, 11, 3, 2, 1, 0, 11, 12, 5, 0, 0, 1, 12, 13, 6,
		0, -1, 0, 13, 1, 1, 0, 0, 0, 14, 15, 6, 1, -1, 0, 15, 16, 5, 30, 0, 0,
		16, 17, 3, 2, 1, 7, 17, 18, 6, 1, -1, 0, 18, 34, 1, 0, 0, 0, 19, 20, 3,
		6, 3, 0, 20, 21, 3, 8, 4, 0, 21, 22, 3, 6, 3, 0, 22, 23, 6, 1, -1, 0, 23,
		34, 1, 0, 0, 0, 24, 25, 5, 2, 0, 0, 25, 26, 3, 2, 1, 0, 26, 27, 5, 1, 0,
		0, 27, 28, 6, 1, -1, 0, 28, 34, 1, 0, 0, 0, 29, 30, 5, 34, 0, 0, 30, 34,
		6, 1, -1, 0, 31, 32, 5, 35, 0, 0, 32, 34, 6, 1, -1, 0, 33, 14, 1, 0, 0,
		0, 33, 19, 1, 0, 0, 0, 33, 24, 1, 0, 0, 0, 33, 29, 1, 0, 0, 0, 33, 31,
		1, 0, 0, 0, 34, 49, 1, 0, 0, 0, 35, 36, 10, 6, 0, 0, 36, 37, 5, 28, 0,
		0, 37, 38, 3, 4, 2, 0, 38, 39, 3, 2, 1, 7, 39, 40, 6, 1, -1, 0, 40, 48,
		1, 0, 0, 0, 41, 42, 10, 5, 0, 0, 42, 43, 5, 29, 0, 0, 43, 44, 3, 4, 2,
		0, 44, 45, 3, 2, 1, 6, 45, 46, 6, 1, -1, 0, 46, 48, 1, 0, 0, 0, 47, 35,
		1, 0, 0, 0, 47, 41, 1, 0, 0, 0, 48, 51, 1, 0, 0, 0, 49, 47, 1, 0, 0, 0,
		49, 50, 1, 0, 0, 0, 50, 3, 1, 0, 0, 0, 51, 49, 1, 0, 0, 0, 52, 53, 6, 2,
		-1, 0, 53, 5, 1, 0, 0, 0, 54, 55, 6, 3, -1, 0, 55, 56, 5, 19, 0, 0, 56,
		57, 3, 6, 3, 6, 57, 58, 6, 3, -1, 0, 58, 69, 1, 0, 0, 0, 59, 60, 5, 2,
		0, 0, 60, 61, 3, 6, 3, 0, 61, 62, 5, 1, 0, 0, 62, 63, 6, 3, -1, 0, 63,
		69, 1, 0, 0, 0, 64, 65, 5, 70, 0, 0, 65, 69, 6, 3, -1, 0, 66, 67, 5, 74,
		0, 0, 67, 69, 6, 3, -1, 0, 68, 54, 1, 0, 0, 0, 68, 59, 1, 0, 0, 0, 68,
		64, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 69, 82, 1, 0, 0, 0, 70, 71, 10, 5,
		0, 0, 71, 72, 7, 0, 0, 0, 72, 73, 3, 6, 3, 6, 73, 74, 6, 3, -1, 0, 74,
		81, 1, 0, 0, 0, 75, 76, 10, 4, 0, 0, 76, 77, 7, 1, 0, 0, 77, 78, 3, 6,
		3, 5, 78, 79, 6, 3, -1, 0, 79, 81, 1, 0, 0, 0, 80, 70, 1, 0, 0, 0, 80,
		75, 1, 0, 0, 0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 82, 83, 1, 0, 0,
		0, 83, 7, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 86, 5, 22, 0, 0, 86, 98,
		6, 4, -1, 0, 87, 88, 5, 23, 0, 0, 88, 98, 6, 4, -1, 0, 89, 90, 5, 26, 0,
		0, 90, 98, 6, 4, -1, 0, 91, 92, 5, 27, 0, 0, 92, 98, 6, 4, -1, 0, 93, 94,
		5, 24, 0, 0, 94, 98, 6, 4, -1, 0, 95, 96, 5, 25, 0, 0, 96, 98, 6, 4, -1,
		0, 97, 85, 1, 0, 0, 0, 97, 87, 1, 0, 0, 0, 97, 89, 1, 0, 0, 0, 97, 91,
		1, 0, 0, 0, 97, 93, 1, 0, 0, 0, 97, 95, 1, 0, 0, 0, 98, 9, 1, 0, 0, 0,
		7, 33, 47, 49, 68, 80, 82, 97,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// Tswift_GrammarParserInit initializes any static state used to implement Tswift_GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTswift_GrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Tswift_GrammarParserInit() {
	staticData := &Tswift_GrammarParserStaticData
	staticData.once.Do(tswift_grammarParserInit)
}

// NewTswift_GrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTswift_GrammarParser(input antlr.TokenStream) *Tswift_GrammarParser {
	Tswift_GrammarParserInit()
	this := new(Tswift_GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Tswift_GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Tswift_Grammar.g4"

	return this
}

// Note that '@members' cannot be changed now, but this should have been 'globals'
// If you are looking to have variables for each instance, use '@structmembers'

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

	// Getter signatures
	EOF() antlr.TerminalNode
	Cond() ICondContext

	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	c1     ICondContext
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

func (s *SContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterS(s)
	}
}

func (s *SContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitS(s)
	}
}

func (p *Tswift_GrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Tswift_GrammarParserRULE_s)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)

		var _x = p.cond(0)

		localctx.(*SContext).c1 = _x
	}
	{
		p.SetState(11)
		p.Match(Tswift_GrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	NOT() antlr.TerminalNode
	AllCond() []ICondContext
	Cond(i int) ICondContext
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	Oprel() IOprelContext
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode
	Marcador() IMarcadorContext
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsCondContext differentiates from other interfaces.
	IsCondContext()
}

type CondContext struct {
	antlr.BaseParserRuleContext
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
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_cond
	return p
}

func InitEmptyCondContext(p *CondContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_cond
}

func (*CondContext) IsCondContext() {}

func NewCondContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondContext {
	var p = new(CondContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICondContext); ok {
			len++
		}
	}

	tst := make([]ICondContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICondContext); ok {
			tst[i] = t.(ICondContext)
			i++
		}
	}

	return tst
}

func (s *CondContext) Cond(i int) ICondContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondContext)
}

func (s *CondContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *CondContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CondContext) Oprel() IOprelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOprelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMarcadorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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

func (s *CondContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterCond(s)
	}
}

func (s *CondContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitCond(s)
	}
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
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(15)

			var _m = p.Match(Tswift_GrammarParserNOT)

			localctx.(*CondContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(25)

			var _x = p.cond(0)

			localctx.(*CondContext).c = _x
		}
		{
			p.SetState(26)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*CondContext).SetEV(localctx.(*CondContext).GetC().GetEV())
		localctx.(*CondContext).SetEF(localctx.(*CondContext).GetC().GetEF())

	case 4:
		{
			p.SetState(29)
			p.Match(Tswift_GrammarParserTRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*CondContext).SetEV(append(localctx.(*CondContext).EV, gen.NewEtq()))
		localctx.(*CondContext).SetEF(append(localctx.(*CondContext).EF, gen.NewEtq()))
		gen.Gen("goto " + localctx.(*CondContext).EV[0])

	case 5:
		{
			p.SetState(31)
			p.Match(Tswift_GrammarParserFALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		localctx.(*CondContext).SetEV(append(localctx.(*CondContext).EV, gen.NewEtq()))
		localctx.(*CondContext).SetEF(append(localctx.(*CondContext).EF, gen.NewEtq()))
		gen.Gen("goto " + localctx.(*CondContext).EF[0])

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(49)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(47)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCondContext(p, _parentctx, _parentState)
				localctx.(*CondContext).c1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_cond)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(36)

					var _m = p.Match(Tswift_GrammarParserAND)

					localctx.(*CondContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
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
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(42)

					var _m = p.Match(Tswift_GrammarParserOR)

					localctx.(*CondContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
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

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(51)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
	ES     []string
}

func NewEmptyMarcadorContext() *MarcadorContext {
	var p = new(MarcadorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_marcador
	return p
}

func InitEmptyMarcadorContext(p *MarcadorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_marcador
}

func (*MarcadorContext) IsMarcadorContext() {}

func NewMarcadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, ES []string) *MarcadorContext {
	var p = new(MarcadorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *MarcadorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterMarcador(s)
	}
}

func (s *MarcadorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitMarcador(s)
	}
}

func (p *Tswift_GrammarParser) Marcador(ES []string) (localctx IMarcadorContext) {
	localctx = NewMarcadorContext(p, p.GetParserRuleContext(), p.GetState(), ES)
	p.EnterRule(localctx, 4, Tswift_GrammarParserRULE_marcador)
	p.EnterOuterAlt(localctx, 1)

	gen.ImprimirEtq(ES)

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	MENOS() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	PARIZQ() antlr.TerminalNode
	PARDER() antlr.TerminalNode
	ENTERO() antlr.TerminalNode
	ID() antlr.TerminalNode
	POR() antlr.TerminalNode
	DIV() antlr.TerminalNode
	MAS() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
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
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

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

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitExpr(s)
	}
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserMENOS:
		{
			p.SetState(55)

			var _m = p.Match(Tswift_GrammarParserMENOS)

			localctx.(*ExprContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(60)

			var _x = p.expr(0)

			localctx.(*ExprContext).e1 = _x
		}
		{
			p.SetState(61)
			p.Match(Tswift_GrammarParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).SetDir(localctx.(*ExprContext).GetE1().GetDir())

	case Tswift_GrammarParserENTERO:
		{
			p.SetState(64)

			var _m = p.Match(Tswift_GrammarParserENTERO)

			localctx.(*ExprContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*ExprContext).SetDir((func() string {
			if localctx.(*ExprContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*ExprContext).GetId().GetText()
			}
		}()))

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(82)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(80)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				localctx.(*ExprContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarParserRULE_expr)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
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
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
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

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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

	// Getter signatures
	IGUALIGUAL() antlr.TerminalNode
	DIFERENTE() antlr.TerminalNode
	MAYOR() antlr.TerminalNode
	MENOR() antlr.TerminalNode
	MAYORIGUAL() antlr.TerminalNode
	MENORIGUAL() antlr.TerminalNode

	// IsOprelContext differentiates from other interfaces.
	IsOprelContext()
}

type OprelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
	op     string
	ope    antlr.Token
}

func NewEmptyOprelContext() *OprelContext {
	var p = new(OprelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_oprel
	return p
}

func InitEmptyOprelContext(p *OprelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarParserRULE_oprel
}

func (*OprelContext) IsOprelContext() {}

func NewOprelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OprelContext {
	var p = new(OprelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

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

func (s *OprelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.EnterOprel(s)
	}
}

func (s *OprelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarListener); ok {
		listenerT.ExitOprel(s)
	}
}

func (p *Tswift_GrammarParser) Oprel() (localctx IOprelContext) {
	localctx = NewOprelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Tswift_GrammarParserRULE_oprel)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarParserIGUALIGUAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)

			var _m = p.Match(Tswift_GrammarParserIGUALIGUAL)

			localctx.(*OprelContext).ope = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		localctx.(*OprelContext).SetOp((func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}()))

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
