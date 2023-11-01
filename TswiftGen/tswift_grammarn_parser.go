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
		"s", "l_sentencias", "sentencia", "trans_sentencia", "print_sentencia",
		"declaracion_sentencia", "asignacion_sentencia", "tipo", "if_sentencia",
		"guard_sentencia", "switch_sentencia", "lcasos", "cdefault", "while_sentencia",
		"for_sentencia", "dec_vector", "def_vector", "func_vector", "dec_matriz",
		"def_matriz", "listavalores_matriz", "simple_matriz", "condicion", "e",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 77, 428, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 1, 5, 1, 52, 8, 1,
		10, 1, 12, 1, 55, 9, 1, 1, 2, 1, 2, 3, 2, 59, 8, 2, 1, 2, 1, 2, 3, 2, 63,
		8, 2, 1, 2, 1, 2, 3, 2, 67, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 74,
		8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 80, 8, 2, 1, 2, 1, 2, 3, 2, 84, 8,
		2, 1, 2, 1, 2, 3, 2, 88, 8, 2, 3, 2, 90, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		3, 3, 96, 8, 3, 3, 3, 98, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 105,
		8, 4, 10, 4, 12, 4, 108, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		3, 5, 129, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 137, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 149, 8,
		7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3,
		8, 162, 8, 8, 3, 8, 164, 8, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9,
		1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 178, 8, 10, 10, 10, 12, 10, 181,
		9, 10, 1, 10, 3, 10, 184, 8, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 3, 14, 221,
		8, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 3, 15, 241,
		8, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 247, 8, 16, 10, 16, 12, 16, 250,
		9, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 257, 8, 16, 1, 17, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 280,
		8, 17, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 286, 8, 18, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 3, 19, 293, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20,
		299, 8, 20, 10, 20, 12, 20, 302, 9, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1,
		20, 1, 20, 5, 20, 310, 8, 20, 10, 20, 12, 20, 313, 9, 20, 1, 20, 1, 20,
		3, 20, 317, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 345, 8,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 3, 22, 359, 8, 22, 1, 22, 1, 22, 1, 22, 5, 22, 364, 8, 22,
		10, 22, 12, 22, 367, 9, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 4, 23, 387, 8, 23, 11, 23, 12, 23, 388, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 3, 23, 406, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 423,
		8, 23, 10, 23, 12, 23, 426, 9, 23, 1, 23, 0, 2, 44, 46, 24, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 0, 8, 1, 0, 32, 33, 1, 0, 14, 15, 1, 0, 22, 27, 1, 0, 34, 35, 1,
		0, 28, 29, 2, 0, 19, 19, 30, 30, 2, 0, 17, 17, 21, 21, 1, 0, 19, 20, 474,
		0, 48, 1, 0, 0, 0, 2, 53, 1, 0, 0, 0, 4, 89, 1, 0, 0, 0, 6, 97, 1, 0, 0,
		0, 8, 99, 1, 0, 0, 0, 10, 128, 1, 0, 0, 0, 12, 136, 1, 0, 0, 0, 14, 148,
		1, 0, 0, 0, 16, 150, 1, 0, 0, 0, 18, 165, 1, 0, 0, 0, 20, 173, 1, 0, 0,
		0, 22, 187, 1, 0, 0, 0, 24, 192, 1, 0, 0, 0, 26, 196, 1, 0, 0, 0, 28, 220,
		1, 0, 0, 0, 30, 240, 1, 0, 0, 0, 32, 256, 1, 0, 0, 0, 34, 279, 1, 0, 0,
		0, 36, 281, 1, 0, 0, 0, 38, 292, 1, 0, 0, 0, 40, 316, 1, 0, 0, 0, 42, 344,
		1, 0, 0, 0, 44, 358, 1, 0, 0, 0, 46, 405, 1, 0, 0, 0, 48, 49, 3, 2, 1,
		0, 49, 1, 1, 0, 0, 0, 50, 52, 3, 4, 2, 0, 51, 50, 1, 0, 0, 0, 52, 55, 1,
		0, 0, 0, 53, 51, 1, 0, 0, 0, 53, 54, 1, 0, 0, 0, 54, 3, 1, 0, 0, 0, 55,
		53, 1, 0, 0, 0, 56, 58, 3, 8, 4, 0, 57, 59, 5, 8, 0, 0, 58, 57, 1, 0, 0,
		0, 58, 59, 1, 0, 0, 0, 59, 90, 1, 0, 0, 0, 60, 62, 3, 10, 5, 0, 61, 63,
		5, 8, 0, 0, 62, 61, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 90, 1, 0, 0, 0,
		64, 66, 3, 12, 6, 0, 65, 67, 5, 8, 0, 0, 66, 65, 1, 0, 0, 0, 66, 67, 1,
		0, 0, 0, 67, 90, 1, 0, 0, 0, 68, 90, 3, 16, 8, 0, 69, 90, 3, 20, 10, 0,
		70, 90, 3, 18, 9, 0, 71, 73, 3, 6, 3, 0, 72, 74, 5, 8, 0, 0, 73, 72, 1,
		0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 90, 1, 0, 0, 0, 75, 90, 3, 26, 13, 0,
		76, 90, 3, 28, 14, 0, 77, 79, 3, 30, 15, 0, 78, 80, 5, 8, 0, 0, 79, 78,
		1, 0, 0, 0, 79, 80, 1, 0, 0, 0, 80, 90, 1, 0, 0, 0, 81, 83, 3, 34, 17,
		0, 82, 84, 5, 8, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 90,
		1, 0, 0, 0, 85, 87, 3, 36, 18, 0, 86, 88, 5, 8, 0, 0, 87, 86, 1, 0, 0,
		0, 87, 88, 1, 0, 0, 0, 88, 90, 1, 0, 0, 0, 89, 56, 1, 0, 0, 0, 89, 60,
		1, 0, 0, 0, 89, 64, 1, 0, 0, 0, 89, 68, 1, 0, 0, 0, 89, 69, 1, 0, 0, 0,
		89, 70, 1, 0, 0, 0, 89, 71, 1, 0, 0, 0, 89, 75, 1, 0, 0, 0, 89, 76, 1,
		0, 0, 0, 89, 77, 1, 0, 0, 0, 89, 81, 1, 0, 0, 0, 89, 85, 1, 0, 0, 0, 90,
		5, 1, 0, 0, 0, 91, 98, 5, 48, 0, 0, 92, 98, 5, 46, 0, 0, 93, 95, 5, 47,
		0, 0, 94, 96, 3, 46, 23, 0, 95, 94, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96,
		98, 1, 0, 0, 0, 97, 91, 1, 0, 0, 0, 97, 92, 1, 0, 0, 0, 97, 93, 1, 0, 0,
		0, 98, 7, 1, 0, 0, 0, 99, 100, 5, 31, 0, 0, 100, 101, 5, 2, 0, 0, 101,
		106, 3, 46, 23, 0, 102, 103, 5, 10, 0, 0, 103, 105, 3, 46, 23, 0, 104,
		102, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107,
		1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 110, 5, 1,
		0, 0, 110, 9, 1, 0, 0, 0, 111, 112, 7, 0, 0, 0, 112, 113, 5, 74, 0, 0,
		113, 114, 5, 7, 0, 0, 114, 115, 3, 14, 7, 0, 115, 116, 5, 16, 0, 0, 116,
		117, 3, 46, 23, 0, 117, 129, 1, 0, 0, 0, 118, 119, 7, 0, 0, 0, 119, 120,
		5, 74, 0, 0, 120, 121, 5, 16, 0, 0, 121, 129, 3, 46, 23, 0, 122, 123, 7,
		0, 0, 0, 123, 124, 5, 74, 0, 0, 124, 125, 5, 7, 0, 0, 125, 126, 3, 14,
		7, 0, 126, 127, 5, 9, 0, 0, 127, 129, 1, 0, 0, 0, 128, 111, 1, 0, 0, 0,
		128, 118, 1, 0, 0, 0, 128, 122, 1, 0, 0, 0, 129, 11, 1, 0, 0, 0, 130, 131,
		5, 74, 0, 0, 131, 132, 5, 16, 0, 0, 132, 137, 3, 46, 23, 0, 133, 134, 5,
		74, 0, 0, 134, 135, 7, 1, 0, 0, 135, 137, 3, 46, 23, 0, 136, 130, 1, 0,
		0, 0, 136, 133, 1, 0, 0, 0, 137, 13, 1, 0, 0, 0, 138, 149, 5, 64, 0, 0,
		139, 149, 5, 65, 0, 0, 140, 149, 5, 68, 0, 0, 141, 149, 5, 66, 0, 0, 142,
		149, 5, 67, 0, 0, 143, 149, 5, 74, 0, 0, 144, 145, 5, 5, 0, 0, 145, 146,
		3, 14, 7, 0, 146, 147, 5, 6, 0, 0, 147, 149, 1, 0, 0, 0, 148, 138, 1, 0,
		0, 0, 148, 139, 1, 0, 0, 0, 148, 140, 1, 0, 0, 0, 148, 141, 1, 0, 0, 0,
		148, 142, 1, 0, 0, 0, 148, 143, 1, 0, 0, 0, 148, 144, 1, 0, 0, 0, 149,
		15, 1, 0, 0, 0, 150, 151, 5, 36, 0, 0, 151, 152, 3, 44, 22, 0, 152, 153,
		5, 3, 0, 0, 153, 154, 3, 2, 1, 0, 154, 163, 5, 4, 0, 0, 155, 161, 5, 37,
		0, 0, 156, 162, 3, 16, 8, 0, 157, 158, 5, 3, 0, 0, 158, 159, 3, 2, 1, 0,
		159, 160, 5, 4, 0, 0, 160, 162, 1, 0, 0, 0, 161, 156, 1, 0, 0, 0, 161,
		157, 1, 0, 0, 0, 162, 164, 1, 0, 0, 0, 163, 155, 1, 0, 0, 0, 163, 164,
		1, 0, 0, 0, 164, 17, 1, 0, 0, 0, 165, 166, 5, 45, 0, 0, 166, 167, 3, 44,
		22, 0, 167, 168, 5, 37, 0, 0, 168, 169, 5, 3, 0, 0, 169, 170, 3, 2, 1,
		0, 170, 171, 3, 6, 3, 0, 171, 172, 5, 4, 0, 0, 172, 19, 1, 0, 0, 0, 173,
		174, 5, 38, 0, 0, 174, 175, 3, 46, 23, 0, 175, 179, 5, 3, 0, 0, 176, 178,
		3, 22, 11, 0, 177, 176, 1, 0, 0, 0, 178, 181, 1, 0, 0, 0, 179, 177, 1,
		0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 183, 1, 0, 0, 0, 181, 179, 1, 0, 0,
		0, 182, 184, 3, 24, 12, 0, 183, 182, 1, 0, 0, 0, 183, 184, 1, 0, 0, 0,
		184, 185, 1, 0, 0, 0, 185, 186, 5, 4, 0, 0, 186, 21, 1, 0, 0, 0, 187, 188,
		5, 39, 0, 0, 188, 189, 3, 46, 23, 0, 189, 190, 5, 7, 0, 0, 190, 191, 3,
		2, 1, 0, 191, 23, 1, 0, 0, 0, 192, 193, 5, 40, 0, 0, 193, 194, 5, 7, 0,
		0, 194, 195, 3, 2, 1, 0, 195, 25, 1, 0, 0, 0, 196, 197, 5, 41, 0, 0, 197,
		198, 3, 44, 22, 0, 198, 199, 5, 3, 0, 0, 199, 200, 3, 2, 1, 0, 200, 201,
		5, 4, 0, 0, 201, 27, 1, 0, 0, 0, 202, 203, 5, 42, 0, 0, 203, 204, 5, 74,
		0, 0, 204, 205, 5, 43, 0, 0, 205, 206, 3, 46, 23, 0, 206, 207, 5, 44, 0,
		0, 207, 208, 3, 46, 23, 0, 208, 209, 5, 3, 0, 0, 209, 210, 3, 2, 1, 0,
		210, 211, 5, 4, 0, 0, 211, 221, 1, 0, 0, 0, 212, 213, 5, 42, 0, 0, 213,
		214, 5, 74, 0, 0, 214, 215, 5, 43, 0, 0, 215, 216, 3, 46, 23, 0, 216, 217,
		5, 3, 0, 0, 217, 218, 3, 2, 1, 0, 218, 219, 5, 4, 0, 0, 219, 221, 1, 0,
		0, 0, 220, 202, 1, 0, 0, 0, 220, 212, 1, 0, 0, 0, 221, 29, 1, 0, 0, 0,
		222, 223, 7, 0, 0, 0, 223, 224, 5, 74, 0, 0, 224, 225, 5, 7, 0, 0, 225,
		226, 5, 5, 0, 0, 226, 227, 3, 14, 7, 0, 227, 228, 5, 6, 0, 0, 228, 229,
		5, 16, 0, 0, 229, 230, 3, 32, 16, 0, 230, 241, 1, 0, 0, 0, 231, 232, 7,
		0, 0, 0, 232, 233, 5, 74, 0, 0, 233, 234, 5, 16, 0, 0, 234, 235, 5, 5,
		0, 0, 235, 236, 3, 14, 7, 0, 236, 237, 5, 6, 0, 0, 237, 238, 5, 2, 0, 0,
		238, 239, 5, 1, 0, 0, 239, 241, 1, 0, 0, 0, 240, 222, 1, 0, 0, 0, 240,
		231, 1, 0, 0, 0, 241, 31, 1, 0, 0, 0, 242, 243, 5, 5, 0, 0, 243, 248, 3,
		46, 23, 0, 244, 245, 5, 10, 0, 0, 245, 247, 3, 46, 23, 0, 246, 244, 1,
		0, 0, 0, 247, 250, 1, 0, 0, 0, 248, 246, 1, 0, 0, 0, 248, 249, 1, 0, 0,
		0, 249, 251, 1, 0, 0, 0, 250, 248, 1, 0, 0, 0, 251, 252, 5, 6, 0, 0, 252,
		257, 1, 0, 0, 0, 253, 254, 5, 5, 0, 0, 254, 257, 5, 6, 0, 0, 255, 257,
		5, 74, 0, 0, 256, 242, 1, 0, 0, 0, 256, 253, 1, 0, 0, 0, 256, 255, 1, 0,
		0, 0, 257, 33, 1, 0, 0, 0, 258, 259, 5, 74, 0, 0, 259, 260, 5, 11, 0, 0,
		260, 261, 5, 50, 0, 0, 261, 262, 5, 2, 0, 0, 262, 263, 3, 46, 23, 0, 263,
		264, 5, 1, 0, 0, 264, 280, 1, 0, 0, 0, 265, 266, 5, 74, 0, 0, 266, 267,
		5, 11, 0, 0, 267, 268, 5, 51, 0, 0, 268, 269, 5, 2, 0, 0, 269, 280, 5,
		1, 0, 0, 270, 271, 5, 74, 0, 0, 271, 272, 5, 11, 0, 0, 272, 273, 5, 52,
		0, 0, 273, 274, 5, 2, 0, 0, 274, 275, 5, 53, 0, 0, 275, 276, 5, 7, 0, 0,
		276, 277, 3, 46, 23, 0, 277, 278, 5, 1, 0, 0, 278, 280, 1, 0, 0, 0, 279,
		258, 1, 0, 0, 0, 279, 265, 1, 0, 0, 0, 279, 270, 1, 0, 0, 0, 280, 35, 1,
		0, 0, 0, 281, 282, 7, 0, 0, 0, 282, 285, 5, 74, 0, 0, 283, 284, 5, 7, 0,
		0, 284, 286, 3, 14, 7, 0, 285, 283, 1, 0, 0, 0, 285, 286, 1, 0, 0, 0, 286,
		287, 1, 0, 0, 0, 287, 288, 5, 16, 0, 0, 288, 289, 3, 38, 19, 0, 289, 37,
		1, 0, 0, 0, 290, 293, 3, 40, 20, 0, 291, 293, 3, 42, 21, 0, 292, 290, 1,
		0, 0, 0, 292, 291, 1, 0, 0, 0, 293, 39, 1, 0, 0, 0, 294, 295, 5, 5, 0,
		0, 295, 300, 3, 40, 20, 0, 296, 297, 5, 10, 0, 0, 297, 299, 3, 40, 20,
		0, 298, 296, 1, 0, 0, 0, 299, 302, 1, 0, 0, 0, 300, 298, 1, 0, 0, 0, 300,
		301, 1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 300, 1, 0, 0, 0, 303, 304,
		5, 6, 0, 0, 304, 317, 1, 0, 0, 0, 305, 306, 5, 5, 0, 0, 306, 311, 3, 46,
		23, 0, 307, 308, 5, 10, 0, 0, 308, 310, 3, 46, 23, 0, 309, 307, 1, 0, 0,
		0, 310, 313, 1, 0, 0, 0, 311, 309, 1, 0, 0, 0, 311, 312, 1, 0, 0, 0, 312,
		314, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0, 314, 315, 5, 6, 0, 0, 315, 317,
		1, 0, 0, 0, 316, 294, 1, 0, 0, 0, 316, 305, 1, 0, 0, 0, 317, 41, 1, 0,
		0, 0, 318, 319, 5, 5, 0, 0, 319, 320, 3, 14, 7, 0, 320, 321, 5, 6, 0, 0,
		321, 322, 5, 2, 0, 0, 322, 323, 5, 57, 0, 0, 323, 324, 5, 7, 0, 0, 324,
		325, 3, 42, 21, 0, 325, 326, 5, 10, 0, 0, 326, 327, 5, 55, 0, 0, 327, 328,
		5, 7, 0, 0, 328, 329, 3, 46, 23, 0, 329, 330, 5, 1, 0, 0, 330, 345, 1,
		0, 0, 0, 331, 332, 5, 5, 0, 0, 332, 333, 3, 14, 7, 0, 333, 334, 5, 6, 0,
		0, 334, 335, 5, 2, 0, 0, 335, 336, 5, 57, 0, 0, 336, 337, 5, 7, 0, 0, 337,
		338, 3, 46, 23, 0, 338, 339, 5, 10, 0, 0, 339, 340, 5, 55, 0, 0, 340, 341,
		5, 7, 0, 0, 341, 342, 3, 46, 23, 0, 342, 343, 5, 1, 0, 0, 343, 345, 1,
		0, 0, 0, 344, 318, 1, 0, 0, 0, 344, 331, 1, 0, 0, 0, 345, 43, 1, 0, 0,
		0, 346, 347, 6, 22, -1, 0, 347, 348, 5, 30, 0, 0, 348, 359, 3, 44, 22,
		5, 349, 350, 3, 46, 23, 0, 350, 351, 7, 2, 0, 0, 351, 352, 3, 46, 23, 0,
		352, 359, 1, 0, 0, 0, 353, 359, 7, 3, 0, 0, 354, 355, 5, 2, 0, 0, 355,
		356, 3, 44, 22, 0, 356, 357, 5, 1, 0, 0, 357, 359, 1, 0, 0, 0, 358, 346,
		1, 0, 0, 0, 358, 349, 1, 0, 0, 0, 358, 353, 1, 0, 0, 0, 358, 354, 1, 0,
		0, 0, 359, 365, 1, 0, 0, 0, 360, 361, 10, 3, 0, 0, 361, 362, 7, 4, 0, 0,
		362, 364, 3, 44, 22, 4, 363, 360, 1, 0, 0, 0, 364, 367, 1, 0, 0, 0, 365,
		363, 1, 0, 0, 0, 365, 366, 1, 0, 0, 0, 366, 45, 1, 0, 0, 0, 367, 365, 1,
		0, 0, 0, 368, 369, 6, 23, -1, 0, 369, 370, 7, 5, 0, 0, 370, 406, 3, 46,
		23, 18, 371, 406, 7, 3, 0, 0, 372, 406, 5, 49, 0, 0, 373, 374, 5, 74, 0,
		0, 374, 375, 5, 5, 0, 0, 375, 376, 3, 46, 23, 0, 376, 377, 5, 6, 0, 0,
		377, 406, 1, 0, 0, 0, 378, 379, 5, 74, 0, 0, 379, 380, 5, 5, 0, 0, 380,
		381, 3, 46, 23, 0, 381, 386, 5, 6, 0, 0, 382, 383, 5, 5, 0, 0, 383, 384,
		3, 46, 23, 0, 384, 385, 5, 6, 0, 0, 385, 387, 1, 0, 0, 0, 386, 382, 1,
		0, 0, 0, 387, 388, 1, 0, 0, 0, 388, 386, 1, 0, 0, 0, 388, 389, 1, 0, 0,
		0, 389, 406, 1, 0, 0, 0, 390, 391, 5, 74, 0, 0, 391, 392, 5, 11, 0, 0,
		392, 406, 5, 54, 0, 0, 393, 394, 5, 74, 0, 0, 394, 395, 5, 11, 0, 0, 395,
		406, 5, 55, 0, 0, 396, 406, 5, 74, 0, 0, 397, 406, 5, 70, 0, 0, 398, 406,
		5, 71, 0, 0, 399, 406, 5, 73, 0, 0, 400, 406, 5, 72, 0, 0, 401, 402, 5,
		2, 0, 0, 402, 403, 3, 46, 23, 0, 403, 404, 5, 1, 0, 0, 404, 406, 1, 0,
		0, 0, 405, 368, 1, 0, 0, 0, 405, 371, 1, 0, 0, 0, 405, 372, 1, 0, 0, 0,
		405, 373, 1, 0, 0, 0, 405, 378, 1, 0, 0, 0, 405, 390, 1, 0, 0, 0, 405,
		393, 1, 0, 0, 0, 405, 396, 1, 0, 0, 0, 405, 397, 1, 0, 0, 0, 405, 398,
		1, 0, 0, 0, 405, 399, 1, 0, 0, 0, 405, 400, 1, 0, 0, 0, 405, 401, 1, 0,
		0, 0, 406, 424, 1, 0, 0, 0, 407, 408, 10, 17, 0, 0, 408, 409, 7, 6, 0,
		0, 409, 423, 3, 46, 23, 18, 410, 411, 10, 16, 0, 0, 411, 412, 7, 7, 0,
		0, 412, 423, 3, 46, 23, 17, 413, 414, 10, 15, 0, 0, 414, 415, 5, 18, 0,
		0, 415, 423, 3, 46, 23, 16, 416, 417, 10, 14, 0, 0, 417, 418, 7, 2, 0,
		0, 418, 423, 3, 46, 23, 15, 419, 420, 10, 13, 0, 0, 420, 421, 7, 4, 0,
		0, 421, 423, 3, 46, 23, 14, 422, 407, 1, 0, 0, 0, 422, 410, 1, 0, 0, 0,
		422, 413, 1, 0, 0, 0, 422, 416, 1, 0, 0, 0, 422, 419, 1, 0, 0, 0, 423,
		426, 1, 0, 0, 0, 424, 422, 1, 0, 0, 0, 424, 425, 1, 0, 0, 0, 425, 47, 1,
		0, 0, 0, 426, 424, 1, 0, 0, 0, 36, 53, 58, 62, 66, 73, 79, 83, 87, 89,
		95, 97, 106, 128, 136, 148, 161, 163, 179, 183, 220, 240, 248, 256, 279,
		285, 292, 300, 311, 316, 344, 358, 365, 388, 405, 422, 424,
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
	Tswift_GrammarNParserCHAR          = 67
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
	Tswift_GrammarNParserRULE_s                     = 0
	Tswift_GrammarNParserRULE_l_sentencias          = 1
	Tswift_GrammarNParserRULE_sentencia             = 2
	Tswift_GrammarNParserRULE_trans_sentencia       = 3
	Tswift_GrammarNParserRULE_print_sentencia       = 4
	Tswift_GrammarNParserRULE_declaracion_sentencia = 5
	Tswift_GrammarNParserRULE_asignacion_sentencia  = 6
	Tswift_GrammarNParserRULE_tipo                  = 7
	Tswift_GrammarNParserRULE_if_sentencia          = 8
	Tswift_GrammarNParserRULE_guard_sentencia       = 9
	Tswift_GrammarNParserRULE_switch_sentencia      = 10
	Tswift_GrammarNParserRULE_lcasos                = 11
	Tswift_GrammarNParserRULE_cdefault              = 12
	Tswift_GrammarNParserRULE_while_sentencia       = 13
	Tswift_GrammarNParserRULE_for_sentencia         = 14
	Tswift_GrammarNParserRULE_dec_vector            = 15
	Tswift_GrammarNParserRULE_def_vector            = 16
	Tswift_GrammarNParserRULE_func_vector           = 17
	Tswift_GrammarNParserRULE_dec_matriz            = 18
	Tswift_GrammarNParserRULE_def_matriz            = 19
	Tswift_GrammarNParserRULE_listavalores_matriz   = 20
	Tswift_GrammarNParserRULE_simple_matriz         = 21
	Tswift_GrammarNParserRULE_condicion             = 22
	Tswift_GrammarNParserRULE_e                     = 23
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
		p.SetState(48)
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
	var _alt int

	localctx = NewL_SentenciaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(50)
				p.Sentencia()
			}

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
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

