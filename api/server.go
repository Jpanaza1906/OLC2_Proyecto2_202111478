package api

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"OLC2_Proyecto2_202111478/TswiftGen"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antlr4-go/antlr/v4"
	"github.com/gorilla/mux"
)

type Entrada struct {
	Entrada string `json:"entrada"`
}

type Salida struct {
	Salida string `json:"salida"`
}

func analizarNodos(input string) string {
	//Se hace un stream
	stream := antlr.NewInputStream(input)
	//Se hace un lexer
	lexer := TswiftGen.NewTswift_GrammarNLexer(stream)
	// se hace un token
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// se hace un parser
	parser := TswiftGen.NewTswift_GrammarNParser(tokenStream)
	// se hace un arbol
	tree := parser.S()
	// se hace un visitor
	visitor := NewVisitor()
	// se hace un visit
	raiz := visitor.Visit(tree).(compilador.CAbstractExpr)
	// se hace un contexto
	ctx := compilador.NewContexto()
	fmt.Println("Compilando---------------------------------------------------------")
	// se hace un visit
	raiz.Compilar(ctx)

	return ctx.Consola
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
		consola := analizarNodos(entry)
		var salida Salida
		salida.Salida = consola

		w.Header().Set("Content-Type", "application/json")

		if err := json.NewEncoder(w).Encode(salida); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
