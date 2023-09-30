// Code generated from Tswift_GrammarN.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftGen // Tswift_GrammarN
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type Tswift_GrammarNParser struct {
	*antlr.BaseParser
}

var Tswift_GrammarNParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tswift_grammarnParserInit() {
	staticData := &Tswift_GrammarNParserStaticData
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
		"s", "l_sentencias", "sentencia", "print_sentencia", "if_sentencia",
		"condicion", "e",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 77, 108, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 1, 0, 1, 0, 1, 1, 5, 1, 18, 8, 1, 10, 1, 12,
		1, 21, 9, 1, 1, 2, 1, 2, 3, 2, 25, 8, 2, 1, 2, 1, 2, 3, 2, 29, 8, 2, 1,
		2, 3, 2, 32, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 39, 8, 3, 10, 3,
		12, 3, 42, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 57, 8, 4, 3, 4, 59, 8, 4, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 73, 8, 5,
		1, 5, 1, 5, 1, 5, 5, 5, 78, 8, 5, 10, 5, 12, 5, 81, 9, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 92, 8, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 5, 6, 103, 8, 6, 10, 6, 12, 6, 106,
		9, 6, 1, 6, 0, 2, 10, 12, 7, 0, 2, 4, 6, 8, 10, 12, 0, 5, 1, 0, 22, 27,
		1, 0, 34, 35, 1, 0, 28, 29, 2, 0, 17, 17, 21, 21, 1, 0, 19, 20, 118, 0,
		14, 1, 0, 0, 0, 2, 19, 1, 0, 0, 0, 4, 31, 1, 0, 0, 0, 6, 33, 1, 0, 0, 0,
		8, 45, 1, 0, 0, 0, 10, 72, 1, 0, 0, 0, 12, 91, 1, 0, 0, 0, 14, 15, 3, 2,
		1, 0, 15, 1, 1, 0, 0, 0, 16, 18, 3, 4, 2, 0, 17, 16, 1, 0, 0, 0, 18, 21,
		1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 3, 1, 0, 0, 0,
		21, 19, 1, 0, 0, 0, 22, 24, 3, 6, 3, 0, 23, 25, 5, 8, 0, 0, 24, 23, 1,
		0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 32, 1, 0, 0, 0, 26, 28, 3, 10, 5, 0, 27,
		29, 5, 8, 0, 0, 28, 27, 1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 32, 1, 0, 0,
		0, 30, 32, 3, 8, 4, 0, 31, 22, 1, 0, 0, 0, 31, 26, 1, 0, 0, 0, 31, 30,
		1, 0, 0, 0, 32, 5, 1, 0, 0, 0, 33, 34, 5, 31, 0, 0, 34, 35, 5, 2, 0, 0,
		35, 40, 3, 12, 6, 0, 36, 37, 5, 10, 0, 0, 37, 39, 3, 12, 6, 0, 38, 36,
		1, 0, 0, 0, 39, 42, 1, 0, 0, 0, 40, 38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0,
		41, 43, 1, 0, 0, 0, 42, 40, 1, 0, 0, 0, 43, 44, 5, 1, 0, 0, 44, 7, 1, 0,
		0, 0, 45, 46, 5, 36, 0, 0, 46, 47, 3, 10, 5, 0, 47, 48, 5, 3, 0, 0, 48,
		49, 3, 2, 1, 0, 49, 58, 5, 4, 0, 0, 50, 56, 5, 37, 0, 0, 51, 57, 3, 8,
		4, 0, 52, 53, 5, 3, 0, 0, 53, 54, 3, 2, 1, 0, 54, 55, 5, 4, 0, 0, 55, 57,
		1, 0, 0, 0, 56, 51, 1, 0, 0, 0, 56, 52, 1, 0, 0, 0, 57, 59, 1, 0, 0, 0,
		58, 50, 1, 0, 0, 0, 58, 59, 1, 0, 0, 0, 59, 9, 1, 0, 0, 0, 60, 61, 6, 5,
		-1, 0, 61, 62, 5, 30, 0, 0, 62, 73, 3, 10, 5, 5, 63, 64, 3, 12, 6, 0, 64,
		65, 7, 0, 0, 0, 65, 66, 3, 12, 6, 0, 66, 73, 1, 0, 0, 0, 67, 73, 7, 1,
		0, 0, 68, 69, 5, 2, 0, 0, 69, 70, 3, 10, 5, 0, 70, 71, 5, 1, 0, 0, 71,
		73, 1, 0, 0, 0, 72, 60, 1, 0, 0, 0, 72, 63, 1, 0, 0, 0, 72, 67, 1, 0, 0,
		0, 72, 68, 1, 0, 0, 0, 73, 79, 1, 0, 0, 0, 74, 75, 10, 3, 0, 0, 75, 76,
		7, 2, 0, 0, 76, 78, 3, 10, 5, 4, 77, 74, 1, 0, 0, 0, 78, 81, 1, 0, 0, 0,
		79, 77, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 11, 1, 0, 0, 0, 81, 79, 1,
		0, 0, 0, 82, 83, 6, 6, -1, 0, 83, 84, 5, 19, 0, 0, 84, 92, 3, 12, 6, 7,
		85, 92, 5, 74, 0, 0, 86, 92, 5, 70, 0, 0, 87, 88, 5, 2, 0, 0, 88, 89, 3,
		12, 6, 0, 89, 90, 5, 1, 0, 0, 90, 92, 1, 0, 0, 0, 91, 82, 1, 0, 0, 0, 91,
		85, 1, 0, 0, 0, 91, 86, 1, 0, 0, 0, 91, 87, 1, 0, 0, 0, 92, 104, 1, 0,
		0, 0, 93, 94, 10, 6, 0, 0, 94, 95, 7, 3, 0, 0, 95, 103, 3, 12, 6, 7, 96,
		97, 10, 5, 0, 0, 97, 98, 7, 4, 0, 0, 98, 103, 3, 12, 6, 6, 99, 100, 10,
		4, 0, 0, 100, 101, 5, 18, 0, 0, 101, 103, 3, 12, 6, 5, 102, 93, 1, 0, 0,
		0, 102, 96, 1, 0, 0, 0, 102, 99, 1, 0, 0, 0, 103, 106, 1, 0, 0, 0, 104,
		102, 1, 0, 0, 0, 104, 105, 1, 0, 0, 0, 105, 13, 1, 0, 0, 0, 106, 104, 1,
		0, 0, 0, 12, 19, 24, 28, 31, 40, 56, 58, 72, 79, 91, 102, 104,
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

// Tswift_GrammarNParserInit initializes any static state used to implement Tswift_GrammarNParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTswift_GrammarNParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Tswift_GrammarNParserInit() {
	staticData := &Tswift_GrammarNParserStaticData
	staticData.once.Do(tswift_grammarnParserInit)
}