type S_ForContext struct {
	SentenciaContext
}

func NewS_ForContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_ForContext {
	var p = new(S_ForContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_ForContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_ForContext) For_sentencia() IFor_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFor_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFor_sentenciaContext)
}

func (s *S_ForContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_For(s)
	}
}

func (s *S_ForContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_For(s)
	}
}

func (s *S_ForContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_For(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Declaracion_MatrizContext struct {
	SentenciaContext
}

func NewS_Declaracion_MatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Declaracion_MatrizContext {
	var p = new(S_Declaracion_MatrizContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Declaracion_MatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Declaracion_MatrizContext) Dec_matriz() IDec_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDec_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDec_matrizContext)
}

func (s *S_Declaracion_MatrizContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPTCOMA, 0)
}

func (s *S_Declaracion_MatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Declaracion_Matriz(s)
	}
}

func (s *S_Declaracion_MatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Declaracion_Matriz(s)
	}
}

func (s *S_Declaracion_MatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Declaracion_Matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_SwitchContext struct {
	SentenciaContext
}

func NewS_SwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_SwitchContext {
	var p = new(S_SwitchContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_SwitchContext) Switch_sentencia() ISwitch_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitch_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitch_sentenciaContext)
}

func (s *S_SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Switch(s)
	}
}

