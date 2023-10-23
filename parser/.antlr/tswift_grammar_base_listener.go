// Code generated from /home/josep/USAC/6to_Semestre/Lab_Compi/Proyecto2/OLC2_Proyecto2_202111478/parser/Tswift_Grammar.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Tswift_Grammar

import "github.com/antlr4-go/antlr/v4"

// BaseTswift_GrammarListener is a complete listener for a parse tree produced by Tswift_GrammarParser.
type BaseTswift_GrammarListener struct{}

var _ Tswift_GrammarListener = &BaseTswift_GrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTswift_GrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTswift_GrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTswift_GrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTswift_GrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterS is called when production s is entered.
func (s *BaseTswift_GrammarListener) EnterS(ctx *SContext) {}

// ExitS is called when production s is exited.
func (s *BaseTswift_GrammarListener) ExitS(ctx *SContext) {}

// EnterCond is called when production cond is entered.
func (s *BaseTswift_GrammarListener) EnterCond(ctx *CondContext) {}

// ExitCond is called when production cond is exited.
func (s *BaseTswift_GrammarListener) ExitCond(ctx *CondContext) {}

// EnterMarcador is called when production marcador is entered.
func (s *BaseTswift_GrammarListener) EnterMarcador(ctx *MarcadorContext) {}

// ExitMarcador is called when production marcador is exited.
func (s *BaseTswift_GrammarListener) ExitMarcador(ctx *MarcadorContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseTswift_GrammarListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseTswift_GrammarListener) ExitExpr(ctx *ExprContext) {}

// EnterOprel is called when production oprel is entered.
func (s *BaseTswift_GrammarListener) EnterOprel(ctx *OprelContext) {}

// ExitOprel is called when production oprel is exited.
func (s *BaseTswift_GrammarListener) ExitOprel(ctx *OprelContext) {}