// NewTswift_GrammarNParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTswift_GrammarNParser(input antlr.TokenStream) *Tswift_GrammarNParser {
	Tswift_GrammarNParserInit()
	this := new(Tswift_GrammarNParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Tswift_GrammarNParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Tswift_GrammarN.g4"

	return this
}

// Tswift_GrammarNParser tokens.
const (
	Tswift_GrammarNParserEOF           = antlr.TokenEOF
	Tswift_GrammarNParserPARDER        = 1
	Tswift_GrammarNParserPARIZQ        = 2
	Tswift_GrammarNParserLLAVEIZQ      = 3
	Tswift_GrammarNParserLLAVEDER      = 4
	Tswift_GrammarNParserCORCHETEIZQ   = 5
	Tswift_GrammarNParserCORCHETEDER   = 6
	Tswift_GrammarNParserDOSPT         = 7
	Tswift_GrammarNParserPTCOMA        = 8
	Tswift_GrammarNParserINTERROGACION = 9
	Tswift_GrammarNParserCOMA          = 10
	Tswift_GrammarNParserPUNTO         = 11
	Tswift_GrammarNParserGUIONBAJO     = 12
	Tswift_GrammarNParserDIR           = 13
	Tswift_GrammarNParserMASIGUAL      = 14
	Tswift_GrammarNParserMENOSIGUAL    = 15
	Tswift_GrammarNParserIGUAL         = 16
	Tswift_GrammarNParserDIV           = 17
	Tswift_GrammarNParserMOD           = 18
	Tswift_GrammarNParserMENOS         = 19
	Tswift_GrammarNParserMAS           = 20
	Tswift_GrammarNParserPOR           = 21
	Tswift_GrammarNParserIGUALIGUAL    = 22
	Tswift_GrammarNParserDIFERENTE     = 23
	Tswift_GrammarNParserMAYORIGUAL    = 24
	Tswift_GrammarNParserMENORIGUAL    = 25
	Tswift_GrammarNParserMAYOR         = 26
	Tswift_GrammarNParserMENOR         = 27
	Tswift_GrammarNParserAND           = 28
	Tswift_GrammarNParserOR            = 29
	Tswift_GrammarNParserNOT           = 30
	Tswift_GrammarNParserPRINT         = 31
	Tswift_GrammarNParserVAR           = 32
	Tswift_GrammarNParserLET           = 33
	Tswift_GrammarNParserTRUE          = 34
	Tswift_GrammarNParserFALSE         = 35
	Tswift_GrammarNParserIF            = 36
	Tswift_GrammarNParserELSE          = 37
	Tswift_GrammarNParserSWITCH        = 38
	Tswift_GrammarNParserCASE          = 39
	Tswift_GrammarNParserDEFAULT       = 40
	Tswift_GrammarNParserWHILE         = 41
	Tswift_GrammarNParserFOR           = 42
	Tswift_GrammarNParserIN            = 43
	Tswift_GrammarNParserRANGO         = 44
	Tswift_GrammarNParserGUARD         = 45
	Tswift_GrammarNParserCONTINUE      = 46
	Tswift_GrammarNParserRETURN        = 47
	Tswift_GrammarNParserBREAK         = 48
	Tswift_GrammarNParserNIL           = 49
	Tswift_GrammarNParserAPPEND        = 50
	Tswift_GrammarNParserREMOVELAST    = 51
	Tswift_GrammarNParserREMOVE        = 52
	Tswift_GrammarNParserAT            = 53
	Tswift_GrammarNParserISEMPTY       = 54
	Tswift_GrammarNParserCOUNT         = 55
	Tswift_GrammarNParserFUNC          = 56
	Tswift_GrammarNParserREPEATING     = 57
	Tswift_GrammarNParserSTRUCT        = 58
	Tswift_GrammarNParserMUTATING      = 59
	Tswift_GrammarNParserINOUT         = 60
	Tswift_GrammarNParserATOI          = 61
	Tswift_GrammarNParserIOTA          = 62
	Tswift_GrammarNParserSELF          = 63
	Tswift_GrammarNParserINT           = 64
	Tswift_GrammarNParserFLOAT         = 65
	Tswift_GrammarNParserBOOL          = 66
	Tswift_GrammarNParserCHARACTER     = 67
	Tswift_GrammarNParserSTRING        = 68
	Tswift_GrammarNParserENBLANCO      = 69
	Tswift_GrammarNParserENTERO        = 70
	Tswift_GrammarNParserDECIMAL       = 71
	Tswift_GrammarNParserCARACTER      = 72
	Tswift_GrammarNParserCADENA        = 73
	Tswift_GrammarNParserID            = 74
	Tswift_GrammarNParserUL_C          = 75
	Tswift_GrammarNParserML_C          = 76
	Tswift_GrammarNParserERROR         = 77
)