func (s *S_SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Switch(s)
	}
}

func (s *S_SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Switch(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_AsignacionContext struct {
	SentenciaContext
}

func NewS_AsignacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_AsignacionContext {
	var p = new(S_AsignacionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_AsignacionContext) Asignacion_sentencia() IAsignacion_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacion_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacion_sentenciaContext)
}

func (s *S_AsignacionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPTCOMA, 0)
}

func (s *S_AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Asignacion(s)
	}
}

func (s *S_AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Asignacion(s)
	}
}

func (s *S_AsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Asignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_GuardContext struct {
	SentenciaContext
}

func NewS_GuardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_GuardContext {
	var p = new(S_GuardContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_GuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_GuardContext) Guard_sentencia() IGuard_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGuard_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGuard_sentenciaContext)
}

func (s *S_GuardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Guard(s)
	}
}

func (s *S_GuardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Guard(s)
	}
}

func (s *S_GuardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Guard(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_WhileContext struct {
	SentenciaContext
}

func NewS_WhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_WhileContext {
	var p = new(S_WhileContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_WhileContext) While_sentencia() IWhile_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhile_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhile_sentenciaContext)
}

func (s *S_WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_While(s)
	}
}

func (s *S_WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_While(s)
	}
}

func (s *S_WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_While(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_TransContext struct {
	SentenciaContext
}

func NewS_TransContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_TransContext {
	var p = new(S_TransContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_TransContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_TransContext) Trans_sentencia() ITrans_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrans_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrans_sentenciaContext)
}

func (s *S_TransContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPTCOMA, 0)
}

func (s *S_TransContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Trans(s)
	}
}

func (s *S_TransContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Trans(s)
	}
}

func (s *S_TransContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Trans(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Funcion_VectorContext struct {
	SentenciaContext
}

func NewS_Funcion_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Funcion_VectorContext {
	var p = new(S_Funcion_VectorContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Funcion_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Funcion_VectorContext) Func_vector() IFunc_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunc_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunc_vectorContext)
}

func (s *S_Funcion_VectorContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPTCOMA, 0)
}

func (s *S_Funcion_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Funcion_Vector(s)
	}
}

func (s *S_Funcion_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Funcion_Vector(s)
	}
}

func (s *S_Funcion_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Funcion_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type S_Declaracion_VectorContext struct {
	SentenciaContext
}

func NewS_Declaracion_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_Declaracion_VectorContext {
	var p = new(S_Declaracion_VectorContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_Declaracion_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_Declaracion_VectorContext) Dec_vector() IDec_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDec_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDec_vectorContext)
}

func (s *S_Declaracion_VectorContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPTCOMA, 0)
}

func (s *S_Declaracion_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Declaracion_Vector(s)
	}
}

func (s *S_Declaracion_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Declaracion_Vector(s)
	}
}

func (s *S_Declaracion_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Declaracion_Vector(s)

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

type S_DeclaracionContext struct {
	SentenciaContext
}

func NewS_DeclaracionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *S_DeclaracionContext {
	var p = new(S_DeclaracionContext)

	InitEmptySentenciaContext(&p.SentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*SentenciaContext))

	return p
}

func (s *S_DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *S_DeclaracionContext) Declaracion_sentencia() IDeclaracion_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracion_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_sentenciaContext)
}

func (s *S_DeclaracionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPTCOMA, 0)
}

func (s *S_DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterS_Declaracion(s)
	}
}

func (s *S_DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitS_Declaracion(s)
	}
}

func (s *S_DeclaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitS_Declaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Sentencia() (localctx ISentenciaContext) {
	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Tswift_GrammarNParserRULE_sentencia)
	var _la int

	p.SetState(89)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewS_PrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(56)
			p.Print_sentencia()
		}
		p.SetState(58)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarNParserPTCOMA {
			{
				p.SetState(57)
				p.Match(Tswift_GrammarNParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 2:
		localctx = NewS_DeclaracionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(60)
			p.Declaracion_sentencia()
		}
		p.SetState(62)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarNParserPTCOMA {
			{
				p.SetState(61)
				p.Match(Tswift_GrammarNParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 3:
		localctx = NewS_AsignacionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Asignacion_sentencia()
		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarNParserPTCOMA {
			{
				p.SetState(65)
				p.Match(Tswift_GrammarNParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 4:
		localctx = NewS_IfContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(68)
			p.If_sentencia()
		}

	case 5:
		localctx = NewS_SwitchContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)
			p.Switch_sentencia()
		}

	case 6:
		localctx = NewS_GuardContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(70)
			p.Guard_sentencia()
		}

	case 7:
		localctx = NewS_TransContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(71)
			p.Trans_sentencia()
		}
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarNParserPTCOMA {
			{
				p.SetState(72)
				p.Match(Tswift_GrammarNParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 8:
		localctx = NewS_WhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(75)
			p.While_sentencia()
		}

	case 9:
		localctx = NewS_ForContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(76)
			p.For_sentencia()
		}

	case 10:
		localctx = NewS_Declaracion_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(77)
			p.Dec_vector()
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarNParserPTCOMA {
			{
				p.SetState(78)
				p.Match(Tswift_GrammarNParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 11:
		localctx = NewS_Funcion_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(81)
			p.Func_vector()
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarNParserPTCOMA {
			{
				p.SetState(82)
				p.Match(Tswift_GrammarNParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case 12:
		localctx = NewS_Declaracion_MatrizContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(85)
			p.Dec_matriz()
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Tswift_GrammarNParserPTCOMA {
			{
				p.SetState(86)
				p.Match(Tswift_GrammarNParserPTCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}

	case antlr.ATNInvalidAltNumber:
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

// ITrans_sentenciaContext is an interface to support dynamic dispatch.
type ITrans_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTrans_sentenciaContext differentiates from other interfaces.
	IsTrans_sentenciaContext()
}

type Trans_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTrans_sentenciaContext() *Trans_sentenciaContext {
	var p = new(Trans_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_trans_sentencia
	return p
}

func InitEmptyTrans_sentenciaContext(p *Trans_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_trans_sentencia
}

func (*Trans_sentenciaContext) IsTrans_sentenciaContext() {}

func NewTrans_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Trans_sentenciaContext {
	var p = new(Trans_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_trans_sentencia

	return p
}

func (s *Trans_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Trans_sentenciaContext) CopyAll(ctx *Trans_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Trans_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Trans_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReturnContext struct {
	Trans_sentenciaContext
}

func NewReturnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnContext {
	var p = new(ReturnContext)

	InitEmptyTrans_sentenciaContext(&p.Trans_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Trans_sentenciaContext))

	return p
}

func (s *ReturnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnContext) RETURN() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserRETURN, 0)
}

func (s *ReturnContext) E() IEContext {
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

func (s *ReturnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterReturn(s)
	}
}

func (s *ReturnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitReturn(s)
	}
}

func (s *ReturnContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitReturn(s)

	default:
		return t.VisitChildren(s)
	}
}

type BreakContext struct {
	Trans_sentenciaContext
}

func NewBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakContext {
	var p = new(BreakContext)

	InitEmptyTrans_sentenciaContext(&p.Trans_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Trans_sentenciaContext))

	return p
}

func (s *BreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakContext) BREAK() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserBREAK, 0)
}

func (s *BreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterBreak(s)
	}
}

func (s *BreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitBreak(s)
	}
}

func (s *BreakContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitBreak(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueContext struct {
	Trans_sentenciaContext
}

func NewContinueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueContext {
	var p = new(ContinueContext)

	InitEmptyTrans_sentenciaContext(&p.Trans_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Trans_sentenciaContext))

	return p
}

func (s *ContinueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCONTINUE, 0)
}

func (s *ContinueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterContinue(s)
	}
}

func (s *ContinueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitContinue(s)
	}
}

func (s *ContinueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitContinue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Trans_sentencia() (localctx ITrans_sentenciaContext) {
	localctx = NewTrans_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Tswift_GrammarNParserRULE_trans_sentencia)
	p.SetState(97)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarNParserBREAK:
		localctx = NewBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.Match(Tswift_GrammarNParserBREAK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserCONTINUE:
		localctx = NewContinueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.Match(Tswift_GrammarNParserCONTINUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserRETURN:
		localctx = NewReturnContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(93)
			p.Match(Tswift_GrammarNParserRETURN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(94)
				p.e(0)
			}

		} else if p.HasError() { // JIM
			goto errorExit
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
	p.EnterRule(localctx, 8, Tswift_GrammarNParserRULE_print_sentencia)
	var _la int

	localctx = NewPrintContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(Tswift_GrammarNParserPRINT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(Tswift_GrammarNParserPARIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(101)
		p.e(0)
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarNParserCOMA {
		{
			p.SetState(102)
			p.Match(Tswift_GrammarNParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(103)
			p.e(0)
		}

		p.SetState(108)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(109)
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

// IDeclaracion_sentenciaContext is an interface to support dynamic dispatch.
type IDeclaracion_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclaracion_sentenciaContext differentiates from other interfaces.
	IsDeclaracion_sentenciaContext()
}

type Declaracion_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracion_sentenciaContext() *Declaracion_sentenciaContext {
	var p = new(Declaracion_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_declaracion_sentencia
	return p
}

func InitEmptyDeclaracion_sentenciaContext(p *Declaracion_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_declaracion_sentencia
}

func (*Declaracion_sentenciaContext) IsDeclaracion_sentenciaContext() {}

func NewDeclaracion_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_sentenciaContext {
	var p = new(Declaracion_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_declaracion_sentencia

	return p
}

func (s *Declaracion_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_sentenciaContext) CopyAll(ctx *Declaracion_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Declaracion_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_Tipo_ValContext struct {
	Declaracion_sentenciaContext
	tip antlr.Token
}

func NewDeclaracion_Tipo_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_Tipo_ValContext {
	var p = new(Declaracion_Tipo_ValContext)

	InitEmptyDeclaracion_sentenciaContext(&p.Declaracion_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_sentenciaContext))

	return p
}

func (s *Declaracion_Tipo_ValContext) GetTip() antlr.Token { return s.tip }

func (s *Declaracion_Tipo_ValContext) SetTip(v antlr.Token) { s.tip = v }

func (s *Declaracion_Tipo_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_Tipo_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Declaracion_Tipo_ValContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDOSPT, 0)
}

func (s *Declaracion_Tipo_ValContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_Tipo_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIGUAL, 0)
}

func (s *Declaracion_Tipo_ValContext) E() IEContext {
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

func (s *Declaracion_Tipo_ValContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserVAR, 0)
}

func (s *Declaracion_Tipo_ValContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLET, 0)
}

func (s *Declaracion_Tipo_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDeclaracion_Tipo_Val(s)
	}
}

func (s *Declaracion_Tipo_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDeclaracion_Tipo_Val(s)
	}
}

func (s *Declaracion_Tipo_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDeclaracion_Tipo_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declaracion_Tipo_noValContext struct {
	Declaracion_sentenciaContext
	tip antlr.Token
}

func NewDeclaracion_Tipo_noValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_Tipo_noValContext {
	var p = new(Declaracion_Tipo_noValContext)

	InitEmptyDeclaracion_sentenciaContext(&p.Declaracion_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_sentenciaContext))

	return p
}

func (s *Declaracion_Tipo_noValContext) GetTip() antlr.Token { return s.tip }

func (s *Declaracion_Tipo_noValContext) SetTip(v antlr.Token) { s.tip = v }

func (s *Declaracion_Tipo_noValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_Tipo_noValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Declaracion_Tipo_noValContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDOSPT, 0)
}

func (s *Declaracion_Tipo_noValContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_Tipo_noValContext) INTERROGACION() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserINTERROGACION, 0)
}

func (s *Declaracion_Tipo_noValContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserVAR, 0)
}

func (s *Declaracion_Tipo_noValContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLET, 0)
}

func (s *Declaracion_Tipo_noValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDeclaracion_Tipo_noVal(s)
	}
}

func (s *Declaracion_Tipo_noValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDeclaracion_Tipo_noVal(s)
	}
}

