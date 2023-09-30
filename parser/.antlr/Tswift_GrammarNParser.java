// Generated from /home/josep-ubu/Lab_Compiladores2/OLC2_Proyecto2_202111478/parser/Tswift_GrammarN.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Tswift_GrammarNParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

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
		RULE_if_sentencia = 4, RULE_condicion = 5, RULE_e = 6;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "print_sentencia", "if_sentencia", 
			"condicion", "e"
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
	public static class SLSentenciasContext extends SContext {
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public SLSentenciasContext(SContext ctx) { copyFrom(ctx); }
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			_localctx = new SLSentenciasContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(14);
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
	public static class L_SentenciaContext extends L_sentenciasContext {
		public List<SentenciaContext> sentencia() {
			return getRuleContexts(SentenciaContext.class);
		}
		public SentenciaContext sentencia(int i) {
			return getRuleContext(SentenciaContext.class,i);
		}
		public L_SentenciaContext(L_sentenciasContext ctx) { copyFrom(ctx); }
	}

	public final L_sentenciasContext l_sentencias() throws RecognitionException {
		L_sentenciasContext _localctx = new L_sentenciasContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_l_sentencias);
		int _la;
		try {
			_localctx = new L_SentenciaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(19);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PARIZQ) | (1L << MENOS) | (1L << NOT) | (1L << PRINT) | (1L << TRUE) | (1L << FALSE) | (1L << IF))) != 0) || _la==ENTERO || _la==ID) {
				{
				{
				setState(16);
				sentencia();
				}
				}
				setState(21);
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
	public static class S_CondicionContext extends SentenciaContext {
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_CondicionContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_IfContext extends SentenciaContext {
		public If_sentenciaContext if_sentencia() {
			return getRuleContext(If_sentenciaContext.class,0);
		}
		public S_IfContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	public static class S_PrintContext extends SentenciaContext {
		public Print_sentenciaContext print_sentencia() {
			return getRuleContext(Print_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_PrintContext(SentenciaContext ctx) { copyFrom(ctx); }
	}

	public final SentenciaContext sentencia() throws RecognitionException {
		SentenciaContext _localctx = new SentenciaContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sentencia);
		int _la;
		try {
			setState(31);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new S_PrintContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(22);
				print_sentencia();
				setState(24);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(23);
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
				_localctx = new S_CondicionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(26);
				condicion(0);
				setState(28);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(27);
					match(PTCOMA);
					}
				}

				}
				break;
			case IF:
				_localctx = new S_IfContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(30);
				if_sentencia();
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
	}

	public final Print_sentenciaContext print_sentencia() throws RecognitionException {
		Print_sentenciaContext _localctx = new Print_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_print_sentencia);
		int _la;
		try {
			_localctx = new PrintContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(33);
			match(PRINT);
			setState(34);
			match(PARIZQ);
			setState(35);
			e(0);
			setState(40);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(36);
				match(COMA);
				setState(37);
				e(0);
				}
				}
				setState(42);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(43);
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

	public static class If_sentenciaContext extends ParserRuleContext {
		public If_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_sentencia; }
	 
		public If_sentenciaContext() { }
		public void copyFrom(If_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class IfContext extends If_sentenciaContext {
		public TerminalNode IF() { return getToken(Tswift_GrammarNParser.IF, 0); }
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public List<TerminalNode> LLAVEIZQ() { return getTokens(Tswift_GrammarNParser.LLAVEIZQ); }
		public TerminalNode LLAVEIZQ(int i) {
			return getToken(Tswift_GrammarNParser.LLAVEIZQ, i);
		}
		public List<L_sentenciasContext> l_sentencias() {
			return getRuleContexts(L_sentenciasContext.class);
		}
		public L_sentenciasContext l_sentencias(int i) {
			return getRuleContext(L_sentenciasContext.class,i);
		}
		public List<TerminalNode> LLAVEDER() { return getTokens(Tswift_GrammarNParser.LLAVEDER); }
		public TerminalNode LLAVEDER(int i) {
			return getToken(Tswift_GrammarNParser.LLAVEDER, i);
		}
		public TerminalNode ELSE() { return getToken(Tswift_GrammarNParser.ELSE, 0); }
		public If_sentenciaContext if_sentencia() {
			return getRuleContext(If_sentenciaContext.class,0);
		}
		public IfContext(If_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final If_sentenciaContext if_sentencia() throws RecognitionException {
		If_sentenciaContext _localctx = new If_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_if_sentencia);
		int _la;
		try {
			_localctx = new IfContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(45);
			match(IF);
			setState(46);
			condicion(0);
			setState(47);
			match(LLAVEIZQ);
			setState(48);
			l_sentencias();
			setState(49);
			match(LLAVEDER);
			setState(58);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(50);
				match(ELSE);
				setState(56);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case IF:
					{
					setState(51);
					if_sentencia();
					}
					break;
				case LLAVEIZQ:
					{
					setState(52);
					match(LLAVEIZQ);
					setState(53);
					l_sentencias();
					setState(54);
					match(LLAVEDER);
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
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
	public static class Cond_ParContext extends CondicionContext {
		public CondicionContext c;
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public Cond_ParContext(CondicionContext ctx) { copyFrom(ctx); }
	}
	public static class Cond_NegContext extends CondicionContext {
		public Token op;
		public CondicionContext c;
		public TerminalNode NOT() { return getToken(Tswift_GrammarNParser.NOT, 0); }
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public Cond_NegContext(CondicionContext ctx) { copyFrom(ctx); }
	}
	public static class Cond_RelContext extends CondicionContext {
		public EContext e1;
		public Token op;
		public EContext e2;
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
	}
	public static class Cond_BooleanoContext extends CondicionContext {
		public Token op;
		public TerminalNode TRUE() { return getToken(Tswift_GrammarNParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Tswift_GrammarNParser.FALSE, 0); }
		public Cond_BooleanoContext(CondicionContext ctx) { copyFrom(ctx); }
	}
	public static class Cond_LogicaContext extends CondicionContext {
		public CondicionContext c1;
		public Token op;
		public CondicionContext c2;
		public List<CondicionContext> condicion() {
			return getRuleContexts(CondicionContext.class);
		}
		public CondicionContext condicion(int i) {
			return getRuleContext(CondicionContext.class,i);
		}
		public TerminalNode AND() { return getToken(Tswift_GrammarNParser.AND, 0); }
		public TerminalNode OR() { return getToken(Tswift_GrammarNParser.OR, 0); }
		public Cond_LogicaContext(CondicionContext ctx) { copyFrom(ctx); }
	}

	public final CondicionContext condicion() throws RecognitionException {
		return condicion(0);
	}

	private CondicionContext condicion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		CondicionContext _localctx = new CondicionContext(_ctx, _parentState);
		CondicionContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_condicion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				_localctx = new Cond_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(61);
				((Cond_NegContext)_localctx).op = match(NOT);
				setState(62);
				((Cond_NegContext)_localctx).c = condicion(5);
				}
				break;
			case 2:
				{
				_localctx = new Cond_RelContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(63);
				((Cond_RelContext)_localctx).e1 = e(0);
				setState(64);
				((Cond_RelContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IGUALIGUAL) | (1L << DIFERENTE) | (1L << MAYORIGUAL) | (1L << MENORIGUAL) | (1L << MAYOR) | (1L << MENOR))) != 0)) ) {
					((Cond_RelContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(65);
				((Cond_RelContext)_localctx).e2 = e(0);
				}
				break;
			case 3:
				{
				_localctx = new Cond_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(67);
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
				setState(68);
				match(PARIZQ);
				setState(69);
				((Cond_ParContext)_localctx).c = condicion(0);
				setState(70);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(79);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Cond_LogicaContext(new CondicionContext(_parentctx, _parentState));
					((Cond_LogicaContext)_localctx).c1 = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_condicion);
					setState(74);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(75);
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
					setState(76);
					((Cond_LogicaContext)_localctx).c2 = condicion(4);
					}
					} 
				}
				setState(81);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
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
	public static class Expr_SumResContext extends EContext {
		public EContext e1;
		public Token op;
		public EContext e2;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode MAS() { return getToken(Tswift_GrammarNParser.MAS, 0); }
		public TerminalNode MENOS() { return getToken(Tswift_GrammarNParser.MENOS, 0); }
		public Expr_SumResContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_IdContext extends EContext {
		public Token id;
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public Expr_IdContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_ModContext extends EContext {
		public EContext e1;
		public Token op;
		public EContext e2;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode MOD() { return getToken(Tswift_GrammarNParser.MOD, 0); }
		public Expr_ModContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_NegContext extends EContext {
		public Token op;
		public EContext e1;
		public TerminalNode MENOS() { return getToken(Tswift_GrammarNParser.MENOS, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public Expr_NegContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_MulDivContext extends EContext {
		public EContext e1;
		public Token op;
		public EContext e2;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode POR() { return getToken(Tswift_GrammarNParser.POR, 0); }
		public TerminalNode DIV() { return getToken(Tswift_GrammarNParser.DIV, 0); }
		public Expr_MulDivContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_ParContext extends EContext {
		public EContext e1;
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public Expr_ParContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class Expr_EnteroContext extends EContext {
		public Token n;
		public TerminalNode ENTERO() { return getToken(Tswift_GrammarNParser.ENTERO, 0); }
		public Expr_EnteroContext(EContext ctx) { copyFrom(ctx); }
	}

	public final EContext e() throws RecognitionException {
		return e(0);
	}

	private EContext e(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		EContext _localctx = new EContext(_ctx, _parentState);
		EContext _prevctx = _localctx;
		int _startState = 12;
		enterRecursionRule(_localctx, 12, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(91);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MENOS:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(83);
				((Expr_NegContext)_localctx).op = match(MENOS);
				setState(84);
				((Expr_NegContext)_localctx).e1 = e(7);
				}
				break;
			case ID:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(85);
				((Expr_IdContext)_localctx).id = match(ID);
				}
				break;
			case ENTERO:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(86);
				((Expr_EnteroContext)_localctx).n = match(ENTERO);
				}
				break;
			case PARIZQ:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(87);
				match(PARIZQ);
				setState(88);
				((Expr_ParContext)_localctx).e1 = e(0);
				setState(89);
				match(PARDER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(104);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(102);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						((Expr_MulDivContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(93);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(94);
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
						setState(95);
						((Expr_MulDivContext)_localctx).e2 = e(7);
						}
						break;
					case 2:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						((Expr_SumResContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(96);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(97);
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
						setState(98);
						((Expr_SumResContext)_localctx).e2 = e(6);
						}
						break;
					case 3:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						((Expr_ModContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(99);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(100);
						((Expr_ModContext)_localctx).op = match(MOD);
						setState(101);
						((Expr_ModContext)_localctx).e2 = e(5);
						}
						break;
					}
					} 
				}
				setState(106);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
		case 5:
			return condicion_sempred((CondicionContext)_localctx, predIndex);
		case 6:
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3On\4\2\t\2\4\3\t\3"+
		"\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\3\2\3\2\3\3\7\3\24\n\3\f\3\16"+
		"\3\27\13\3\3\4\3\4\5\4\33\n\4\3\4\3\4\5\4\37\n\4\3\4\5\4\"\n\4\3\5\3\5"+
		"\3\5\3\5\3\5\7\5)\n\5\f\5\16\5,\13\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\5\6;\n\6\5\6=\n\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\7\5\7K\n\7\3\7\3\7\3\7\7\7P\n\7\f\7\16\7S\13\7\3\b\3\b\3"+
		"\b\3\b\3\b\3\b\3\b\3\b\3\b\5\b^\n\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3"+
		"\b\7\bi\n\b\f\b\16\bl\13\b\3\b\2\4\f\16\t\2\4\6\b\n\f\16\2\7\3\2\30\35"+
		"\3\2$%\3\2\36\37\4\2\23\23\27\27\3\2\25\26\2x\2\20\3\2\2\2\4\25\3\2\2"+
		"\2\6!\3\2\2\2\b#\3\2\2\2\n/\3\2\2\2\fJ\3\2\2\2\16]\3\2\2\2\20\21\5\4\3"+
		"\2\21\3\3\2\2\2\22\24\5\6\4\2\23\22\3\2\2\2\24\27\3\2\2\2\25\23\3\2\2"+
		"\2\25\26\3\2\2\2\26\5\3\2\2\2\27\25\3\2\2\2\30\32\5\b\5\2\31\33\7\n\2"+
		"\2\32\31\3\2\2\2\32\33\3\2\2\2\33\"\3\2\2\2\34\36\5\f\7\2\35\37\7\n\2"+
		"\2\36\35\3\2\2\2\36\37\3\2\2\2\37\"\3\2\2\2 \"\5\n\6\2!\30\3\2\2\2!\34"+
		"\3\2\2\2! \3\2\2\2\"\7\3\2\2\2#$\7!\2\2$%\7\4\2\2%*\5\16\b\2&\'\7\f\2"+
		"\2\')\5\16\b\2(&\3\2\2\2),\3\2\2\2*(\3\2\2\2*+\3\2\2\2+-\3\2\2\2,*\3\2"+
		"\2\2-.\7\3\2\2.\t\3\2\2\2/\60\7&\2\2\60\61\5\f\7\2\61\62\7\5\2\2\62\63"+
		"\5\4\3\2\63<\7\6\2\2\64:\7\'\2\2\65;\5\n\6\2\66\67\7\5\2\2\678\5\4\3\2"+
		"89\7\6\2\29;\3\2\2\2:\65\3\2\2\2:\66\3\2\2\2;=\3\2\2\2<\64\3\2\2\2<=\3"+
		"\2\2\2=\13\3\2\2\2>?\b\7\1\2?@\7 \2\2@K\5\f\7\7AB\5\16\b\2BC\t\2\2\2C"+
		"D\5\16\b\2DK\3\2\2\2EK\t\3\2\2FG\7\4\2\2GH\5\f\7\2HI\7\3\2\2IK\3\2\2\2"+
		"J>\3\2\2\2JA\3\2\2\2JE\3\2\2\2JF\3\2\2\2KQ\3\2\2\2LM\f\5\2\2MN\t\4\2\2"+
		"NP\5\f\7\6OL\3\2\2\2PS\3\2\2\2QO\3\2\2\2QR\3\2\2\2R\r\3\2\2\2SQ\3\2\2"+
		"\2TU\b\b\1\2UV\7\25\2\2V^\5\16\b\tW^\7L\2\2X^\7H\2\2YZ\7\4\2\2Z[\5\16"+
		"\b\2[\\\7\3\2\2\\^\3\2\2\2]T\3\2\2\2]W\3\2\2\2]X\3\2\2\2]Y\3\2\2\2^j\3"+
		"\2\2\2_`\f\b\2\2`a\t\5\2\2ai\5\16\b\tbc\f\7\2\2cd\t\6\2\2di\5\16\b\be"+
		"f\f\6\2\2fg\7\24\2\2gi\5\16\b\7h_\3\2\2\2hb\3\2\2\2he\3\2\2\2il\3\2\2"+
		"\2jh\3\2\2\2jk\3\2\2\2k\17\3\2\2\2lj\3\2\2\2\16\25\32\36!*:<JQ]hj";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}