// Tswift_GrammarNParser rules.
const (
	Tswift_GrammarNParserRULE_s               = 0
	Tswift_GrammarNParserRULE_l_sentencias    = 1
	Tswift_GrammarNParserRULE_sentencia       = 2
	Tswift_GrammarNParserRULE_print_sentencia = 3
	Tswift_GrammarNParserRULE_if_sentencia    = 4
	Tswift_GrammarNParserRULE_condicion       = 5
	Tswift_GrammarNParserRULE_e               = 6
)

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) CopyAll(ctx *SContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SLSentenciasContext struct {
	SContext
}

func NewSLSentenciasContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SLSentenciasContext {
	var p = new(SLSentenciasContext)

	InitEmptySContext(&p.SContext)
	p.parser = parser
	p.CopyAll(ctx.(*SContext))

	return p
}

func (s *SLSentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SLSentenciasContext) L_sentencias() IL_sentenciasContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IL_sentenciasContext)
}

func (s *SLSentenciasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterSLSentencias(s)
	}
}

func (s *SLSentenciasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitSLSentencias(s)
	}
}

func (s *SLSentenciasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitSLSentencias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Tswift_GrammarNParserRULE_s)
	localctx = NewSLSentenciasContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)
		p.L_sentencias()
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

// IL_sentenciasContext is an interface to support dynamic dispatch.
type IL_sentenciasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsL_sentenciasContext differentiates from other interfaces.
	IsL_sentenciasContext()
}

type L_sentenciasContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyL_sentenciasContext() *L_sentenciasContext {
	var p = new(L_sentenciasContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_l_sentencias
	return p
}

func InitEmptyL_sentenciasContext(p *L_sentenciasContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_l_sentencias
}

func (*L_sentenciasContext) IsL_sentenciasContext() {}

func NewL_sentenciasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *L_sentenciasContext {
	var p = new(L_sentenciasContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_l_sentencias

	return p
}

func (s *L_sentenciasContext) GetParser() antlr.Parser { return s.parser }

func (s *L_sentenciasContext) CopyAll(ctx *L_sentenciasContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *L_sentenciasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_sentenciasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type L_SentenciaContext struct {
	L_sentenciasContext
}

func NewL_SentenciaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *L_SentenciaContext {
	var p = new(L_SentenciaContext)

	InitEmptyL_sentenciasContext(&p.L_sentenciasContext)
	p.parser = parser
	p.CopyAll(ctx.(*L_sentenciasContext))

	return p
}

func (s *L_SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *L_SentenciaContext) AllSentencia() []ISentenciaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentenciaContext); ok {
			len++
		}
	}

	tst := make([]ISentenciaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentenciaContext); ok {
			tst[i] = t.(ISentenciaContext)
			i++
		}
	}

	return tst
}

func (s *L_SentenciaContext) Sentencia(i int) ISentenciaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaContext); ok {
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

	return t.(ISentenciaContext)
}

func (s *L_SentenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterL_Sentencia(s)
	}
}

func (s *L_SentenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitL_Sentencia(s)
	}
}