func (s *Declaracion_Tipo_noValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDeclaracion_Tipo_noVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declaracion_ValContext struct {
	Declaracion_sentenciaContext
	tip antlr.Token
}

func NewDeclaracion_ValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_ValContext {
	var p = new(Declaracion_ValContext)

	InitEmptyDeclaracion_sentenciaContext(&p.Declaracion_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Declaracion_sentenciaContext))

	return p
}

func (s *Declaracion_ValContext) GetTip() antlr.Token { return s.tip }

func (s *Declaracion_ValContext) SetTip(v antlr.Token) { s.tip = v }

func (s *Declaracion_ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_ValContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Declaracion_ValContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIGUAL, 0)
}

func (s *Declaracion_ValContext) E() IEContext {
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

func (s *Declaracion_ValContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserVAR, 0)
}

func (s *Declaracion_ValContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLET, 0)
}

func (s *Declaracion_ValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDeclaracion_Val(s)
	}
}

func (s *Declaracion_ValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDeclaracion_Val(s)
	}
}

func (s *Declaracion_ValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDeclaracion_Val(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Declaracion_sentencia() (localctx IDeclaracion_sentenciaContext) {
	localctx = NewDeclaracion_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Tswift_GrammarNParserRULE_declaracion_sentencia)
	var _la int

	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclaracion_Tipo_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(111)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Declaracion_Tipo_ValContext).tip = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarNParserVAR || _la == Tswift_GrammarNParserLET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Declaracion_Tipo_ValContext).tip = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(112)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(113)
			p.Match(Tswift_GrammarNParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(114)
			p.Tipo()
		}
		{
			p.SetState(115)
			p.Match(Tswift_GrammarNParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(116)
			p.e(0)
		}

	case 2:
		localctx = NewDeclaracion_ValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Declaracion_ValContext).tip = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarNParserVAR || _la == Tswift_GrammarNParserLET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Declaracion_ValContext).tip = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(119)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(120)
			p.Match(Tswift_GrammarNParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(121)
			p.e(0)
		}

	case 3:
		localctx = NewDeclaracion_Tipo_noValContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(122)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Declaracion_Tipo_noValContext).tip = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarNParserVAR || _la == Tswift_GrammarNParserLET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Declaracion_Tipo_noValContext).tip = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(123)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(124)
			p.Match(Tswift_GrammarNParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.Tipo()
		}
		{
			p.SetState(126)
			p.Match(Tswift_GrammarNParserINTERROGACION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IAsignacion_sentenciaContext is an interface to support dynamic dispatch.
type IAsignacion_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAsignacion_sentenciaContext differentiates from other interfaces.
	IsAsignacion_sentenciaContext()
}

type Asignacion_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacion_sentenciaContext() *Asignacion_sentenciaContext {
	var p = new(Asignacion_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_asignacion_sentencia
	return p
}

func InitEmptyAsignacion_sentenciaContext(p *Asignacion_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_asignacion_sentencia
}

func (*Asignacion_sentenciaContext) IsAsignacion_sentenciaContext() {}

func NewAsignacion_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asignacion_sentenciaContext {
	var p = new(Asignacion_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_asignacion_sentencia

	return p
}

func (s *Asignacion_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Asignacion_sentenciaContext) CopyAll(ctx *Asignacion_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Asignacion_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AsignacionContext struct {
	Asignacion_sentenciaContext
}

func NewAsignacionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsignacionContext {
	var p = new(AsignacionContext)

	InitEmptyAsignacion_sentenciaContext(&p.Asignacion_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asignacion_sentenciaContext))

	return p
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *AsignacionContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIGUAL, 0)
}

func (s *AsignacionContext) E() IEContext {
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

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (s *AsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

type Asignacion_SumaRestaContext struct {
	Asignacion_sentenciaContext
	op antlr.Token
}

func NewAsignacion_SumaRestaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Asignacion_SumaRestaContext {
	var p = new(Asignacion_SumaRestaContext)

	InitEmptyAsignacion_sentenciaContext(&p.Asignacion_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Asignacion_sentenciaContext))

	return p
}

func (s *Asignacion_SumaRestaContext) GetOp() antlr.Token { return s.op }

func (s *Asignacion_SumaRestaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Asignacion_SumaRestaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_SumaRestaContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Asignacion_SumaRestaContext) E() IEContext {
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

func (s *Asignacion_SumaRestaContext) MASIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMASIGUAL, 0)
}

func (s *Asignacion_SumaRestaContext) MENOSIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMENOSIGUAL, 0)
}

func (s *Asignacion_SumaRestaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterAsignacion_SumaResta(s)
	}
}

func (s *Asignacion_SumaRestaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitAsignacion_SumaResta(s)
	}
}

func (s *Asignacion_SumaRestaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitAsignacion_SumaResta(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Asignacion_sentencia() (localctx IAsignacion_sentenciaContext) {
	localctx = NewAsignacion_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Tswift_GrammarNParserRULE_asignacion_sentencia)
	var _la int

	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAsignacionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(131)
			p.Match(Tswift_GrammarNParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(132)
			p.e(0)
		}

	case 2:
		localctx = NewAsignacion_SumaRestaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(133)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(134)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Asignacion_SumaRestaContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarNParserMASIGUAL || _la == Tswift_GrammarNParserMENOSIGUAL) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Asignacion_SumaRestaContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(135)
			p.e(0)
		}

	case antlr.ATNInvalidAltNumber:
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

// ITipoContext is an interface to support dynamic dispatch.
type ITipoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTipoContext differentiates from other interfaces.
	IsTipoContext()
}

type TipoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTipoContext() *TipoContext {
	var p = new(TipoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_tipo
	return p
}

func InitEmptyTipoContext(p *TipoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_tipo
}

func (*TipoContext) IsTipoContext() {}

func NewTipoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TipoContext {
	var p = new(TipoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_tipo

	return p
}

func (s *TipoContext) GetParser() antlr.Parser { return s.parser }

func (s *TipoContext) CopyAll(ctx *TipoContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TipoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TipoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Tipo_VectorContext struct {
	TipoContext
}

func NewTipo_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_VectorContext {
	var p = new(Tipo_VectorContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Tipo_VectorContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Tipo_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Tipo_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterTipo_Vector(s)
	}
}

func (s *Tipo_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitTipo_Vector(s)
	}
}

func (s *Tipo_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitTipo_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_StructContext struct {
	TipoContext
}

func NewTipo_StructContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_StructContext {
	var p = new(Tipo_StructContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_StructContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_StructContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Tipo_StructContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterTipo_Struct(s)
	}
}

func (s *Tipo_StructContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitTipo_Struct(s)
	}
}

func (s *Tipo_StructContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitTipo_Struct(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_FloatContext struct {
	TipoContext
}

func NewTipo_FloatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_FloatContext {
	var p = new(Tipo_FloatContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_FloatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_FloatContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserFLOAT, 0)
}

func (s *Tipo_FloatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterTipo_Float(s)
	}
}

func (s *Tipo_FloatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitTipo_Float(s)
	}
}

func (s *Tipo_FloatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitTipo_Float(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_CharacterContext struct {
	TipoContext
}

func NewTipo_CharacterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_CharacterContext {
	var p = new(Tipo_CharacterContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_CharacterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_CharacterContext) CHAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCHAR, 0)
}

func (s *Tipo_CharacterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterTipo_Character(s)
	}
}

func (s *Tipo_CharacterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitTipo_Character(s)
	}
}

func (s *Tipo_CharacterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitTipo_Character(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_StringContext struct {
	TipoContext
}

func NewTipo_StringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_StringContext {
	var p = new(Tipo_StringContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_StringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_StringContext) STRING() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserSTRING, 0)
}

func (s *Tipo_StringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterTipo_String(s)
	}
}

func (s *Tipo_StringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitTipo_String(s)
	}
}

func (s *Tipo_StringContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitTipo_String(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_BoolContext struct {
	TipoContext
}

func NewTipo_BoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_BoolContext {
	var p = new(Tipo_BoolContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_BoolContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserBOOL, 0)
}

func (s *Tipo_BoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterTipo_Bool(s)
	}
}

func (s *Tipo_BoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitTipo_Bool(s)
	}
}

func (s *Tipo_BoolContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitTipo_Bool(s)

	default:
		return t.VisitChildren(s)
	}
}

type Tipo_IntContext struct {
	TipoContext
}

func NewTipo_IntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Tipo_IntContext {
	var p = new(Tipo_IntContext)

	InitEmptyTipoContext(&p.TipoContext)
	p.parser = parser
	p.CopyAll(ctx.(*TipoContext))

	return p
}

func (s *Tipo_IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_IntContext) INT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserINT, 0)
}

func (s *Tipo_IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterTipo_Int(s)
	}
}

func (s *Tipo_IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitTipo_Int(s)
	}
}

