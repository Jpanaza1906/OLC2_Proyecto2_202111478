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
		BOOL=66, CHAR=67, STRING=68, ENBLANCO=69, ENTERO=70, DECIMAL=71, CARACTER=72, 
		CADENA=73, ID=74, UL_C=75, ML_C=76, ERROR=77;
	public static final int
		RULE_s = 0, RULE_l_sentencias = 1, RULE_sentencia = 2, RULE_trans_sentencia = 3, 
		RULE_print_sentencia = 4, RULE_declaracion_sentencia = 5, RULE_asignacion_sentencia = 6, 
		RULE_tipo = 7, RULE_if_sentencia = 8, RULE_guard_sentencia = 9, RULE_switch_sentencia = 10, 
		RULE_lcasos = 11, RULE_cdefault = 12, RULE_while_sentencia = 13, RULE_for_sentencia = 14, 
		RULE_condicion = 15, RULE_e = 16;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "trans_sentencia", "print_sentencia", 
			"declaracion_sentencia", "asignacion_sentencia", "tipo", "if_sentencia", 
			"guard_sentencia", "switch_sentencia", "lcasos", "cdefault", "while_sentencia", 
			"for_sentencia", "condicion", "e"
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
			"'self'", "'Int'", "'Float'", "'Bool'", "'Char'", "'String'"
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
			"INOUT", "ATOI", "IOTA", "SELF", "INT", "FLOAT", "BOOL", "CHAR", "STRING", 
			"ENBLANCO", "ENTERO", "DECIMAL", "CARACTER", "CADENA", "ID", "UL_C", 
			"ML_C", "ERROR"
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
			setState(34);
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
		try {
			int _alt;
			_localctx = new L_SentenciaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(36);
					sentencia();
					}
					} 
				}
				setState(41);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_For(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_For(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_SwitchContext extends SentenciaContext {
		public Switch_sentenciaContext switch_sentencia() {
			return getRuleContext(Switch_sentenciaContext.class,0);
		}
		public S_SwitchContext(SentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_Switch(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_Switch(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_AsignacionContext extends SentenciaContext {
		public Asignacion_sentenciaContext asignacion_sentencia() {
			return getRuleContext(Asignacion_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_AsignacionContext(SentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_Asignacion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_Asignacion(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_GuardContext extends SentenciaContext {
		public Guard_sentenciaContext guard_sentencia() {
			return getRuleContext(Guard_sentenciaContext.class,0);
		}
		public S_GuardContext(SentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_Guard(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_Guard(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_WhileContext extends SentenciaContext {
		public While_sentenciaContext while_sentencia() {
			return getRuleContext(While_sentenciaContext.class,0);
		}
		public S_WhileContext(SentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_While(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_While(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_TransContext extends SentenciaContext {
		public Trans_sentenciaContext trans_sentencia() {
			return getRuleContext(Trans_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_TransContext(SentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_Trans(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_Trans(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_IfContext extends SentenciaContext {
		public If_sentenciaContext if_sentencia() {
			return getRuleContext(If_sentenciaContext.class,0);
		}
		public S_IfContext(SentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_If(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_If(this);
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
	@SuppressWarnings("CheckReturnValue")
	public static class S_DeclaracionContext extends SentenciaContext {
		public Declaracion_sentenciaContext declaracion_sentencia() {
			return getRuleContext(Declaracion_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_DeclaracionContext(SentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterS_Declaracion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitS_Declaracion(this);
		}
	}

	public final SentenciaContext sentencia() throws RecognitionException {
		SentenciaContext _localctx = new SentenciaContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sentencia);
		int _la;
		try {
			setState(63);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new S_PrintContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(42);
				print_sentencia();
				setState(44);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(43);
					match(PTCOMA);
					}
				}

				}
				break;
			case VAR:
			case LET:
				_localctx = new S_DeclaracionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(46);
				declaracion_sentencia();
				setState(48);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(47);
					match(PTCOMA);
					}
				}

				}
				break;
			case ID:
				_localctx = new S_AsignacionContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(50);
				asignacion_sentencia();
				setState(52);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(51);
					match(PTCOMA);
					}
				}

				}
				break;
			case IF:
				_localctx = new S_IfContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(54);
				if_sentencia();
				}
				break;
			case SWITCH:
				_localctx = new S_SwitchContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(55);
				switch_sentencia();
				}
				break;
			case GUARD:
				_localctx = new S_GuardContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(56);
				guard_sentencia();
				}
				break;
			case CONTINUE:
			case RETURN:
			case BREAK:
				_localctx = new S_TransContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(57);
				trans_sentencia();
				setState(59);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(58);
					match(PTCOMA);
					}
				}

				}
				break;
			case WHILE:
				_localctx = new S_WhileContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(61);
				while_sentencia();
				}
				break;
			case FOR:
				_localctx = new S_ForContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(62);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterReturn(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitReturn(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakContext extends Trans_sentenciaContext {
		public TerminalNode BREAK() { return getToken(Tswift_GrammarNParser.BREAK, 0); }
		public BreakContext(Trans_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterBreak(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitBreak(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueContext extends Trans_sentenciaContext {
		public TerminalNode CONTINUE() { return getToken(Tswift_GrammarNParser.CONTINUE, 0); }
		public ContinueContext(Trans_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterContinue(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitContinue(this);
		}
	}

	public final Trans_sentenciaContext trans_sentencia() throws RecognitionException {
		Trans_sentenciaContext _localctx = new Trans_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_trans_sentencia);
		try {
			setState(71);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BREAK:
				_localctx = new BreakContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
				match(BREAK);
				}
				break;
			case CONTINUE:
				_localctx = new ContinueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				match(CONTINUE);
				}
				break;
			case RETURN:
				_localctx = new ReturnContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(67);
				match(RETURN);
				setState(69);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
				case 1:
					{
					setState(68);
					e(0);
					}
					break;
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
		enterRule(_localctx, 8, RULE_print_sentencia);
		int _la;
		try {
			_localctx = new PrintContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(73);
			match(PRINT);
			setState(74);
			match(PARIZQ);
			setState(75);
			e(0);
			setState(80);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(76);
				match(COMA);
				setState(77);
				e(0);
				}
				}
				setState(82);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(83);
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
	public static class Declaracion_sentenciaContext extends ParserRuleContext {
		public Declaracion_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_sentencia; }
	 
		public Declaracion_sentenciaContext() { }
		public void copyFrom(Declaracion_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Declaracion_Tipo_ValContext extends Declaracion_sentenciaContext {
		public Token tip;
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarNParser.DOSPT, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarNParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode VAR() { return getToken(Tswift_GrammarNParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarNParser.LET, 0); }
		public Declaracion_Tipo_ValContext(Declaracion_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterDeclaracion_Tipo_Val(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitDeclaracion_Tipo_Val(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Declaracion_Tipo_noValContext extends Declaracion_sentenciaContext {
		public Token tip;
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarNParser.DOSPT, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode INTERROGACION() { return getToken(Tswift_GrammarNParser.INTERROGACION, 0); }
		public TerminalNode VAR() { return getToken(Tswift_GrammarNParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarNParser.LET, 0); }
		public Declaracion_Tipo_noValContext(Declaracion_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterDeclaracion_Tipo_noVal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitDeclaracion_Tipo_noVal(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Declaracion_ValContext extends Declaracion_sentenciaContext {
		public Token tip;
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarNParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode VAR() { return getToken(Tswift_GrammarNParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarNParser.LET, 0); }
		public Declaracion_ValContext(Declaracion_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterDeclaracion_Val(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitDeclaracion_Val(this);
		}
	}

	public final Declaracion_sentenciaContext declaracion_sentencia() throws RecognitionException {
		Declaracion_sentenciaContext _localctx = new Declaracion_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_declaracion_sentencia);
		int _la;
		try {
			setState(102);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				_localctx = new Declaracion_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				((Declaracion_Tipo_ValContext)_localctx).tip = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
					((Declaracion_Tipo_ValContext)_localctx).tip = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(86);
				match(ID);
				setState(87);
				match(DOSPT);
				setState(88);
				tipo();
				setState(89);
				match(IGUAL);
				setState(90);
				e(0);
				}
				break;
			case 2:
				_localctx = new Declaracion_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(92);
				((Declaracion_ValContext)_localctx).tip = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
					((Declaracion_ValContext)_localctx).tip = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(93);
				match(ID);
				setState(94);
				match(IGUAL);
				setState(95);
				e(0);
				}
				break;
			case 3:
				_localctx = new Declaracion_Tipo_noValContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(96);
				((Declaracion_Tipo_noValContext)_localctx).tip = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
					((Declaracion_Tipo_noValContext)_localctx).tip = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(97);
				match(ID);
				setState(98);
				match(DOSPT);
				setState(99);
				tipo();
				setState(100);
				match(INTERROGACION);
				}
				break;
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
	public static class Asignacion_sentenciaContext extends ParserRuleContext {
		public Asignacion_sentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion_sentencia; }
	 
		public Asignacion_sentenciaContext() { }
		public void copyFrom(Asignacion_sentenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionContext extends Asignacion_sentenciaContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarNParser.IGUAL, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public AsignacionContext(Asignacion_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterAsignacion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitAsignacion(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Asignacion_SumaRestaContext extends Asignacion_sentenciaContext {
		public Token op;
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode MASIGUAL() { return getToken(Tswift_GrammarNParser.MASIGUAL, 0); }
		public TerminalNode MENOSIGUAL() { return getToken(Tswift_GrammarNParser.MENOSIGUAL, 0); }
		public Asignacion_SumaRestaContext(Asignacion_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterAsignacion_SumaResta(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitAsignacion_SumaResta(this);
		}
	}

	public final Asignacion_sentenciaContext asignacion_sentencia() throws RecognitionException {
		Asignacion_sentenciaContext _localctx = new Asignacion_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_asignacion_sentencia);
		int _la;
		try {
			setState(110);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				_localctx = new AsignacionContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(104);
				match(ID);
				setState(105);
				match(IGUAL);
				setState(106);
				e(0);
				}
				break;
			case 2:
				_localctx = new Asignacion_SumaRestaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(107);
				match(ID);
				setState(108);
				((Asignacion_SumaRestaContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==MASIGUAL || _la==MENOSIGUAL) ) {
					((Asignacion_SumaRestaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(109);
				e(0);
				}
				break;
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
	public static class TipoContext extends ParserRuleContext {
		public TipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo; }
	 
		public TipoContext() { }
		public void copyFrom(TipoContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_VectorContext extends TipoContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public Tipo_VectorContext(TipoContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterTipo_Vector(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitTipo_Vector(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_StructContext extends TipoContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public Tipo_StructContext(TipoContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterTipo_Struct(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitTipo_Struct(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_FloatContext extends TipoContext {
		public TerminalNode FLOAT() { return getToken(Tswift_GrammarNParser.FLOAT, 0); }
		public Tipo_FloatContext(TipoContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterTipo_Float(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitTipo_Float(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_CharacterContext extends TipoContext {
		public TerminalNode CHAR() { return getToken(Tswift_GrammarNParser.CHAR, 0); }
		public Tipo_CharacterContext(TipoContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterTipo_Character(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitTipo_Character(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_StringContext extends TipoContext {
		public TerminalNode STRING() { return getToken(Tswift_GrammarNParser.STRING, 0); }
		public Tipo_StringContext(TipoContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterTipo_String(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitTipo_String(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_BoolContext extends TipoContext {
		public TerminalNode BOOL() { return getToken(Tswift_GrammarNParser.BOOL, 0); }
		public Tipo_BoolContext(TipoContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterTipo_Bool(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitTipo_Bool(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_IntContext extends TipoContext {
		public TerminalNode INT() { return getToken(Tswift_GrammarNParser.INT, 0); }
		public Tipo_IntContext(TipoContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterTipo_Int(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitTipo_Int(this);
		}
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_tipo);
		try {
			setState(122);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new Tipo_IntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(112);
				match(INT);
				}
				break;
			case FLOAT:
				_localctx = new Tipo_FloatContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(113);
				match(FLOAT);
				}
				break;
			case STRING:
				_localctx = new Tipo_StringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(114);
				match(STRING);
				}
				break;
			case BOOL:
				_localctx = new Tipo_BoolContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(115);
				match(BOOL);
				}
				break;
			case CHAR:
				_localctx = new Tipo_CharacterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(116);
				match(CHAR);
				}
				break;
			case ID:
				_localctx = new Tipo_StructContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(117);
				match(ID);
				}
				break;
			case CORCHETEIZQ:
				_localctx = new Tipo_VectorContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(118);
				match(CORCHETEIZQ);
				setState(119);
				tipo();
				setState(120);
				match(CORCHETEDER);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterIf(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitIf(this);
		}
	}

	public final If_sentenciaContext if_sentencia() throws RecognitionException {
		If_sentenciaContext _localctx = new If_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_if_sentencia);
		int _la;
		try {
			_localctx = new IfContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(124);
			match(IF);
			setState(125);
			condicion(0);
			setState(126);
			match(LLAVEIZQ);
			setState(127);
			l_sentencias();
			setState(128);
			match(LLAVEDER);
			setState(137);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(129);
				match(ELSE);
				setState(135);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case IF:
					{
					setState(130);
					if_sentencia();
					}
					break;
				case LLAVEIZQ:
					{
					setState(131);
					match(LLAVEIZQ);
					setState(132);
					l_sentencias();
					setState(133);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterGuard(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitGuard(this);
		}
	}

	public final Guard_sentenciaContext guard_sentencia() throws RecognitionException {
		Guard_sentenciaContext _localctx = new Guard_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_guard_sentencia);
		try {
			_localctx = new GuardContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			match(GUARD);
			setState(140);
			condicion(0);
			setState(141);
			match(ELSE);
			setState(142);
			match(LLAVEIZQ);
			setState(143);
			l_sentencias();
			setState(144);
			trans_sentencia();
			setState(145);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterSwitch(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitSwitch(this);
		}
	}

	public final Switch_sentenciaContext switch_sentencia() throws RecognitionException {
		Switch_sentenciaContext _localctx = new Switch_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_switch_sentencia);
		int _la;
		try {
			_localctx = new SwitchContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(147);
			match(SWITCH);
			setState(148);
			e(0);
			setState(149);
			match(LLAVEIZQ);
			setState(153);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(150);
				lcasos();
				}
				}
				setState(155);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(156);
				cdefault();
				}
			}

			setState(159);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterCase(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitCase(this);
		}
	}

	public final LcasosContext lcasos() throws RecognitionException {
		LcasosContext _localctx = new LcasosContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_lcasos);
		try {
			_localctx = new CaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(161);
			match(CASE);
			setState(162);
			e(0);
			setState(163);
			match(DOSPT);
			setState(164);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterDefault(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitDefault(this);
		}
	}

	public final CdefaultContext cdefault() throws RecognitionException {
		CdefaultContext _localctx = new CdefaultContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_cdefault);
		try {
			_localctx = new DefaultContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(166);
			match(DEFAULT);
			setState(167);
			match(DOSPT);
			setState(168);
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
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterWhile(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitWhile(this);
		}
	}

	public final While_sentenciaContext while_sentencia() throws RecognitionException {
		While_sentenciaContext _localctx = new While_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_while_sentencia);
		try {
			_localctx = new WhileContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			match(WHILE);
			setState(171);
			condicion(0);
			setState(172);
			match(LLAVEIZQ);
			setState(173);
			l_sentencias();
			setState(174);
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
	public static class ForIntContext extends For_sentenciaContext {
		public TerminalNode FOR() { return getToken(Tswift_GrammarNParser.FOR, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode IN() { return getToken(Tswift_GrammarNParser.IN, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode RANGO() { return getToken(Tswift_GrammarNParser.RANGO, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarNParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarNParser.LLAVEDER, 0); }
		public ForIntContext(For_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterForInt(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitForInt(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForListContext extends For_sentenciaContext {
		public TerminalNode FOR() { return getToken(Tswift_GrammarNParser.FOR, 0); }
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode IN() { return getToken(Tswift_GrammarNParser.IN, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Tswift_GrammarNParser.LLAVEIZQ, 0); }
		public L_sentenciasContext l_sentencias() {
			return getRuleContext(L_sentenciasContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Tswift_GrammarNParser.LLAVEDER, 0); }
		public ForListContext(For_sentenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterForList(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitForList(this);
		}
	}

	public final For_sentenciaContext for_sentencia() throws RecognitionException {
		For_sentenciaContext _localctx = new For_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_for_sentencia);
		try {
			setState(194);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				_localctx = new ForIntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(176);
				match(FOR);
				setState(177);
				match(ID);
				setState(178);
				match(IN);
				setState(179);
				e(0);
				setState(180);
				match(RANGO);
				setState(181);
				e(0);
				setState(182);
				match(LLAVEIZQ);
				setState(183);
				l_sentencias();
				setState(184);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new ForListContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(186);
				match(FOR);
				setState(187);
				match(ID);
				setState(188);
				match(IN);
				setState(189);
				e(0);
				setState(190);
				match(LLAVEIZQ);
				setState(191);
				l_sentencias();
				setState(192);
				match(LLAVEDER);
				}
				break;
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
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_condicion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(208);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
			case 1:
				{
				_localctx = new Cond_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(197);
				((Cond_NegContext)_localctx).op = match(NOT);
				setState(198);
				((Cond_NegContext)_localctx).c = condicion(5);
				}
				break;
			case 2:
				{
				_localctx = new Cond_RelContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(199);
				((Cond_RelContext)_localctx).e1 = e(0);
				setState(200);
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
				setState(201);
				((Cond_RelContext)_localctx).e2 = e(0);
				}
				break;
			case 3:
				{
				_localctx = new Cond_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(203);
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
				setState(204);
				match(PARIZQ);
				setState(205);
				((Cond_ParContext)_localctx).c = condicion(0);
				setState(206);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(215);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Cond_LogicaContext(new CondicionContext(_parentctx, _parentState));
					((Cond_LogicaContext)_localctx).c1 = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_condicion);
					setState(210);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(211);
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
					setState(212);
					((Cond_LogicaContext)_localctx).c2 = condicion(4);
					}
					} 
				}
				setState(217);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,18,_ctx);
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
	public static class Expr_DecimalContext extends EContext {
		public Token n;
		public TerminalNode DECIMAL() { return getToken(Tswift_GrammarNParser.DECIMAL, 0); }
		public Expr_DecimalContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Decimal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Decimal(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_CaracterContext extends EContext {
		public Token n;
		public TerminalNode CARACTER() { return getToken(Tswift_GrammarNParser.CARACTER, 0); }
		public Expr_CaracterContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Caracter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Caracter(this);
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
		public Token id;
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
		public EContext e1;
		public TerminalNode MENOS() { return getToken(Tswift_GrammarNParser.MENOS, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
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
	public static class Expr_NilContext extends EContext {
		public Token n;
		public TerminalNode NIL() { return getToken(Tswift_GrammarNParser.NIL, 0); }
		public Expr_NilContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Nil(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Nil(this);
		}
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
	public static class Expr_BooleanoContext extends EContext {
		public Token n;
		public TerminalNode TRUE() { return getToken(Tswift_GrammarNParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Tswift_GrammarNParser.FALSE, 0); }
		public Expr_BooleanoContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Booleano(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Booleano(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_EnteroContext extends EContext {
		public Token n;
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
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_CadenaContext extends EContext {
		public Token n;
		public TerminalNode CADENA() { return getToken(Tswift_GrammarNParser.CADENA, 0); }
		public Expr_CadenaContext(EContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).enterExpr_Cadena(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof Tswift_GrammarNListener ) ((Tswift_GrammarNListener)listener).exitExpr_Cadena(this);
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
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(232);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MENOS:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(219);
				((Expr_NegContext)_localctx).op = match(MENOS);
				setState(220);
				((Expr_NegContext)_localctx).e1 = e(12);
				}
				break;
			case TRUE:
			case FALSE:
				{
				_localctx = new Expr_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(221);
				((Expr_BooleanoContext)_localctx).n = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==TRUE || _la==FALSE) ) {
					((Expr_BooleanoContext)_localctx).n = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case NIL:
				{
				_localctx = new Expr_NilContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(222);
				((Expr_NilContext)_localctx).n = match(NIL);
				}
				break;
			case ID:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(223);
				((Expr_IdContext)_localctx).id = match(ID);
				}
				break;
			case ENTERO:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(224);
				((Expr_EnteroContext)_localctx).n = match(ENTERO);
				}
				break;
			case DECIMAL:
				{
				_localctx = new Expr_DecimalContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(225);
				((Expr_DecimalContext)_localctx).n = match(DECIMAL);
				}
				break;
			case CADENA:
				{
				_localctx = new Expr_CadenaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(226);
				((Expr_CadenaContext)_localctx).n = match(CADENA);
				}
				break;
			case CARACTER:
				{
				_localctx = new Expr_CaracterContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(227);
				((Expr_CaracterContext)_localctx).n = match(CARACTER);
				}
				break;
			case PARIZQ:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(228);
				match(PARIZQ);
				setState(229);
				((Expr_ParContext)_localctx).e1 = e(0);
				setState(230);
				match(PARDER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(245);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(243);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						((Expr_MulDivContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(234);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(235);
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
						setState(236);
						((Expr_MulDivContext)_localctx).e2 = e(12);
						}
						break;
					case 2:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						((Expr_SumResContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(237);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(238);
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
						setState(239);
						((Expr_SumResContext)_localctx).e2 = e(11);
						}
						break;
					case 3:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						((Expr_ModContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(240);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(241);
						((Expr_ModContext)_localctx).op = match(MOD);
						setState(242);
						((Expr_ModContext)_localctx).e2 = e(10);
						}
						break;
					}
					} 
				}
				setState(247);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
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
		case 15:
			return condicion_sempred((CondicionContext)_localctx, predIndex);
		case 16:
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
			return precpred(_ctx, 11);
		case 2:
			return precpred(_ctx, 10);
		case 3:
			return precpred(_ctx, 9);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001M\u00f9\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0001\u0000\u0001\u0000\u0001\u0001\u0005\u0001"+
		"&\b\u0001\n\u0001\f\u0001)\t\u0001\u0001\u0002\u0001\u0002\u0003\u0002"+
		"-\b\u0002\u0001\u0002\u0001\u0002\u0003\u00021\b\u0002\u0001\u0002\u0001"+
		"\u0002\u0003\u00025\b\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002<\b\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002@\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003"+
		"\u0003F\b\u0003\u0003\u0003H\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0005\u0004O\b\u0004\n\u0004\f\u0004R\t\u0004"+
		"\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0003\u0005g\b\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006o\b\u0006\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007{\b\u0007\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0003\b\u0088\b\b\u0003\b\u008a\b\b\u0001\t\u0001\t\u0001\t\u0001\t"+
		"\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0005"+
		"\n\u0098\b\n\n\n\f\n\u009b\t\n\u0001\n\u0003\n\u009e\b\n\u0001\n\u0001"+
		"\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0003\u000e\u00c3\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0003\u000f\u00d1\b\u000f\u0001\u000f\u0001\u000f"+
		"\u0001\u000f\u0005\u000f\u00d6\b\u000f\n\u000f\f\u000f\u00d9\t\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0003\u0010\u00e9\b\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0005\u0010\u00f4\b\u0010\n\u0010\f\u0010\u00f7\t\u0010\u0001\u0010"+
		"\u0000\u0002\u001e \u0011\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0014\u0016\u0018\u001a\u001c\u001e \u0000\u0007\u0001\u0000 !\u0001"+
		"\u0000\u000e\u000f\u0001\u0000\u0016\u001b\u0001\u0000\"#\u0001\u0000"+
		"\u001c\u001d\u0002\u0000\u0011\u0011\u0015\u0015\u0001\u0000\u0013\u0014"+
		"\u0115\u0000\"\u0001\u0000\u0000\u0000\u0002\'\u0001\u0000\u0000\u0000"+
		"\u0004?\u0001\u0000\u0000\u0000\u0006G\u0001\u0000\u0000\u0000\bI\u0001"+
		"\u0000\u0000\u0000\nf\u0001\u0000\u0000\u0000\fn\u0001\u0000\u0000\u0000"+
		"\u000ez\u0001\u0000\u0000\u0000\u0010|\u0001\u0000\u0000\u0000\u0012\u008b"+
		"\u0001\u0000\u0000\u0000\u0014\u0093\u0001\u0000\u0000\u0000\u0016\u00a1"+
		"\u0001\u0000\u0000\u0000\u0018\u00a6\u0001\u0000\u0000\u0000\u001a\u00aa"+
		"\u0001\u0000\u0000\u0000\u001c\u00c2\u0001\u0000\u0000\u0000\u001e\u00d0"+
		"\u0001\u0000\u0000\u0000 \u00e8\u0001\u0000\u0000\u0000\"#\u0003\u0002"+
		"\u0001\u0000#\u0001\u0001\u0000\u0000\u0000$&\u0003\u0004\u0002\u0000"+
		"%$\u0001\u0000\u0000\u0000&)\u0001\u0000\u0000\u0000\'%\u0001\u0000\u0000"+
		"\u0000\'(\u0001\u0000\u0000\u0000(\u0003\u0001\u0000\u0000\u0000)\'\u0001"+
		"\u0000\u0000\u0000*,\u0003\b\u0004\u0000+-\u0005\b\u0000\u0000,+\u0001"+
		"\u0000\u0000\u0000,-\u0001\u0000\u0000\u0000-@\u0001\u0000\u0000\u0000"+
		".0\u0003\n\u0005\u0000/1\u0005\b\u0000\u00000/\u0001\u0000\u0000\u0000"+
		"01\u0001\u0000\u0000\u00001@\u0001\u0000\u0000\u000024\u0003\f\u0006\u0000"+
		"35\u0005\b\u0000\u000043\u0001\u0000\u0000\u000045\u0001\u0000\u0000\u0000"+
		"5@\u0001\u0000\u0000\u00006@\u0003\u0010\b\u00007@\u0003\u0014\n\u0000"+
		"8@\u0003\u0012\t\u00009;\u0003\u0006\u0003\u0000:<\u0005\b\u0000\u0000"+
		";:\u0001\u0000\u0000\u0000;<\u0001\u0000\u0000\u0000<@\u0001\u0000\u0000"+
		"\u0000=@\u0003\u001a\r\u0000>@\u0003\u001c\u000e\u0000?*\u0001\u0000\u0000"+
		"\u0000?.\u0001\u0000\u0000\u0000?2\u0001\u0000\u0000\u0000?6\u0001\u0000"+
		"\u0000\u0000?7\u0001\u0000\u0000\u0000?8\u0001\u0000\u0000\u0000?9\u0001"+
		"\u0000\u0000\u0000?=\u0001\u0000\u0000\u0000?>\u0001\u0000\u0000\u0000"+
		"@\u0005\u0001\u0000\u0000\u0000AH\u00050\u0000\u0000BH\u0005.\u0000\u0000"+
		"CE\u0005/\u0000\u0000DF\u0003 \u0010\u0000ED\u0001\u0000\u0000\u0000E"+
		"F\u0001\u0000\u0000\u0000FH\u0001\u0000\u0000\u0000GA\u0001\u0000\u0000"+
		"\u0000GB\u0001\u0000\u0000\u0000GC\u0001\u0000\u0000\u0000H\u0007\u0001"+
		"\u0000\u0000\u0000IJ\u0005\u001f\u0000\u0000JK\u0005\u0002\u0000\u0000"+
		"KP\u0003 \u0010\u0000LM\u0005\n\u0000\u0000MO\u0003 \u0010\u0000NL\u0001"+
		"\u0000\u0000\u0000OR\u0001\u0000\u0000\u0000PN\u0001\u0000\u0000\u0000"+
		"PQ\u0001\u0000\u0000\u0000QS\u0001\u0000\u0000\u0000RP\u0001\u0000\u0000"+
		"\u0000ST\u0005\u0001\u0000\u0000T\t\u0001\u0000\u0000\u0000UV\u0007\u0000"+
		"\u0000\u0000VW\u0005J\u0000\u0000WX\u0005\u0007\u0000\u0000XY\u0003\u000e"+
		"\u0007\u0000YZ\u0005\u0010\u0000\u0000Z[\u0003 \u0010\u0000[g\u0001\u0000"+
		"\u0000\u0000\\]\u0007\u0000\u0000\u0000]^\u0005J\u0000\u0000^_\u0005\u0010"+
		"\u0000\u0000_g\u0003 \u0010\u0000`a\u0007\u0000\u0000\u0000ab\u0005J\u0000"+
		"\u0000bc\u0005\u0007\u0000\u0000cd\u0003\u000e\u0007\u0000de\u0005\t\u0000"+
		"\u0000eg\u0001\u0000\u0000\u0000fU\u0001\u0000\u0000\u0000f\\\u0001\u0000"+
		"\u0000\u0000f`\u0001\u0000\u0000\u0000g\u000b\u0001\u0000\u0000\u0000"+
		"hi\u0005J\u0000\u0000ij\u0005\u0010\u0000\u0000jo\u0003 \u0010\u0000k"+
		"l\u0005J\u0000\u0000lm\u0007\u0001\u0000\u0000mo\u0003 \u0010\u0000nh"+
		"\u0001\u0000\u0000\u0000nk\u0001\u0000\u0000\u0000o\r\u0001\u0000\u0000"+
		"\u0000p{\u0005@\u0000\u0000q{\u0005A\u0000\u0000r{\u0005D\u0000\u0000"+
		"s{\u0005B\u0000\u0000t{\u0005C\u0000\u0000u{\u0005J\u0000\u0000vw\u0005"+
		"\u0005\u0000\u0000wx\u0003\u000e\u0007\u0000xy\u0005\u0006\u0000\u0000"+
		"y{\u0001\u0000\u0000\u0000zp\u0001\u0000\u0000\u0000zq\u0001\u0000\u0000"+
		"\u0000zr\u0001\u0000\u0000\u0000zs\u0001\u0000\u0000\u0000zt\u0001\u0000"+
		"\u0000\u0000zu\u0001\u0000\u0000\u0000zv\u0001\u0000\u0000\u0000{\u000f"+
		"\u0001\u0000\u0000\u0000|}\u0005$\u0000\u0000}~\u0003\u001e\u000f\u0000"+
		"~\u007f\u0005\u0003\u0000\u0000\u007f\u0080\u0003\u0002\u0001\u0000\u0080"+
		"\u0089\u0005\u0004\u0000\u0000\u0081\u0087\u0005%\u0000\u0000\u0082\u0088"+
		"\u0003\u0010\b\u0000\u0083\u0084\u0005\u0003\u0000\u0000\u0084\u0085\u0003"+
		"\u0002\u0001\u0000\u0085\u0086\u0005\u0004\u0000\u0000\u0086\u0088\u0001"+
		"\u0000\u0000\u0000\u0087\u0082\u0001\u0000\u0000\u0000\u0087\u0083\u0001"+
		"\u0000\u0000\u0000\u0088\u008a\u0001\u0000\u0000\u0000\u0089\u0081\u0001"+
		"\u0000\u0000\u0000\u0089\u008a\u0001\u0000\u0000\u0000\u008a\u0011\u0001"+
		"\u0000\u0000\u0000\u008b\u008c\u0005-\u0000\u0000\u008c\u008d\u0003\u001e"+
		"\u000f\u0000\u008d\u008e\u0005%\u0000\u0000\u008e\u008f\u0005\u0003\u0000"+
		"\u0000\u008f\u0090\u0003\u0002\u0001\u0000\u0090\u0091\u0003\u0006\u0003"+
		"\u0000\u0091\u0092\u0005\u0004\u0000\u0000\u0092\u0013\u0001\u0000\u0000"+
		"\u0000\u0093\u0094\u0005&\u0000\u0000\u0094\u0095\u0003 \u0010\u0000\u0095"+
		"\u0099\u0005\u0003\u0000\u0000\u0096\u0098\u0003\u0016\u000b\u0000\u0097"+
		"\u0096\u0001\u0000\u0000\u0000\u0098\u009b\u0001\u0000\u0000\u0000\u0099"+
		"\u0097\u0001\u0000\u0000\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u009a"+
		"\u009d\u0001\u0000\u0000\u0000\u009b\u0099\u0001\u0000\u0000\u0000\u009c"+
		"\u009e\u0003\u0018\f\u0000\u009d\u009c\u0001\u0000\u0000\u0000\u009d\u009e"+
		"\u0001\u0000\u0000\u0000\u009e\u009f\u0001\u0000\u0000\u0000\u009f\u00a0"+
		"\u0005\u0004\u0000\u0000\u00a0\u0015\u0001\u0000\u0000\u0000\u00a1\u00a2"+
		"\u0005\'\u0000\u0000\u00a2\u00a3\u0003 \u0010\u0000\u00a3\u00a4\u0005"+
		"\u0007\u0000\u0000\u00a4\u00a5\u0003\u0002\u0001\u0000\u00a5\u0017\u0001"+
		"\u0000\u0000\u0000\u00a6\u00a7\u0005(\u0000\u0000\u00a7\u00a8\u0005\u0007"+
		"\u0000\u0000\u00a8\u00a9\u0003\u0002\u0001\u0000\u00a9\u0019\u0001\u0000"+
		"\u0000\u0000\u00aa\u00ab\u0005)\u0000\u0000\u00ab\u00ac\u0003\u001e\u000f"+
		"\u0000\u00ac\u00ad\u0005\u0003\u0000\u0000\u00ad\u00ae\u0003\u0002\u0001"+
		"\u0000\u00ae\u00af\u0005\u0004\u0000\u0000\u00af\u001b\u0001\u0000\u0000"+
		"\u0000\u00b0\u00b1\u0005*\u0000\u0000\u00b1\u00b2\u0005J\u0000\u0000\u00b2"+
		"\u00b3\u0005+\u0000\u0000\u00b3\u00b4\u0003 \u0010\u0000\u00b4\u00b5\u0005"+
		",\u0000\u0000\u00b5\u00b6\u0003 \u0010\u0000\u00b6\u00b7\u0005\u0003\u0000"+
		"\u0000\u00b7\u00b8\u0003\u0002\u0001\u0000\u00b8\u00b9\u0005\u0004\u0000"+
		"\u0000\u00b9\u00c3\u0001\u0000\u0000\u0000\u00ba\u00bb\u0005*\u0000\u0000"+
		"\u00bb\u00bc\u0005J\u0000\u0000\u00bc\u00bd\u0005+\u0000\u0000\u00bd\u00be"+
		"\u0003 \u0010\u0000\u00be\u00bf\u0005\u0003\u0000\u0000\u00bf\u00c0\u0003"+
		"\u0002\u0001\u0000\u00c0\u00c1\u0005\u0004\u0000\u0000\u00c1\u00c3\u0001"+
		"\u0000\u0000\u0000\u00c2\u00b0\u0001\u0000\u0000\u0000\u00c2\u00ba\u0001"+
		"\u0000\u0000\u0000\u00c3\u001d\u0001\u0000\u0000\u0000\u00c4\u00c5\u0006"+
		"\u000f\uffff\uffff\u0000\u00c5\u00c6\u0005\u001e\u0000\u0000\u00c6\u00d1"+
		"\u0003\u001e\u000f\u0005\u00c7\u00c8\u0003 \u0010\u0000\u00c8\u00c9\u0007"+
		"\u0002\u0000\u0000\u00c9\u00ca\u0003 \u0010\u0000\u00ca\u00d1\u0001\u0000"+
		"\u0000\u0000\u00cb\u00d1\u0007\u0003\u0000\u0000\u00cc\u00cd\u0005\u0002"+
		"\u0000\u0000\u00cd\u00ce\u0003\u001e\u000f\u0000\u00ce\u00cf\u0005\u0001"+
		"\u0000\u0000\u00cf\u00d1\u0001\u0000\u0000\u0000\u00d0\u00c4\u0001\u0000"+
		"\u0000\u0000\u00d0\u00c7\u0001\u0000\u0000\u0000\u00d0\u00cb\u0001\u0000"+
		"\u0000\u0000\u00d0\u00cc\u0001\u0000\u0000\u0000\u00d1\u00d7\u0001\u0000"+
		"\u0000\u0000\u00d2\u00d3\n\u0003\u0000\u0000\u00d3\u00d4\u0007\u0004\u0000"+
		"\u0000\u00d4\u00d6\u0003\u001e\u000f\u0004\u00d5\u00d2\u0001\u0000\u0000"+
		"\u0000\u00d6\u00d9\u0001\u0000\u0000\u0000\u00d7\u00d5\u0001\u0000\u0000"+
		"\u0000\u00d7\u00d8\u0001\u0000\u0000\u0000\u00d8\u001f\u0001\u0000\u0000"+
		"\u0000\u00d9\u00d7\u0001\u0000\u0000\u0000\u00da\u00db\u0006\u0010\uffff"+
		"\uffff\u0000\u00db\u00dc\u0005\u0013\u0000\u0000\u00dc\u00e9\u0003 \u0010"+
		"\f\u00dd\u00e9\u0007\u0003\u0000\u0000\u00de\u00e9\u00051\u0000\u0000"+
		"\u00df\u00e9\u0005J\u0000\u0000\u00e0\u00e9\u0005F\u0000\u0000\u00e1\u00e9"+
		"\u0005G\u0000\u0000\u00e2\u00e9\u0005I\u0000\u0000\u00e3\u00e9\u0005H"+
		"\u0000\u0000\u00e4\u00e5\u0005\u0002\u0000\u0000\u00e5\u00e6\u0003 \u0010"+
		"\u0000\u00e6\u00e7\u0005\u0001\u0000\u0000\u00e7\u00e9\u0001\u0000\u0000"+
		"\u0000\u00e8\u00da\u0001\u0000\u0000\u0000\u00e8\u00dd\u0001\u0000\u0000"+
		"\u0000\u00e8\u00de\u0001\u0000\u0000\u0000\u00e8\u00df\u0001\u0000\u0000"+
		"\u0000\u00e8\u00e0\u0001\u0000\u0000\u0000\u00e8\u00e1\u0001\u0000\u0000"+
		"\u0000\u00e8\u00e2\u0001\u0000\u0000\u0000\u00e8\u00e3\u0001\u0000\u0000"+
		"\u0000\u00e8\u00e4\u0001\u0000\u0000\u0000\u00e9\u00f5\u0001\u0000\u0000"+
		"\u0000\u00ea\u00eb\n\u000b\u0000\u0000\u00eb\u00ec\u0007\u0005\u0000\u0000"+
		"\u00ec\u00f4\u0003 \u0010\f\u00ed\u00ee\n\n\u0000\u0000\u00ee\u00ef\u0007"+
		"\u0006\u0000\u0000\u00ef\u00f4\u0003 \u0010\u000b\u00f0\u00f1\n\t\u0000"+
		"\u0000\u00f1\u00f2\u0005\u0012\u0000\u0000\u00f2\u00f4\u0003 \u0010\n"+
		"\u00f3\u00ea\u0001\u0000\u0000\u0000\u00f3\u00ed\u0001\u0000\u0000\u0000"+
		"\u00f3\u00f0\u0001\u0000\u0000\u0000\u00f4\u00f7\u0001\u0000\u0000\u0000"+
		"\u00f5\u00f3\u0001\u0000\u0000\u0000\u00f5\u00f6\u0001\u0000\u0000\u0000"+
		"\u00f6!\u0001\u0000\u0000\u0000\u00f7\u00f5\u0001\u0000\u0000\u0000\u0016"+
		"\',04;?EGPfnz\u0087\u0089\u0099\u009d\u00c2\u00d0\u00d7\u00e8\u00f3\u00f5";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}