func (s *L_SentenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitL_Sentencia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) L_sentencias() (localctx IL_sentenciasContext) {
	localctx = NewL_sentenciasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Tswift_GrammarNParserRULE_l_sentencias)
	var _la int

	localctx = NewL_SentenciaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&123480834052) != 0) || _la == Tswift_GrammarNParserENTERO || _la == Tswift_GrammarNParserID {
		{
			p.SetState(16)
			p.Sentencia()
		}

		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

// ISentenciaContext is an interface to support dynamic dispatch.
type ISentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSentenciaContext differentiates from other interfaces.
	IsSentenciaContext()
}

type SentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaContext() *SentenciaContext {
	var p = new(SentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_sentencia
	return p
}

func InitEmptySentenciaContext(p *SentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_sentencia
}

func (*SentenciaContext) IsSentenciaContext() {}

func NewSentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaContext {
	var p = new(SentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_sentencia

	return p
}

func (s *SentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaContext) CopyAll(ctx *SentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type S_CondicionContext struct {
	SentenciaContext
}

func NewS_CondicionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_CondicionContext {
	var p = new(S_CondicionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_CondicionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_CondicionContext) Condicion() ICondicionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondicionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondicionContext)
}

func (s *S_CondicionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPTCOMA, 0)
}

func (s *S_CondicionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Condicion(s)
	}
}

func (s *S_CondicionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Condicion(s)
	}
}

func (s *S_CondicionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Condicion(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_IfContext struct {
	SentenciaContext
}

func NewS_IfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_IfContext {
	var p = new(S_IfContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_IfContext) If_sentencia() IIf_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_sentenciaContext)
}

func (s *S_IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_If(s)
	}
}

func (s *S_IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_If(s)
	}
}

func (s *S_IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_If(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_PrintContext struct {
	SentenciaContext
}

func NewS_PrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_PrintContext {
	var p = new(S_PrintContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_PrintContext) Print_sentencia() IPrint_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrint_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrint_sentenciaContext)
}

func (s *S_PrintContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPTCOMA, 0)
}

func (s *S_PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Print(s)
	}
}

func (s *S_PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Print(s)
	}
}

func (s *S_PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Print(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Sentencia() (localctx ISentenciaContext) {
	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Tswift_GrammarNParserRULE_sentencia)
	var _la int

	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarNParserPRINT:
		localctx = NewS_PrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(22)
			p.Print_sentencia()
		}
		p.SetState(24)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarNParserPTCOMA {
			{
				p.SetState(23)
				p.Match(Tswift_GrammarNParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case Tswift_GrammarNParserPARIZQ, Tswift_GrammarNParserMENOS, Tswift_GrammarNParserNOT, Tswift_GrammarNParserTRUE, Tswift_GrammarNParserFALSE, Tswift_GrammarNParserENTERO, Tswift_GrammarNParserID:
		localctx = NewS_CondicionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(26)
			p.condicion(0)
		}
		p.SetState(28)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarNParserPTCOMA {
			{
				p.SetState(27)
				p.Match(Tswift_GrammarNParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case Tswift_GrammarNParserIF:
		localctx = NewS_IfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(30)
			p.If_sentencia()
		}

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

// IPrint_sentenciaContext is an interface to support dynamic dispatch.
type IPrint_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPrint_sentenciaContext differentiates from other interfaces.
	IsPrint_sentenciaContext()
}

type Print_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrint_sentenciaContext() *Print_sentenciaContext {
	var p = new(Print_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_print_sentencia
	return p
}

func InitEmptyPrint_sentenciaContext(p *Print_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_print_sentencia
}

func (*Print_sentenciaContext) IsPrint_sentenciaContext() {}

func NewPrint_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Print_sentenciaContext {
	var p = new(Print_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_print_sentencia

	return p
}

func (s *Print_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Print_sentenciaContext) CopyAll(ctx *Print_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Print_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Print_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintContext struct {
	Print_sentenciaContext
}

func NewPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintContext {
	var p = new(PrintContext)

	InitEmptyPrint_sentenciaContext(&p.Print_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Print_sentenciaContext))

	return p
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPRINT, 0)
}

func (s *PrintContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARIZQ, 0)
}

func (s *PrintContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *PrintContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
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

	return t.(IEContext)
}

func (s *PrintContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARDER, 0)
}

func (s *PrintContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserCOMA)
}

func (s *PrintContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCOMA, i)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitPrint(s)
	}
}

