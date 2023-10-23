// Generated from /home/josep/USAC/6to_Semestre/Lab_Compi/Proyecto2/OLC2_Proyecto2_202111478/parser/Tswift_Grammar.g4 by ANTLR 4.13.1

    //Es el codigo que incrusta al principio del archivo
    //Importacionies
    import "OLC2_Proyecto2_202111478/gen"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Tswift_GrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PARDER=1, PARIZQ=2, LLAVEIZQ=3, LLAVEDER=4, CORCHETEIZQ=5, CORCHETEDER=6, 
		DOSPT=7, PTCOMA=8, INTERROGACION=9, COMA=10, PUNTO=11, GUIONBAJO=12, DIR=13, 
		MASIGUAL=14, MENOSIGUAL=15, IGUAL=16, DIV=17, MOD=18, MENOS=19, MAS=20, 
		POR=21, IGUALIGUAL=22, DIFERENTE=23, MAYORIGUAL=24, MENORIGUAL=25, MAYOR=26, 
		MENOR=27, AND=28, OR=29, NOT=30, PRINT=31, VAR=32, LET=33, TRUE=34, FALSE=35, 
		IF=36, ELSE=37, SWITCH=38, CASE=39, DEFAULT=40, WHILE=41, FOR=42, IN=43, 
		RANGO=44, GUARD=45, CONTINUE=46, RETURN=47, BREAK=48, NIL=49, APPEND=50, 
		REMOVELAST=51, REMOVE=52, AT=53, ISEMPTY=54, COUNT=55, FUNC=56, REPEATING=57, 
		STRUCT=58, MUTATING=59, INOUT=60, ATOI=61, IOTA=62, SELF=63, INT=64, FLOAT=65, 
		BOOL=66, CHARACTER=67, STRING=68, ENBLANCO=69, ENTERO=70, DECIMAL=71, 
		CARACTER=72, CADENA=73, ID=74, UL_C=75, ML_C=76, ERROR=77;
	public static final int
		RULE_s = 0, RULE_cond = 1, RULE_marcador = 2, RULE_expr = 3, RULE_oprel = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "cond", "marcador", "expr", "oprel"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "')'", "'('", "'{'", "'}'", "'['", "']'", "':'", "';'", "'?'", 
			"','", "'.'", "'_'", "'&'", "'+='", "'-='", "'='", "'/'", "'%'", "'-'", 
			"'+'", "'*'", "'=='", "'!='", "'>='", "'<='", "'>'", "'<'", "'&&'", "'||'", 
			"'!'", "'print'", "'var'", "'let'", "'true'", "'false'", "'if'", "'else'", 
			"'switch'", "'case'", "'default'", "'while'", "'for'", "'in'", "'...'", 
			"'guard'", "'continue'", "'return'", "'break'", "'nil'", "'append'", 
			"'removeLast'", "'remove'", "'at'", "'IsEmpty'", "'count'", "'func'", 
			"'repeating'", "'struct'", "'mutating'", "'inout'", "'atoi'", "'iota'", 
			"'self'", "'Int'", "'Float'", "'Bool'", "'Character'", "'String'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PARDER", "PARIZQ", "LLAVEIZQ", "LLAVEDER", "CORCHETEIZQ", "CORCHETEDER", 
			"DOSPT", "PTCOMA", "INTERROGACION", "COMA", "PUNTO", "GUIONBAJO", "DIR", 
			"MASIGUAL", "MENOSIGUAL", "IGUAL", "DIV", "MOD", "MENOS", "MAS", "POR", 
			"IGUALIGUAL", "DIFERENTE", "MAYORIGUAL", "MENORIGUAL", "MAYOR", "MENOR", 
			"AND", "OR", "NOT", "PRINT", "VAR", "LET", "TRUE", "FALSE", "IF", "ELSE", 
			"SWITCH", "CASE", "DEFAULT", "WHILE", "FOR", "IN", "RANGO", "GUARD", 
			"CONTINUE", "RETURN", "BREAK", "NIL", "APPEND", "REMOVELAST", "REMOVE", 
			"AT", "ISEMPTY", "COUNT", "FUNC", "REPEATING", "STRUCT", "MUTATING", 
			"INOUT", "ATOI", "IOTA", "SELF", "INT", "FLOAT", "BOOL", "CHARACTER", 
			"STRING", "ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID", 
			"UL_C", "ML_C", "ERROR"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Tswift_Grammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }


	    //instancias de variables
	    //contador de temportales
	    var temp int = 1;

	public Tswift_GrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SContext extends ParserRuleContext {
		public CondContext c1;
		public TerminalNode EOF() { return getToken(Tswift_GrammarParser.EOF, 0); }
		public CondContext cond() {
			return getRuleContext(CondContext.class,0);
		}
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(10);
			((SContext)_localctx).c1 = cond(0);
			setState(11);
			match(EOF);

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CondContext extends ParserRuleContext {
		public []string EV;
		public []string EF;
		public CondContext c1;
		public Token op;
		public CondContext c;
		public ExprContext e1;
		public OprelContext opr;
		public ExprContext e2;
		public CondContext c2;
		public TerminalNode NOT() { return getToken(Tswift_GrammarParser.NOT, 0); }
		public List<CondContext> cond() {
			return getRuleContexts(CondContext.class);
		}
		public CondContext cond(int i) {
			return getRuleContext(CondContext.class,i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public OprelContext oprel() {
			return getRuleContext(OprelContext.class,0);
		}
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public TerminalNode TRUE() { return getToken(Tswift_GrammarParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Tswift_GrammarParser.FALSE, 0); }
		public MarcadorContext marcador() {
			return getRuleContext(MarcadorContext.class,0);
		}
		public TerminalNode AND() { return getToken(Tswift_GrammarParser.AND, 0); }
		public TerminalNode OR() { return getToken(Tswift_GrammarParser.OR, 0); }
		public CondContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cond; }
	}

	public final CondContext cond() throws RecognitionException {
		return cond(0);
	}

	private CondContext cond(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		CondContext _localctx = new CondContext(_ctx, _parentState);
		CondContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_cond, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(33);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,0,_ctx) ) {
			case 1:
				{
				setState(15);
				((CondContext)_localctx).op = match(NOT);
				setState(16);
				((CondContext)_localctx).c = cond(7);
				 
				                                        ((CondContext)_localctx).EV =  ((CondContext)_localctx).c.EF; 
				                                        ((CondContext)_localctx).EF =  ((CondContext)_localctx).c.EV; 
				                                    
				}
				break;
			case 2:
				{
				setState(19);
				((CondContext)_localctx).e1 = expr(0);
				setState(20);
				((CondContext)_localctx).opr = oprel();
				setState(21);
				((CondContext)_localctx).e2 = expr(0);
				 
				                                        ((CondContext)_localctx).EV =  append(_localctx.EV, gen.NewEtq());
				                                        ((CondContext)_localctx).EF =  append(_localctx.EF, gen.NewEtq());
				                                        var cad string;
				                                        cad = ((CondContext)_localctx).e1.dir + " " + ((CondContext)_localctx).opr.op + " " + ((CondContext)_localctx).e2.dir;
				                                        gen.Gen("if " + cad + " then goto " + _localctx.EV[0]);
				                                        gen.Gen("goto " + _localctx.EF[0]);

				                                    
				}
				break;
			case 3:
				{
				setState(24);
				match(PARIZQ);
				setState(25);
				((CondContext)_localctx).c = cond(0);
				setState(26);
				match(PARDER);
				 
				                                        ((CondContext)_localctx).EV =  ((CondContext)_localctx).c.EV; 
				                                        ((CondContext)_localctx).EF =  ((CondContext)_localctx).c.EF; 
				                                    
				}
				break;
			case 4:
				{
				setState(29);
				match(TRUE);
				 
				                                        ((CondContext)_localctx).EV =  append(_localctx.EV, gen.NewEtq());
				                                        ((CondContext)_localctx).EF =  append(_localctx.EF, gen.NewEtq());
				                                        gen.Gen("goto " + _localctx.EV[0]);
				                                    
				}
				break;
			case 5:
				{
				setState(31);
				match(FALSE);
				 
				                                        ((CondContext)_localctx).EV =  append(_localctx.EV, gen.NewEtq());
				                                        ((CondContext)_localctx).EF =  append(_localctx.EF, gen.NewEtq());
				                                        gen.Gen("goto " + _localctx.EF[0]);
				                                    
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(49);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(47);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new CondContext(_parentctx, _parentState);
						_localctx.c1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_cond);
						setState(35);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(36);
						((CondContext)_localctx).op = match(AND);
						setState(37);
						marcador(((CondContext)_localctx).c1.EV);
						setState(38);
						((CondContext)_localctx).c2 = cond(7);
						 
						                                                          ((CondContext)_localctx).EV =  ((CondContext)_localctx).c2.EV;
						                                                          ((CondContext)_localctx).EF =  gen.Unir(((CondContext)_localctx).c1.EF, ((CondContext)_localctx).c2.EF);
						                                                      
						}
						break;
					case 2:
						{
						_localctx = new CondContext(_parentctx, _parentState);
						_localctx.c1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_cond);
						setState(41);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(42);
						((CondContext)_localctx).op = match(OR);
						setState(43);
						marcador(((CondContext)_localctx).c1.EF);
						setState(44);
						((CondContext)_localctx).c2 = cond(6);
						 
						                                                          ((CondContext)_localctx).EV =  gen.Unir(((CondContext)_localctx).c1.EV, ((CondContext)_localctx).c2.EV);
						                                                          ((CondContext)_localctx).EF =  ((CondContext)_localctx).c2.EF;
						                                                      
						}
						break;
					}
					} 
				}
				setState(51);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MarcadorContext extends ParserRuleContext {
		public []string ES;
		public MarcadorContext(ParserRuleContext parent, int invokingState) { super(parent, invokingState); }
		public MarcadorContext(ParserRuleContext parent, int invokingState, []string ES) {
			super(parent, invokingState);
			this.ES = ES;
		}
		@Override public int getRuleIndex() { return RULE_marcador; }
	}

	public final MarcadorContext marcador([]string ES) throws RecognitionException {
		MarcadorContext _localctx = new MarcadorContext(_ctx, getState(), ES);
		enterRule(_localctx, 4, RULE_marcador);
		try {
			enterOuterAlt(_localctx, 1);
			{

			    gen.ImprimirEtq(ES);

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public string dir;
		public ExprContext e1;
		public Token op;
		public Token n;
		public Token id;
		public ExprContext e2;
		public TerminalNode MENOS() { return getToken(Tswift_GrammarParser.MENOS, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarParser.PARDER, 0); }
		public TerminalNode ENTERO() { return getToken(Tswift_GrammarParser.ENTERO, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarParser.ID, 0); }
		public TerminalNode POR() { return getToken(Tswift_GrammarParser.POR, 0); }
		public TerminalNode DIV() { return getToken(Tswift_GrammarParser.DIV, 0); }
		public TerminalNode MAS() { return getToken(Tswift_GrammarParser.MAS, 0); }
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(68);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MENOS:
				{
				setState(55);
				((ExprContext)_localctx).op = match(MENOS);
				setState(56);
				((ExprContext)_localctx).e1 = expr(6);

				                                        ((ExprContext)_localctx).dir =  gen.NewTemp();
				                                        gen.Gen(_localctx.dir + " = -1");
				                                        gen.Gen(_localctx.dir + " = " + _localctx.dir + " * " + ((ExprContext)_localctx).e1.dir);
				                                    
				}
				break;
			case PARIZQ:
				{
				setState(59);
				match(PARIZQ);
				setState(60);
				((ExprContext)_localctx).e1 = expr(0);
				setState(61);
				match(PARDER);
				((ExprContext)_localctx).dir =  ((ExprContext)_localctx).e1.dir;
				}
				break;
			case ENTERO:
				{
				setState(64);
				((ExprContext)_localctx).n = match(ENTERO);
				((ExprContext)_localctx).dir =  (((ExprContext)_localctx).n!=null?((ExprContext)_localctx).n.getText():null);
				}
				break;
			case ID:
				{
				setState(66);
				((ExprContext)_localctx).id = match(ID);
				((ExprContext)_localctx).dir =  (((ExprContext)_localctx).id!=null?((ExprContext)_localctx).id.getText():null);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(82);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(80);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
					case 1:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(70);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(71);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIV || _la==POR) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(72);
						((ExprContext)_localctx).e2 = expr(6);

						                                                  ((ExprContext)_localctx).dir =  gen.NewTemp();
						                                                  gen.Gen(_localctx.dir + " = " + ((ExprContext)_localctx).e1.dir + " " + (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null) + " " + ((ExprContext)_localctx).e2.dir);
						                                              
						}
						break;
					case 2:
						{
						_localctx = new ExprContext(_parentctx, _parentState);
						_localctx.e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(75);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(76);
						((ExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MENOS || _la==MAS) ) {
							((ExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(77);
						((ExprContext)_localctx).e2 = expr(5);

						                                                  ((ExprContext)_localctx).dir =  gen.NewTemp();
						                                                  gen.Gen(_localctx.dir + " = " + ((ExprContext)_localctx).e1.dir + " " + (((ExprContext)_localctx).op!=null?((ExprContext)_localctx).op.getText():null) + " " + ((ExprContext)_localctx).e2.dir);
						                                              
						}
						break;
					}
					} 
				}
				setState(84);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,5,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OprelContext extends ParserRuleContext {
		public string op;
		public Token ope;
		public TerminalNode IGUALIGUAL() { return getToken(Tswift_GrammarParser.IGUALIGUAL, 0); }
		public TerminalNode DIFERENTE() { return getToken(Tswift_GrammarParser.DIFERENTE, 0); }
		public TerminalNode MAYOR() { return getToken(Tswift_GrammarParser.MAYOR, 0); }
		public TerminalNode MENOR() { return getToken(Tswift_GrammarParser.MENOR, 0); }
		public TerminalNode MAYORIGUAL() { return getToken(Tswift_GrammarParser.MAYORIGUAL, 0); }
		public TerminalNode MENORIGUAL() { return getToken(Tswift_GrammarParser.MENORIGUAL, 0); }
		public OprelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_oprel; }
	}

	public final OprelContext oprel() throws RecognitionException {
		OprelContext _localctx = new OprelContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_oprel);
		try {
			setState(97);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IGUALIGUAL:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				((OprelContext)_localctx).ope = match(IGUALIGUAL);
				 ((OprelContext)_localctx).op =  (((OprelContext)_localctx).ope!=null?((OprelContext)_localctx).ope.getText():null) ;
				}
				break;
			case DIFERENTE:
				enterOuterAlt(_localctx, 2);
				{
				setState(87);
				((OprelContext)_localctx).ope = match(DIFERENTE);
				 ((OprelContext)_localctx).op =  (((OprelContext)_localctx).ope!=null?((OprelContext)_localctx).ope.getText():null) ;
				}
				break;
			case MAYOR:
				enterOuterAlt(_localctx, 3);
				{
				setState(89);
				((OprelContext)_localctx).ope = match(MAYOR);
				 ((OprelContext)_localctx).op =  (((OprelContext)_localctx).ope!=null?((OprelContext)_localctx).ope.getText():null); 
				}
				break;
			case MENOR:
				enterOuterAlt(_localctx, 4);
				{
				setState(91);
				((OprelContext)_localctx).ope = match(MENOR);
				 ((OprelContext)_localctx).op =  (((OprelContext)_localctx).ope!=null?((OprelContext)_localctx).ope.getText():null) ;
				}
				break;
			case MAYORIGUAL:
				enterOuterAlt(_localctx, 5);
				{
				setState(93);
				((OprelContext)_localctx).ope = match(MAYORIGUAL);
				 ((OprelContext)_localctx).op =  (((OprelContext)_localctx).ope!=null?((OprelContext)_localctx).ope.getText():null) ;
				}
				break;
			case MENORIGUAL:
				enterOuterAlt(_localctx, 6);
				{
				setState(95);
				((OprelContext)_localctx).ope = match(MENORIGUAL);
				 ((OprelContext)_localctx).op =  (((OprelContext)_localctx).ope!=null?((OprelContext)_localctx).ope.getText():null) ;
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return cond_sempred((CondContext)_localctx, predIndex);
		case 3:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean cond_sempred(CondContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001Md\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002\u0002"+
		"\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001\"\b\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0005\u00010\b\u0001\n\u0001\f\u00013\t\u0001\u0001\u0002\u0001\u0002"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0003\u0003E\b\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0005\u0003Q\b\u0003\n\u0003\f\u0003T\t\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004b\b\u0004\u0001\u0004\u0000\u0002\u0002\u0006\u0005\u0000"+
		"\u0002\u0004\u0006\b\u0000\u0002\u0002\u0000\u0011\u0011\u0015\u0015\u0001"+
		"\u0000\u0013\u0014n\u0000\n\u0001\u0000\u0000\u0000\u0002!\u0001\u0000"+
		"\u0000\u0000\u00044\u0001\u0000\u0000\u0000\u0006D\u0001\u0000\u0000\u0000"+
		"\ba\u0001\u0000\u0000\u0000\n\u000b\u0003\u0002\u0001\u0000\u000b\f\u0005"+
		"\u0000\u0000\u0001\f\r\u0006\u0000\uffff\uffff\u0000\r\u0001\u0001\u0000"+
		"\u0000\u0000\u000e\u000f\u0006\u0001\uffff\uffff\u0000\u000f\u0010\u0005"+
		"\u001e\u0000\u0000\u0010\u0011\u0003\u0002\u0001\u0007\u0011\u0012\u0006"+
		"\u0001\uffff\uffff\u0000\u0012\"\u0001\u0000\u0000\u0000\u0013\u0014\u0003"+
		"\u0006\u0003\u0000\u0014\u0015\u0003\b\u0004\u0000\u0015\u0016\u0003\u0006"+
		"\u0003\u0000\u0016\u0017\u0006\u0001\uffff\uffff\u0000\u0017\"\u0001\u0000"+
		"\u0000\u0000\u0018\u0019\u0005\u0002\u0000\u0000\u0019\u001a\u0003\u0002"+
		"\u0001\u0000\u001a\u001b\u0005\u0001\u0000\u0000\u001b\u001c\u0006\u0001"+
		"\uffff\uffff\u0000\u001c\"\u0001\u0000\u0000\u0000\u001d\u001e\u0005\""+
		"\u0000\u0000\u001e\"\u0006\u0001\uffff\uffff\u0000\u001f \u0005#\u0000"+
		"\u0000 \"\u0006\u0001\uffff\uffff\u0000!\u000e\u0001\u0000\u0000\u0000"+
		"!\u0013\u0001\u0000\u0000\u0000!\u0018\u0001\u0000\u0000\u0000!\u001d"+
		"\u0001\u0000\u0000\u0000!\u001f\u0001\u0000\u0000\u0000\"1\u0001\u0000"+
		"\u0000\u0000#$\n\u0006\u0000\u0000$%\u0005\u001c\u0000\u0000%&\u0003\u0004"+
		"\u0002\u0000&\'\u0003\u0002\u0001\u0007\'(\u0006\u0001\uffff\uffff\u0000"+
		"(0\u0001\u0000\u0000\u0000)*\n\u0005\u0000\u0000*+\u0005\u001d\u0000\u0000"+
		"+,\u0003\u0004\u0002\u0000,-\u0003\u0002\u0001\u0006-.\u0006\u0001\uffff"+
		"\uffff\u0000.0\u0001\u0000\u0000\u0000/#\u0001\u0000\u0000\u0000/)\u0001"+
		"\u0000\u0000\u000003\u0001\u0000\u0000\u00001/\u0001\u0000\u0000\u0000"+
		"12\u0001\u0000\u0000\u00002\u0003\u0001\u0000\u0000\u000031\u0001\u0000"+
		"\u0000\u000045\u0006\u0002\uffff\uffff\u00005\u0005\u0001\u0000\u0000"+
		"\u000067\u0006\u0003\uffff\uffff\u000078\u0005\u0013\u0000\u000089\u0003"+
		"\u0006\u0003\u00069:\u0006\u0003\uffff\uffff\u0000:E\u0001\u0000\u0000"+
		"\u0000;<\u0005\u0002\u0000\u0000<=\u0003\u0006\u0003\u0000=>\u0005\u0001"+
		"\u0000\u0000>?\u0006\u0003\uffff\uffff\u0000?E\u0001\u0000\u0000\u0000"+
		"@A\u0005F\u0000\u0000AE\u0006\u0003\uffff\uffff\u0000BC\u0005J\u0000\u0000"+
		"CE\u0006\u0003\uffff\uffff\u0000D6\u0001\u0000\u0000\u0000D;\u0001\u0000"+
		"\u0000\u0000D@\u0001\u0000\u0000\u0000DB\u0001\u0000\u0000\u0000ER\u0001"+
		"\u0000\u0000\u0000FG\n\u0005\u0000\u0000GH\u0007\u0000\u0000\u0000HI\u0003"+
		"\u0006\u0003\u0006IJ\u0006\u0003\uffff\uffff\u0000JQ\u0001\u0000\u0000"+
		"\u0000KL\n\u0004\u0000\u0000LM\u0007\u0001\u0000\u0000MN\u0003\u0006\u0003"+
		"\u0005NO\u0006\u0003\uffff\uffff\u0000OQ\u0001\u0000\u0000\u0000PF\u0001"+
		"\u0000\u0000\u0000PK\u0001\u0000\u0000\u0000QT\u0001\u0000\u0000\u0000"+
		"RP\u0001\u0000\u0000\u0000RS\u0001\u0000\u0000\u0000S\u0007\u0001\u0000"+
		"\u0000\u0000TR\u0001\u0000\u0000\u0000UV\u0005\u0016\u0000\u0000Vb\u0006"+
		"\u0004\uffff\uffff\u0000WX\u0005\u0017\u0000\u0000Xb\u0006\u0004\uffff"+
		"\uffff\u0000YZ\u0005\u001a\u0000\u0000Zb\u0006\u0004\uffff\uffff\u0000"+
		"[\\\u0005\u001b\u0000\u0000\\b\u0006\u0004\uffff\uffff\u0000]^\u0005\u0018"+
		"\u0000\u0000^b\u0006\u0004\uffff\uffff\u0000_`\u0005\u0019\u0000\u0000"+
		"`b\u0006\u0004\uffff\uffff\u0000aU\u0001\u0000\u0000\u0000aW\u0001\u0000"+
		"\u0000\u0000aY\u0001\u0000\u0000\u0000a[\u0001\u0000\u0000\u0000a]\u0001"+
		"\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000b\t\u0001\u0000\u0000\u0000"+
		"\u0007!/1DPRa";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}