func (s *Tipo_IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitTipo_Int(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Tipo() (localctx ITipoContext) {
	localctx = NewTipoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Tswift_GrammarNParserRULE_tipo)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Tswift_GrammarNParserINT:
		localctx = NewTipo_IntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.Match(Tswift_GrammarNParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserFLOAT:
		localctx = NewTipo_FloatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.Match(Tswift_GrammarNParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserSTRING:
		localctx = NewTipo_StringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.Match(Tswift_GrammarNParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserBOOL:
		localctx = NewTipo_BoolContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(141)
			p.Match(Tswift_GrammarNParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserCHAR:
		localctx = NewTipo_CharacterContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(142)
			p.Match(Tswift_GrammarNParserCHAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserID:
		localctx = NewTipo_StructContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(143)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Tswift_GrammarNParserCORCHETEIZQ:
		localctx = NewTipo_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(144)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(145)
			p.Tipo()
		}
		{
			p.SetState(146)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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
	p.EnterRule(localctx, 16, Tswift_GrammarNParserRULE_if_sentencia)
	var _la int

	localctx = NewIfContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(Tswift_GrammarNParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.condicion(0)
	}
	{
		p.SetState(152)
		p.Match(Tswift_GrammarNParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.L_sentencias()
	}
	{
		p.SetState(154)
		p.Match(Tswift_GrammarNParserLLAVEDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarNParserELSE {
		{
			p.SetState(155)
			p.Match(Tswift_GrammarNParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Tswift_GrammarNParserIF:
			{
				p.SetState(156)
				p.If_sentencia()
			}

		case Tswift_GrammarNParserLLAVEIZQ:
			{
				p.SetState(157)
				p.Match(Tswift_GrammarNParserLLAVEIZQ)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(158)
				p.L_sentencias()
			}
			{
				p.SetState(159)
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

// IGuard_sentenciaContext is an interface to support dynamic dispatch.
type IGuard_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsGuard_sentenciaContext differentiates from other interfaces.
	IsGuard_sentenciaContext()
}

type Guard_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGuard_sentenciaContext() *Guard_sentenciaContext {
	var p = new(Guard_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_guard_sentencia
	return p
}

func InitEmptyGuard_sentenciaContext(p *Guard_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_guard_sentencia
}

func (*Guard_sentenciaContext) IsGuard_sentenciaContext() {}

func NewGuard_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Guard_sentenciaContext {
	var p = new(Guard_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_guard_sentencia

	return p
}

func (s *Guard_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Guard_sentenciaContext) CopyAll(ctx *Guard_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Guard_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Guard_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GuardContext struct {
	Guard_sentenciaContext
}

func NewGuardContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GuardContext {
	var p = new(GuardContext)

	InitEmptyGuard_sentenciaContext(&p.Guard_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Guard_sentenciaContext))

	return p
}

func (s *GuardContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GuardContext) GUARD() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserGUARD, 0)
}

func (s *GuardContext) Condicion() ICondicionContext {
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

func (s *GuardContext) ELSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserELSE, 0)
}

func (s *GuardContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEIZQ, 0)
}

func (s *GuardContext) L_sentencias() IL_sentenciasContext {
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

func (s *GuardContext) Trans_sentencia() ITrans_sentenciaContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITrans_sentenciaContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITrans_sentenciaContext)
}

func (s *GuardContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEDER, 0)
}

func (s *GuardContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterGuard(s)
	}
}

func (s *GuardContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitGuard(s)
	}
}

func (s *GuardContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitGuard(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Guard_sentencia() (localctx IGuard_sentenciaContext) {
	localctx = NewGuard_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Tswift_GrammarNParserRULE_guard_sentencia)
	localctx = NewGuardContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(165)
		p.Match(Tswift_GrammarNParserGUARD)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(166)
		p.condicion(0)
	}
	{
		p.SetState(167)
		p.Match(Tswift_GrammarNParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(Tswift_GrammarNParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.L_sentencias()
	}
	{
		p.SetState(170)
		p.Trans_sentencia()
	}
	{
		p.SetState(171)
		p.Match(Tswift_GrammarNParserLLAVEDER)
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

// ISwitch_sentenciaContext is an interface to support dynamic dispatch.
type ISwitch_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSwitch_sentenciaContext differentiates from other interfaces.
	IsSwitch_sentenciaContext()
}

type Switch_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitch_sentenciaContext() *Switch_sentenciaContext {
	var p = new(Switch_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_switch_sentencia
	return p
}

func InitEmptySwitch_sentenciaContext(p *Switch_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_switch_sentencia
}

func (*Switch_sentenciaContext) IsSwitch_sentenciaContext() {}

func NewSwitch_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Switch_sentenciaContext {
	var p = new(Switch_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_switch_sentencia

	return p
}

func (s *Switch_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *Switch_sentenciaContext) CopyAll(ctx *Switch_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Switch_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Switch_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SwitchContext struct {
	Switch_sentenciaContext
}

func NewSwitchContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchContext {
	var p = new(SwitchContext)

	InitEmptySwitch_sentenciaContext(&p.Switch_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*Switch_sentenciaContext))

	return p
}

func (s *SwitchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserSWITCH, 0)
}

func (s *SwitchContext) E() IEContext {
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

func (s *SwitchContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEIZQ, 0)
}

func (s *SwitchContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEDER, 0)
}

func (s *SwitchContext) AllLcasos() []ILcasosContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILcasosContext); ok {
			len++
		}
	}

	tst := make([]ILcasosContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILcasosContext); ok {
			tst[i] = t.(ILcasosContext)
			i++
		}
	}

	return tst
}

func (s *SwitchContext) Lcasos(i int) ILcasosContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILcasosContext); ok {
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

	return t.(ILcasosContext)
}

func (s *SwitchContext) Cdefault() ICdefaultContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICdefaultContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICdefaultContext)
}

func (s *SwitchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterSwitch(s)
	}
}

func (s *SwitchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitSwitch(s)
	}
}

func (s *SwitchContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitSwitch(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Switch_sentencia() (localctx ISwitch_sentenciaContext) {
	localctx = NewSwitch_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Tswift_GrammarNParserRULE_switch_sentencia)
	var _la int

	localctx = NewSwitchContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(173)
		p.Match(Tswift_GrammarNParserSWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(174)
		p.e(0)
	}
	{
		p.SetState(175)
		p.Match(Tswift_GrammarNParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Tswift_GrammarNParserCASE {
		{
			p.SetState(176)
			p.Lcasos()
		}

		p.SetState(181)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarNParserDEFAULT {
		{
			p.SetState(182)
			p.Cdefault()
		}

	}
	{
		p.SetState(185)
		p.Match(Tswift_GrammarNParserLLAVEDER)
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

// ILcasosContext is an interface to support dynamic dispatch.
type ILcasosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLcasosContext differentiates from other interfaces.
	IsLcasosContext()
}

type LcasosContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLcasosContext() *LcasosContext {
	var p = new(LcasosContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_lcasos
	return p
}

func InitEmptyLcasosContext(p *LcasosContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_lcasos
}

func (*LcasosContext) IsLcasosContext() {}

func NewLcasosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LcasosContext {
	var p = new(LcasosContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_lcasos

	return p
}

func (s *LcasosContext) GetParser() antlr.Parser { return s.parser }

func (s *LcasosContext) CopyAll(ctx *LcasosContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LcasosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LcasosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CaseContext struct {
	LcasosContext
}

func NewCaseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CaseContext {
	var p = new(CaseContext)

	InitEmptyLcasosContext(&p.LcasosContext)
	p.parser = parser
	p.CopyAll(ctx.(*LcasosContext))

	return p
}

func (s *CaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseContext) CASE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCASE, 0)
}

func (s *CaseContext) E() IEContext {
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

func (s *CaseContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDOSPT, 0)
}

func (s *CaseContext) L_sentencias() IL_sentenciasContext {
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

func (s *CaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterCase(s)
	}
}

func (s *CaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitCase(s)
	}
}

func (s *CaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitCase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Lcasos() (localctx ILcasosContext) {
	localctx = NewLcasosContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Tswift_GrammarNParserRULE_lcasos)
	localctx = NewCaseContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Match(Tswift_GrammarNParserCASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		p.e(0)
	}
	{
		p.SetState(189)
		p.Match(Tswift_GrammarNParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(190)
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

// ICdefaultContext is an interface to support dynamic dispatch.
type ICdefaultContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsCdefaultContext differentiates from other interfaces.
	IsCdefaultContext()
}

type CdefaultContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCdefaultContext() *CdefaultContext {
	var p = new(CdefaultContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_cdefault
	return p
}

func InitEmptyCdefaultContext(p *CdefaultContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_cdefault
}

func (*CdefaultContext) IsCdefaultContext() {}

func NewCdefaultContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CdefaultContext {
	var p = new(CdefaultContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_cdefault

	return p
}

func (s *CdefaultContext) GetParser() antlr.Parser { return s.parser }

func (s *CdefaultContext) CopyAll(ctx *CdefaultContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *CdefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CdefaultContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DefaultContext struct {
	CdefaultContext
}

func NewDefaultContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefaultContext {
	var p = new(DefaultContext)

	InitEmptyCdefaultContext(&p.CdefaultContext)
	p.parser = parser
	p.CopyAll(ctx.(*CdefaultContext))

	return p
}

func (s *DefaultContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDEFAULT, 0)
}

func (s *DefaultContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDOSPT, 0)
}

func (s *DefaultContext) L_sentencias() IL_sentenciasContext {
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

func (s *DefaultContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDefault(s)
	}
}

func (s *DefaultContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDefault(s)
	}
}

func (s *DefaultContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDefault(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Cdefault() (localctx ICdefaultContext) {
	localctx = NewCdefaultContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Tswift_GrammarNParserRULE_cdefault)
	localctx = NewDefaultContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(Tswift_GrammarNParserDEFAULT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(Tswift_GrammarNParserDOSPT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
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

// IWhile_sentenciaContext is an interface to support dynamic dispatch.
type IWhile_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsWhile_sentenciaContext differentiates from other interfaces.
	IsWhile_sentenciaContext()
}

type While_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhile_sentenciaContext() *While_sentenciaContext {
	var p = new(While_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_while_sentencia
	return p
}

func InitEmptyWhile_sentenciaContext(p *While_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_while_sentencia
}

func (*While_sentenciaContext) IsWhile_sentenciaContext() {}

func NewWhile_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *While_sentenciaContext {
	var p = new(While_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_while_sentencia

	return p
}

func (s *While_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *While_sentenciaContext) CopyAll(ctx *While_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *While_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *While_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type WhileContext struct {
	While_sentenciaContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	InitEmptyWhile_sentenciaContext(&p.While_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*While_sentenciaContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) WHILE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserWHILE, 0)
}

func (s *WhileContext) Condicion() ICondicionContext {
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

func (s *WhileContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEIZQ, 0)
}

func (s *WhileContext) L_sentencias() IL_sentenciasContext {
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

func (s *WhileContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEDER, 0)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitWhile(s)
	}
}

func (s *WhileContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitWhile(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) While_sentencia() (localctx IWhile_sentenciaContext) {
	localctx = NewWhile_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Tswift_GrammarNParserRULE_while_sentencia)
	localctx = NewWhileContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(Tswift_GrammarNParserWHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(197)
		p.condicion(0)
	}
	{
		p.SetState(198)
		p.Match(Tswift_GrammarNParserLLAVEIZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.L_sentencias()
	}
	{
		p.SetState(200)
		p.Match(Tswift_GrammarNParserLLAVEDER)
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

// IFor_sentenciaContext is an interface to support dynamic dispatch.
type IFor_sentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFor_sentenciaContext differentiates from other interfaces.
	IsFor_sentenciaContext()
}

type For_sentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_sentenciaContext() *For_sentenciaContext {
	var p = new(For_sentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_for_sentencia
	return p
}

func InitEmptyFor_sentenciaContext(p *For_sentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_for_sentencia
}

func (*For_sentenciaContext) IsFor_sentenciaContext() {}

func NewFor_sentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_sentenciaContext {
	var p = new(For_sentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_for_sentencia

	return p
}

func (s *For_sentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *For_sentenciaContext) CopyAll(ctx *For_sentenciaContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *For_sentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_sentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ForIntContext struct {
	For_sentenciaContext
}

func NewForIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForIntContext {
	var p = new(ForIntContext)

	InitEmptyFor_sentenciaContext(&p.For_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_sentenciaContext))

	return p
}

func (s *ForIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForIntContext) FOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserFOR, 0)
}

func (s *ForIntContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *ForIntContext) IN() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIN, 0)
}

func (s *ForIntContext) AllE() []IEContext {
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

func (s *ForIntContext) E(i int) IEContext {
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

func (s *ForIntContext) RANGO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserRANGO, 0)
}

func (s *ForIntContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEIZQ, 0)
}

func (s *ForIntContext) L_sentencias() IL_sentenciasContext {
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

func (s *ForIntContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEDER, 0)
}

func (s *ForIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterForInt(s)
	}
}

func (s *ForIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitForInt(s)
	}
}

func (s *ForIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitForInt(s)

	default:
		return t.VisitChildren(s)
	}
}

type ForListContext struct {
	For_sentenciaContext
}

func NewForListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ForListContext {
	var p = new(ForListContext)

	InitEmptyFor_sentenciaContext(&p.For_sentenciaContext)
	p.parser = parser
	p.CopyAll(ctx.(*For_sentenciaContext))

	return p
}

func (s *ForListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForListContext) FOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserFOR, 0)
}

func (s *ForListContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *ForListContext) IN() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIN, 0)
}

func (s *ForListContext) E() IEContext {
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

func (s *ForListContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEIZQ, 0)
}

func (s *ForListContext) L_sentencias() IL_sentenciasContext {
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

func (s *ForListContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLLAVEDER, 0)
}

func (s *ForListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterForList(s)
	}
}

func (s *ForListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitForList(s)
	}
}

func (s *ForListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitForList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) For_sentencia() (localctx IFor_sentenciaContext) {
	localctx = NewFor_sentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Tswift_GrammarNParserRULE_for_sentencia)
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		localctx = NewForIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(202)
			p.Match(Tswift_GrammarNParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(203)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(204)
			p.Match(Tswift_GrammarNParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(205)
			p.e(0)
		}
		{
			p.SetState(206)
			p.Match(Tswift_GrammarNParserRANGO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(207)
			p.e(0)
		}
		{
			p.SetState(208)
			p.Match(Tswift_GrammarNParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(209)
			p.L_sentencias()
		}
		{
			p.SetState(210)
			p.Match(Tswift_GrammarNParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewForListContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(212)
			p.Match(Tswift_GrammarNParserFOR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(213)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(214)
			p.Match(Tswift_GrammarNParserIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.e(0)
		}
		{
			p.SetState(216)
			p.Match(Tswift_GrammarNParserLLAVEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(217)
			p.L_sentencias()
		}
		{
			p.SetState(218)
			p.Match(Tswift_GrammarNParserLLAVEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IDec_vectorContext is an interface to support dynamic dispatch.
type IDec_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDec_vectorContext differentiates from other interfaces.
	IsDec_vectorContext()
}

type Dec_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDec_vectorContext() *Dec_vectorContext {
	var p = new(Dec_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_dec_vector
	return p
}

func InitEmptyDec_vectorContext(p *Dec_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_dec_vector
}

func (*Dec_vectorContext) IsDec_vectorContext() {}

func NewDec_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_vectorContext {
	var p = new(Dec_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_dec_vector

	return p
}

func (s *Dec_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_vectorContext) CopyAll(ctx *Dec_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Dec_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_VectorContext struct {
	Dec_vectorContext
	tipod antlr.Token
}

func NewDeclaracion_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_VectorContext {
	var p = new(Declaracion_VectorContext)

	InitEmptyDec_vectorContext(&p.Dec_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Dec_vectorContext))

	return p
}

func (s *Declaracion_VectorContext) GetTipod() antlr.Token { return s.tipod }

func (s *Declaracion_VectorContext) SetTipod(v antlr.Token) { s.tipod = v }

func (s *Declaracion_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_VectorContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Declaracion_VectorContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDOSPT, 0)
}

func (s *Declaracion_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Declaracion_VectorContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Declaracion_VectorContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIGUAL, 0)
}

func (s *Declaracion_VectorContext) Def_vector() IDef_vectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDef_vectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDef_vectorContext)
}

func (s *Declaracion_VectorContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserVAR, 0)
}

func (s *Declaracion_VectorContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLET, 0)
}

func (s *Declaracion_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDeclaracion_Vector(s)
	}
}

func (s *Declaracion_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDeclaracion_Vector(s)
	}
}

func (s *Declaracion_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDeclaracion_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Declaracion_AlternaContext struct {
	Dec_vectorContext
	tipod antlr.Token
}

func NewDeclaracion_AlternaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_AlternaContext {
	var p = new(Declaracion_AlternaContext)

	InitEmptyDec_vectorContext(&p.Dec_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Dec_vectorContext))

	return p
}

func (s *Declaracion_AlternaContext) GetTipod() antlr.Token { return s.tipod }

func (s *Declaracion_AlternaContext) SetTipod(v antlr.Token) { s.tipod = v }

func (s *Declaracion_AlternaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_AlternaContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Declaracion_AlternaContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIGUAL, 0)
}

func (s *Declaracion_AlternaContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Declaracion_AlternaContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_AlternaContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Declaracion_AlternaContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARIZQ, 0)
}

func (s *Declaracion_AlternaContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARDER, 0)
}

func (s *Declaracion_AlternaContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserVAR, 0)
}

