// Code generated from /home/josep/USAC/6to_Semestre/Lab_Compi/Proyecto2/OLC2_Proyecto2_202111478/parser/Tswift_Grammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tswift_Grammar

import "github.com/antlr4-go/antlr/v4"

// Tswift_GrammarListener is a complete listener for a parse tree produced by Tswift_GrammarParser.
type Tswift_GrammarListener interface {
	antlr.ParseTreeListener

	// EnterS is called when entering the s production.
	EnterS(c *SContext)

	// EnterCond is called when entering the cond production.
	EnterCond(c *CondContext)

	// EnterMarcador is called when entering the marcador production.
	EnterMarcador(c *MarcadorContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterOprel is called when entering the oprel production.
	EnterOprel(c *OprelContext)

	// ExitS is called when exiting the s production.
	ExitS(c *SContext)

	// ExitCond is called when exiting the cond production.
	ExitCond(c *CondContext)

	// ExitMarcador is called when exiting the marcador production.
	ExitMarcador(c *MarcadorContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitOprel is called when exiting the oprel production.
	ExitOprel(c *OprelContext)
}
