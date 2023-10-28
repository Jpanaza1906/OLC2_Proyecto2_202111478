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
		RULE_dec_vector = 15, RULE_def_vector = 16, RULE_asig_vector = 17, RULE_func_vector = 18, 
		RULE_condicion = 19, RULE_e = 20;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "trans_sentencia", "print_sentencia", 
			"declaracion_sentencia", "asignacion_sentencia", "tipo", "if_sentencia", 
			"guard_sentencia", "switch_sentencia", "lcasos", "cdefault", "while_sentencia", 
			"for_sentencia", "dec_vector", "def_vector", "asig_vector", "func_vector", 
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
	}

	public final SContext s() throws RecognitionException {
		SContext _localctx = new SContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_s);
		try {
			_localctx = new SLSentenciasContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(42);
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
			setState(47);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(44);
					sentencia();
					}
					} 
				}
				setState(49);
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
	public static class S_Asignacion_VectorContext extends SentenciaContext {
		public Asig_vectorContext asig_vector() {
			return getRuleContext(Asig_vectorContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_Asignacion_VectorContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_AsignacionContext extends SentenciaContext {
		public Asignacion_sentenciaContext asignacion_sentencia() {
			return getRuleContext(Asignacion_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_AsignacionContext(SentenciaContext ctx) { copyFrom(ctx); }
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
	public static class S_Funcion_VectorContext extends SentenciaContext {
		public Func_vectorContext func_vector() {
			return getRuleContext(Func_vectorContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_Funcion_VectorContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_Declaracion_VectorContext extends SentenciaContext {
		public Dec_vectorContext dec_vector() {
			return getRuleContext(Dec_vectorContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_Declaracion_VectorContext(SentenciaContext ctx) { copyFrom(ctx); }
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
	@SuppressWarnings("CheckReturnValue")
	public static class S_DeclaracionContext extends SentenciaContext {
		public Declaracion_sentenciaContext declaracion_sentencia() {
			return getRuleContext(Declaracion_sentenciaContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_DeclaracionContext(SentenciaContext ctx) { copyFrom(ctx); }
	}

	public final SentenciaContext sentencia() throws RecognitionException {
		SentenciaContext _localctx = new SentenciaContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_sentencia);
		int _la;
		try {
			setState(83);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				_localctx = new S_PrintContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(50);
				print_sentencia();
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
			case 2:
				_localctx = new S_DeclaracionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(54);
				declaracion_sentencia();
				setState(56);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(55);
					match(PTCOMA);
					}
				}

				}
				break;
			case 3:
				_localctx = new S_AsignacionContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(58);
				asignacion_sentencia();
				setState(60);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(59);
					match(PTCOMA);
					}
				}

				}
				break;
			case 4:
				_localctx = new S_IfContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(62);
				if_sentencia();
				}
				break;
			case 5:
				_localctx = new S_SwitchContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(63);
				switch_sentencia();
				}
				break;
			case 6:
				_localctx = new S_GuardContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(64);
				guard_sentencia();
				}
				break;
			case 7:
				_localctx = new S_TransContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(65);
				trans_sentencia();
				setState(67);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(66);
					match(PTCOMA);
					}
				}

				}
				break;
			case 8:
				_localctx = new S_WhileContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(69);
				while_sentencia();
				}
				break;
			case 9:
				_localctx = new S_ForContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(70);
				for_sentencia();
				}
				break;
			case 10:
				_localctx = new S_Declaracion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(71);
				dec_vector();
				setState(73);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(72);
					match(PTCOMA);
					}
				}

				}
				break;
			case 11:
				_localctx = new S_Funcion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(75);
				func_vector();
				setState(77);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(76);
					match(PTCOMA);
					}
				}

				}
				break;
			case 12:
				_localctx = new S_Asignacion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(79);
				asig_vector();
				setState(81);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(80);
					match(PTCOMA);
					}
				}

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
		try {
			setState(91);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BREAK:
				_localctx = new BreakContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				match(BREAK);
				}
				break;
			case CONTINUE:
				_localctx = new ContinueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				match(CONTINUE);
				}
				break;
			case RETURN:
				_localctx = new ReturnContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(87);
				match(RETURN);
				setState(89);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
				case 1:
					{
					setState(88);
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
	}

	public final Print_sentenciaContext print_sentencia() throws RecognitionException {
		Print_sentenciaContext _localctx = new Print_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_print_sentencia);
		int _la;
		try {
			_localctx = new PrintContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(93);
			match(PRINT);
			setState(94);
			match(PARIZQ);
			setState(95);
			e(0);
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(96);
				match(COMA);
				setState(97);
				e(0);
				}
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(103);
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
	}

	public final Declaracion_sentenciaContext declaracion_sentencia() throws RecognitionException {
		Declaracion_sentenciaContext _localctx = new Declaracion_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_declaracion_sentencia);
		int _la;
		try {
			setState(122);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new Declaracion_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(105);
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
				setState(106);
				match(ID);
				setState(107);
				match(DOSPT);
				setState(108);
				tipo();
				setState(109);
				match(IGUAL);
				setState(110);
				e(0);
				}
				break;
			case 2:
				_localctx = new Declaracion_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
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
				setState(113);
				match(ID);
				setState(114);
				match(IGUAL);
				setState(115);
				e(0);
				}
				break;
			case 3:
				_localctx = new Declaracion_Tipo_noValContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(116);
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
				setState(117);
				match(ID);
				setState(118);
				match(DOSPT);
				setState(119);
				tipo();
				setState(120);
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
	}

	public final Asignacion_sentenciaContext asignacion_sentencia() throws RecognitionException {
		Asignacion_sentenciaContext _localctx = new Asignacion_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_asignacion_sentencia);
		int _la;
		try {
			setState(130);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				_localctx = new AsignacionContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(124);
				match(ID);
				setState(125);
				match(IGUAL);
				setState(126);
				e(0);
				}
				break;
			case 2:
				_localctx = new Asignacion_SumaRestaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(127);
				match(ID);
				setState(128);
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
				setState(129);
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_StructContext extends TipoContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public Tipo_StructContext(TipoContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_FloatContext extends TipoContext {
		public TerminalNode FLOAT() { return getToken(Tswift_GrammarNParser.FLOAT, 0); }
		public Tipo_FloatContext(TipoContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_CharacterContext extends TipoContext {
		public TerminalNode CHAR() { return getToken(Tswift_GrammarNParser.CHAR, 0); }
		public Tipo_CharacterContext(TipoContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_StringContext extends TipoContext {
		public TerminalNode STRING() { return getToken(Tswift_GrammarNParser.STRING, 0); }
		public Tipo_StringContext(TipoContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_BoolContext extends TipoContext {
		public TerminalNode BOOL() { return getToken(Tswift_GrammarNParser.BOOL, 0); }
		public Tipo_BoolContext(TipoContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Tipo_IntContext extends TipoContext {
		public TerminalNode INT() { return getToken(Tswift_GrammarNParser.INT, 0); }
		public Tipo_IntContext(TipoContext ctx) { copyFrom(ctx); }
	}

	public final TipoContext tipo() throws RecognitionException {
		TipoContext _localctx = new TipoContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_tipo);
		try {
			setState(142);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new Tipo_IntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(132);
				match(INT);
				}
				break;
			case FLOAT:
				_localctx = new Tipo_FloatContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(133);
				match(FLOAT);
				}
				break;
			case STRING:
				_localctx = new Tipo_StringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(134);
				match(STRING);
				}
				break;
			case BOOL:
				_localctx = new Tipo_BoolContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(135);
				match(BOOL);
				}
				break;
			case CHAR:
				_localctx = new Tipo_CharacterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(136);
				match(CHAR);
				}
				break;
			case ID:
				_localctx = new Tipo_StructContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(137);
				match(ID);
				}
				break;
			case CORCHETEIZQ:
				_localctx = new Tipo_VectorContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(138);
				match(CORCHETEIZQ);
				setState(139);
				tipo();
				setState(140);
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
	}

	public final If_sentenciaContext if_sentencia() throws RecognitionException {
		If_sentenciaContext _localctx = new If_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_if_sentencia);
		int _la;
		try {
			_localctx = new IfContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			match(IF);
			setState(145);
			condicion(0);
			setState(146);
			match(LLAVEIZQ);
			setState(147);
			l_sentencias();
			setState(148);
			match(LLAVEDER);
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(149);
				match(ELSE);
				setState(155);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case IF:
					{
					setState(150);
					if_sentencia();
					}
					break;
				case LLAVEIZQ:
					{
					setState(151);
					match(LLAVEIZQ);
					setState(152);
					l_sentencias();
					setState(153);
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
		enterRule(_localctx, 18, RULE_guard_sentencia);
		try {
			_localctx = new GuardContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(159);
			match(GUARD);
			setState(160);
			condicion(0);
			setState(161);
			match(ELSE);
			setState(162);
			match(LLAVEIZQ);
			setState(163);
			l_sentencias();
			setState(164);
			trans_sentencia();
			setState(165);
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
		enterRule(_localctx, 20, RULE_switch_sentencia);
		int _la;
		try {
			_localctx = new SwitchContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(167);
			match(SWITCH);
			setState(168);
			e(0);
			setState(169);
			match(LLAVEIZQ);
			setState(173);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(170);
				lcasos();
				}
				}
				setState(175);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(177);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(176);
				cdefault();
				}
			}

			setState(179);
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
		enterRule(_localctx, 22, RULE_lcasos);
		try {
			_localctx = new CaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(181);
			match(CASE);
			setState(182);
			e(0);
			setState(183);
			match(DOSPT);
			setState(184);
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
		enterRule(_localctx, 24, RULE_cdefault);
		try {
			_localctx = new DefaultContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(186);
			match(DEFAULT);
			setState(187);
			match(DOSPT);
			setState(188);
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
		enterRule(_localctx, 26, RULE_while_sentencia);
		try {
			_localctx = new WhileContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(190);
			match(WHILE);
			setState(191);
			condicion(0);
			setState(192);
			match(LLAVEIZQ);
			setState(193);
			l_sentencias();
			setState(194);
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
	}

	public final For_sentenciaContext for_sentencia() throws RecognitionException {
		For_sentenciaContext _localctx = new For_sentenciaContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_for_sentencia);
		try {
			setState(214);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				_localctx = new ForIntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(196);
				match(FOR);
				setState(197);
				match(ID);
				setState(198);
				match(IN);
				setState(199);
				e(0);
				setState(200);
				match(RANGO);
				setState(201);
				e(0);
				setState(202);
				match(LLAVEIZQ);
				setState(203);
				l_sentencias();
				setState(204);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new ForListContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(206);
				match(FOR);
				setState(207);
				match(ID);
				setState(208);
				match(IN);
				setState(209);
				e(0);
				setState(210);
				match(LLAVEIZQ);
				setState(211);
				l_sentencias();
				setState(212);
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
	public static class Dec_vectorContext extends ParserRuleContext {
		public Dec_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dec_vector; }
	 
		public Dec_vectorContext() { }
		public void copyFrom(Dec_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Declaracion_VectorContext extends Dec_vectorContext {
		public Token tipod;
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarNParser.DOSPT, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarNParser.IGUAL, 0); }
		public Def_vectorContext def_vector() {
			return getRuleContext(Def_vectorContext.class,0);
		}
		public TerminalNode VAR() { return getToken(Tswift_GrammarNParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarNParser.LET, 0); }
		public Declaracion_VectorContext(Dec_vectorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Declaracion_AlternaContext extends Dec_vectorContext {
		public Token tipod;
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarNParser.IGUAL, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public TerminalNode VAR() { return getToken(Tswift_GrammarNParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarNParser.LET, 0); }
		public Declaracion_AlternaContext(Dec_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Dec_vectorContext dec_vector() throws RecognitionException {
		Dec_vectorContext _localctx = new Dec_vectorContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_dec_vector);
		int _la;
		try {
			setState(234);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				_localctx = new Declaracion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(216);
				((Declaracion_VectorContext)_localctx).tipod = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
					((Declaracion_VectorContext)_localctx).tipod = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(217);
				match(ID);
				setState(218);
				match(DOSPT);
				setState(219);
				match(CORCHETEIZQ);
				setState(220);
				tipo();
				setState(221);
				match(CORCHETEDER);
				setState(222);
				match(IGUAL);
				setState(223);
				def_vector();
				}
				break;
			case 2:
				_localctx = new Declaracion_AlternaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(225);
				((Declaracion_AlternaContext)_localctx).tipod = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==VAR || _la==LET) ) {
					((Declaracion_AlternaContext)_localctx).tipod = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(226);
				match(ID);
				setState(227);
				match(IGUAL);
				setState(228);
				match(CORCHETEIZQ);
				setState(229);
				tipo();
				setState(230);
				match(CORCHETEDER);
				setState(231);
				match(PARIZQ);
				setState(232);
				match(PARDER);
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
	public static class Def_vectorContext extends ParserRuleContext {
		public Def_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_def_vector; }
	 
		public Def_vectorContext() { }
		public void copyFrom(Def_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Def_Vector_VacioContext extends Def_vectorContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public Def_Vector_VacioContext(Def_vectorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Def_VectorContext extends Def_vectorContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public List<TerminalNode> COMA() { return getTokens(Tswift_GrammarNParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(Tswift_GrammarNParser.COMA, i);
		}
		public Def_VectorContext(Def_vectorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Def_Vector_IdContext extends Def_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public Def_Vector_IdContext(Def_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Def_vectorContext def_vector() throws RecognitionException {
		Def_vectorContext _localctx = new Def_vectorContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_def_vector);
		int _la;
		try {
			setState(250);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				_localctx = new Def_VectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(236);
				match(CORCHETEIZQ);
				setState(237);
				e(0);
				setState(242);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMA) {
					{
					{
					setState(238);
					match(COMA);
					setState(239);
					e(0);
					}
					}
					setState(244);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(245);
				match(CORCHETEDER);
				}
				break;
			case 2:
				_localctx = new Def_Vector_VacioContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(247);
				match(CORCHETEIZQ);
				setState(248);
				match(CORCHETEDER);
				}
				break;
			case 3:
				_localctx = new Def_Vector_IdContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(249);
				match(ID);
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
	public static class Asig_vectorContext extends ParserRuleContext {
		public Asig_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asig_vector; }
	 
		public Asig_vectorContext() { }
		public void copyFrom(Asig_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Asig_VectorContext extends Asig_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarNParser.IGUAL, 0); }
		public Asig_VectorContext(Asig_vectorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ResAsg_VectorContext extends Asig_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public TerminalNode MENOSIGUAL() { return getToken(Tswift_GrammarNParser.MENOSIGUAL, 0); }
		public ResAsg_VectorContext(Asig_vectorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SumAsg_VectorContext extends Asig_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public TerminalNode MASIGUAL() { return getToken(Tswift_GrammarNParser.MASIGUAL, 0); }
		public SumAsg_VectorContext(Asig_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Asig_vectorContext asig_vector() throws RecognitionException {
		Asig_vectorContext _localctx = new Asig_vectorContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_asig_vector);
		try {
			setState(273);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				_localctx = new Asig_VectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(252);
				match(ID);
				setState(253);
				match(CORCHETEIZQ);
				setState(254);
				e(0);
				setState(255);
				match(CORCHETEDER);
				setState(256);
				match(IGUAL);
				setState(257);
				e(0);
				}
				break;
			case 2:
				_localctx = new SumAsg_VectorContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(259);
				match(ID);
				setState(260);
				match(CORCHETEIZQ);
				setState(261);
				e(0);
				setState(262);
				match(CORCHETEDER);
				setState(263);
				match(MASIGUAL);
				setState(264);
				e(0);
				}
				break;
			case 3:
				_localctx = new ResAsg_VectorContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(266);
				match(ID);
				setState(267);
				match(CORCHETEIZQ);
				setState(268);
				e(0);
				setState(269);
				match(CORCHETEDER);
				setState(270);
				match(MENOSIGUAL);
				setState(271);
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
	public static class Func_vectorContext extends ParserRuleContext {
		public Func_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_vector; }
	 
		public Func_vectorContext() { }
		public void copyFrom(Func_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Func_Vector_AppendContext extends Func_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarNParser.PUNTO, 0); }
		public TerminalNode APPEND() { return getToken(Tswift_GrammarNParser.APPEND, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public Func_Vector_AppendContext(Func_vectorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Func_Vector_RemoveContext extends Func_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarNParser.PUNTO, 0); }
		public TerminalNode REMOVE() { return getToken(Tswift_GrammarNParser.REMOVE, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public TerminalNode AT() { return getToken(Tswift_GrammarNParser.AT, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarNParser.DOSPT, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public Func_Vector_RemoveContext(Func_vectorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Func_Vector_RemoveLastContext extends Func_vectorContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarNParser.PUNTO, 0); }
		public TerminalNode REMOVELAST() { return getToken(Tswift_GrammarNParser.REMOVELAST, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public Func_Vector_RemoveLastContext(Func_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Func_vectorContext func_vector() throws RecognitionException {
		Func_vectorContext _localctx = new Func_vectorContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_func_vector);
		try {
			setState(296);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				_localctx = new Func_Vector_AppendContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(275);
				match(ID);
				setState(276);
				match(PUNTO);
				setState(277);
				match(APPEND);
				setState(278);
				match(PARIZQ);
				setState(279);
				e(0);
				setState(280);
				match(PARDER);
				}
				break;
			case 2:
				_localctx = new Func_Vector_RemoveLastContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(282);
				match(ID);
				setState(283);
				match(PUNTO);
				setState(284);
				match(REMOVELAST);
				setState(285);
				match(PARIZQ);
				setState(286);
				match(PARDER);
				}
				break;
			case 3:
				_localctx = new Func_Vector_RemoveContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(287);
				match(ID);
				setState(288);
				match(PUNTO);
				setState(289);
				match(REMOVE);
				setState(290);
				match(PARIZQ);
				setState(291);
				match(AT);
				setState(292);
				match(DOSPT);
				setState(293);
				e(0);
				setState(294);
				match(PARDER);
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
		int _startState = 38;
		enterRecursionRule(_localctx, 38, RULE_condicion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(310);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				{
				_localctx = new Cond_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(299);
				((Cond_NegContext)_localctx).op = match(NOT);
				setState(300);
				((Cond_NegContext)_localctx).c = condicion(5);
				}
				break;
			case 2:
				{
				_localctx = new Cond_RelContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(301);
				((Cond_RelContext)_localctx).e1 = e(0);
				setState(302);
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
				setState(303);
				((Cond_RelContext)_localctx).e2 = e(0);
				}
				break;
			case 3:
				{
				_localctx = new Cond_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(305);
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
				setState(306);
				match(PARIZQ);
				setState(307);
				((Cond_ParContext)_localctx).c = condicion(0);
				setState(308);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(317);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Cond_LogicaContext(new CondicionContext(_parentctx, _parentState));
					((Cond_LogicaContext)_localctx).c1 = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_condicion);
					setState(312);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(313);
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
					setState(314);
					((Cond_LogicaContext)_localctx).c2 = condicion(4);
					}
					} 
				}
				setState(319);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,26,_ctx);
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
	public static class Expr_RelContext extends EContext {
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
		public Expr_RelContext(EContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_DecimalContext extends EContext {
		public Token n;
		public TerminalNode DECIMAL() { return getToken(Tswift_GrammarNParser.DECIMAL, 0); }
		public Expr_DecimalContext(EContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_CaracterContext extends EContext {
		public Token n;
		public TerminalNode CARACTER() { return getToken(Tswift_GrammarNParser.CARACTER, 0); }
		public Expr_CaracterContext(EContext ctx) { copyFrom(ctx); }
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
	public static class Expr_NegContext extends EContext {
		public Token op;
		public EContext e1;
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode MENOS() { return getToken(Tswift_GrammarNParser.MENOS, 0); }
		public TerminalNode NOT() { return getToken(Tswift_GrammarNParser.NOT, 0); }
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
	public static class Expr_NilContext extends EContext {
		public Token n;
		public TerminalNode NIL() { return getToken(Tswift_GrammarNParser.NIL, 0); }
		public Expr_NilContext(EContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_CadenaContext extends EContext {
		public Token n;
		public TerminalNode CADENA() { return getToken(Tswift_GrammarNParser.CADENA, 0); }
		public Expr_CadenaContext(EContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_CountContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarNParser.PUNTO, 0); }
		public TerminalNode COUNT() { return getToken(Tswift_GrammarNParser.COUNT, 0); }
		public Expr_CountContext(EContext ctx) { copyFrom(ctx); }
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
	public static class Expr_LogicaContext extends EContext {
		public Token op;
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode AND() { return getToken(Tswift_GrammarNParser.AND, 0); }
		public TerminalNode OR() { return getToken(Tswift_GrammarNParser.OR, 0); }
		public Expr_LogicaContext(EContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_IsEmptyContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode PUNTO() { return getToken(Tswift_GrammarNParser.PUNTO, 0); }
		public TerminalNode ISEMPTY() { return getToken(Tswift_GrammarNParser.ISEMPTY, 0); }
		public Expr_IsEmptyContext(EContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expr_BooleanoContext extends EContext {
		public Token n;
		public TerminalNode TRUE() { return getToken(Tswift_GrammarNParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Tswift_GrammarNParser.FALSE, 0); }
		public Expr_BooleanoContext(EContext ctx) { copyFrom(ctx); }
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
		int _startState = 40;
		enterRecursionRule(_localctx, 40, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(340);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(321);
				((Expr_NegContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==MENOS || _la==NOT) ) {
					((Expr_NegContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(322);
				((Expr_NegContext)_localctx).e1 = e(16);
				}
				break;
			case 2:
				{
				_localctx = new Expr_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(323);
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
			case 3:
				{
				_localctx = new Expr_NilContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(324);
				((Expr_NilContext)_localctx).n = match(NIL);
				}
				break;
			case 4:
				{
				_localctx = new Expr_IsEmptyContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(325);
				match(ID);
				setState(326);
				match(PUNTO);
				setState(327);
				match(ISEMPTY);
				}
				break;
			case 5:
				{
				_localctx = new Expr_CountContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(328);
				match(ID);
				setState(329);
				match(PUNTO);
				setState(330);
				match(COUNT);
				}
				break;
			case 6:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(331);
				((Expr_IdContext)_localctx).id = match(ID);
				}
				break;
			case 7:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(332);
				((Expr_EnteroContext)_localctx).n = match(ENTERO);
				}
				break;
			case 8:
				{
				_localctx = new Expr_DecimalContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(333);
				((Expr_DecimalContext)_localctx).n = match(DECIMAL);
				}
				break;
			case 9:
				{
				_localctx = new Expr_CadenaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(334);
				((Expr_CadenaContext)_localctx).n = match(CADENA);
				}
				break;
			case 10:
				{
				_localctx = new Expr_CaracterContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(335);
				((Expr_CaracterContext)_localctx).n = match(CARACTER);
				}
				break;
			case 11:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(336);
				match(PARIZQ);
				setState(337);
				((Expr_ParContext)_localctx).e1 = e(0);
				setState(338);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(359);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(357);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						((Expr_MulDivContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(342);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(343);
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
						setState(344);
						((Expr_MulDivContext)_localctx).e2 = e(16);
						}
						break;
					case 2:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						((Expr_SumResContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(345);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(346);
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
						setState(347);
						((Expr_SumResContext)_localctx).e2 = e(15);
						}
						break;
					case 3:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						((Expr_ModContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(348);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(349);
						((Expr_ModContext)_localctx).op = match(MOD);
						setState(350);
						((Expr_ModContext)_localctx).e2 = e(14);
						}
						break;
					case 4:
						{
						_localctx = new Expr_RelContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(351);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(352);
						((Expr_RelContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 264241152L) != 0)) ) {
							((Expr_RelContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(353);
						e(13);
						}
						break;
					case 5:
						{
						_localctx = new Expr_LogicaContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(354);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(355);
						((Expr_LogicaContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==AND || _la==OR) ) {
							((Expr_LogicaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(356);
						e(12);
						}
						break;
					}
					} 
				}
				setState(361);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,29,_ctx);
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
		case 19:
			return condicion_sempred((CondicionContext)_localctx, predIndex);
		case 20:
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
			return precpred(_ctx, 15);
		case 2:
			return precpred(_ctx, 14);
		case 3:
			return precpred(_ctx, 13);
		case 4:
			return precpred(_ctx, 12);
		case 5:
			return precpred(_ctx, 11);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001M\u016b\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0005\u0001.\b\u0001\n\u0001\f\u00011\t\u0001\u0001\u0002"+
		"\u0001\u0002\u0003\u00025\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"9\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002=\b\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002D\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002J\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002N\b\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002R\b\u0002\u0003\u0002T\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0003\u0003Z\b\u0003\u0003\u0003\\\b\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004c\b\u0004\n\u0004"+
		"\f\u0004f\t\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0003\u0005{\b\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u0083"+
		"\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u008f"+
		"\b\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0003\b\u009c\b\b\u0003\b\u009e\b\b\u0001\t"+
		"\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0005\n\u00ac\b\n\n\n\f\n\u00af\t\n\u0001\n\u0003\n"+
		"\u00b2\b\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0003\u000e\u00d7\b\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00eb\b\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u00f1\b\u0010\n"+
		"\u0010\f\u0010\u00f4\t\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0003\u0010\u00fb\b\u0010\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0003\u0011\u0112\b\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003"+
		"\u0012\u0129\b\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0003\u0013\u0137\b\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0005\u0013\u013c\b\u0013\n\u0013\f\u0013\u013f\t\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0003\u0014\u0155\b\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0005\u0014\u0166\b\u0014\n\u0014\f\u0014\u0169\t\u0014\u0001\u0014\u0000"+
		"\u0002&(\u0015\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016"+
		"\u0018\u001a\u001c\u001e \"$&(\u0000\b\u0001\u0000 !\u0001\u0000\u000e"+
		"\u000f\u0001\u0000\u0016\u001b\u0001\u0000\"#\u0001\u0000\u001c\u001d"+
		"\u0002\u0000\u0013\u0013\u001e\u001e\u0002\u0000\u0011\u0011\u0015\u0015"+
		"\u0001\u0000\u0013\u0014\u0195\u0000*\u0001\u0000\u0000\u0000\u0002/\u0001"+
		"\u0000\u0000\u0000\u0004S\u0001\u0000\u0000\u0000\u0006[\u0001\u0000\u0000"+
		"\u0000\b]\u0001\u0000\u0000\u0000\nz\u0001\u0000\u0000\u0000\f\u0082\u0001"+
		"\u0000\u0000\u0000\u000e\u008e\u0001\u0000\u0000\u0000\u0010\u0090\u0001"+
		"\u0000\u0000\u0000\u0012\u009f\u0001\u0000\u0000\u0000\u0014\u00a7\u0001"+
		"\u0000\u0000\u0000\u0016\u00b5\u0001\u0000\u0000\u0000\u0018\u00ba\u0001"+
		"\u0000\u0000\u0000\u001a\u00be\u0001\u0000\u0000\u0000\u001c\u00d6\u0001"+
		"\u0000\u0000\u0000\u001e\u00ea\u0001\u0000\u0000\u0000 \u00fa\u0001\u0000"+
		"\u0000\u0000\"\u0111\u0001\u0000\u0000\u0000$\u0128\u0001\u0000\u0000"+
		"\u0000&\u0136\u0001\u0000\u0000\u0000(\u0154\u0001\u0000\u0000\u0000*"+
		"+\u0003\u0002\u0001\u0000+\u0001\u0001\u0000\u0000\u0000,.\u0003\u0004"+
		"\u0002\u0000-,\u0001\u0000\u0000\u0000.1\u0001\u0000\u0000\u0000/-\u0001"+
		"\u0000\u0000\u0000/0\u0001\u0000\u0000\u00000\u0003\u0001\u0000\u0000"+
		"\u00001/\u0001\u0000\u0000\u000024\u0003\b\u0004\u000035\u0005\b\u0000"+
		"\u000043\u0001\u0000\u0000\u000045\u0001\u0000\u0000\u00005T\u0001\u0000"+
		"\u0000\u000068\u0003\n\u0005\u000079\u0005\b\u0000\u000087\u0001\u0000"+
		"\u0000\u000089\u0001\u0000\u0000\u00009T\u0001\u0000\u0000\u0000:<\u0003"+
		"\f\u0006\u0000;=\u0005\b\u0000\u0000<;\u0001\u0000\u0000\u0000<=\u0001"+
		"\u0000\u0000\u0000=T\u0001\u0000\u0000\u0000>T\u0003\u0010\b\u0000?T\u0003"+
		"\u0014\n\u0000@T\u0003\u0012\t\u0000AC\u0003\u0006\u0003\u0000BD\u0005"+
		"\b\u0000\u0000CB\u0001\u0000\u0000\u0000CD\u0001\u0000\u0000\u0000DT\u0001"+
		"\u0000\u0000\u0000ET\u0003\u001a\r\u0000FT\u0003\u001c\u000e\u0000GI\u0003"+
		"\u001e\u000f\u0000HJ\u0005\b\u0000\u0000IH\u0001\u0000\u0000\u0000IJ\u0001"+
		"\u0000\u0000\u0000JT\u0001\u0000\u0000\u0000KM\u0003$\u0012\u0000LN\u0005"+
		"\b\u0000\u0000ML\u0001\u0000\u0000\u0000MN\u0001\u0000\u0000\u0000NT\u0001"+
		"\u0000\u0000\u0000OQ\u0003\"\u0011\u0000PR\u0005\b\u0000\u0000QP\u0001"+
		"\u0000\u0000\u0000QR\u0001\u0000\u0000\u0000RT\u0001\u0000\u0000\u0000"+
		"S2\u0001\u0000\u0000\u0000S6\u0001\u0000\u0000\u0000S:\u0001\u0000\u0000"+
		"\u0000S>\u0001\u0000\u0000\u0000S?\u0001\u0000\u0000\u0000S@\u0001\u0000"+
		"\u0000\u0000SA\u0001\u0000\u0000\u0000SE\u0001\u0000\u0000\u0000SF\u0001"+
		"\u0000\u0000\u0000SG\u0001\u0000\u0000\u0000SK\u0001\u0000\u0000\u0000"+
		"SO\u0001\u0000\u0000\u0000T\u0005\u0001\u0000\u0000\u0000U\\\u00050\u0000"+
		"\u0000V\\\u0005.\u0000\u0000WY\u0005/\u0000\u0000XZ\u0003(\u0014\u0000"+
		"YX\u0001\u0000\u0000\u0000YZ\u0001\u0000\u0000\u0000Z\\\u0001\u0000\u0000"+
		"\u0000[U\u0001\u0000\u0000\u0000[V\u0001\u0000\u0000\u0000[W\u0001\u0000"+
		"\u0000\u0000\\\u0007\u0001\u0000\u0000\u0000]^\u0005\u001f\u0000\u0000"+
		"^_\u0005\u0002\u0000\u0000_d\u0003(\u0014\u0000`a\u0005\n\u0000\u0000"+
		"ac\u0003(\u0014\u0000b`\u0001\u0000\u0000\u0000cf\u0001\u0000\u0000\u0000"+
		"db\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000eg\u0001\u0000\u0000"+
		"\u0000fd\u0001\u0000\u0000\u0000gh\u0005\u0001\u0000\u0000h\t\u0001\u0000"+
		"\u0000\u0000ij\u0007\u0000\u0000\u0000jk\u0005J\u0000\u0000kl\u0005\u0007"+
		"\u0000\u0000lm\u0003\u000e\u0007\u0000mn\u0005\u0010\u0000\u0000no\u0003"+
		"(\u0014\u0000o{\u0001\u0000\u0000\u0000pq\u0007\u0000\u0000\u0000qr\u0005"+
		"J\u0000\u0000rs\u0005\u0010\u0000\u0000s{\u0003(\u0014\u0000tu\u0007\u0000"+
		"\u0000\u0000uv\u0005J\u0000\u0000vw\u0005\u0007\u0000\u0000wx\u0003\u000e"+
		"\u0007\u0000xy\u0005\t\u0000\u0000y{\u0001\u0000\u0000\u0000zi\u0001\u0000"+
		"\u0000\u0000zp\u0001\u0000\u0000\u0000zt\u0001\u0000\u0000\u0000{\u000b"+
		"\u0001\u0000\u0000\u0000|}\u0005J\u0000\u0000}~\u0005\u0010\u0000\u0000"+
		"~\u0083\u0003(\u0014\u0000\u007f\u0080\u0005J\u0000\u0000\u0080\u0081"+
		"\u0007\u0001\u0000\u0000\u0081\u0083\u0003(\u0014\u0000\u0082|\u0001\u0000"+
		"\u0000\u0000\u0082\u007f\u0001\u0000\u0000\u0000\u0083\r\u0001\u0000\u0000"+
		"\u0000\u0084\u008f\u0005@\u0000\u0000\u0085\u008f\u0005A\u0000\u0000\u0086"+
		"\u008f\u0005D\u0000\u0000\u0087\u008f\u0005B\u0000\u0000\u0088\u008f\u0005"+
		"C\u0000\u0000\u0089\u008f\u0005J\u0000\u0000\u008a\u008b\u0005\u0005\u0000"+
		"\u0000\u008b\u008c\u0003\u000e\u0007\u0000\u008c\u008d\u0005\u0006\u0000"+
		"\u0000\u008d\u008f\u0001\u0000\u0000\u0000\u008e\u0084\u0001\u0000\u0000"+
		"\u0000\u008e\u0085\u0001\u0000\u0000\u0000\u008e\u0086\u0001\u0000\u0000"+
		"\u0000\u008e\u0087\u0001\u0000\u0000\u0000\u008e\u0088\u0001\u0000\u0000"+
		"\u0000\u008e\u0089\u0001\u0000\u0000\u0000\u008e\u008a\u0001\u0000\u0000"+
		"\u0000\u008f\u000f\u0001\u0000\u0000\u0000\u0090\u0091\u0005$\u0000\u0000"+
		"\u0091\u0092\u0003&\u0013\u0000\u0092\u0093\u0005\u0003\u0000\u0000\u0093"+
		"\u0094\u0003\u0002\u0001\u0000\u0094\u009d\u0005\u0004\u0000\u0000\u0095"+
		"\u009b\u0005%\u0000\u0000\u0096\u009c\u0003\u0010\b\u0000\u0097\u0098"+
		"\u0005\u0003\u0000\u0000\u0098\u0099\u0003\u0002\u0001\u0000\u0099\u009a"+
		"\u0005\u0004\u0000\u0000\u009a\u009c\u0001\u0000\u0000\u0000\u009b\u0096"+
		"\u0001\u0000\u0000\u0000\u009b\u0097\u0001\u0000\u0000\u0000\u009c\u009e"+
		"\u0001\u0000\u0000\u0000\u009d\u0095\u0001\u0000\u0000\u0000\u009d\u009e"+
		"\u0001\u0000\u0000\u0000\u009e\u0011\u0001\u0000\u0000\u0000\u009f\u00a0"+
		"\u0005-\u0000\u0000\u00a0\u00a1\u0003&\u0013\u0000\u00a1\u00a2\u0005%"+
		"\u0000\u0000\u00a2\u00a3\u0005\u0003\u0000\u0000\u00a3\u00a4\u0003\u0002"+
		"\u0001\u0000\u00a4\u00a5\u0003\u0006\u0003\u0000\u00a5\u00a6\u0005\u0004"+
		"\u0000\u0000\u00a6\u0013\u0001\u0000\u0000\u0000\u00a7\u00a8\u0005&\u0000"+
		"\u0000\u00a8\u00a9\u0003(\u0014\u0000\u00a9\u00ad\u0005\u0003\u0000\u0000"+
		"\u00aa\u00ac\u0003\u0016\u000b\u0000\u00ab\u00aa\u0001\u0000\u0000\u0000"+
		"\u00ac\u00af\u0001\u0000\u0000\u0000\u00ad\u00ab\u0001\u0000\u0000\u0000"+
		"\u00ad\u00ae\u0001\u0000\u0000\u0000\u00ae\u00b1\u0001\u0000\u0000\u0000"+
		"\u00af\u00ad\u0001\u0000\u0000\u0000\u00b0\u00b2\u0003\u0018\f\u0000\u00b1"+
		"\u00b0\u0001\u0000\u0000\u0000\u00b1\u00b2\u0001\u0000\u0000\u0000\u00b2"+
		"\u00b3\u0001\u0000\u0000\u0000\u00b3\u00b4\u0005\u0004\u0000\u0000\u00b4"+
		"\u0015\u0001\u0000\u0000\u0000\u00b5\u00b6\u0005\'\u0000\u0000\u00b6\u00b7"+
		"\u0003(\u0014\u0000\u00b7\u00b8\u0005\u0007\u0000\u0000\u00b8\u00b9\u0003"+
		"\u0002\u0001\u0000\u00b9\u0017\u0001\u0000\u0000\u0000\u00ba\u00bb\u0005"+
		"(\u0000\u0000\u00bb\u00bc\u0005\u0007\u0000\u0000\u00bc\u00bd\u0003\u0002"+
		"\u0001\u0000\u00bd\u0019\u0001\u0000\u0000\u0000\u00be\u00bf\u0005)\u0000"+
		"\u0000\u00bf\u00c0\u0003&\u0013\u0000\u00c0\u00c1\u0005\u0003\u0000\u0000"+
		"\u00c1\u00c2\u0003\u0002\u0001\u0000\u00c2\u00c3\u0005\u0004\u0000\u0000"+
		"\u00c3\u001b\u0001\u0000\u0000\u0000\u00c4\u00c5\u0005*\u0000\u0000\u00c5"+
		"\u00c6\u0005J\u0000\u0000\u00c6\u00c7\u0005+\u0000\u0000\u00c7\u00c8\u0003"+
		"(\u0014\u0000\u00c8\u00c9\u0005,\u0000\u0000\u00c9\u00ca\u0003(\u0014"+
		"\u0000\u00ca\u00cb\u0005\u0003\u0000\u0000\u00cb\u00cc\u0003\u0002\u0001"+
		"\u0000\u00cc\u00cd\u0005\u0004\u0000\u0000\u00cd\u00d7\u0001\u0000\u0000"+
		"\u0000\u00ce\u00cf\u0005*\u0000\u0000\u00cf\u00d0\u0005J\u0000\u0000\u00d0"+
		"\u00d1\u0005+\u0000\u0000\u00d1\u00d2\u0003(\u0014\u0000\u00d2\u00d3\u0005"+
		"\u0003\u0000\u0000\u00d3\u00d4\u0003\u0002\u0001\u0000\u00d4\u00d5\u0005"+
		"\u0004\u0000\u0000\u00d5\u00d7\u0001\u0000\u0000\u0000\u00d6\u00c4\u0001"+
		"\u0000\u0000\u0000\u00d6\u00ce\u0001\u0000\u0000\u0000\u00d7\u001d\u0001"+
		"\u0000\u0000\u0000\u00d8\u00d9\u0007\u0000\u0000\u0000\u00d9\u00da\u0005"+
		"J\u0000\u0000\u00da\u00db\u0005\u0007\u0000\u0000\u00db\u00dc\u0005\u0005"+
		"\u0000\u0000\u00dc\u00dd\u0003\u000e\u0007\u0000\u00dd\u00de\u0005\u0006"+
		"\u0000\u0000\u00de\u00df\u0005\u0010\u0000\u0000\u00df\u00e0\u0003 \u0010"+
		"\u0000\u00e0\u00eb\u0001\u0000\u0000\u0000\u00e1\u00e2\u0007\u0000\u0000"+
		"\u0000\u00e2\u00e3\u0005J\u0000\u0000\u00e3\u00e4\u0005\u0010\u0000\u0000"+
		"\u00e4\u00e5\u0005\u0005\u0000\u0000\u00e5\u00e6\u0003\u000e\u0007\u0000"+
		"\u00e6\u00e7\u0005\u0006\u0000\u0000\u00e7\u00e8\u0005\u0002\u0000\u0000"+
		"\u00e8\u00e9\u0005\u0001\u0000\u0000\u00e9\u00eb\u0001\u0000\u0000\u0000"+
		"\u00ea\u00d8\u0001\u0000\u0000\u0000\u00ea\u00e1\u0001\u0000\u0000\u0000"+
		"\u00eb\u001f\u0001\u0000\u0000\u0000\u00ec\u00ed\u0005\u0005\u0000\u0000"+
		"\u00ed\u00f2\u0003(\u0014\u0000\u00ee\u00ef\u0005\n\u0000\u0000\u00ef"+
		"\u00f1\u0003(\u0014\u0000\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f1\u00f4"+
		"\u0001\u0000\u0000\u0000\u00f2\u00f0\u0001\u0000\u0000\u0000\u00f2\u00f3"+
		"\u0001\u0000\u0000\u0000\u00f3\u00f5\u0001\u0000\u0000\u0000\u00f4\u00f2"+
		"\u0001\u0000\u0000\u0000\u00f5\u00f6\u0005\u0006\u0000\u0000\u00f6\u00fb"+
		"\u0001\u0000\u0000\u0000\u00f7\u00f8\u0005\u0005\u0000\u0000\u00f8\u00fb"+
		"\u0005\u0006\u0000\u0000\u00f9\u00fb\u0005J\u0000\u0000\u00fa\u00ec\u0001"+
		"\u0000\u0000\u0000\u00fa\u00f7\u0001\u0000\u0000\u0000\u00fa\u00f9\u0001"+
		"\u0000\u0000\u0000\u00fb!\u0001\u0000\u0000\u0000\u00fc\u00fd\u0005J\u0000"+
		"\u0000\u00fd\u00fe\u0005\u0005\u0000\u0000\u00fe\u00ff\u0003(\u0014\u0000"+
		"\u00ff\u0100\u0005\u0006\u0000\u0000\u0100\u0101\u0005\u0010\u0000\u0000"+
		"\u0101\u0102\u0003(\u0014\u0000\u0102\u0112\u0001\u0000\u0000\u0000\u0103"+
		"\u0104\u0005J\u0000\u0000\u0104\u0105\u0005\u0005\u0000\u0000\u0105\u0106"+
		"\u0003(\u0014\u0000\u0106\u0107\u0005\u0006\u0000\u0000\u0107\u0108\u0005"+
		"\u000e\u0000\u0000\u0108\u0109\u0003(\u0014\u0000\u0109\u0112\u0001\u0000"+
		"\u0000\u0000\u010a\u010b\u0005J\u0000\u0000\u010b\u010c\u0005\u0005\u0000"+
		"\u0000\u010c\u010d\u0003(\u0014\u0000\u010d\u010e\u0005\u0006\u0000\u0000"+
		"\u010e\u010f\u0005\u000f\u0000\u0000\u010f\u0110\u0003(\u0014\u0000\u0110"+
		"\u0112\u0001\u0000\u0000\u0000\u0111\u00fc\u0001\u0000\u0000\u0000\u0111"+
		"\u0103\u0001\u0000\u0000\u0000\u0111\u010a\u0001\u0000\u0000\u0000\u0112"+
		"#\u0001\u0000\u0000\u0000\u0113\u0114\u0005J\u0000\u0000\u0114\u0115\u0005"+
		"\u000b\u0000\u0000\u0115\u0116\u00052\u0000\u0000\u0116\u0117\u0005\u0002"+
		"\u0000\u0000\u0117\u0118\u0003(\u0014\u0000\u0118\u0119\u0005\u0001\u0000"+
		"\u0000\u0119\u0129\u0001\u0000\u0000\u0000\u011a\u011b\u0005J\u0000\u0000"+
		"\u011b\u011c\u0005\u000b\u0000\u0000\u011c\u011d\u00053\u0000\u0000\u011d"+
		"\u011e\u0005\u0002\u0000\u0000\u011e\u0129\u0005\u0001\u0000\u0000\u011f"+
		"\u0120\u0005J\u0000\u0000\u0120\u0121\u0005\u000b\u0000\u0000\u0121\u0122"+
		"\u00054\u0000\u0000\u0122\u0123\u0005\u0002\u0000\u0000\u0123\u0124\u0005"+
		"5\u0000\u0000\u0124\u0125\u0005\u0007\u0000\u0000\u0125\u0126\u0003(\u0014"+
		"\u0000\u0126\u0127\u0005\u0001\u0000\u0000\u0127\u0129\u0001\u0000\u0000"+
		"\u0000\u0128\u0113\u0001\u0000\u0000\u0000\u0128\u011a\u0001\u0000\u0000"+
		"\u0000\u0128\u011f\u0001\u0000\u0000\u0000\u0129%\u0001\u0000\u0000\u0000"+
		"\u012a\u012b\u0006\u0013\uffff\uffff\u0000\u012b\u012c\u0005\u001e\u0000"+
		"\u0000\u012c\u0137\u0003&\u0013\u0005\u012d\u012e\u0003(\u0014\u0000\u012e"+
		"\u012f\u0007\u0002\u0000\u0000\u012f\u0130\u0003(\u0014\u0000\u0130\u0137"+
		"\u0001\u0000\u0000\u0000\u0131\u0137\u0007\u0003\u0000\u0000\u0132\u0133"+
		"\u0005\u0002\u0000\u0000\u0133\u0134\u0003&\u0013\u0000\u0134\u0135\u0005"+
		"\u0001\u0000\u0000\u0135\u0137\u0001\u0000\u0000\u0000\u0136\u012a\u0001"+
		"\u0000\u0000\u0000\u0136\u012d\u0001\u0000\u0000\u0000\u0136\u0131\u0001"+
		"\u0000\u0000\u0000\u0136\u0132\u0001\u0000\u0000\u0000\u0137\u013d\u0001"+
		"\u0000\u0000\u0000\u0138\u0139\n\u0003\u0000\u0000\u0139\u013a\u0007\u0004"+
		"\u0000\u0000\u013a\u013c\u0003&\u0013\u0004\u013b\u0138\u0001\u0000\u0000"+
		"\u0000\u013c\u013f\u0001\u0000\u0000\u0000\u013d\u013b\u0001\u0000\u0000"+
		"\u0000\u013d\u013e\u0001\u0000\u0000\u0000\u013e\'\u0001\u0000\u0000\u0000"+
		"\u013f\u013d\u0001\u0000\u0000\u0000\u0140\u0141\u0006\u0014\uffff\uffff"+
		"\u0000\u0141\u0142\u0007\u0005\u0000\u0000\u0142\u0155\u0003(\u0014\u0010"+
		"\u0143\u0155\u0007\u0003\u0000\u0000\u0144\u0155\u00051\u0000\u0000\u0145"+
		"\u0146\u0005J\u0000\u0000\u0146\u0147\u0005\u000b\u0000\u0000\u0147\u0155"+
		"\u00056\u0000\u0000\u0148\u0149\u0005J\u0000\u0000\u0149\u014a\u0005\u000b"+
		"\u0000\u0000\u014a\u0155\u00057\u0000\u0000\u014b\u0155\u0005J\u0000\u0000"+
		"\u014c\u0155\u0005F\u0000\u0000\u014d\u0155\u0005G\u0000\u0000\u014e\u0155"+
		"\u0005I\u0000\u0000\u014f\u0155\u0005H\u0000\u0000\u0150\u0151\u0005\u0002"+
		"\u0000\u0000\u0151\u0152\u0003(\u0014\u0000\u0152\u0153\u0005\u0001\u0000"+
		"\u0000\u0153\u0155\u0001\u0000\u0000\u0000\u0154\u0140\u0001\u0000\u0000"+
		"\u0000\u0154\u0143\u0001\u0000\u0000\u0000\u0154\u0144\u0001\u0000\u0000"+
		"\u0000\u0154\u0145\u0001\u0000\u0000\u0000\u0154\u0148\u0001\u0000\u0000"+
		"\u0000\u0154\u014b\u0001\u0000\u0000\u0000\u0154\u014c\u0001\u0000\u0000"+
		"\u0000\u0154\u014d\u0001\u0000\u0000\u0000\u0154\u014e\u0001\u0000\u0000"+
		"\u0000\u0154\u014f\u0001\u0000\u0000\u0000\u0154\u0150\u0001\u0000\u0000"+
		"\u0000\u0155\u0167\u0001\u0000\u0000\u0000\u0156\u0157\n\u000f\u0000\u0000"+
		"\u0157\u0158\u0007\u0006\u0000\u0000\u0158\u0166\u0003(\u0014\u0010\u0159"+
		"\u015a\n\u000e\u0000\u0000\u015a\u015b\u0007\u0007\u0000\u0000\u015b\u0166"+
		"\u0003(\u0014\u000f\u015c\u015d\n\r\u0000\u0000\u015d\u015e\u0005\u0012"+
		"\u0000\u0000\u015e\u0166\u0003(\u0014\u000e\u015f\u0160\n\f\u0000\u0000"+
		"\u0160\u0161\u0007\u0002\u0000\u0000\u0161\u0166\u0003(\u0014\r\u0162"+
		"\u0163\n\u000b\u0000\u0000\u0163\u0164\u0007\u0004\u0000\u0000\u0164\u0166"+
		"\u0003(\u0014\f\u0165\u0156\u0001\u0000\u0000\u0000\u0165\u0159\u0001"+
		"\u0000\u0000\u0000\u0165\u015c\u0001\u0000\u0000\u0000\u0165\u015f\u0001"+
		"\u0000\u0000\u0000\u0165\u0162\u0001\u0000\u0000\u0000\u0166\u0169\u0001"+
		"\u0000\u0000\u0000\u0167\u0165\u0001\u0000\u0000\u0000\u0167\u0168\u0001"+
		"\u0000\u0000\u0000\u0168)\u0001\u0000\u0000\u0000\u0169\u0167\u0001\u0000"+
		"\u0000\u0000\u001e/48<CIMQSY[dz\u0082\u008e\u009b\u009d\u00ad\u00b1\u00d6"+
		"\u00ea\u00f2\u00fa\u0111\u0128\u0136\u013d\u0154\u0165\u0167";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}