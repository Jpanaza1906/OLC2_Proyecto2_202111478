// Code generated from Tswift_GrammarN.g4 by ANTLR 4.13.0. DO NOT EDIT.

package TswiftGen

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type Tswift_GrammarNLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var Tswift_GrammarNLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tswift_grammarnlexerLexerInit() {
	staticData := &Tswift_GrammarNLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "')'", "'('", "'{'", "'}'", "'['", "']'", "':'", "';'", "'?'", "','",
		"'.'", "'_'", "'&'", "'+='", "'-='", "'='", "'/'", "'%'", "'-'", "'+'",
		"'*'", "'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "'&&'", "'||'",
		"'!'", "'print'", "'var'", "'let'", "'true'", "'false'", "'if'", "'else'",
		"'switch'", "'case'", "'default'", "'while'", "'for'", "'in'", "'...'",
		"'guard'", "'continue'", "'return'", "'break'", "'nil'", "'append'",
		"'removeLast'", "'remove'", "'at'", "'IsEmpty'", "'count'", "'func'",
		"'repeating'", "'struct'", "'mutating'", "'inout'", "'atoi'", "'iota'",
		"'self'", "'Int'", "'Float'", "'Bool'", "'Char'", "'String'",
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
		"INOUT", "ATOI", "IOTA", "SELF", "INT", "FLOAT", "BOOL", "CHAR", "STRING",
		"ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID", "UL_C",
		"ML_C", "ERROR",
	}
	staticData.RuleNames = []string{
		"PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER",
		"DOSPT", "PTCOMA", "INTERROGACION", "COMA", "PUNTO", "GUIONBAJO", "DIR",
		"MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR",
		"IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR",
		"AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "IF", "ELSE",
		"SWITCH", "CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", "GUARD",
		"CONTINUE", "RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", "REMOVE",
		"AT", "ISEMPTY", "COUNT", "FUNC", "REPEATING", "STRUCT", "MUTATING",
		"INOUT", "ATOI", "IOTA", "SELF", "INT", "FLOAT", "BOOL", "CHAR", "STRING",
		"ENBLANCO", "DIG", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID",
		"UL_C", "ML_C", "ERROR",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 77, 536, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7,
		41, 2, 42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46,
		2, 47, 7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2,
		52, 7, 52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57,
		7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7,
		62, 2, 63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67,
		2, 68, 7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2,
		73, 7, 73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 2, 77, 7, 77, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6,
		1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16,
		1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1,
		21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31,
		1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		38, 1, 38, 1, 38, 1, 38, 1, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		1, 39, 1, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 41, 1, 41, 1,
		41, 1, 41, 1, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44,
		1, 44, 1, 44, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1,
		45, 1, 45, 1, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47,
		1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1, 48, 1, 48, 1, 49, 1,
		49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50,
		1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1, 51, 1, 51, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53,
		1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 54, 1, 55, 1,
		55, 1, 55, 1, 55, 1, 55, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56, 1, 56,
		1, 56, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1, 57, 1,
		58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 58, 1, 59, 1, 59,
		1, 59, 1, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 61, 1,
		61, 1, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62, 1, 62, 1, 63, 1, 63,
		1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 65, 1, 65, 1,
		65, 1, 65, 1, 65, 1, 66, 1, 66, 1, 66, 1, 66, 1, 66, 1, 67, 1, 67, 1, 67,
		1, 67, 1, 67, 1, 67, 1, 67, 1, 68, 4, 68, 450, 8, 68, 11, 68, 12, 68, 451,
		1, 68, 1, 68, 1, 69, 1, 69, 1, 70, 4, 70, 459, 8, 70, 11, 70, 12, 70, 460,
		1, 71, 4, 71, 464, 8, 71, 11, 71, 12, 71, 465, 1, 71, 1, 71, 4, 71, 470,
		8, 71, 11, 71, 12, 71, 471, 1, 72, 1, 72, 1, 72, 1, 72, 3, 72, 478, 8,
		72, 1, 72, 1, 72, 1, 73, 1, 73, 1, 73, 1, 73, 5, 73, 486, 8, 73, 10, 73,
		12, 73, 489, 9, 73, 1, 73, 1, 73, 1, 74, 1, 74, 5, 74, 495, 8, 74, 10,
		74, 12, 74, 498, 9, 74, 1, 74, 1, 74, 4, 74, 502, 8, 74, 11, 74, 12, 74,
		503, 3, 74, 506, 8, 74, 1, 75, 1, 75, 1, 75, 1, 75, 5, 75, 512, 8, 75,
		10, 75, 12, 75, 515, 9, 75, 1, 75, 1, 75, 1, 76, 1, 76, 1, 76, 1, 76, 5,
		76, 523, 8, 76, 10, 76, 12, 76, 526, 9, 76, 1, 76, 1, 76, 1, 76, 1, 76,
		1, 76, 1, 77, 1, 77, 1, 77, 1, 77, 1, 524, 0, 78, 1, 1, 3, 2, 5, 3, 7,
		4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27,
		14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45,
		23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63,
		32, 65, 33, 67, 34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81,
		41, 83, 42, 85, 43, 87, 44, 89, 45, 91, 46, 93, 47, 95, 48, 97, 49, 99,
		50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111, 56, 113, 57, 115,
		58, 117, 59, 119, 60, 121, 61, 123, 62, 125, 63, 127, 64, 129, 65, 131,
		66, 133, 67, 135, 68, 137, 69, 139, 0, 141, 70, 143, 71, 145, 72, 147,
		73, 149, 74, 151, 75, 153, 76, 155, 77, 1, 0, 7, 3, 0, 9, 10, 13, 13, 32,
		32, 1, 0, 48, 57, 2, 0, 39, 39, 92, 92, 2, 0, 34, 34, 92, 92, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 10, 10,
		13, 13, 546, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7,
		1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0,
		15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0,
		0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0,
		0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0,
		0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1,
		0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
		1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0,
		61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0,
		0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1, 0, 0,
		0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0,
		0, 0, 0, 85, 1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1,
		0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0, 0, 0, 0, 97, 1, 0, 0, 0, 0, 99,
		1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0, 0,
		0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1,
		0, 0, 0, 0, 115, 1, 0, 0, 0, 0, 117, 1, 0, 0, 0, 0, 119, 1, 0, 0, 0, 0,
		121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 0, 125, 1, 0, 0, 0, 0, 127, 1, 0,
		0, 0, 0, 129, 1, 0, 0, 0, 0, 131, 1, 0, 0, 0, 0, 133, 1, 0, 0, 0, 0, 135,
		1, 0, 0, 0, 0, 137, 1, 0, 0, 0, 0, 141, 1, 0, 0, 0, 0, 143, 1, 0, 0, 0,
		0, 145, 1, 0, 0, 0, 0, 147, 1, 0, 0, 0, 0, 149, 1, 0, 0, 0, 0, 151, 1,
		0, 0, 0, 0, 153, 1, 0, 0, 0, 0, 155, 1, 0, 0, 0, 1, 157, 1, 0, 0, 0, 3,
		159, 1, 0, 0, 0, 5, 161, 1, 0, 0, 0, 7, 163, 1, 0, 0, 0, 9, 165, 1, 0,
		0, 0, 11, 167, 1, 0, 0, 0, 13, 169, 1, 0, 0, 0, 15, 171, 1, 0, 0, 0, 17,
		173, 1, 0, 0, 0, 19, 175, 1, 0, 0, 0, 21, 177, 1, 0, 0, 0, 23, 179, 1,
		0, 0, 0, 25, 181, 1, 0, 0, 0, 27, 183, 1, 0, 0, 0, 29, 186, 1, 0, 0, 0,
		31, 189, 1, 0, 0, 0, 33, 191, 1, 0, 0, 0, 35, 193, 1, 0, 0, 0, 37, 195,
		1, 0, 0, 0, 39, 197, 1, 0, 0, 0, 41, 199, 1, 0, 0, 0, 43, 201, 1, 0, 0,
		0, 45, 204, 1, 0, 0, 0, 47, 207, 1, 0, 0, 0, 49, 210, 1, 0, 0, 0, 51, 213,
		1, 0, 0, 0, 53, 215, 1, 0, 0, 0, 55, 217, 1, 0, 0, 0, 57, 220, 1, 0, 0,
		0, 59, 223, 1, 0, 0, 0, 61, 225, 1, 0, 0, 0, 63, 231, 1, 0, 0, 0, 65, 235,
		1, 0, 0, 0, 67, 239, 1, 0, 0, 0, 69, 244, 1, 0, 0, 0, 71, 250, 1, 0, 0,
		0, 73, 253, 1, 0, 0, 0, 75, 258, 1, 0, 0, 0, 77, 265, 1, 0, 0, 0, 79, 270,
		1, 0, 0, 0, 81, 278, 1, 0, 0, 0, 83, 284, 1, 0, 0, 0, 85, 288, 1, 0, 0,
		0, 87, 291, 1, 0, 0, 0, 89, 295, 1, 0, 0, 0, 91, 301, 1, 0, 0, 0, 93, 310,
		1, 0, 0, 0, 95, 317, 1, 0, 0, 0, 97, 323, 1, 0, 0, 0, 99, 327, 1, 0, 0,
		0, 101, 334, 1, 0, 0, 0, 103, 345, 1, 0, 0, 0, 105, 352, 1, 0, 0, 0, 107,
		355, 1, 0, 0, 0, 109, 363, 1, 0, 0, 0, 111, 369, 1, 0, 0, 0, 113, 374,
		1, 0, 0, 0, 115, 384, 1, 0, 0, 0, 117, 391, 1, 0, 0, 0, 119, 400, 1, 0,
		0, 0, 121, 406, 1, 0, 0, 0, 123, 411, 1, 0, 0, 0, 125, 416, 1, 0, 0, 0,
		127, 421, 1, 0, 0, 0, 129, 425, 1, 0, 0, 0, 131, 431, 1, 0, 0, 0, 133,
		436, 1, 0, 0, 0, 135, 441, 1, 0, 0, 0, 137, 449, 1, 0, 0, 0, 139, 455,
		1, 0, 0, 0, 141, 458, 1, 0, 0, 0, 143, 463, 1, 0, 0, 0, 145, 473, 1, 0,
		0, 0, 147, 481, 1, 0, 0, 0, 149, 505, 1, 0, 0, 0, 151, 507, 1, 0, 0, 0,
		153, 518, 1, 0, 0, 0, 155, 532, 1, 0, 0, 0, 157, 158, 5, 41, 0, 0, 158,
		2, 1, 0, 0, 0, 159, 160, 5, 40, 0, 0, 160, 4, 1, 0, 0, 0, 161, 162, 5,
		123, 0, 0, 162, 6, 1, 0, 0, 0, 163, 164, 5, 125, 0, 0, 164, 8, 1, 0, 0,
		0, 165, 166, 5, 91, 0, 0, 166, 10, 1, 0, 0, 0, 167, 168, 5, 93, 0, 0, 168,
		12, 1, 0, 0, 0, 169, 170, 5, 58, 0, 0, 170, 14, 1, 0, 0, 0, 171, 172, 5,
		59, 0, 0, 172, 16, 1, 0, 0, 0, 173, 174, 5, 63, 0, 0, 174, 18, 1, 0, 0,
		0, 175, 176, 5, 44, 0, 0, 176, 20, 1, 0, 0, 0, 177, 178, 5, 46, 0, 0, 178,
		22, 1, 0, 0, 0, 179, 180, 5, 95, 0, 0, 180, 24, 1, 0, 0, 0, 181, 182, 5,
		38, 0, 0, 182, 26, 1, 0, 0, 0, 183, 184, 5, 43, 0, 0, 184, 185, 5, 61,
		0, 0, 185, 28, 1, 0, 0, 0, 186, 187, 5, 45, 0, 0, 187, 188, 5, 61, 0, 0,
		188, 30, 1, 0, 0, 0, 189, 190, 5, 61, 0, 0, 190, 32, 1, 0, 0, 0, 191, 192,
		5, 47, 0, 0, 192, 34, 1, 0, 0, 0, 193, 194, 5, 37, 0, 0, 194, 36, 1, 0,
		0, 0, 195, 196, 5, 45, 0, 0, 196, 38, 1, 0, 0, 0, 197, 198, 5, 43, 0, 0,
		198, 40, 1, 0, 0, 0, 199, 200, 5, 42, 0, 0, 200, 42, 1, 0, 0, 0, 201, 202,
		5, 61, 0, 0, 202, 203, 5, 61, 0, 0, 203, 44, 1, 0, 0, 0, 204, 205, 5, 33,
		0, 0, 205, 206, 5, 61, 0, 0, 206, 46, 1, 0, 0, 0, 207, 208, 5, 62, 0, 0,
		208, 209, 5, 61, 0, 0, 209, 48, 1, 0, 0, 0, 210, 211, 5, 60, 0, 0, 211,
		212, 5, 61, 0, 0, 212, 50, 1, 0, 0, 0, 213, 214, 5, 62, 0, 0, 214, 52,
		1, 0, 0, 0, 215, 216, 5, 60, 0, 0, 216, 54, 1, 0, 0, 0, 217, 218, 5, 38,
		0, 0, 218, 219, 5, 38, 0, 0, 219, 56, 1, 0, 0, 0, 220, 221, 5, 124, 0,
		0, 221, 222, 5, 124, 0, 0, 222, 58, 1, 0, 0, 0, 223, 224, 5, 33, 0, 0,
		224, 60, 1, 0, 0, 0, 225, 226, 5, 112, 0, 0, 226, 227, 5, 114, 0, 0, 227,
		228, 5, 105, 0, 0, 228, 229, 5, 110, 0, 0, 229, 230, 5, 116, 0, 0, 230,
		62, 1, 0, 0, 0, 231, 232, 5, 118, 0, 0, 232, 233, 5, 97, 0, 0, 233, 234,
		5, 114, 0, 0, 234, 64, 1, 0, 0, 0, 235, 236, 5, 108, 0, 0, 236, 237, 5,
		101, 0, 0, 237, 238, 5, 116, 0, 0, 238, 66, 1, 0, 0, 0, 239, 240, 5, 116,
		0, 0, 240, 241, 5, 114, 0, 0, 241, 242, 5, 117, 0, 0, 242, 243, 5, 101,
		0, 0, 243, 68, 1, 0, 0, 0, 244, 245, 5, 102, 0, 0, 245, 246, 5, 97, 0,
		0, 246, 247, 5, 108, 0, 0, 247, 248, 5, 115, 0, 0, 248, 249, 5, 101, 0,
		0, 249, 70, 1, 0, 0, 0, 250, 251, 5, 105, 0, 0, 251, 252, 5, 102, 0, 0,
		252, 72, 1, 0, 0, 0, 253, 254, 5, 101, 0, 0, 254, 255, 5, 108, 0, 0, 255,
		256, 5, 115, 0, 0, 256, 257, 5, 101, 0, 0, 257, 74, 1, 0, 0, 0, 258, 259,
		5, 115, 0, 0, 259, 260, 5, 119, 0, 0, 260, 261, 5, 105, 0, 0, 261, 262,
		5, 116, 0, 0, 262, 263, 5, 99, 0, 0, 263, 264, 5, 104, 0, 0, 264, 76, 1,
		0, 0, 0, 265, 266, 5, 99, 0, 0, 266, 267, 5, 97, 0, 0, 267, 268, 5, 115,
		0, 0, 268, 269, 5, 101, 0, 0, 269, 78, 1, 0, 0, 0, 270, 271, 5, 100, 0,
		0, 271, 272, 5, 101, 0, 0, 272, 273, 5, 102, 0, 0, 273, 274, 5, 97, 0,
		0, 274, 275, 5, 117, 0, 0, 275, 276, 5, 108, 0, 0, 276, 277, 5, 116, 0,
		0, 277, 80, 1, 0, 0, 0, 278, 279, 5, 119, 0, 0, 279, 280, 5, 104, 0, 0,
		280, 281, 5, 105, 0, 0, 281, 282, 5, 108, 0, 0, 282, 283, 5, 101, 0, 0,
		283, 82, 1, 0, 0, 0, 284, 285, 5, 102, 0, 0, 285, 286, 5, 111, 0, 0, 286,
		287, 5, 114, 0, 0, 287, 84, 1, 0, 0, 0, 288, 289, 5, 105, 0, 0, 289, 290,
		5, 110, 0, 0, 290, 86, 1, 0, 0, 0, 291, 292, 5, 46, 0, 0, 292, 293, 5,
		46, 0, 0, 293, 294, 5, 46, 0, 0, 294, 88, 1, 0, 0, 0, 295, 296, 5, 103,
		0, 0, 296, 297, 5, 117, 0, 0, 297, 298, 5, 97, 0, 0, 298, 299, 5, 114,
		0, 0, 299, 300, 5, 100, 0, 0, 300, 90, 1, 0, 0, 0, 301, 302, 5, 99, 0,
		0, 302, 303, 5, 111, 0, 0, 303, 304, 5, 110, 0, 0, 304, 305, 5, 116, 0,
		0, 305, 306, 5, 105, 0, 0, 306, 307, 5, 110, 0, 0, 307, 308, 5, 117, 0,
		0, 308, 309, 5, 101, 0, 0, 309, 92, 1, 0, 0, 0, 310, 311, 5, 114, 0, 0,
		311, 312, 5, 101, 0, 0, 312, 313, 5, 116, 0, 0, 313, 314, 5, 117, 0, 0,
		314, 315, 5, 114, 0, 0, 315, 316, 5, 110, 0, 0, 316, 94, 1, 0, 0, 0, 317,
		318, 5, 98, 0, 0, 318, 319, 5, 114, 0, 0, 319, 320, 5, 101, 0, 0, 320,
		321, 5, 97, 0, 0, 321, 322, 5, 107, 0, 0, 322, 96, 1, 0, 0, 0, 323, 324,
		5, 110, 0, 0, 324, 325, 5, 105, 0, 0, 325, 326, 5, 108, 0, 0, 326, 98,
		1, 0, 0, 0, 327, 328, 5, 97, 0, 0, 328, 329, 5, 112, 0, 0, 329, 330, 5,
		112, 0, 0, 330, 331, 5, 101, 0, 0, 331, 332, 5, 110, 0, 0, 332, 333, 5,
		100, 0, 0, 333, 100, 1, 0, 0, 0, 334, 335, 5, 114, 0, 0, 335, 336, 5, 101,
		0, 0, 336, 337, 5, 109, 0, 0, 337, 338, 5, 111, 0, 0, 338, 339, 5, 118,
		0, 0, 339, 340, 5, 101, 0, 0, 340, 341, 5, 76, 0, 0, 341, 342, 5, 97, 0,
		0, 342, 343, 5, 115, 0, 0, 343, 344, 5, 116, 0, 0, 344, 102, 1, 0, 0, 0,
		345, 346, 5, 114, 0, 0, 346, 347, 5, 101, 0, 0, 347, 348, 5, 109, 0, 0,
		348, 349, 5, 111, 0, 0, 349, 350, 5, 118, 0, 0, 350, 351, 5, 101, 0, 0,
		351, 104, 1, 0, 0, 0, 352, 353, 5, 97, 0, 0, 353, 354, 5, 116, 0, 0, 354,
		106, 1, 0, 0, 0, 355, 356, 5, 73, 0, 0, 356, 357, 5, 115, 0, 0, 357, 358,
		5, 69, 0, 0, 358, 359, 5, 109, 0, 0, 359, 360, 5, 112, 0, 0, 360, 361,
		5, 116, 0, 0, 361, 362, 5, 121, 0, 0, 362, 108, 1, 0, 0, 0, 363, 364, 5,
		99, 0, 0, 364, 365, 5, 111, 0, 0, 365, 366, 5, 117, 0, 0, 366, 367, 5,
		110, 0, 0, 367, 368, 5, 116, 0, 0, 368, 110, 1, 0, 0, 0, 369, 370, 5, 102,
		0, 0, 370, 371, 5, 117, 0, 0, 371, 372, 5, 110, 0, 0, 372, 373, 5, 99,
		0, 0, 373, 112, 1, 0, 0, 0, 374, 375, 5, 114, 0, 0, 375, 376, 5, 101, 0,
		0, 376, 377, 5, 112, 0, 0, 377, 378, 5, 101, 0, 0, 378, 379, 5, 97, 0,
		0, 379, 380, 5, 116, 0, 0, 380, 381, 5, 105, 0, 0, 381, 382, 5, 110, 0,
		0, 382, 383, 5, 103, 0, 0, 383, 114, 1, 0, 0, 0, 384, 385, 5, 115, 0, 0,
		385, 386, 5, 116, 0, 0, 386, 387, 5, 114, 0, 0, 387, 388, 5, 117, 0, 0,
		388, 389, 5, 99, 0, 0, 389, 390, 5, 116, 0, 0, 390, 116, 1, 0, 0, 0, 391,
		392, 5, 109, 0, 0, 392, 393, 5, 117, 0, 0, 393, 394, 5, 116, 0, 0, 394,
		395, 5, 97, 0, 0, 395, 396, 5, 116, 0, 0, 396, 397, 5, 105, 0, 0, 397,
		398, 5, 110, 0, 0, 398, 399, 5, 103, 0, 0, 399, 118, 1, 0, 0, 0, 400, 401,
		5, 105, 0, 0, 401, 402, 5, 110, 0, 0, 402, 403, 5, 111, 0, 0, 403, 404,
		5, 117, 0, 0, 404, 405, 5, 116, 0, 0, 405, 120, 1, 0, 0, 0, 406, 407, 5,
		97, 0, 0, 407, 408, 5, 116, 0, 0, 408, 409, 5, 111, 0, 0, 409, 410, 5,
		105, 0, 0, 410, 122, 1, 0, 0, 0, 411, 412, 5, 105, 0, 0, 412, 413, 5, 111,
		0, 0, 413, 414, 5, 116, 0, 0, 414, 415, 5, 97, 0, 0, 415, 124, 1, 0, 0,
		0, 416, 417, 5, 115, 0, 0, 417, 418, 5, 101, 0, 0, 418, 419, 5, 108, 0,
		0, 419, 420, 5, 102, 0, 0, 420, 126, 1, 0, 0, 0, 421, 422, 5, 73, 0, 0,
		422, 423, 5, 110, 0, 0, 423, 424, 5, 116, 0, 0, 424, 128, 1, 0, 0, 0, 425,
		426, 5, 70, 0, 0, 426, 427, 5, 108, 0, 0, 427, 428, 5, 111, 0, 0, 428,
		429, 5, 97, 0, 0, 429, 430, 5, 116, 0, 0, 430, 130, 1, 0, 0, 0, 431, 432,
		5, 66, 0, 0, 432, 433, 5, 111, 0, 0, 433, 434, 5, 111, 0, 0, 434, 435,
		5, 108, 0, 0, 435, 132, 1, 0, 0, 0, 436, 437, 5, 67, 0, 0, 437, 438, 5,
		104, 0, 0, 438, 439, 5, 97, 0, 0, 439, 440, 5, 114, 0, 0, 440, 134, 1,
		0, 0, 0, 441, 442, 5, 83, 0, 0, 442, 443, 5, 116, 0, 0, 443, 444, 5, 114,
		0, 0, 444, 445, 5, 105, 0, 0, 445, 446, 5, 110, 0, 0, 446, 447, 5, 103,
		0, 0, 447, 136, 1, 0, 0, 0, 448, 450, 7, 0, 0, 0, 449, 448, 1, 0, 0, 0,
		450, 451, 1, 0, 0, 0, 451, 449, 1, 0, 0, 0, 451, 452, 1, 0, 0, 0, 452,
		453, 1, 0, 0, 0, 453, 454, 6, 68, 0, 0, 454, 138, 1, 0, 0, 0, 455, 456,
		7, 1, 0, 0, 456, 140, 1, 0, 0, 0, 457, 459, 3, 139, 69, 0, 458, 457, 1,
		0, 0, 0, 459, 460, 1, 0, 0, 0, 460, 458, 1, 0, 0, 0, 460, 461, 1, 0, 0,
		0, 461, 142, 1, 0, 0, 0, 462, 464, 3, 139, 69, 0, 463, 462, 1, 0, 0, 0,
		464, 465, 1, 0, 0, 0, 465, 463, 1, 0, 0, 0, 465, 466, 1, 0, 0, 0, 466,
		467, 1, 0, 0, 0, 467, 469, 5, 46, 0, 0, 468, 470, 3, 139, 69, 0, 469, 468,
		1, 0, 0, 0, 470, 471, 1, 0, 0, 0, 471, 469, 1, 0, 0, 0, 471, 472, 1, 0,
		0, 0, 472, 144, 1, 0, 0, 0, 473, 477, 5, 39, 0, 0, 474, 478, 8, 2, 0, 0,
		475, 476, 5, 92, 0, 0, 476, 478, 9, 0, 0, 0, 477, 474, 1, 0, 0, 0, 477,
		475, 1, 0, 0, 0, 478, 479, 1, 0, 0, 0, 479, 480, 5, 39, 0, 0, 480, 146,
		1, 0, 0, 0, 481, 487, 5, 34, 0, 0, 482, 486, 8, 3, 0, 0, 483, 484, 5, 92,
		0, 0, 484, 486, 9, 0, 0, 0, 485, 482, 1, 0, 0, 0, 485, 483, 1, 0, 0, 0,
		486, 489, 1, 0, 0, 0, 487, 485, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488,
		490, 1, 0, 0, 0, 489, 487, 1, 0, 0, 0, 490, 491, 5, 34, 0, 0, 491, 148,
		1, 0, 0, 0, 492, 496, 7, 4, 0, 0, 493, 495, 7, 5, 0, 0, 494, 493, 1, 0,
		0, 0, 495, 498, 1, 0, 0, 0, 496, 494, 1, 0, 0, 0, 496, 497, 1, 0, 0, 0,
		497, 506, 1, 0, 0, 0, 498, 496, 1, 0, 0, 0, 499, 501, 5, 95, 0, 0, 500,
		502, 7, 5, 0, 0, 501, 500, 1, 0, 0, 0, 502, 503, 1, 0, 0, 0, 503, 501,
		1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504, 506, 1, 0, 0, 0, 505, 492, 1, 0,
		0, 0, 505, 499, 1, 0, 0, 0, 506, 150, 1, 0, 0, 0, 507, 508, 5, 47, 0, 0,
		508, 509, 5, 47, 0, 0, 509, 513, 1, 0, 0, 0, 510, 512, 8, 6, 0, 0, 511,
		510, 1, 0, 0, 0, 512, 515, 1, 0, 0, 0, 513, 511, 1, 0, 0, 0, 513, 514,
		1, 0, 0, 0, 514, 516, 1, 0, 0, 0, 515, 513, 1, 0, 0, 0, 516, 517, 6, 75,
		1, 0, 517, 152, 1, 0, 0, 0, 518, 519, 5, 47, 0, 0, 519, 520, 5, 42, 0,
		0, 520, 524, 1, 0, 0, 0, 521, 523, 9, 0, 0, 0, 522, 521, 1, 0, 0, 0, 523,
		526, 1, 0, 0, 0, 524, 525, 1, 0, 0, 0, 524, 522, 1, 0, 0, 0, 525, 527,
		1, 0, 0, 0, 526, 524, 1, 0, 0, 0, 527, 528, 5, 42, 0, 0, 528, 529, 5, 47,
		0, 0, 529, 530, 1, 0, 0, 0, 530, 531, 6, 76, 1, 0, 531, 154, 1, 0, 0, 0,
		532, 533, 9, 0, 0, 0, 533, 534, 1, 0, 0, 0, 534, 535, 6, 77, 1, 0, 535,
		156, 1, 0, 0, 0, 13, 0, 451, 460, 465, 471, 477, 485, 487, 496, 503, 505,
		513, 524, 2, 6, 0, 0, 0, 1, 0,
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

// Tswift_GrammarNLexerInit initializes any static state used to implement Tswift_GrammarNLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTswift_GrammarNLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func Tswift_GrammarNLexerInit() {
	staticData := &Tswift_GrammarNLexerLexerStaticData
	staticData.once.Do(tswift_grammarnlexerLexerInit)
}