func (s *PrintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitPrint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Print_sentencia() (localctx IPrint_sentenciaContext) {
	localctx = NewPrint_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Tswift_GrammarNParserRULE_print_sentencia)
	var _la int

	localctx = NewPrintContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(33)
		p.Match(Tswift_GrammarNParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(34)
		p.Match(Tswift_GrammarNParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.e(0)
	}
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarNParserCOMA {
		{
			p.SetState(36)
			p.Match(Tswift_GrammarNParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.e(0)
		}

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(43)
		p.Match(Tswift_GrammarNParserPARDER)
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

// IIf_sentenciaContext is an interface to support dynamic dispatch.
type IIf_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIf_sentenciaContext differentiates from other interfaces.
	IsIf_sentenciaContext()
}

type If_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIf_sentenciaContext() *If_sentenciaContext {
	var p = new(If_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_if_sentencia
	return p
}

func InitEmptyIf_sentenciaContext(p *If_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_if_sentencia
}

func (*If_sentenciaContext) IsIf_sentenciaContext() {}

func NewIf_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *If_sentenciaContext {
	var p = new(If_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_if_sentencia

	return p
}

func (s *If_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *If_sentenciaContext) CopyAll(ctx *If_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *If_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *If_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfContext struct {
	If_sentenciaContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	InitEmptyIf_sentenciaContext(&p.If_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*If_sentenciaContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IF() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIF, 0)
}

func (s *IfContext) Condicion() ICondicionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondicionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondicionContext)
}

func (s *IfContext) AllLLAVEIZQ() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserLLAVEIZQ)
}

func (s *IfContext) LLAVEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEIZQ, i)
}

func (s *IfContext) AllL_sentencias() []IL_sentenciasContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IL_sentenciasContext); ok {
			len++
		}
	}

	tst := make([]IL_sentenciasContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IL_sentenciasContext); ok {
			tst[i] = t.(IL_sentenciasContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) L_sentencias(i int) IL_sentenciasContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IL_sentenciasContext); ok {
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

	return t.(IL_sentenciasContext)
}

func (s *IfContext) AllLLAVEDER() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserLLAVEDER)
}

func (s *IfContext) LLAVEDER(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEDER, i)
}

func (s *IfContext) ELSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserELSE, 0)
}

func (s *IfContext) If_sentencia() IIf_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIf_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIf_sentenciaContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitIf(s)
	}
}

func (s *IfContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitIf(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) If_sentencia() (localctx IIf_sentenciaContext) {
	localctx = NewIf_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Tswift_GrammarNParserRULE_if_sentencia)
	var _la int

	localctx = NewIfContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		p.Match(Tswift_GrammarNParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.condicion(0)
	}
	{
		p.SetState(47)
		p.Match(Tswift_GrammarNParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(48)
		p.L_sentencias()
	}
	{
		p.SetState(49)
		p.Match(Tswift_GrammarNParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarNParserELSE {
		{
			p.SetState(50)
			p.Match(Tswift_GrammarNParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(56)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Tswift_GrammarNParserIF:
			{
				p.SetState(51)
				p.If_sentencia()
			}

		case Tswift_GrammarNParserLLAVEIZQ:
			{
				p.SetState(52)
				p.Match(Tswift_GrammarNParserLLAVEIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(53)
				p.L_sentencias()
			}
			{
				p.SetState(54)
				p.Match(Tswift_GrammarNParserLLAVEDER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// ICondicionContext is an interface to support dynamic dispatch.
type ICondicionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCondicionContext differentiates from other interfaces.
	IsCondicionContext()
}

type CondicionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondicionContext() *CondicionContext {
	var p = new(CondicionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_condicion
	return p
}

func InitEmptyCondicionContext(p *CondicionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_condicion
}

func (*CondicionContext) IsCondicionContext() {}

func NewCondicionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CondicionContext {
	var p = new(CondicionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_condicion

	return p
}

func (s *CondicionContext) GetParser() antlr.Parser { return s.parser }

func (s *CondicionContext) CopyAll(ctx *CondicionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CondicionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CondicionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Cond_ParContext struct {
	CondicionContext
	c ICondicionContext
}

func NewCond_ParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Cond_ParContext {
	var p = new(Cond_ParContext)

	InitEmptyCondicionContext(&p.CondicionContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondicionContext))

	return p
}

func (s *Cond_ParContext) GetC() ICondicionContext { return s.c }

func (s *Cond_ParContext) SetC(v ICondicionContext) { s.c = v }

func (s *Cond_ParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cond_ParContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARIZQ, 0)
}

func (s *Cond_ParContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARDER, 0)
}

func (s *Cond_ParContext) Condicion() ICondicionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondicionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondicionContext)
}

func (s *Cond_ParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterCond_Par(s)
	}
}

func (s *Cond_ParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitCond_Par(s)
	}
}

func (s *Cond_ParContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitCond_Par(s)

	default:
		return t.VisitChildren(s)
	}
}

type Cond_NegContext struct {
	CondicionContext
	op antlr.Token
	c  ICondicionContext
}

func NewCond_NegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Cond_NegContext {
	var p = new(Cond_NegContext)

	InitEmptyCondicionContext(&p.CondicionContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondicionContext))

	return p
}

func (s *Cond_NegContext) GetOp() antlr.Token { return s.op }

func (s *Cond_NegContext) SetOp(v antlr.Token) { s.op = v }

func (s *Cond_NegContext) GetC() ICondicionContext { return s.c }

func (s *Cond_NegContext) SetC(v ICondicionContext) { s.c = v }

func (s *Cond_NegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cond_NegContext) NOT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserNOT, 0)
}

func (s *Cond_NegContext) Condicion() ICondicionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondicionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICondicionContext)
}