func (s *Declaracion_AlternaContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLET, 0)
}

func (s *Declaracion_AlternaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDeclaracion_Alterna(s)
	}
}

func (s *Declaracion_AlternaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDeclaracion_Alterna(s)
	}
}

func (s *Declaracion_AlternaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDeclaracion_Alterna(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Dec_vector() (localctx IDec_vectorContext) {
	localctx = NewDec_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Tswift_GrammarNParserRULE_dec_vector)
	var _la int

	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclaracion_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(222)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Declaracion_VectorContext).tipod = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarNParserVAR || _la == Tswift_GrammarNParserLET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Declaracion_VectorContext).tipod = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(223)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(224)
			p.Match(Tswift_GrammarNParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(225)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(226)
			p.Tipo()
		}
		{
			p.SetState(227)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(228)
			p.Match(Tswift_GrammarNParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.Def_vector()
		}

	case 2:
		localctx = NewDeclaracion_AlternaContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(231)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Declaracion_AlternaContext).tipod = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarNParserVAR || _la == Tswift_GrammarNParserLET) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Declaracion_AlternaContext).tipod = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(232)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.Match(Tswift_GrammarNParserIGUAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(234)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.Tipo()
		}
		{
			p.SetState(236)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.Match(Tswift_GrammarNParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IDef_vectorContext is an interface to support dynamic dispatch.
type IDef_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDef_vectorContext differentiates from other interfaces.
	IsDef_vectorContext()
}

type Def_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDef_vectorContext() *Def_vectorContext {
	var p = new(Def_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_def_vector
	return p
}

func InitEmptyDef_vectorContext(p *Def_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_def_vector
}

func (*Def_vectorContext) IsDef_vectorContext() {}

func NewDef_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Def_vectorContext {
	var p = new(Def_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_def_vector

	return p
}

func (s *Def_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Def_vectorContext) CopyAll(ctx *Def_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Def_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Def_Vector_VacioContext struct {
	Def_vectorContext
}

func NewDef_Vector_VacioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Vector_VacioContext {
	var p = new(Def_Vector_VacioContext)

	InitEmptyDef_vectorContext(&p.Def_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_vectorContext))

	return p
}

func (s *Def_Vector_VacioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Vector_VacioContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Def_Vector_VacioContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Def_Vector_VacioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDef_Vector_Vacio(s)
	}
}

func (s *Def_Vector_VacioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDef_Vector_Vacio(s)
	}
}

func (s *Def_Vector_VacioContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDef_Vector_Vacio(s)

	default:
		return t.VisitChildren(s)
	}
}

type Def_VectorContext struct {
	Def_vectorContext
}

func NewDef_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_VectorContext {
	var p = new(Def_VectorContext)

	InitEmptyDef_vectorContext(&p.Def_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_vectorContext))

	return p
}

func (s *Def_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Def_VectorContext) AllE() []IEContext {
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

func (s *Def_VectorContext) E(i int) IEContext {
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

func (s *Def_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Def_VectorContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserCOMA)
}

func (s *Def_VectorContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCOMA, i)
}

func (s *Def_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDef_Vector(s)
	}
}

func (s *Def_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDef_Vector(s)
	}
}

func (s *Def_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDef_Vector(s)

	default:
		return t.VisitChildren(s)
	}
}

type Def_Vector_IdContext struct {
	Def_vectorContext
}

func NewDef_Vector_IdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Vector_IdContext {
	var p = new(Def_Vector_IdContext)

	InitEmptyDef_vectorContext(&p.Def_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_vectorContext))

	return p
}

func (s *Def_Vector_IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Vector_IdContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Def_Vector_IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDef_Vector_Id(s)
	}
}

func (s *Def_Vector_IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDef_Vector_Id(s)
	}
}

func (s *Def_Vector_IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDef_Vector_Id(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Def_vector() (localctx IDef_vectorContext) {
	localctx = NewDef_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Tswift_GrammarNParserRULE_def_vector)
	var _la int

	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDef_VectorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.e(0)
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Tswift_GrammarNParserCOMA {
			{
				p.SetState(244)
				p.Match(Tswift_GrammarNParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(245)
				p.e(0)
			}

			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(251)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDef_Vector_VacioContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(254)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewDef_Vector_IdContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(255)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IFunc_vectorContext is an interface to support dynamic dispatch.
type IFunc_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFunc_vectorContext differentiates from other interfaces.
	IsFunc_vectorContext()
}

type Func_vectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunc_vectorContext() *Func_vectorContext {
	var p = new(Func_vectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_func_vector
	return p
}

func InitEmptyFunc_vectorContext(p *Func_vectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_func_vector
}

func (*Func_vectorContext) IsFunc_vectorContext() {}

func NewFunc_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Func_vectorContext {
	var p = new(Func_vectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_func_vector

	return p
}

func (s *Func_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Func_vectorContext) CopyAll(ctx *Func_vectorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Func_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Func_Vector_AppendContext struct {
	Func_vectorContext
}

func NewFunc_Vector_AppendContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_Vector_AppendContext {
	var p = new(Func_Vector_AppendContext)

	InitEmptyFunc_vectorContext(&p.Func_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_vectorContext))

	return p
}

func (s *Func_Vector_AppendContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Vector_AppendContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Func_Vector_AppendContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPUNTO, 0)
}

func (s *Func_Vector_AppendContext) APPEND() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserAPPEND, 0)
}

func (s *Func_Vector_AppendContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARIZQ, 0)
}

func (s *Func_Vector_AppendContext) E() IEContext {
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

func (s *Func_Vector_AppendContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARDER, 0)
}

func (s *Func_Vector_AppendContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterFunc_Vector_Append(s)
	}
}

func (s *Func_Vector_AppendContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitFunc_Vector_Append(s)
	}
}

func (s *Func_Vector_AppendContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitFunc_Vector_Append(s)

	default:
		return t.VisitChildren(s)
	}
}

type Func_Vector_RemoveContext struct {
	Func_vectorContext
}

func NewFunc_Vector_RemoveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_Vector_RemoveContext {
	var p = new(Func_Vector_RemoveContext)

	InitEmptyFunc_vectorContext(&p.Func_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_vectorContext))

	return p
}

func (s *Func_Vector_RemoveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Vector_RemoveContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Func_Vector_RemoveContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPUNTO, 0)
}

func (s *Func_Vector_RemoveContext) REMOVE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserREMOVE, 0)
}

func (s *Func_Vector_RemoveContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARIZQ, 0)
}

func (s *Func_Vector_RemoveContext) AT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserAT, 0)
}

func (s *Func_Vector_RemoveContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDOSPT, 0)
}

func (s *Func_Vector_RemoveContext) E() IEContext {
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

func (s *Func_Vector_RemoveContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARDER, 0)
}

func (s *Func_Vector_RemoveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterFunc_Vector_Remove(s)
	}
}

func (s *Func_Vector_RemoveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitFunc_Vector_Remove(s)
	}
}

func (s *Func_Vector_RemoveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitFunc_Vector_Remove(s)

	default:
		return t.VisitChildren(s)
	}
}

type Func_Vector_RemoveLastContext struct {
	Func_vectorContext
}

func NewFunc_Vector_RemoveLastContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Func_Vector_RemoveLastContext {
	var p = new(Func_Vector_RemoveLastContext)

	InitEmptyFunc_vectorContext(&p.Func_vectorContext)
	p.parser = parser
	p.CopyAll(ctx.(*Func_vectorContext))

	return p
}

func (s *Func_Vector_RemoveLastContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Func_Vector_RemoveLastContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Func_Vector_RemoveLastContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPUNTO, 0)
}

func (s *Func_Vector_RemoveLastContext) REMOVELAST() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserREMOVELAST, 0)
}

func (s *Func_Vector_RemoveLastContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARIZQ, 0)
}

func (s *Func_Vector_RemoveLastContext) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARDER, 0)
}

func (s *Func_Vector_RemoveLastContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterFunc_Vector_RemoveLast(s)
	}
}

func (s *Func_Vector_RemoveLastContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitFunc_Vector_RemoveLast(s)
	}
}

func (s *Func_Vector_RemoveLastContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitFunc_Vector_RemoveLast(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Func_vector() (localctx IFunc_vectorContext) {
	localctx = NewFunc_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Tswift_GrammarNParserRULE_func_vector)
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		localctx = NewFunc_Vector_AppendContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(258)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(259)
			p.Match(Tswift_GrammarNParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(260)
			p.Match(Tswift_GrammarNParserAPPEND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(261)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(262)
			p.e(0)
		}
		{
			p.SetState(263)
			p.Match(Tswift_GrammarNParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFunc_Vector_RemoveLastContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(266)
			p.Match(Tswift_GrammarNParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(267)
			p.Match(Tswift_GrammarNParserREMOVELAST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(269)
			p.Match(Tswift_GrammarNParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewFunc_Vector_RemoveContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(270)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(271)
			p.Match(Tswift_GrammarNParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(272)
			p.Match(Tswift_GrammarNParserREMOVE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(273)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(274)
			p.Match(Tswift_GrammarNParserAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.Match(Tswift_GrammarNParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(276)
			p.e(0)
		}
		{
			p.SetState(277)
			p.Match(Tswift_GrammarNParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IDec_matrizContext is an interface to support dynamic dispatch.
type IDec_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDec_matrizContext differentiates from other interfaces.
	IsDec_matrizContext()
}

type Dec_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDec_matrizContext() *Dec_matrizContext {
	var p = new(Dec_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_dec_matriz
	return p
}

func InitEmptyDec_matrizContext(p *Dec_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_dec_matriz
}

func (*Dec_matrizContext) IsDec_matrizContext() {}

func NewDec_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Dec_matrizContext {
	var p = new(Dec_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_dec_matriz

	return p
}

func (s *Dec_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Dec_matrizContext) CopyAll(ctx *Dec_matrizContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Dec_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Dec_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Declaracion_MatrizContext struct {
	Dec_matrizContext
	tipod antlr.Token
}

func NewDeclaracion_MatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Declaracion_MatrizContext {
	var p = new(Declaracion_MatrizContext)

	InitEmptyDec_matrizContext(&p.Dec_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Dec_matrizContext))

	return p
}

func (s *Declaracion_MatrizContext) GetTipod() antlr.Token { return s.tipod }

func (s *Declaracion_MatrizContext) SetTipod(v antlr.Token) { s.tipod = v }

func (s *Declaracion_MatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_MatrizContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Declaracion_MatrizContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIGUAL, 0)
}

func (s *Declaracion_MatrizContext) Def_matriz() IDef_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDef_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDef_matrizContext)
}

func (s *Declaracion_MatrizContext) VAR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserVAR, 0)
}

func (s *Declaracion_MatrizContext) LET() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserLET, 0)
}

func (s *Declaracion_MatrizContext) DOSPT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDOSPT, 0)
}

func (s *Declaracion_MatrizContext) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Declaracion_MatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDeclaracion_Matriz(s)
	}
}

func (s *Declaracion_MatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDeclaracion_Matriz(s)
	}
}

func (s *Declaracion_MatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDeclaracion_Matriz(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Dec_matriz() (localctx IDec_matrizContext) {
	localctx = NewDec_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Tswift_GrammarNParserRULE_dec_matriz)
	var _la int

	localctx = NewDeclaracion_MatrizContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(281)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*Declaracion_MatrizContext).tipod = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == Tswift_GrammarNParserVAR || _la == Tswift_GrammarNParserLET) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*Declaracion_MatrizContext).tipod = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(282)
		p.Match(Tswift_GrammarNParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(285)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Tswift_GrammarNParserDOSPT {
		{
			p.SetState(283)
			p.Match(Tswift_GrammarNParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(284)
			p.Tipo()
		}

	}
	{
		p.SetState(287)
		p.Match(Tswift_GrammarNParserIGUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(288)
		p.Def_matriz()
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

// IDef_matrizContext is an interface to support dynamic dispatch.
type IDef_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDef_matrizContext differentiates from other interfaces.
	IsDef_matrizContext()
}

type Def_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDef_matrizContext() *Def_matrizContext {
	var p = new(Def_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_def_matriz
	return p
}

func InitEmptyDef_matrizContext(p *Def_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_def_matriz
}

func (*Def_matrizContext) IsDef_matrizContext() {}

func NewDef_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Def_matrizContext {
	var p = new(Def_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_def_matriz

	return p
}

func (s *Def_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Def_matrizContext) CopyAll(ctx *Def_matrizContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Def_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Def_Matriz_ListaContext struct {
	Def_matrizContext
}

func NewDef_Matriz_ListaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Matriz_ListaContext {
	var p = new(Def_Matriz_ListaContext)

	InitEmptyDef_matrizContext(&p.Def_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_matrizContext))

	return p
}

func (s *Def_Matriz_ListaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Matriz_ListaContext) Listavalores_matriz() IListavalores_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavalores_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListavalores_matrizContext)
}

func (s *Def_Matriz_ListaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDef_Matriz_Lista(s)
	}
}

func (s *Def_Matriz_ListaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDef_Matriz_Lista(s)
	}
}

func (s *Def_Matriz_ListaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDef_Matriz_Lista(s)

	default:
		return t.VisitChildren(s)
	}
}

type Def_Matriz_SimpleContext struct {
	Def_matrizContext
}

func NewDef_Matriz_SimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Matriz_SimpleContext {
	var p = new(Def_Matriz_SimpleContext)

	InitEmptyDef_matrizContext(&p.Def_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Def_matrizContext))

	return p
}