// NewTswift_GrammarNLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTswift_GrammarNLexer(input antlr.CharStream) *Tswift_GrammarNLexer {
	Tswift_GrammarNLexerInit()
	l := new(Tswift_GrammarNLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &Tswift_GrammarNLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Tswift_GrammarN.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Tswift_GrammarNLexer tokens.
const (
	Tswift_GrammarNLexerPARDER        = 1
	Tswift_GrammarNLexerPARIZQ        = 2
	Tswift_GrammarNLexerLLAVEIZQ      = 3
	Tswift_GrammarNLexerLLAVEDER      = 4
	Tswift_GrammarNLexerCORCHETEIZQ   = 5
	Tswift_GrammarNLexerCORCHETEDER   = 6
	Tswift_GrammarNLexerDOSPT         = 7
	Tswift_GrammarNLexerPTCOMA        = 8
	Tswift_GrammarNLexerINTERROGACION = 9
	Tswift_GrammarNLexerCOMA          = 10
	Tswift_GrammarNLexerPUNTO         = 11
	Tswift_GrammarNLexerGUIONBAJO     = 12
	Tswift_GrammarNLexerDIR           = 13
	Tswift_GrammarNLexerMASIGUAL      = 14
	Tswift_GrammarNLexerMENOSIGUAL    = 15
	Tswift_GrammarNLexerIGUAL         = 16
	Tswift_GrammarNLexerDIV           = 17
	Tswift_GrammarNLexerMOD           = 18
	Tswift_GrammarNLexerMENOS         = 19
	Tswift_GrammarNLexerMAS           = 20
	Tswift_GrammarNLexerPOR           = 21
	Tswift_GrammarNLexerIGUALIGUAL    = 22
	Tswift_GrammarNLexerDIFERENTE     = 23
	Tswift_GrammarNLexerMAYORIGUAL    = 24
	Tswift_GrammarNLexerMENORIGUAL    = 25
	Tswift_GrammarNLexerMAYOR         = 26
	Tswift_GrammarNLexerMENOR         = 27
	Tswift_GrammarNLexerAND           = 28
	Tswift_GrammarNLexerOR            = 29
	Tswift_GrammarNLexerNOT           = 30
	Tswift_GrammarNLexerPRINT         = 31
	Tswift_GrammarNLexerVAR           = 32
	Tswift_GrammarNLexerLET           = 33
	Tswift_GrammarNLexerTRUE          = 34
	Tswift_GrammarNLexerFALSE         = 35
	Tswift_GrammarNLexerIF            = 36
	Tswift_GrammarNLexerELSE          = 37
	Tswift_GrammarNLexerSWITCH        = 38
	Tswift_GrammarNLexerCASE          = 39
	Tswift_GrammarNLexerDEFAULT       = 40
	Tswift_GrammarNLexerWHILE         = 41
	Tswift_GrammarNLexerFOR           = 42
	Tswift_GrammarNLexerIN            = 43
	Tswift_GrammarNLexerRANGO         = 44
	Tswift_GrammarNLexerGUARD         = 45
	Tswift_GrammarNLexerCONTINUE      = 46
	Tswift_GrammarNLexerRETURN        = 47
	Tswift_GrammarNLexerBREAK         = 48
	Tswift_GrammarNLexerNIL           = 49
	Tswift_GrammarNLexerAPPEND        = 50
	Tswift_GrammarNLexerREMOVELAST    = 51
	Tswift_GrammarNLexerREMOVE        = 52
	Tswift_GrammarNLexerAT            = 53
	Tswift_GrammarNLexerISEMPTY       = 54
	Tswift_GrammarNLexerCOUNT         = 55
	Tswift_GrammarNLexerFUNC          = 56
	Tswift_GrammarNLexerREPEATING     = 57
	Tswift_GrammarNLexerSTRUCT        = 58
	Tswift_GrammarNLexerMUTATING      = 59
	Tswift_GrammarNLexerINOUT         = 60
	Tswift_GrammarNLexerATOI          = 61
	Tswift_GrammarNLexerIOTA          = 62
	Tswift_GrammarNLexerSELF          = 63
	Tswift_GrammarNLexerINT           = 64
	Tswift_GrammarNLexerFLOAT         = 65
	Tswift_GrammarNLexerBOOL          = 66
	Tswift_GrammarNLexerCHAR          = 67
	Tswift_GrammarNLexerSTRING        = 68
	Tswift_GrammarNLexerENBLANCO      = 69
	Tswift_GrammarNLexerENTERO        = 70
	Tswift_GrammarNLexerDECIMAL       = 71
	Tswift_GrammarNLexerCARACTER      = 72
	Tswift_GrammarNLexerCADENA        = 73
	Tswift_GrammarNLexerID            = 74
	Tswift_GrammarNLexerUL_C          = 75
	Tswift_GrammarNLexerML_C          = 76
	Tswift_GrammarNLexerERROR         = 77
)
