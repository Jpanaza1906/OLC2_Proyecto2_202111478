package api

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"OLC2_Proyecto2_202111478/TswiftGen"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gorilla/mux"
)

type Entrada struct {
	Entrada string `json:"entrada"`
}

type Salida struct {
	Salida    string `json:"salida"`
	Consola   string `json:"consola"`
	Simbolos  string `json:"simbolos"`
	Errores   string `json:"errores"`
	Funciones string `json:"funciones"`
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []string
}

func (el *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	errorMessage := fmt.Sprintf("Error gramatical (%d,%d) : %s", line, column, msg)
	el.Errors = append(el.Errors, errorMessage)
}

func analizarNodos(input string) (string, string, string, string, string) {
	//Se hace un stream
	stream := antlr.NewInputStream(input)
	//Se hace un lexer
	lexer := TswiftGen.NewTswift_GrammarNLexer(stream)
	// se hace un token
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// se hace un parser
	parser := TswiftGen.NewTswift_GrammarNParser(tokenStream)

	// Agrega tu manejador de errores personalizado
	errorListener := &CustomErrorListener{}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorListener)
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	// se hace un arbol
	tree := parser.S()
	// se hace un visitor
	visitor := NewVisitor()
	// se hace un visit
	raiz := visitor.Visit(tree).(compilador.CAbstractExpr)
	// se hace un contexto
	ctx := compilador.NewContexto()

	//agregar ambito global
	ctx.PushAmbito()

	fmt.Println("Compilando---------------------------------------------------------")
	// se hace un visit
	raiz.Compilar(ctx)

	// Captura los errores generados por ANTLR
	antlrErrors := errorListener.Errors
	errorString := strings.Join(antlrErrors, "\n")

	ctx.Errores += errorString
	fmt.Print(ctx.Errores)

	var inicio string

	inicio = "#include <stdio.h>\n"
	inicio += "float stack[10000];\n"
	inicio += "float heap[10000];\n"
	inicio += "//Declaracion de variables\n"
	inicio += "float P;\n"
	inicio += "float H;\n"

	//se obtienen los temporales
	arraytemp := ctx.GetTemplist()
	if len(arraytemp) > 0 {
		inicio += "//Declaracion de temporales\n"
		inicio += "float "
		for _, t := range arraytemp {
			//se es el ultimo se agrega el ; en lugar de la ,
			if t == arraytemp[len(arraytemp)-1] {
				inicio += t + ";\n"
			} else {
				inicio += t + ", "
			}
		}
	}
	inicio += ctx.Nativas
	//funcion main
	inicio += "\nint main(){\n"
	//inicializan punteros P y H
	inicio += "\t//Inicializacion de punteros\n"
	inicio += "\tP = 0;\n"
	inicio += "\tH = 0;\n\n"

	fin := "\treturn 0;\n}"
	ctx.Consola = inicio + ctx.Consola + fin

	console := "Compilacion Exitosa"
	//si hay errores que devuelva consola vacia
	if ctx.Errores != "" {
		ctx.Consola = ""
		console = "Compilacion sin exito, ver consola de errores"
	}

	ctx.GetSimbolosCadena()

	return ctx.Consola, ctx.Errores, ctx.Simbolos, ctx.Funciones, console
}

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.Routes()
	return s
}

func (s *Server) Routes() {
	s.HandleFunc("/compilar", s.compile()).Methods("POST")
}

func (s *Server) compile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input Entrada
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		entry := input.Entrada
		consola, errores, simbolos, funciones, consoleout := analizarNodos(entry)
		var salida Salida
		salida.Salida = consola
		salida.Errores = errores
		salida.Simbolos = simbolos
		salida.Funciones = funciones
		salida.Consola = consoleout

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(salida); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
