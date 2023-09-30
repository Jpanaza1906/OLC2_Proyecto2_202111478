// Generated from Tswift_GrammarN.g4 by ANTLR 4.13.0
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class Tswift_GrammarNParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.0", RuntimeMetaData.VERSION); }

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
		RULE_s = 0, RULE_l_sentencias = 1, RULE_sentencia = 2, RULE_print_sentencia = 3, 
		RULE_condicion = 4, RULE_e = 5;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "print_sentencia", "condicion", "e"
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
	public String getGrammarFileName() { return "Tswift_GrammarN.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Tswift_GrammarNParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SContext extends ParserRuleContext {
		public SContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_s; }
	 
		public SContext() { }
		public void copyFrom(SContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SLSentenciasContext extends SContext {
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public SLSentenciasContext(SContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterSLSentencias(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitSLSentencias(this);
		}
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			_localctx = new SLSentenciasContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(12);
			l_sentencias();
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
	public static class L_sentenciasContext extends ParserRuleContext {
		public L_sentenciasContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l_sentencias; }
	 
		public L_sentenciasContext() { }
		public void copyFrom(L_sentenciasContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class L_SentenciaContext extends L_sentenciasContext {
		public List<SentenciaContext> sentencia() {
			return getRuleContexts(SentenciaContext.class);
		}
		public SentenciaContext sentencia(int i) {
			return getRuleContext(SentenciaContext.class,i);
		}
		public L_SentenciaContext(L_sentenciasContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterL_Sentencia(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitL_Sentencia(this);
		}
	}

	public final L_sentenciasContext l_sentencias() throws RecognitionException {
		L_sentenciasContext _localctx = new L_sentenciasContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_l_sentencias);
		int _la;
		try {
			_localctx = new L_SentenciaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(17);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 54761357316L) != 0) || _la==ENTERO || _la==ID) {
				{
				{
				setState(14);
				sentencia();
				}
				}
				setState(19);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
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
	public static class SentenciaContext extends ParserRuleContext {
		public SentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentencia; }
	 
		public SentenciaContext() { }
		public void copyFrom(SentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_eContext extends SentenciaContext {
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_eContext(SentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_e(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_e(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_PrintContext extends SentenciaContext {
		public Print_sentenciaContext print_sentencia() {
			return getRuleContext(Print_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_PrintContext(SentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_Print(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_Print(this);
		}
	}

	public final SentenciaContext sentencia() throws RecognitionException {
		SentenciaContext _localctx = new SentenciaContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sentencia);
		int _la;
		try {
			setState(28);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new S_PrintContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(20);
				print_sentencia();
				setState(22);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(21);
					match(PTCOMA);
					}
				}

				}
				break;
			case PARIZQ:
			case MENOS:
			case NOT:
			case TRUE:
			case FALSE:
			case ENTERO:
			case ID:
				_localctx = new S_eContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(24);
				condicion(0);
				setState(26);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(25);
					match(PTCOMA);
					}
				}

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

	@SuppressWarnings("CheckReturnValue")
	public static class Print_sentenciaContext extends ParserRuleContext {
		public Print_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_print_sentencia; }
	 
		public Print_sentenciaContext() { }
		public void copyFrom(Print_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintContext extends Print_sentenciaContext {
		public TerminalNode PRINT() { return getToken(Tswift_GrammarNParser.PRINT, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public List<TerminalNode> COMA() { return getTokens(Tswift_GrammarNParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(Tswift_GrammarNParser.COMA, i);
		}
		public PrintContext(Print_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterPrint(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitPrint(this);
		}
	}

	public final Print_sentenciaContext print_sentencia() throws RecognitionException {
		Print_sentenciaContext _localctx = new Print_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_print_sentencia);
		int _la;
		try {
			_localctx = new PrintContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
			match(PRINT);
			setState(31);
			match(PARIZQ);
			setState(32);
			e(0);
			setState(37);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(33);
				match(COMA);
				setState(34);
				e(0);
				}
				}
				setState(39);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(40);
			match(PARDER);
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
	public static class CondicionContext extends ParserRuleContext {
		public CondicionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_condicion; }
	 
		public CondicionContext() { }
		public void copyFrom(CondicionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Cond_ParContext extends CondicionContext {
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public Cond_ParContext(CondicionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterCond_Par(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitCond_Par(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Cond_NegContext extends CondicionContext {
		public Token op;
		public CondicionContext c;
		public TerminalNode NOT() { return getToken(Tswift_GrammarNParser.NOT, 0); }
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public Cond_NegContext(CondicionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterCond_Neg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitCond_Neg(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Cond_RelContext extends CondicionContext {
		public Token op;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode IGUALIGUAL() { return getToken(Tswift_GrammarNParser.IGUALIGUAL, 0); }
		public TerminalNode DIFERENTE() { return getToken(Tswift_GrammarNParser.DIFERENTE, 0); }
		public TerminalNode MAYORIGUAL() { return getToken(Tswift_GrammarNParser.MAYORIGUAL, 0); }
		public TerminalNode MAYOR() { return getToken(Tswift_GrammarNParser.MAYOR, 0); }
		public TerminalNode MENORIGUAL() { return getToken(Tswift_GrammarNParser.MENORIGUAL, 0); }
		public TerminalNode MENOR() { return getToken(Tswift_GrammarNParser.MENOR, 0); }
		public Cond_RelContext(CondicionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterCond_Rel(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitCond_Rel(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Cond_BooleanoContext extends CondicionContext {
		public Token op;
		public TerminalNode TRUE() { return getToken(Tswift_GrammarNParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Tswift_GrammarNParser.FALSE, 0); }
		public Cond_BooleanoContext(CondicionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterCond_Booleano(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitCond_Booleano(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Cond_LogicaContext extends CondicionContext {
		public Token op;
		public List<CondicionContext> condicion() {
			return getRuleContexts(CondicionContext.class);
		}
		public CondicionContext condicion(int i) {
			return getRuleContext(CondicionContext.class,i);
		}
		public TerminalNode AND() { return getToken(Tswift_GrammarNParser.AND, 0); }
		public TerminalNode OR() { return getToken(Tswift_GrammarNParser.OR, 0); }
		public Cond_LogicaContext(CondicionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterCond_Logica(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitCond_Logica(this);
		}
	}

	public final CondicionContext condicion() throws RecognitionException {
		return condicion(0);
	}

	private CondicionContext condicion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		CondicionContext _localctx = new CondicionContext(_ctx, _parentState);
		CondicionContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_condicion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				{
				_localctx = new Cond_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(43);
				((Cond_NegContext)_localctx).op = match(NOT);
				setState(44);
				((Cond_NegContext)_localctx).c = condicion(5);
				}
				break;
			case 2:
				{
				_localctx = new Cond_RelContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(45);
				e(0);
				setState(46);
				((Cond_RelContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 264241152L) != 0)) ) {
					((Cond_RelContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(47);
				e(0);
				}
				break;
			case 3:
				{
				_localctx = new Cond_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(49);
				((Cond_BooleanoContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==TRUE || _la==FALSE) ) {
					((Cond_BooleanoContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 4:
				{
				_localctx = new Cond_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(50);
				match(PARIZQ);
				setState(51);
				condicion(0);
				setState(52);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(61);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Cond_LogicaContext(new CondicionContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_condicion);
					setState(56);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(57);
					((Cond_LogicaContext)_localctx).op = _input.LT(1);
					_la = _input.LA(1);
					if ( !(_la==AND || _la==OR) ) {
						((Cond_LogicaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(58);
					condicion(4);
					}
					} 
				}
				setState(63);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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
	public static class EContext extends ParserRuleContext {
		public EContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_e; }
	 
		public EContext() { }
		public void copyFrom(EContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_SumResContext extends EContext {
		public Token op;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode MAS() { return getToken(Tswift_GrammarNParser.MAS, 0); }
		public TerminalNode MENOS() { return getToken(Tswift_GrammarNParser.MENOS, 0); }
		public Expr_SumResContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_SumRes(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_SumRes(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_IdContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public Expr_IdContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Id(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Id(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_ModContext extends EContext {
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode MOD() { return getToken(Tswift_GrammarNParser.MOD, 0); }
		public Expr_ModContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Mod(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Mod(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_NegContext extends EContext {
		public Token op;
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode MENOS() { return getToken(Tswift_GrammarNParser.MENOS, 0); }
		public Expr_NegContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Neg(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Neg(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_MulDivContext extends EContext {
		public Token op;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode POR() { return getToken(Tswift_GrammarNParser.POR, 0); }
		public TerminalNode DIV() { return getToken(Tswift_GrammarNParser.DIV, 0); }
		public Expr_MulDivContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_MulDiv(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_MulDiv(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_ParContext extends EContext {
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public Expr_ParContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Par(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Par(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_EnteroContext extends EContext {
		public TerminalNode ENTERO() { return getToken(Tswift_GrammarNParser.ENTERO, 0); }
		public Expr_EnteroContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Entero(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Entero(this);
		}
	}

	public final EContext e() throws RecognitionException {
		return e(0);
	}

	private EContext e(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		EContext _localctx = new EContext(_ctx, _parentState);
		EContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MENOS:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(65);
				((Expr_NegContext)_localctx).op = match(MENOS);
				setState(66);
				e(7);
				}
				break;
			case ID:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(67);
				match(ID);
				}
				break;
			case ENTERO:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(68);
				match(ENTERO);
				}
				break;
			case PARIZQ:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(69);
				match(PARIZQ);
				setState(70);
				e(0);
				setState(71);
				match(PARDER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(86);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(84);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(75);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(76);
						((Expr_MulDivContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==DIV || _la==POR) ) {
							((Expr_MulDivContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(77);
						e(7);
						}
						break;
					case 2:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(78);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(79);
						((Expr_SumResContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==MENOS || _la==MAS) ) {
							((Expr_SumResContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(80);
						e(6);
						}
						break;
					case 3:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(81);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(82);
						match(MOD);
						setState(83);
						e(5);
						}
						break;
					}
					} 
				}
				setState(88);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 4:
			return condicion_sempred((CondicionContext)_localctx, predIndex);
		case 5:
			return e_sempred((EContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean condicion_sempred(CondicionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean e_sempred(EContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 6);
		case 2:
			return precpred(_ctx, 5);
		case 3:
			return precpred(_ctx, 4);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001MZ\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002\u0002"+
		"\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002\u0005"+
		"\u0007\u0005\u0001\u0000\u0001\u0000\u0001\u0001\u0005\u0001\u0010\b\u0001"+
		"\n\u0001\f\u0001\u0013\t\u0001\u0001\u0002\u0001\u0002\u0003\u0002\u0017"+
		"\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u001b\b\u0002\u0003\u0002"+
		"\u001d\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0005\u0003$\b\u0003\n\u0003\f\u0003\'\t\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u00047\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004"+
		"<\b\u0004\n\u0004\f\u0004?\t\u0004\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0003\u0005J\b\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005"+
		"U\b\u0005\n\u0005\f\u0005X\t\u0005\u0001\u0005\u0000\u0002\b\n\u0006\u0000"+
		"\u0002\u0004\u0006\b\n\u0000\u0005\u0001\u0000\u0016\u001b\u0001\u0000"+
		"\"#\u0001\u0000\u001c\u001d\u0002\u0000\u0011\u0011\u0015\u0015\u0001"+
		"\u0000\u0013\u0014b\u0000\f\u0001\u0000\u0000\u0000\u0002\u0011\u0001"+
		"\u0000\u0000\u0000\u0004\u001c\u0001\u0000\u0000\u0000\u0006\u001e\u0001"+
		"\u0000\u0000\u0000\b6\u0001\u0000\u0000\u0000\nI\u0001\u0000\u0000\u0000"+
		"\f\r\u0003\u0002\u0001\u0000\r\u0001\u0001\u0000\u0000\u0000\u000e\u0010"+
		"\u0003\u0004\u0002\u0000\u000f\u000e\u0001\u0000\u0000\u0000\u0010\u0013"+
		"\u0001\u0000\u0000\u0000\u0011\u000f\u0001\u0000\u0000\u0000\u0011\u0012"+
		"\u0001\u0000\u0000\u0000\u0012\u0003\u0001\u0000\u0000\u0000\u0013\u0011"+
		"\u0001\u0000\u0000\u0000\u0014\u0016\u0003\u0006\u0003\u0000\u0015\u0017"+
		"\u0005\b\u0000\u0000\u0016\u0015\u0001\u0000\u0000\u0000\u0016\u0017\u0001"+
		"\u0000\u0000\u0000\u0017\u001d\u0001\u0000\u0000\u0000\u0018\u001a\u0003"+
		"\b\u0004\u0000\u0019\u001b\u0005\b\u0000\u0000\u001a\u0019\u0001\u0000"+
		"\u0000\u0000\u001a\u001b\u0001\u0000\u0000\u0000\u001b\u001d\u0001\u0000"+
		"\u0000\u0000\u001c\u0014\u0001\u0000\u0000\u0000\u001c\u0018\u0001\u0000"+
		"\u0000\u0000\u001d\u0005\u0001\u0000\u0000\u0000\u001e\u001f\u0005\u001f"+
		"\u0000\u0000\u001f \u0005\u0002\u0000\u0000 %\u0003\n\u0005\u0000!\"\u0005"+
		"\n\u0000\u0000\"$\u0003\n\u0005\u0000#!\u0001\u0000\u0000\u0000$\'\u0001"+
		"\u0000\u0000\u0000%#\u0001\u0000\u0000\u0000%&\u0001\u0000\u0000\u0000"+
		"&(\u0001\u0000\u0000\u0000\'%\u0001\u0000\u0000\u0000()\u0005\u0001\u0000"+
		"\u0000)\u0007\u0001\u0000\u0000\u0000*+\u0006\u0004\uffff\uffff\u0000"+
		"+,\u0005\u001e\u0000\u0000,7\u0003\b\u0004\u0005-.\u0003\n\u0005\u0000"+
		"./\u0007\u0000\u0000\u0000/0\u0003\n\u0005\u000007\u0001\u0000\u0000\u0000"+
		"17\u0007\u0001\u0000\u000023\u0005\u0002\u0000\u000034\u0003\b\u0004\u0000"+
		"45\u0005\u0001\u0000\u000057\u0001\u0000\u0000\u00006*\u0001\u0000\u0000"+
		"\u00006-\u0001\u0000\u0000\u000061\u0001\u0000\u0000\u000062\u0001\u0000"+
		"\u0000\u00007=\u0001\u0000\u0000\u000089\n\u0003\u0000\u00009:\u0007\u0002"+
		"\u0000\u0000:<\u0003\b\u0004\u0004;8\u0001\u0000\u0000\u0000<?\u0001\u0000"+
		"\u0000\u0000=;\u0001\u0000\u0000\u0000=>\u0001\u0000\u0000\u0000>\t\u0001"+
		"\u0000\u0000\u0000?=\u0001\u0000\u0000\u0000@A\u0006\u0005\uffff\uffff"+
		"\u0000AB\u0005\u0013\u0000\u0000BJ\u0003\n\u0005\u0007CJ\u0005J\u0000"+
		"\u0000DJ\u0005F\u0000\u0000EF\u0005\u0002\u0000\u0000FG\u0003\n\u0005"+
		"\u0000GH\u0005\u0001\u0000\u0000HJ\u0001\u0000\u0000\u0000I@\u0001\u0000"+
		"\u0000\u0000IC\u0001\u0000\u0000\u0000ID\u0001\u0000\u0000\u0000IE\u0001"+
		"\u0000\u0000\u0000JV\u0001\u0000\u0000\u0000KL\n\u0006\u0000\u0000LM\u0007"+
		"\u0003\u0000\u0000MU\u0003\n\u0005\u0007NO\n\u0005\u0000\u0000OP\u0007"+
		"\u0004\u0000\u0000PU\u0003\n\u0005\u0006QR\n\u0004\u0000\u0000RS\u0005"+
		"\u0012\u0000\u0000SU\u0003\n\u0005\u0005TK\u0001\u0000\u0000\u0000TN\u0001"+
		"\u0000\u0000\u0000TQ\u0001\u0000\u0000\u0000UX\u0001\u0000\u0000\u0000"+
		"VT\u0001\u0000\u0000\u0000VW\u0001\u0000\u0000\u0000W\u000b\u0001\u0000"+
		"\u0000\u0000XV\u0001\u0000\u0000\u0000\n\u0011\u0016\u001a\u001c%6=IT"+
		"V";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}