func (s *Def_Matriz_SimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Matriz_SimpleContext) Simple_matriz() ISimple_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_matrizContext)
}

func (s *Def_Matriz_SimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDef_Matriz_Simple(s)
	}
}

func (s *Def_Matriz_SimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDef_Matriz_Simple(s)
	}
}

func (s *Def_Matriz_SimpleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDef_Matriz_Simple(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Def_matriz() (localctx IDef_matrizContext) {
	localctx = NewDef_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Tswift_GrammarNParserRULE_def_matriz)
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDef_Matriz_ListaContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(290)
			p.Listavalores_matriz()
		}

	case 2:
		localctx = NewDef_Matriz_SimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(291)
			p.Simple_matriz()
		}

	case antlr.ATNInvalidAltNumber:
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

// IListavalores_matrizContext is an interface to support dynamic dispatch.
type IListavalores_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListavalores_matrizContext differentiates from other interfaces.
	IsListavalores_matrizContext()
}

type Listavalores_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListavalores_matrizContext() *Listavalores_matrizContext {
	var p = new(Listavalores_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_listavalores_matriz
	return p
}

func InitEmptyListavalores_matrizContext(p *Listavalores_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_listavalores_matriz
}

func (*Listavalores_matrizContext) IsListavalores_matrizContext() {}

func NewListavalores_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Listavalores_matrizContext {
	var p = new(Listavalores_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_listavalores_matriz

	return p
}

func (s *Listavalores_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Listavalores_matrizContext) CopyAll(ctx *Listavalores_matrizContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Listavalores_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Listavalores_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Def_Matriz_ExpresionContext struct {
	Listavalores_matrizContext
}

func NewDef_Matriz_ExpresionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Matriz_ExpresionContext {
	var p = new(Def_Matriz_ExpresionContext)

	InitEmptyListavalores_matrizContext(&p.Listavalores_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Listavalores_matrizContext))

	return p
}

func (s *Def_Matriz_ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Matriz_ExpresionContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Def_Matriz_ExpresionContext) AllE() []IEContext {
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

func (s *Def_Matriz_ExpresionContext) E(i int) IEContext {
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

func (s *Def_Matriz_ExpresionContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Def_Matriz_ExpresionContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserCOMA)
}

func (s *Def_Matriz_ExpresionContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCOMA, i)
}

func (s *Def_Matriz_ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDef_Matriz_Expresion(s)
	}
}

func (s *Def_Matriz_ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDef_Matriz_Expresion(s)
	}
}

func (s *Def_Matriz_ExpresionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDef_Matriz_Expresion(s)

	default:
		return t.VisitChildren(s)
	}
}

type Def_Matriz_Lista_ValoresContext struct {
	Listavalores_matrizContext
}

func NewDef_Matriz_Lista_ValoresContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Matriz_Lista_ValoresContext {
	var p = new(Def_Matriz_Lista_ValoresContext)

	InitEmptyListavalores_matrizContext(&p.Listavalores_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Listavalores_matrizContext))

	return p
}

func (s *Def_Matriz_Lista_ValoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Matriz_Lista_ValoresContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Def_Matriz_Lista_ValoresContext) AllListavalores_matriz() []IListavalores_matrizContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IListavalores_matrizContext); ok {
			len++
		}
	}

	tst := make([]IListavalores_matrizContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IListavalores_matrizContext); ok {
			tst[i] = t.(IListavalores_matrizContext)
			i++
		}
	}

	return tst
}

func (s *Def_Matriz_Lista_ValoresContext) Listavalores_matriz(i int) IListavalores_matrizContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListavalores_matrizContext); ok {
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

	return t.(IListavalores_matrizContext)
}

func (s *Def_Matriz_Lista_ValoresContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Def_Matriz_Lista_ValoresContext) AllCOMA() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserCOMA)
}

func (s *Def_Matriz_Lista_ValoresContext) COMA(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCOMA, i)
}

func (s *Def_Matriz_Lista_ValoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDef_Matriz_Lista_Valores(s)
	}
}

func (s *Def_Matriz_Lista_ValoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDef_Matriz_Lista_Valores(s)
	}
}

func (s *Def_Matriz_Lista_ValoresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDef_Matriz_Lista_Valores(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Listavalores_matriz() (localctx IListavalores_matrizContext) {
	localctx = NewListavalores_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Tswift_GrammarNParserRULE_listavalores_matriz)
	var _la int

	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDef_Matriz_Lista_ValoresContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.Listavalores_matriz()
		}
		p.SetState(300)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Tswift_GrammarNParserCOMA {
			{
				p.SetState(296)
				p.Match(Tswift_GrammarNParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(297)
				p.Listavalores_matriz()
			}

			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(303)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDef_Matriz_ExpresionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(305)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(306)
			p.e(0)
		}
		p.SetState(311)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Tswift_GrammarNParserCOMA {
			{
				p.SetState(307)
				p.Match(Tswift_GrammarNParserCOMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(308)
				p.e(0)
			}

			p.SetState(313)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(314)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// ISimple_matrizContext is an interface to support dynamic dispatch.
type ISimple_matrizContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSimple_matrizContext differentiates from other interfaces.
	IsSimple_matrizContext()
}

type Simple_matrizContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_matrizContext() *Simple_matrizContext {
	var p = new(Simple_matrizContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_simple_matriz
	return p
}

func InitEmptySimple_matrizContext(p *Simple_matrizContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Tswift_GrammarNParserRULE_simple_matriz
}

func (*Simple_matrizContext) IsSimple_matrizContext() {}

func NewSimple_matrizContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_matrizContext {
	var p = new(Simple_matrizContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Tswift_GrammarNParserRULE_simple_matriz

	return p
}

func (s *Simple_matrizContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_matrizContext) CopyAll(ctx *Simple_matrizContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *Simple_matrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_matrizContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Def_Matriz_Simple3Context struct {
	Simple_matrizContext
}

func NewDef_Matriz_Simple3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Matriz_Simple3Context {
	var p = new(Def_Matriz_Simple3Context)

	InitEmptySimple_matrizContext(&p.Simple_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Simple_matrizContext))

	return p
}

func (s *Def_Matriz_Simple3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Matriz_Simple3Context) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Def_Matriz_Simple3Context) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Def_Matriz_Simple3Context) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Def_Matriz_Simple3Context) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARIZQ, 0)
}

func (s *Def_Matriz_Simple3Context) REPEATING() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserREPEATING, 0)
}

func (s *Def_Matriz_Simple3Context) AllDOSPT() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserDOSPT)
}

func (s *Def_Matriz_Simple3Context) DOSPT(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDOSPT, i)
}

func (s *Def_Matriz_Simple3Context) AllE() []IEContext {
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

func (s *Def_Matriz_Simple3Context) E(i int) IEContext {
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

func (s *Def_Matriz_Simple3Context) COMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCOMA, 0)
}

func (s *Def_Matriz_Simple3Context) COUNT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCOUNT, 0)
}

func (s *Def_Matriz_Simple3Context) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARDER, 0)
}

func (s *Def_Matriz_Simple3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDef_Matriz_Simple3(s)
	}
}

func (s *Def_Matriz_Simple3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDef_Matriz_Simple3(s)
	}
}

func (s *Def_Matriz_Simple3Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDef_Matriz_Simple3(s)

	default:
		return t.VisitChildren(s)
	}
}

type Def_Matriz_Simple2Context struct {
	Simple_matrizContext
}

func NewDef_Matriz_Simple2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Def_Matriz_Simple2Context {
	var p = new(Def_Matriz_Simple2Context)

	InitEmptySimple_matrizContext(&p.Simple_matrizContext)
	p.parser = parser
	p.CopyAll(ctx.(*Simple_matrizContext))

	return p
}

func (s *Def_Matriz_Simple2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Def_Matriz_Simple2Context) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Def_Matriz_Simple2Context) Tipo() ITipoContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITipoContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITipoContext)
}

func (s *Def_Matriz_Simple2Context) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Def_Matriz_Simple2Context) PARIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARIZQ, 0)
}

func (s *Def_Matriz_Simple2Context) REPEATING() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserREPEATING, 0)
}

func (s *Def_Matriz_Simple2Context) AllDOSPT() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserDOSPT)
}

func (s *Def_Matriz_Simple2Context) DOSPT(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDOSPT, i)
}

func (s *Def_Matriz_Simple2Context) Simple_matriz() ISimple_matrizContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimple_matrizContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimple_matrizContext)
}

func (s *Def_Matriz_Simple2Context) COMA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCOMA, 0)
}

func (s *Def_Matriz_Simple2Context) COUNT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCOUNT, 0)
}

func (s *Def_Matriz_Simple2Context) E() IEContext {
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

func (s *Def_Matriz_Simple2Context) PARDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPARDER, 0)
}

func (s *Def_Matriz_Simple2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterDef_Matriz_Simple2(s)
	}
}

func (s *Def_Matriz_Simple2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitDef_Matriz_Simple2(s)
	}
}