func (s *Cond_NegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterCond_Neg(s)
	}
}

func (s *Cond_NegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitCond_Neg(s)
	}
}

func (s *Cond_NegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitCond_Neg(s)

	default:
		return t.VisitChildren(s)
	}
}

type Cond_RelContext struct {
	CondicionContext
	e1 IEContext
	op antlr.Token
	e2 IEContext
}

func NewCond_RelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Cond_RelContext {
	var p = new(Cond_RelContext)

	InitEmptyCondicionContext(&p.CondicionContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondicionContext))

	return p
}

func (s *Cond_RelContext) GetOp() antlr.Token { return s.op }

func (s *Cond_RelContext) SetOp(v antlr.Token) { s.op = v }

func (s *Cond_RelContext) GetE1() IEContext { return s.e1 }

func (s *Cond_RelContext) GetE2() IEContext { return s.e2 }

func (s *Cond_RelContext) SetE1(v IEContext) { s.e1 = v }

func (s *Cond_RelContext) SetE2(v IEContext) { s.e2 = v }

func (s *Cond_RelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cond_RelContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Cond_RelContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
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

	return t.(IEContext)
}

func (s *Cond_RelContext) IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIGUALIGUAL, 0)
}

func (s *Cond_RelContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDIFERENTE, 0)
}

func (s *Cond_RelContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMAYORIGUAL, 0)
}

func (s *Cond_RelContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMAYOR, 0)
}

func (s *Cond_RelContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMENORIGUAL, 0)
}

func (s *Cond_RelContext) MENOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMENOR, 0)
}

func (s *Cond_RelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterCond_Rel(s)
	}
}

func (s *Cond_RelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitCond_Rel(s)
	}
}

func (s *Cond_RelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitCond_Rel(s)

	default:
		return t.VisitChildren(s)
	}
}

type Cond_BooleanoContext struct {
	CondicionContext
	op antlr.Token
}

func NewCond_BooleanoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Cond_BooleanoContext {
	var p = new(Cond_BooleanoContext)

	InitEmptyCondicionContext(&p.CondicionContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondicionContext))

	return p
}

func (s *Cond_BooleanoContext) GetOp() antlr.Token { return s.op }

func (s *Cond_BooleanoContext) SetOp(v antlr.Token) { s.op = v }

func (s *Cond_BooleanoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cond_BooleanoContext) TRUE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserTRUE, 0)
}

func (s *Cond_BooleanoContext) FALSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserFALSE, 0)
}

func (s *Cond_BooleanoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterCond_Booleano(s)
	}
}

func (s *Cond_BooleanoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitCond_Booleano(s)
	}
}

func (s *Cond_BooleanoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitCond_Booleano(s)

	default:
		return t.VisitChildren(s)
	}
}

type Cond_LogicaContext struct {
	CondicionContext
	c1 ICondicionContext
	op antlr.Token
	c2 ICondicionContext
}

func NewCond_LogicaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Cond_LogicaContext {
	var p = new(Cond_LogicaContext)

	InitEmptyCondicionContext(&p.CondicionContext)
	p.parser = parser
	p.CopyAll(ctx.(*CondicionContext))

	return p
}

func (s *Cond_LogicaContext) GetOp() antlr.Token { return s.op }

func (s *Cond_LogicaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Cond_LogicaContext) GetC1() ICondicionContext { return s.c1 }

func (s *Cond_LogicaContext) GetC2() ICondicionContext { return s.c2 }

func (s *Cond_LogicaContext) SetC1(v ICondicionContext) { s.c1 = v }

func (s *Cond_LogicaContext) SetC2(v ICondicionContext) { s.c2 = v }

func (s *Cond_LogicaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Cond_LogicaContext) AllCondicion() []ICondicionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICondicionContext); ok {
			len++
		}
	}

	tst := make([]ICondicionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICondicionContext); ok {
			tst[i] = t.(ICondicionContext)
			i++
		}
	}

	return tst
}

func (s *Cond_LogicaContext) Condicion(i int) ICondicionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICondicionContext); ok {
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

	return t.(ICondicionContext)
}

func (s *Cond_LogicaContext) AND() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserAND, 0)
}

func (s *Cond_LogicaContext) OR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserOR, 0)
}

func (s *Cond_LogicaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterCond_Logica(s)
	}
}

func (s *Cond_LogicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitCond_Logica(s)
	}
}

