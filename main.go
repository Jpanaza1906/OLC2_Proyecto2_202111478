package main

import (
	compilador "OLC2_Proyecto2_202111478/Compilador"
	"OLC2_Proyecto2_202111478/TswiftG"
	"OLC2_Proyecto2_202111478/TswiftGen"
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

type Listener struct {
	*TswiftG.BaseTswift_GrammarListener
}

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
*/
func analizarNodos(input string) {
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

func main() {
	input := leerEntrada("entrada.txt")
	fmt.Println("SALIDA Nodos:==============")
	analizarNodos(input)
}