func (s *Def_Matriz_Simple2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitDef_Matriz_Simple2(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Tswift_GrammarNParser) Simple_matriz() (localctx ISimple_matrizContext) {
	localctx = NewSimple_matrizContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Tswift_GrammarNParserRULE_simple_matriz)
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDef_Matriz_Simple2Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(318)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(319)
			p.Tipo()
		}
		{
			p.SetState(320)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(322)
			p.Match(Tswift_GrammarNParserREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(323)
			p.Match(Tswift_GrammarNParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(324)
			p.Simple_matriz()
		}
		{
			p.SetState(325)
			p.Match(Tswift_GrammarNParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(326)
			p.Match(Tswift_GrammarNParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(327)
			p.Match(Tswift_GrammarNParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(328)
			p.e(0)
		}
		{
			p.SetState(329)
			p.Match(Tswift_GrammarNParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDef_Matriz_Simple3Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(331)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(332)
			p.Tipo()
		}
		{
			p.SetState(333)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(335)
			p.Match(Tswift_GrammarNParserREPEATING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(336)
			p.Match(Tswift_GrammarNParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(337)
			p.e(0)
		}
		{
			p.SetState(338)
			p.Match(Tswift_GrammarNParserCOMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.Match(Tswift_GrammarNParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(340)
			p.Match(Tswift_GrammarNParserDOSPT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(341)
			p.e(0)
		}
		{
			p.SetState(342)
			p.Match(Tswift_GrammarNParserPARDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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
	_startState := 44
	p.EnterRecursionRule(localctx, 44, Tswift_GrammarNParserRULE_condicion, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCond_NegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(347)

			var _m = p.Match(Tswift_GrammarNParserNOT)

			localctx.(*Cond_NegContext).op = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(348)

			var _x = p.condicion(5)

			localctx.(*Cond_NegContext).c = _x
		}

	case 2:
		localctx = NewCond_RelContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(349)

			var _x = p.e(0)

			localctx.(*Cond_RelContext).e1 = _x
		}
		{
			p.SetState(350)

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
			p.SetState(351)

			var _x = p.e(0)

			localctx.(*Cond_RelContext).e2 = _x
		}

	case 3:
		localctx = NewCond_BooleanoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(353)

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
			p.SetState(354)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)

			var _x = p.condicion(0)

			localctx.(*Cond_ParContext).c = _x
		}
		{
			p.SetState(356)
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
	p.SetState(365)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
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
			p.SetState(360)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(361)

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
				p.SetState(362)

				var _x = p.condicion(4)

				localctx.(*Cond_LogicaContext).c2 = _x
			}

		}
		p.SetState(367)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext())
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

type Expr_RelContext struct {
	EContext
	op antlr.Token
}

func NewExpr_RelContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_RelContext {
	var p = new(Expr_RelContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_RelContext) GetOp() antlr.Token { return s.op }

func (s *Expr_RelContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_RelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_RelContext) AllE() []IEContext {
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

func (s *Expr_RelContext) E(i int) IEContext {
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

func (s *Expr_RelContext) IGUALIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserIGUALIGUAL, 0)
}

func (s *Expr_RelContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDIFERENTE, 0)
}

func (s *Expr_RelContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMAYORIGUAL, 0)
}

func (s *Expr_RelContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMAYOR, 0)
}

func (s *Expr_RelContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMENORIGUAL, 0)
}

func (s *Expr_RelContext) MENOR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMENOR, 0)
}

func (s *Expr_RelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Rel(s)
	}
}

func (s *Expr_RelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Rel(s)
	}
}

func (s *Expr_RelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Rel(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_DecimalContext struct {
	EContext
	n antlr.Token
}

func NewExpr_DecimalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_DecimalContext {
	var p = new(Expr_DecimalContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_DecimalContext) GetN() antlr.Token { return s.n }

func (s *Expr_DecimalContext) SetN(v antlr.Token) { s.n = v }

func (s *Expr_DecimalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_DecimalContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserDECIMAL, 0)
}

func (s *Expr_DecimalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Decimal(s)
	}
}

func (s *Expr_DecimalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Decimal(s)
	}
}

func (s *Expr_DecimalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Decimal(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_CaracterContext struct {
	EContext
	n antlr.Token
}

func NewExpr_CaracterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_CaracterContext {
	var p = new(Expr_CaracterContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_CaracterContext) GetN() antlr.Token { return s.n }

func (s *Expr_CaracterContext) SetN(v antlr.Token) { s.n = v }

func (s *Expr_CaracterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_CaracterContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCARACTER, 0)
}

func (s *Expr_CaracterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Caracter(s)
	}
}

func (s *Expr_CaracterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Caracter(s)
	}
}

func (s *Expr_CaracterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Caracter(s)

	default:
		return t.VisitChildren(s)
	}
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

type Expr_MatrizContext struct {
	EContext
}

func NewExpr_MatrizContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_MatrizContext {
	var p = new(Expr_MatrizContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_MatrizContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_MatrizContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Expr_MatrizContext) AllCORCHETEIZQ() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserCORCHETEIZQ)
}

func (s *Expr_MatrizContext) CORCHETEIZQ(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, i)
}

func (s *Expr_MatrizContext) AllE() []IEContext {
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

func (s *Expr_MatrizContext) E(i int) IEContext {
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

func (s *Expr_MatrizContext) AllCORCHETEDER() []antlr.TerminalNode {
	return s.GetTokens(Tswift_GrammarNParserCORCHETEDER)
}

func (s *Expr_MatrizContext) CORCHETEDER(i int) antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, i)
}

func (s *Expr_MatrizContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Matriz(s)
	}
}

func (s *Expr_MatrizContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Matriz(s)
	}
}

func (s *Expr_MatrizContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Matriz(s)

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

func (s *Expr_NegContext) MENOS() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserMENOS, 0)
}

func (s *Expr_NegContext) NOT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserNOT, 0)
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

type Expr_NilContext struct {
	EContext
	n antlr.Token
}

func NewExpr_NilContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_NilContext {
	var p = new(Expr_NilContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_NilContext) GetN() antlr.Token { return s.n }

func (s *Expr_NilContext) SetN(v antlr.Token) { s.n = v }

func (s *Expr_NilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_NilContext) NIL() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserNIL, 0)
}

func (s *Expr_NilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Nil(s)
	}
}

func (s *Expr_NilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Nil(s)
	}
}

func (s *Expr_NilContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Nil(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_CadenaContext struct {
	EContext
	n antlr.Token
}

func NewExpr_CadenaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_CadenaContext {
	var p = new(Expr_CadenaContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_CadenaContext) GetN() antlr.Token { return s.n }

func (s *Expr_CadenaContext) SetN(v antlr.Token) { s.n = v }

func (s *Expr_CadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_CadenaContext) CADENA() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCADENA, 0)
}

func (s *Expr_CadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Cadena(s)
	}
}

func (s *Expr_CadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Cadena(s)
	}
}

func (s *Expr_CadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Cadena(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_CountContext struct {
	EContext
}

func NewExpr_CountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_CountContext {
	var p = new(Expr_CountContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_CountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_CountContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Expr_CountContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPUNTO, 0)
}

func (s *Expr_CountContext) COUNT() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCOUNT, 0)
}

func (s *Expr_CountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Count(s)
	}
}

func (s *Expr_CountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Count(s)
	}
}

func (s *Expr_CountContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Count(s)

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

type Expr_LogicaContext struct {
	EContext
	op antlr.Token
}

func NewExpr_LogicaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_LogicaContext {
	var p = new(Expr_LogicaContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_LogicaContext) GetOp() antlr.Token { return s.op }

func (s *Expr_LogicaContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_LogicaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_LogicaContext) AllE() []IEContext {
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

func (s *Expr_LogicaContext) E(i int) IEContext {
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

func (s *Expr_LogicaContext) AND() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserAND, 0)
}

func (s *Expr_LogicaContext) OR() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserOR, 0)
}

func (s *Expr_LogicaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Logica(s)
	}
}

func (s *Expr_LogicaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Logica(s)
	}
}

func (s *Expr_LogicaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Logica(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_IsEmptyContext struct {
	EContext
}

func NewExpr_IsEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_IsEmptyContext {
	var p = new(Expr_IsEmptyContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_IsEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_IsEmptyContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Expr_IsEmptyContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserPUNTO, 0)
}

func (s *Expr_IsEmptyContext) ISEMPTY() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserISEMPTY, 0)
}

func (s *Expr_IsEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_IsEmpty(s)
	}
}

func (s *Expr_IsEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_IsEmpty(s)
	}
}

func (s *Expr_IsEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_IsEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_BooleanoContext struct {
	EContext
	n antlr.Token
}

func NewExpr_BooleanoContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_BooleanoContext {
	var p = new(Expr_BooleanoContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_BooleanoContext) GetN() antlr.Token { return s.n }

func (s *Expr_BooleanoContext) SetN(v antlr.Token) { s.n = v }

func (s *Expr_BooleanoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_BooleanoContext) TRUE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserTRUE, 0)
}

func (s *Expr_BooleanoContext) FALSE() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserFALSE, 0)
}

func (s *Expr_BooleanoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Booleano(s)
	}
}

func (s *Expr_BooleanoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Booleano(s)
	}
}

func (s *Expr_BooleanoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Booleano(s)

	default:
		return t.VisitChildren(s)
	}
}

type Expr_VectorContext struct {
	EContext
}

func NewExpr_VectorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Expr_VectorContext {
	var p = new(Expr_VectorContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *Expr_VectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_VectorContext) ID() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserID, 0)
}

func (s *Expr_VectorContext) CORCHETEIZQ() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEIZQ, 0)
}

func (s *Expr_VectorContext) E() IEContext {
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

func (s *Expr_VectorContext) CORCHETEDER() antlr.TerminalNode {
	return s.GetToken(Tswift_GrammarNParserCORCHETEDER, 0)
}

func (s *Expr_VectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.EnterExpr_Vector(s)
	}
}

func (s *Expr_VectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Tswift_GrammarNListener); ok {
		listenerT.ExitExpr_Vector(s)
	}
}

func (s *Expr_VectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Tswift_GrammarNVisitor:
		return t.VisitExpr_Vector(s)

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
	_startState := 46
	p.EnterRecursionRule(localctx, 46, Tswift_GrammarNParserRULE_e, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExpr_NegContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(369)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Expr_NegContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarNParserMENOS || _la == Tswift_GrammarNParserNOT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Expr_NegContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(370)

			var _x = p.e(18)

			localctx.(*Expr_NegContext).e1 = _x
		}

	case 2:
		localctx = NewExpr_BooleanoContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(371)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*Expr_BooleanoContext).n = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == Tswift_GrammarNParserTRUE || _la == Tswift_GrammarNParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*Expr_BooleanoContext).n = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 3:
		localctx = NewExpr_NilContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(372)

			var _m = p.Match(Tswift_GrammarNParserNIL)

			localctx.(*Expr_NilContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewExpr_VectorContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(373)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(374)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(375)
			p.e(0)
		}
		{
			p.SetState(376)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		localctx = NewExpr_MatrizContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(378)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(379)
			p.Match(Tswift_GrammarNParserCORCHETEIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(380)
			p.e(0)
		}
		{
			p.SetState(381)
			p.Match(Tswift_GrammarNParserCORCHETEDER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(382)
					p.Match(Tswift_GrammarNParserCORCHETEIZQ)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(383)
					p.e(0)
				}
				{
					p.SetState(384)
					p.Match(Tswift_GrammarNParserCORCHETEDER)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(388)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 32, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 6:
		localctx = NewExpr_IsEmptyContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(390)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(391)
			p.Match(Tswift_GrammarNParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(392)
			p.Match(Tswift_GrammarNParserISEMPTY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		localctx = NewExpr_CountContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(393)
			p.Match(Tswift_GrammarNParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(394)
			p.Match(Tswift_GrammarNParserPUNTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(395)
			p.Match(Tswift_GrammarNParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		localctx = NewExpr_IdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(396)

			var _m = p.Match(Tswift_GrammarNParserID)

			localctx.(*Expr_IdContext).id = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		localctx = NewExpr_EnteroContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(397)

			var _m = p.Match(Tswift_GrammarNParserENTERO)

			localctx.(*Expr_EnteroContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		localctx = NewExpr_DecimalContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(398)

			var _m = p.Match(Tswift_GrammarNParserDECIMAL)

			localctx.(*Expr_DecimalContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		localctx = NewExpr_CadenaContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(399)

			var _m = p.Match(Tswift_GrammarNParserCADENA)

			localctx.(*Expr_CadenaContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		localctx = NewExpr_CaracterContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(400)

			var _m = p.Match(Tswift_GrammarNParserCARACTER)

			localctx.(*Expr_CaracterContext).n = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		localctx = NewExpr_ParContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(401)
			p.Match(Tswift_GrammarNParserPARIZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(402)

			var _x = p.e(0)

			localctx.(*Expr_ParContext).e1 = _x
		}
		{
			p.SetState(403)
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
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(422)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 34, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_MulDivContext(p, NewEContext(p, _parentctx, _parentState))
				localctx.(*Expr_MulDivContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarNParserRULE_e)
				p.SetState(407)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
					goto errorExit
				}
				{
					p.SetState(408)

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
					p.SetState(409)

					var _x = p.e(18)

					localctx.(*Expr_MulDivContext).e2 = _x
				}

			case 2:
				localctx = NewExpr_SumResContext(p, NewEContext(p, _parentctx, _parentState))
				localctx.(*Expr_SumResContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarNParserRULE_e)
				p.SetState(410)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
					goto errorExit
				}
				{
					p.SetState(411)

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
					p.SetState(412)

					var _x = p.e(17)

					localctx.(*Expr_SumResContext).e2 = _x
				}

			case 3:
				localctx = NewExpr_ModContext(p, NewEContext(p, _parentctx, _parentState))
				localctx.(*Expr_ModContext).e1 = _prevctx

				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarNParserRULE_e)
				p.SetState(413)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
					goto errorExit
				}
				{
					p.SetState(414)

					var _m = p.Match(Tswift_GrammarNParserMOD)

					localctx.(*Expr_ModContext).op = _m
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(415)

					var _x = p.e(16)

					localctx.(*Expr_ModContext).e2 = _x
				}

			case 4:
				localctx = NewExpr_RelContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarNParserRULE_e)
				p.SetState(416)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
					goto errorExit
				}
				{
					p.SetState(417)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_RelContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&264241152) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_RelContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(418)
					p.e(15)
				}

			case 5:
				localctx = NewExpr_LogicaContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, Tswift_GrammarNParserRULE_e)
				p.SetState(419)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(420)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_LogicaContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == Tswift_GrammarNParserAND || _la == Tswift_GrammarNParserOR) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_LogicaContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(421)
					p.e(14)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext())
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
	case 22:
		var t *CondicionContext = nil
		if localctx != nil {
			t = localctx.(*CondicionContext)
		}
		return p.Condicion_Sempred(t, predIndex)

	case 23:
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
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 13)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
