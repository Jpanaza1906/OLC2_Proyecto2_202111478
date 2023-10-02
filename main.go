package main

import (
	"OLC2_Proyecto2_202111478/api"
	"net/http"

	"github.com/rs/cors"
)

/*
	ESQUEMA DE TRADUCCION

	func analizar(input string) {
		//Se hace un stream
		stream := antlr.NewInputStream(input)
		//Se hace un lexer
		lexer := TswiftG.NewTswift_GrammarLexer(stream)
		// se hace un token
		tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		// se hace un parser
		parser := TswiftG.NewTswift_GrammarParser(tokenStream)
		// se hace un arbol
		tree := parser.S()

		listen := Listener{}

		antlr.ParseTreeWalkerDefault.Walk(&listen, tree)

}

	func leerEntrada(name string) string {
		file, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		fileScanner := bufio.NewScanner(file)
		var text string
		for fileScanner.Scan() {
			text += fileScanner.Text() + "\n"
		}
		return text
	}
*/
/*
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
	// se hace un visit
	raiz.Compilar(ctx)

	return ctx.Consola

}
*/
func main() {
	srv := api.NewServer()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})

	http.ListenAndServe(":8080", corsHandler.Handler(srv))
}