func (s *Cond_LogicaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitCond_Logica(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Condicion() (localctx ICondicionContext) {
	return p.condicion(0)
}

func (p *Tswift_GrammarNParser) condicion(_p int) (localctx ICondicionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewCondicionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICondicionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, Tswift_GrammarNParserRULE_condicion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCond_NegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(61)

			var _m = p.Match(Tswift_GrammarNParserNOT)

			localctx.(*Cond_NegContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(62)

			var _x = p.condicion(5)

			localctx.(*Cond_NegContext).c = _x
		}

	case 2:
		localctx = NewCond_RelContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(63)

			var _x = p.e(0)

			localctx.(*Cond_RelContext).e1 = _x
		}
		{
			p.SetState(64)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Cond_RelContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&264241152) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Cond_RelContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(65)

			var _x = p.e(0)

			localctx.(*Cond_RelContext).e2 = _x
		}

	case 3:
		localctx = NewCond_BooleanoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Cond_BooleanoContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarNParserTRUE || _la == Tswift_GrammarNParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Cond_BooleanoContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 4:
		localctx = NewCond_ParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(68)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(69)

			var _x = p.condicion(0)

			localctx.(*Cond_ParContext).c = _x
		}
		{
			p.SetState(70)
			p.Match(Tswift_GrammarNParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewCond_LogicaContext(p, NewCondicionContext(p, _parentctx, _parentState))
			localctx.(*Cond_LogicaContext).c1 = _prevctx

			p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarNParserRULE_condicion)
			p.SetState(74)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(75)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*Cond_LogicaContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == Tswift_GrammarNParserAND || _la == Tswift_GrammarNParserOR) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*Cond_LogicaContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(76)

				var _x = p.condicion(4)

				localctx.(*Cond_LogicaContext).c2 = _x
			}

		}
		p.SetState(81)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext())
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

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_e
	return p
}

func InitEmptyEContext(p *EContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_e
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) CopyAll(ctx *EContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Expr_SumResContext struct {
	EContext
	e1 IEContext
	op antlr.Token
	e2 IEContext
}

func NewExpr_SumResContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_SumResContext {
	var p = new(Expr_SumResContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_SumResContext) GetOp() antlr.Token { return s.op }

func (s *Expr_SumResContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_SumResContext) GetE1() IEContext { return s.e1 }

func (s *Expr_SumResContext) GetE2() IEContext { return s.e2 }

func (s *Expr_SumResContext) SetE1(v IEContext) { s.e1 = v }

func (s *Expr_SumResContext) SetE2(v IEContext) { s.e2 = v }

func (s *Expr_SumResContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_SumResContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_SumResContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
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

	return t.(IEContext)
}

func (s *Expr_SumResContext) MAS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMAS, 0)
}

func (s *Expr_SumResContext) MENOS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMENOS, 0)
}

func (s *Expr_SumResContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_SumRes(s)
	}
}

func (s *Expr_SumResContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_SumRes(s)
	}
}

func (s *Expr_SumResContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_SumRes(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_IdContext struct {
	EContext
	id antlr.Token
}

func NewExpr_IdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_IdContext {
	var p = new(Expr_IdContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_IdContext) GetId() antlr.Token { return s.id }

func (s *Expr_IdContext) SetId(v antlr.Token) { s.id = v }

func (s *Expr_IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_IdContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Expr_IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Id(s)
	}
}

func (s *Expr_IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Id(s)
	}
}

func (s *Expr_IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Id(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_ModContext struct {
	EContext
	e1 IEContext
	op antlr.Token
	e2 IEContext
}

func NewExpr_ModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_ModContext {
	var p = new(Expr_ModContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_ModContext) GetOp() antlr.Token { return s.op }

func (s *Expr_ModContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_ModContext) GetE1() IEContext { return s.e1 }

func (s *Expr_ModContext) GetE2() IEContext { return s.e2 }

func (s *Expr_ModContext) SetE1(v IEContext) { s.e1 = v }

func (s *Expr_ModContext) SetE2(v IEContext) { s.e2 = v }

func (s *Expr_ModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_ModContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_ModContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
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

	return t.(IEContext)
}

func (s *Expr_ModContext) MOD() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMOD, 0)
}

func (s *Expr_ModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Mod(s)
	}
}

func (s *Expr_ModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Mod(s)
	}
}

func (s *Expr_ModContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Mod(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_NegContext struct {
	EContext
	op antlr.Token
	e1 IEContext
}

func NewExpr_NegContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_NegContext {
	var p = new(Expr_NegContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_NegContext) GetOp() antlr.Token { return s.op }

func (s *Expr_NegContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_NegContext) GetE1() IEContext { return s.e1 }

func (s *Expr_NegContext) SetE1(v IEContext) { s.e1 = v }

func (s *Expr_NegContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_NegContext) MENOS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMENOS, 0)
}

func (s *Expr_NegContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_NegContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Neg(s)
	}
}

func (s *Expr_NegContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Neg(s)
	}
}

