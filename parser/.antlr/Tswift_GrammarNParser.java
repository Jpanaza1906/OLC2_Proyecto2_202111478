// Generated from /home/josep/USAC/6to_Semestre/Lab_Compi/Proyecto2/OLC2_Proyecto2_202111478/parser/Tswift_GrammarN.g4 by ANTLR 4.13.1
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
		RULE_s = 0, RULE_l_sentencias = 1, RULE_sentencia = 2, RULE_trans_sentencia = 3, 
		RULE_print_sentencia = 4, RULE_if_sentencia = 5, RULE_guard_sentencia = 6, 
		RULE_switch_sentencia = 7, RULE_lcasos = 8, RULE_cdefault = 9, RULE_while_sentencia = 10, 
		RULE_for_sentencia = 11, RULE_rango_p = 12, RULE_condicion = 13, RULE_e = 14;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "trans_sentencia", "print_sentencia", 
			"if_sentencia", "guard_sentencia", "switch_sentencia", "lcasos", "cdefault", 
			"while_sentencia", "for_sentencia", "rango_p", "condicion", "e"
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
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			_localctx = new SLSentenciasContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(30);
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
	}

	public final L_sentenciasContext l_sentencias() throws RecognitionException {
		L_sentenciasContext _localctx = new L_sentenciasContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_l_sentencias);
		try {
			int _alt;
			_localctx = new L_SentenciaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(35);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(32);
					sentencia();
					}
					} 
				}
				setState(37);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
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
	public static class S_ForContext extends SentenciaContext {
		public For_sentenciaContext for_sentencia() {
			return getRuleContext(For_sentenciaContext.class,0);
		}
		public S_ForContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_SwitchContext extends SentenciaContext {
		public Switch_sentenciaContext switch_sentencia() {
			return getRuleContext(Switch_sentenciaContext.class,0);
		}
		public S_SwitchContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_GuardContext extends SentenciaContext {
		public Guard_sentenciaContext guard_sentencia() {
			return getRuleContext(Guard_sentenciaContext.class,0);
		}
		public S_GuardContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_WhileContext extends SentenciaContext {
		public While_sentenciaContext while_sentencia() {
			return getRuleContext(While_sentenciaContext.class,0);
		}
		public S_WhileContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_TransContext extends SentenciaContext {
		public Trans_sentenciaContext trans_sentencia() {
			return getRuleContext(Trans_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_TransContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_IfContext extends SentenciaContext {
		public If_sentenciaContext if_sentencia() {
			return getRuleContext(If_sentenciaContext.class,0);
		}
		public S_IfContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
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
			setState(51);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new S_PrintContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(38);
				print_sentencia();
				setState(40);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(39);
					match(PTCOMA);
					}
				}

				}
				break;
			case IF:
				_localctx = new S_IfContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(42);
				if_sentencia();
				}
				break;
			case SWITCH:
				_localctx = new S_SwitchContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(43);
				switch_sentencia();
				}
				break;
			case GUARD:
				_localctx = new S_GuardContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(44);
				guard_sentencia();
				}
				break;
			case CONTINUE:
			case RETURN:
			case BREAK:
				_localctx = new S_TransContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(45);
				trans_sentencia();
				setState(47);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(46);
					match(PTCOMA);
					}
				}

				}
				break;
			case WHILE:
				_localctx = new S_WhileContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(49);
				while_sentencia();
				}
				break;
			case FOR:
				_localctx = new S_ForContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(50);
				for_sentencia();
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
	public static class Trans_sentenciaContext extends ParserRuleContext {
		public Trans_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_trans_sentencia; }
	 
		public Trans_sentenciaContext() { }
		public void copyFrom(Trans_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnContext extends Trans_sentenciaContext {
		public TerminalNode RETURN() { return getToken(Tswift_GrammarNParser.RETURN, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public ReturnContext(Trans_sentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakContext extends Trans_sentenciaContext {
		public TerminalNode BREAK() { return getToken(Tswift_GrammarNParser.BREAK, 0); }
		public BreakContext(Trans_sentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueContext extends Trans_sentenciaContext {
		public TerminalNode CONTINUE() { return getToken(Tswift_GrammarNParser.CONTINUE, 0); }
		public ContinueContext(Trans_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final Trans_sentenciaContext trans_sentencia() throws RecognitionException {
		Trans_sentenciaContext _localctx = new Trans_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_trans_sentencia);
		int _la;
		try {
			setState(59);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BREAK:
				_localctx = new BreakContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(53);
				match(BREAK);
				}
				break;
			case CONTINUE:
				_localctx = new ContinueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(54);
				match(CONTINUE);
				}
				break;
			case RETURN:
				_localctx = new ReturnContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(55);
				match(RETURN);
				setState(57);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PARIZQ || _la==MENOS || _la==ENTERO || _la==ID) {
					{
					setState(56);
					e(0);
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
	}

	public final Print_sentenciaContext print_sentencia() throws RecognitionException {
		Print_sentenciaContext _localctx = new Print_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_print_sentencia);
		int _la;
		try {
			_localctx = new PrintContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(61);
			match(PRINT);
			setState(62);
			match(PARIZQ);
			setState(63);
			e(0);
			setState(68);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(64);
				match(COMA);
				setState(65);
				e(0);
				}
				}
				setState(70);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(71);
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
	@SuppressWarnings("CheckReturnValue")
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
		enterRule(_localctx, 10, RULE_if_sentencia);
		int _la;
		try {
			_localctx = new IfContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			match(IF);
			setState(74);
			condicion(0);
			setState(75);
			match(LLAVEIZQ);
			setState(76);
			l_sentencias();
			setState(77);
			match(LLAVEDER);
			setState(86);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(78);
				match(ELSE);
				setState(84);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case IF:
					{
					setState(79);
					if_sentencia();
					}
					break;
				case LLAVEIZQ:
					{
					setState(80);
					match(LLAVEIZQ);
					setState(81);
					l_sentencias();
					setState(82);
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

	@SuppressWarnings("CheckReturnValue")
	public static class Guard_sentenciaContext extends ParserRuleContext {
		public Guard_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guard_sentencia; }
	 
		public Guard_sentenciaContext() { }
		public void copyFrom(Guard_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class GuardContext extends Guard_sentenciaContext {
		public TerminalNode GUARD() { return getToken(Tswift_GrammarNParser.GUARD, 0); }
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(Tswift_GrammarNParser.ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarNParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public Trans_sentenciaContext trans_sentencia() {
			return getRuleContext(Trans_sentenciaContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarNParser.LLAVEDER, 0); }
		public GuardContext(Guard_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final Guard_sentenciaContext guard_sentencia() throws RecognitionException {
		Guard_sentenciaContext _localctx = new Guard_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_guard_sentencia);
		try {
			_localctx = new GuardContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(88);
			match(GUARD);
			setState(89);
			condicion(0);
			setState(90);
			match(ELSE);
			setState(91);
			match(LLAVEIZQ);
			setState(92);
			l_sentencias();
			setState(93);
			trans_sentencia();
			setState(94);
			match(LLAVEDER);
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
	public static class Switch_sentenciaContext extends ParserRuleContext {
		public Switch_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_sentencia; }
	 
		public Switch_sentenciaContext() { }
		public void copyFrom(Switch_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchContext extends Switch_sentenciaContext {
		public TerminalNode SWITCH() { return getToken(Tswift_GrammarNParser.SWITCH, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarNParser.LLAVEIZQ, 0); }
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarNParser.LLAVEDER, 0); }
		public List<LcasosContext> lcasos() {
			return getRuleContexts(LcasosContext.class);
		}
		public LcasosContext lcasos(int i) {
			return getRuleContext(LcasosContext.class,i);
		}
		public CdefaultContext cdefault() {
			return getRuleContext(CdefaultContext.class,0);
		}
		public SwitchContext(Switch_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final Switch_sentenciaContext switch_sentencia() throws RecognitionException {
		Switch_sentenciaContext _localctx = new Switch_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_switch_sentencia);
		int _la;
		try {
			_localctx = new SwitchContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(96);
			match(SWITCH);
			setState(97);
			e(0);
			setState(98);
			match(LLAVEIZQ);
			setState(102);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(99);
				lcasos();
				}
				}
				setState(104);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(105);
				cdefault();
				}
			}

			setState(108);
			match(LLAVEDER);
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
	public static class LcasosContext extends ParserRuleContext {
		public LcasosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_lcasos; }
	 
		public LcasosContext() { }
		public void copyFrom(LcasosContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CaseContext extends LcasosContext {
		public TerminalNode CASE() { return getToken(Tswift_GrammarNParser.CASE, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarNParser.DOSPT, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public CaseContext(LcasosContext ctx) { copyFrom(ctx); }
	}

	public final LcasosContext lcasos() throws RecognitionException {
		LcasosContext _localctx = new LcasosContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_lcasos);
		try {
			_localctx = new CaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			match(CASE);
			setState(111);
			e(0);
			setState(112);
			match(DOSPT);
			setState(113);
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
	public static class CdefaultContext extends ParserRuleContext {
		public CdefaultContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cdefault; }
	 
		public CdefaultContext() { }
		public void copyFrom(CdefaultContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DefaultContext extends CdefaultContext {
		public TerminalNode DEFAULT() { return getToken(Tswift_GrammarNParser.DEFAULT, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarNParser.DOSPT, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public DefaultContext(CdefaultContext ctx) { copyFrom(ctx); }
	}

	public final CdefaultContext cdefault() throws RecognitionException {
		CdefaultContext _localctx = new CdefaultContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_cdefault);
		try {
			_localctx = new DefaultContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(115);
			match(DEFAULT);
			setState(116);
			match(DOSPT);
			setState(117);
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
	public static class While_sentenciaContext extends ParserRuleContext {
		public While_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_sentencia; }
	 
		public While_sentenciaContext() { }
		public void copyFrom(While_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class WhileContext extends While_sentenciaContext {
		public TerminalNode WHILE() { return getToken(Tswift_GrammarNParser.WHILE, 0); }
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarNParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarNParser.LLAVEDER, 0); }
		public WhileContext(While_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final While_sentenciaContext while_sentencia() throws RecognitionException {
		While_sentenciaContext _localctx = new While_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_while_sentencia);
		try {
			_localctx = new WhileContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			match(WHILE);
			setState(120);
			condicion(0);
			setState(121);
			match(LLAVEIZQ);
			setState(122);
			l_sentencias();
			setState(123);
			match(LLAVEDER);
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
	public static class For_sentenciaContext extends ParserRuleContext {
		public For_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_sentencia; }
	 
		public For_sentenciaContext() { }
		public void copyFrom(For_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForContext extends For_sentenciaContext {
		public TerminalNode FOR() { return getToken(Tswift_GrammarNParser.FOR, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode IN() { return getToken(Tswift_GrammarNParser.IN, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarNParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarNParser.LLAVEDER, 0); }
		public Rango_pContext rango_p() {
			return getRuleContext(Rango_pContext.class,0);
		}
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public ForContext(For_sentenciaContext ctx) { copyFrom(ctx); }
	}

	public final For_sentenciaContext for_sentencia() throws RecognitionException {
		For_sentenciaContext _localctx = new For_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_for_sentencia);
		try {
			_localctx = new ForContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			match(FOR);
			setState(126);
			match(ID);
			setState(127);
			match(IN);
			setState(130);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				setState(128);
				rango_p();
				}
				break;
			case 2:
				{
				setState(129);
				e(0);
				}
				break;
			}
			setState(132);
			match(LLAVEIZQ);
			setState(133);
			l_sentencias();
			setState(134);
			match(LLAVEDER);
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
	public static class Rango_pContext extends ParserRuleContext {
		public Rango_pContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rango_p; }
	 
		public Rango_pContext() { }
		public void copyFrom(Rango_pContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RangoContext extends Rango_pContext {
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode RANGO() { return getToken(Tswift_GrammarNParser.RANGO, 0); }
		public RangoContext(Rango_pContext ctx) { copyFrom(ctx); }
	}

	public final Rango_pContext rango_p() throws RecognitionException {
		Rango_pContext _localctx = new Rango_pContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_rango_p);
		try {
			_localctx = new RangoContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			e(0);
			setState(137);
			match(RANGO);
			setState(138);
			e(0);
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
		public CondicionContext c;
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public CondicionContext condicion() {
			return getRuleContext(CondicionContext.class,0);
		}
		public Cond_ParContext(CondicionContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
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
	@SuppressWarnings("CheckReturnValue")
	public static class Cond_BooleanoContext extends CondicionContext {
		public Token op;
		public TerminalNode TRUE() { return getToken(Tswift_GrammarNParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Tswift_GrammarNParser.FALSE, 0); }
		public Cond_BooleanoContext(CondicionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
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
		int _startState = 26;
		enterRecursionRule(_localctx, 26, RULE_condicion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				_localctx = new Cond_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(141);
				((Cond_NegContext)_localctx).op = match(NOT);
				setState(142);
				((Cond_NegContext)_localctx).c = condicion(5);
				}
				break;
			case 2:
				{
				_localctx = new Cond_RelContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(143);
				((Cond_RelContext)_localctx).e1 = e(0);
				setState(144);
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
				setState(145);
				((Cond_RelContext)_localctx).e2 = e(0);
				}
				break;
			case 3:
				{
				_localctx = new Cond_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(147);
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
				setState(148);
				match(PARIZQ);
				setState(149);
				((Cond_ParContext)_localctx).c = condicion(0);
				setState(150);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(159);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Cond_LogicaContext(new CondicionContext(_parentctx, _parentState));
					((Cond_LogicaContext)_localctx).c1 = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_condicion);
					setState(154);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(155);
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
					setState(156);
					((Cond_LogicaContext)_localctx).c2 = condicion(4);
					}
					} 
				}
				setState(161);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_IdContext extends EContext {
		public Token id;
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public Expr_IdContext(EContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
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
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_NegContext extends EContext {
		public Token op;
		public EContext e1;
		public TerminalNode MENOS() { return getToken(Tswift_GrammarNParser.MENOS, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public Expr_NegContext(EContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
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
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_ParContext extends EContext {
		public EContext e1;
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public Expr_ParContext(EContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
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
		int _startState = 28;
		enterRecursionRule(_localctx, 28, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MENOS:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(163);
				((Expr_NegContext)_localctx).op = match(MENOS);
				setState(164);
				((Expr_NegContext)_localctx).e1 = e(7);
				}
				break;
			case ID:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(165);
				((Expr_IdContext)_localctx).id = match(ID);
				}
				break;
			case ENTERO:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(166);
				((Expr_EnteroContext)_localctx).n = match(ENTERO);
				}
				break;
			case PARIZQ:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(167);
				match(PARIZQ);
				setState(168);
				((Expr_ParContext)_localctx).e1 = e(0);
				setState(169);
				match(PARDER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(184);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(182);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						((Expr_MulDivContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(173);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(174);
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
						setState(175);
						((Expr_MulDivContext)_localctx).e2 = e(7);
						}
						break;
					case 2:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						((Expr_SumResContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(176);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(177);
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
						setState(178);
						((Expr_SumResContext)_localctx).e2 = e(6);
						}
						break;
					case 3:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						((Expr_ModContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(179);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(180);
						((Expr_ModContext)_localctx).op = match(MOD);
						setState(181);
						((Expr_ModContext)_localctx).e2 = e(5);
						}
						break;
					}
					} 
				}
				setState(186);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
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
		case 13:
			return condicion_sempred((CondicionContext)_localctx, predIndex);
		case 14:
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
		"\u0004\u0001M\u00bc\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0005\u0001\"\b\u0001\n\u0001\f\u0001%\t\u0001\u0001\u0002"+
		"\u0001\u0002\u0003\u0002)\b\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u00020\b\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u00024\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0003\u0003:\b\u0003\u0003\u0003<\b\u0003\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0005\u0004C\b\u0004\n\u0004\f\u0004F\t"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0003\u0005U\b\u0005\u0003\u0005W\b\u0005\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007"+
		"e\b\u0007\n\u0007\f\u0007h\t\u0007\u0001\u0007\u0003\u0007k\b\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0003\u000b\u0083"+
		"\b\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0003\r\u0099\b\r\u0001\r\u0001"+
		"\r\u0001\r\u0005\r\u009e\b\r\n\r\f\r\u00a1\t\r\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0003\u000e\u00ac\b\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0005\u000e\u00b7\b\u000e\n\u000e\f\u000e\u00ba\t\u000e\u0001\u000e\u0000"+
		"\u0002\u001a\u001c\u000f\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0014\u0016\u0018\u001a\u001c\u0000\u0005\u0001\u0000\u0016\u001b\u0001"+
		"\u0000\"#\u0001\u0000\u001c\u001d\u0002\u0000\u0011\u0011\u0015\u0015"+
		"\u0001\u0000\u0013\u0014\u00c8\u0000\u001e\u0001\u0000\u0000\u0000\u0002"+
		"#\u0001\u0000\u0000\u0000\u00043\u0001\u0000\u0000\u0000\u0006;\u0001"+
		"\u0000\u0000\u0000\b=\u0001\u0000\u0000\u0000\nI\u0001\u0000\u0000\u0000"+
		"\fX\u0001\u0000\u0000\u0000\u000e`\u0001\u0000\u0000\u0000\u0010n\u0001"+
		"\u0000\u0000\u0000\u0012s\u0001\u0000\u0000\u0000\u0014w\u0001\u0000\u0000"+
		"\u0000\u0016}\u0001\u0000\u0000\u0000\u0018\u0088\u0001\u0000\u0000\u0000"+
		"\u001a\u0098\u0001\u0000\u0000\u0000\u001c\u00ab\u0001\u0000\u0000\u0000"+
		"\u001e\u001f\u0003\u0002\u0001\u0000\u001f\u0001\u0001\u0000\u0000\u0000"+
		" \"\u0003\u0004\u0002\u0000! \u0001\u0000\u0000\u0000\"%\u0001\u0000\u0000"+
		"\u0000#!\u0001\u0000\u0000\u0000#$\u0001\u0000\u0000\u0000$\u0003\u0001"+
		"\u0000\u0000\u0000%#\u0001\u0000\u0000\u0000&(\u0003\b\u0004\u0000\')"+
		"\u0005\b\u0000\u0000(\'\u0001\u0000\u0000\u0000()\u0001\u0000\u0000\u0000"+
		")4\u0001\u0000\u0000\u0000*4\u0003\n\u0005\u0000+4\u0003\u000e\u0007\u0000"+
		",4\u0003\f\u0006\u0000-/\u0003\u0006\u0003\u0000.0\u0005\b\u0000\u0000"+
		"/.\u0001\u0000\u0000\u0000/0\u0001\u0000\u0000\u000004\u0001\u0000\u0000"+
		"\u000014\u0003\u0014\n\u000024\u0003\u0016\u000b\u00003&\u0001\u0000\u0000"+
		"\u00003*\u0001\u0000\u0000\u00003+\u0001\u0000\u0000\u00003,\u0001\u0000"+
		"\u0000\u00003-\u0001\u0000\u0000\u000031\u0001\u0000\u0000\u000032\u0001"+
		"\u0000\u0000\u00004\u0005\u0001\u0000\u0000\u00005<\u00050\u0000\u0000"+
		"6<\u0005.\u0000\u000079\u0005/\u0000\u00008:\u0003\u001c\u000e\u00009"+
		"8\u0001\u0000\u0000\u00009:\u0001\u0000\u0000\u0000:<\u0001\u0000\u0000"+
		"\u0000;5\u0001\u0000\u0000\u0000;6\u0001\u0000\u0000\u0000;7\u0001\u0000"+
		"\u0000\u0000<\u0007\u0001\u0000\u0000\u0000=>\u0005\u001f\u0000\u0000"+
		">?\u0005\u0002\u0000\u0000?D\u0003\u001c\u000e\u0000@A\u0005\n\u0000\u0000"+
		"AC\u0003\u001c\u000e\u0000B@\u0001\u0000\u0000\u0000CF\u0001\u0000\u0000"+
		"\u0000DB\u0001\u0000\u0000\u0000DE\u0001\u0000\u0000\u0000EG\u0001\u0000"+
		"\u0000\u0000FD\u0001\u0000\u0000\u0000GH\u0005\u0001\u0000\u0000H\t\u0001"+
		"\u0000\u0000\u0000IJ\u0005$\u0000\u0000JK\u0003\u001a\r\u0000KL\u0005"+
		"\u0003\u0000\u0000LM\u0003\u0002\u0001\u0000MV\u0005\u0004\u0000\u0000"+
		"NT\u0005%\u0000\u0000OU\u0003\n\u0005\u0000PQ\u0005\u0003\u0000\u0000"+
		"QR\u0003\u0002\u0001\u0000RS\u0005\u0004\u0000\u0000SU\u0001\u0000\u0000"+
		"\u0000TO\u0001\u0000\u0000\u0000TP\u0001\u0000\u0000\u0000UW\u0001\u0000"+
		"\u0000\u0000VN\u0001\u0000\u0000\u0000VW\u0001\u0000\u0000\u0000W\u000b"+
		"\u0001\u0000\u0000\u0000XY\u0005-\u0000\u0000YZ\u0003\u001a\r\u0000Z["+
		"\u0005%\u0000\u0000[\\\u0005\u0003\u0000\u0000\\]\u0003\u0002\u0001\u0000"+
		"]^\u0003\u0006\u0003\u0000^_\u0005\u0004\u0000\u0000_\r\u0001\u0000\u0000"+
		"\u0000`a\u0005&\u0000\u0000ab\u0003\u001c\u000e\u0000bf\u0005\u0003\u0000"+
		"\u0000ce\u0003\u0010\b\u0000dc\u0001\u0000\u0000\u0000eh\u0001\u0000\u0000"+
		"\u0000fd\u0001\u0000\u0000\u0000fg\u0001\u0000\u0000\u0000gj\u0001\u0000"+
		"\u0000\u0000hf\u0001\u0000\u0000\u0000ik\u0003\u0012\t\u0000ji\u0001\u0000"+
		"\u0000\u0000jk\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000\u0000lm\u0005"+
		"\u0004\u0000\u0000m\u000f\u0001\u0000\u0000\u0000no\u0005\'\u0000\u0000"+
		"op\u0003\u001c\u000e\u0000pq\u0005\u0007\u0000\u0000qr\u0003\u0002\u0001"+
		"\u0000r\u0011\u0001\u0000\u0000\u0000st\u0005(\u0000\u0000tu\u0005\u0007"+
		"\u0000\u0000uv\u0003\u0002\u0001\u0000v\u0013\u0001\u0000\u0000\u0000"+
		"wx\u0005)\u0000\u0000xy\u0003\u001a\r\u0000yz\u0005\u0003\u0000\u0000"+
		"z{\u0003\u0002\u0001\u0000{|\u0005\u0004\u0000\u0000|\u0015\u0001\u0000"+
		"\u0000\u0000}~\u0005*\u0000\u0000~\u007f\u0005J\u0000\u0000\u007f\u0082"+
		"\u0005+\u0000\u0000\u0080\u0083\u0003\u0018\f\u0000\u0081\u0083\u0003"+
		"\u001c\u000e\u0000\u0082\u0080\u0001\u0000\u0000\u0000\u0082\u0081\u0001"+
		"\u0000\u0000\u0000\u0083\u0084\u0001\u0000\u0000\u0000\u0084\u0085\u0005"+
		"\u0003\u0000\u0000\u0085\u0086\u0003\u0002\u0001\u0000\u0086\u0087\u0005"+
		"\u0004\u0000\u0000\u0087\u0017\u0001\u0000\u0000\u0000\u0088\u0089\u0003"+
		"\u001c\u000e\u0000\u0089\u008a\u0005,\u0000\u0000\u008a\u008b\u0003\u001c"+
		"\u000e\u0000\u008b\u0019\u0001\u0000\u0000\u0000\u008c\u008d\u0006\r\uffff"+
		"\uffff\u0000\u008d\u008e\u0005\u001e\u0000\u0000\u008e\u0099\u0003\u001a"+
		"\r\u0005\u008f\u0090\u0003\u001c\u000e\u0000\u0090\u0091\u0007\u0000\u0000"+
		"\u0000\u0091\u0092\u0003\u001c\u000e\u0000\u0092\u0099\u0001\u0000\u0000"+
		"\u0000\u0093\u0099\u0007\u0001\u0000\u0000\u0094\u0095\u0005\u0002\u0000"+
		"\u0000\u0095\u0096\u0003\u001a\r\u0000\u0096\u0097\u0005\u0001\u0000\u0000"+
		"\u0097\u0099\u0001\u0000\u0000\u0000\u0098\u008c\u0001\u0000\u0000\u0000"+
		"\u0098\u008f\u0001\u0000\u0000\u0000\u0098\u0093\u0001\u0000\u0000\u0000"+
		"\u0098\u0094\u0001\u0000\u0000\u0000\u0099\u009f\u0001\u0000\u0000\u0000"+
		"\u009a\u009b\n\u0003\u0000\u0000\u009b\u009c\u0007\u0002\u0000\u0000\u009c"+
		"\u009e\u0003\u001a\r\u0004\u009d\u009a\u0001\u0000\u0000\u0000\u009e\u00a1"+
		"\u0001\u0000\u0000\u0000\u009f\u009d\u0001\u0000\u0000\u0000\u009f\u00a0"+
		"\u0001\u0000\u0000\u0000\u00a0\u001b\u0001\u0000\u0000\u0000\u00a1\u009f"+
		"\u0001\u0000\u0000\u0000\u00a2\u00a3\u0006\u000e\uffff\uffff\u0000\u00a3"+
		"\u00a4\u0005\u0013\u0000\u0000\u00a4\u00ac\u0003\u001c\u000e\u0007\u00a5"+
		"\u00ac\u0005J\u0000\u0000\u00a6\u00ac\u0005F\u0000\u0000\u00a7\u00a8\u0005"+
		"\u0002\u0000\u0000\u00a8\u00a9\u0003\u001c\u000e\u0000\u00a9\u00aa\u0005"+
		"\u0001\u0000\u0000\u00aa\u00ac\u0001\u0000\u0000\u0000\u00ab\u00a2\u0001"+
		"\u0000\u0000\u0000\u00ab\u00a5\u0001\u0000\u0000\u0000\u00ab\u00a6\u0001"+
		"\u0000\u0000\u0000\u00ab\u00a7\u0001\u0000\u0000\u0000\u00ac\u00b8\u0001"+
		"\u0000\u0000\u0000\u00ad\u00ae\n\u0006\u0000\u0000\u00ae\u00af\u0007\u0003"+
		"\u0000\u0000\u00af\u00b7\u0003\u001c\u000e\u0007\u00b0\u00b1\n\u0005\u0000"+
		"\u0000\u00b1\u00b2\u0007\u0004\u0000\u0000\u00b2\u00b7\u0003\u001c\u000e"+
		"\u0006\u00b3\u00b4\n\u0004\u0000\u0000\u00b4\u00b5\u0005\u0012\u0000\u0000"+
		"\u00b5\u00b7\u0003\u001c\u000e\u0005\u00b6\u00ad\u0001\u0000\u0000\u0000"+
		"\u00b6\u00b0\u0001\u0000\u0000\u0000\u00b6\u00b3\u0001\u0000\u0000\u0000"+
		"\u00b7\u00ba\u0001\u0000\u0000\u0000\u00b8\u00b6\u0001\u0000\u0000\u0000"+
		"\u00b8\u00b9\u0001\u0000\u0000\u0000\u00b9\u001d\u0001\u0000\u0000\u0000"+
		"\u00ba\u00b8\u0001\u0000\u0000\u0000\u0011#(/39;DTVfj\u0082\u0098\u009f"+
		"\u00ab\u00b6\u00b8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}