func (s *Expr_NegContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Neg(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_MulDivContext struct {
	EContext
	e1 IEContext
	op antlr.Token
	e2 IEContext
}

func NewExpr_MulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_MulDivContext {
	var p = new(Expr_MulDivContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_MulDivContext) GetOp() antlr.Token { return s.op }

func (s *Expr_MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_MulDivContext) GetE1() IEContext { return s.e1 }

func (s *Expr_MulDivContext) GetE2() IEContext { return s.e2 }

func (s *Expr_MulDivContext) SetE1(v IEContext) { s.e1 = v }

func (s *Expr_MulDivContext) SetE2(v IEContext) { s.e2 = v }

func (s *Expr_MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_MulDivContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *Expr_MulDivContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
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

	return t.(IEContext)
}

func (s *Expr_MulDivContext) POR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPOR, 0)
}

func (s *Expr_MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDIV, 0)
}

func (s *Expr_MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_MulDiv(s)
	}
}

func (s *Expr_MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_MulDiv(s)
	}
}

func (s *Expr_MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_MulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_ParContext struct {
	EContext
	e1 IEContext
}

func NewExpr_ParContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_ParContext {
	var p = new(Expr_ParContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_ParContext) GetE1() IEContext { return s.e1 }

func (s *Expr_ParContext) SetE1(v IEContext) { s.e1 = v }

func (s *Expr_ParContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_ParContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARIZQ, 0)
}

func (s *Expr_ParContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARDER, 0)
}

func (s *Expr_ParContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *Expr_ParContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Par(s)
	}
}

func (s *Expr_ParContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Par(s)
	}
}

func (s *Expr_ParContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Par(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_EnteroContext struct {
	EContext
	n antlr.Token
}

func NewExpr_EnteroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_EnteroContext {
	var p = new(Expr_EnteroContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_EnteroContext) GetN() antlr.Token { return s.n }

func (s *Expr_EnteroContext) SetN(v antlr.Token) { s.n = v }

func (s *Expr_EnteroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_EnteroContext) ENTERO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserENTERO, 0)
}

func (s *Expr_EnteroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Entero(s)
	}
}

func (s *Expr_EnteroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Entero(s)
	}
}

func (s *Expr_EnteroContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Entero(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) E() (localctx IEContext) {
	return p.e(0)
}

func (p *Tswift_GrammarNParser) e(_p int) (localctx IEContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, Tswift_GrammarNParserRULE_e, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarNParserMENOS:
		localctx = NewExpr_NegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(83)

			var _m = p.Match(Tswift_GrammarNParserMENOS)

			localctx.(*Expr_NegContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(84)

			var _x = p.e(7)

			localctx.(*Expr_NegContext).e1 = _x
		}

	case Tswift_GrammarNParserID:
		localctx = NewExpr_IdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(85)

			var _m = p.Match(Tswift_GrammarNParserID)

			localctx.(*Expr_IdContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserENTERO:
		localctx = NewExpr_EnteroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(86)

			var _m = p.Match(Tswift_GrammarNParserENTERO)

			localctx.(*Expr_EnteroContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserPARIZQ:
		localctx = NewExpr_ParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(87)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(88)

			var _x = p.e(0)

			localctx.(*Expr_ParContext).e1 = _x
		}
		{
			p.SetState(89)
			p.Match(Tswift_GrammarNParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(104)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(102)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_MulDivContext(p, NewEContext(p, _parentctx, _parentState))
				localctx.(*Expr_MulDivContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarNParserRULE_e)
				p.SetState(93)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(94)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarNParserDIV || _la == Tswift_GrammarNParserPOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(95)

					var _x = p.e(7)

					localctx.(*Expr_MulDivContext).e2 = _x
				}

			case 2:
				localctx = NewExpr_SumResContext(p, NewEContext(p, _parentctx, _parentState))
				localctx.(*Expr_SumResContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarNParserRULE_e)
				p.SetState(96)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(97)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_SumResContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarNParserMENOS || _la == Tswift_GrammarNParserMAS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_SumResContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(98)

					var _x = p.e(6)

					localctx.(*Expr_SumResContext).e2 = _x
				}

			case 3:
				localctx = NewExpr_ModContext(p, NewEContext(p, _parentctx, _parentState))
				localctx.(*Expr_ModContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarNParserRULE_e)
				p.SetState(99)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(100)

					var _m = p.Match(Tswift_GrammarNParserMOD)

					localctx.(*Expr_ModContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(101)

					var _x = p.e(5)

					localctx.(*Expr_ModContext).e2 = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(106)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

func (p *Tswift_GrammarNParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *CondicionContext = nil
		if localctx != nil {
			t = localctx.(*CondicionContext)
		}
		return p.Condicion_Sempred(t, predIndex)

	case 6:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Tswift_GrammarNParser) Condicion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Tswift_GrammarNParser) E_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
