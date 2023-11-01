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
		RULE_dec_vector = 15, RULE_def_vector = 16, RULE_func_vector = 17, RULE_dec_matriz = 18, 
		RULE_def_matriz = 19, RULE_listavalores_matriz = 20, RULE_simple_matriz = 21, 
		RULE_condicion = 22, RULE_e = 23;
	private static String[] makeRuleNames() {
		return new String[] {
			"s", "l_sentencias", "sentencia", "trans_sentencia", "print_sentencia", 
			"declaracion_sentencia", "asignacion_sentencia", "tipo", "if_sentencia", 
			"guard_sentencia", "switch_sentencia", "lcasos", "cdefault", "while_sentencia", 
			"for_sentencia", "dec_vector", "def_vector", "func_vector", "dec_matriz", 
			"def_matriz", "listavalores_matriz", "simple_matriz", "condicion", "e"
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
			setState(48);
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
			setState(53);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(50);
					sentencia();
					}
					} 
				}
				setState(55);
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
	public static class S_Declaracion_MatrizContext extends SentenciaContext {
		public Dec_matrizContext dec_matriz() {
			return getRuleContext(Dec_matrizContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Tswift_GrammarNParser.PTCOMA, 0); }
		public S_Declaracion_MatrizContext(SentenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class S_SwitchContext extends SentenciaContext {
		public Switch_sentenciaContext switch_sentencia() {
			return getRuleContext(Switch_sentenciaContext.class,0);
		}
		public S_SwitchContext(SentenciaContext ctx) { copyFrom(ctx); }
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
			setState(89);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				_localctx = new S_PrintContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(56);
				print_sentencia();
				setState(58);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(57);
					match(PTCOMA);
					}
				}

				}
				break;
			case 2:
				_localctx = new S_DeclaracionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(60);
				declaracion_sentencia();
				setState(62);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(61);
					match(PTCOMA);
					}
				}

				}
				break;
			case 3:
				_localctx = new S_AsignacionContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(64);
				asignacion_sentencia();
				setState(66);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(65);
					match(PTCOMA);
					}
				}

				}
				break;
			case 4:
				_localctx = new S_IfContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(68);
				if_sentencia();
				}
				break;
			case 5:
				_localctx = new S_SwitchContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(69);
				switch_sentencia();
				}
				break;
			case 6:
				_localctx = new S_GuardContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(70);
				guard_sentencia();
				}
				break;
			case 7:
				_localctx = new S_TransContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(71);
				trans_sentencia();
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
			case 8:
				_localctx = new S_WhileContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(75);
				while_sentencia();
				}
				break;
			case 9:
				_localctx = new S_ForContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(76);
				for_sentencia();
				}
				break;
			case 10:
				_localctx = new S_Declaracion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(77);
				dec_vector();
				setState(79);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(78);
					match(PTCOMA);
					}
				}

				}
				break;
			case 11:
				_localctx = new S_Funcion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(81);
				func_vector();
				setState(83);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(82);
					match(PTCOMA);
					}
				}

				}
				break;
			case 12:
				_localctx = new S_Declaracion_MatrizContext(_localctx);
				enterOuterAlt(_localctx, 12);
				{
				setState(85);
				dec_matriz();
				setState(87);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PTCOMA) {
					{
					setState(86);
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
			setState(97);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BREAK:
				_localctx = new BreakContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(91);
				match(BREAK);
				}
				break;
			case CONTINUE:
				_localctx = new ContinueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(92);
				match(CONTINUE);
				}
				break;
			case RETURN:
				_localctx = new ReturnContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(93);
				match(RETURN);
				setState(95);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
				case 1:
					{
					setState(94);
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
			setState(99);
			match(PRINT);
			setState(100);
			match(PARIZQ);
			setState(101);
			e(0);
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMA) {
				{
				{
				setState(102);
				match(COMA);
				setState(103);
				e(0);
				}
				}
				setState(108);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(109);
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
			setState(128);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				_localctx = new Declaracion_Tipo_ValContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
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
				setState(112);
				match(ID);
				setState(113);
				match(DOSPT);
				setState(114);
				tipo();
				setState(115);
				match(IGUAL);
				setState(116);
				e(0);
				}
				break;
			case 2:
				_localctx = new Declaracion_ValContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(118);
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
				setState(119);
				match(ID);
				setState(120);
				match(IGUAL);
				setState(121);
				e(0);
				}
				break;
			case 3:
				_localctx = new Declaracion_Tipo_noValContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(122);
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
				setState(123);
				match(ID);
				setState(124);
				match(DOSPT);
				setState(125);
				tipo();
				setState(126);
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
			setState(136);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				_localctx = new AsignacionContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(130);
				match(ID);
				setState(131);
				match(IGUAL);
				setState(132);
				e(0);
				}
				break;
			case 2:
				_localctx = new Asignacion_SumaRestaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(133);
				match(ID);
				setState(134);
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
				setState(135);
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
			setState(148);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				_localctx = new Tipo_IntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(138);
				match(INT);
				}
				break;
			case FLOAT:
				_localctx = new Tipo_FloatContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(139);
				match(FLOAT);
				}
				break;
			case STRING:
				_localctx = new Tipo_StringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(140);
				match(STRING);
				}
				break;
			case BOOL:
				_localctx = new Tipo_BoolContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(141);
				match(BOOL);
				}
				break;
			case CHAR:
				_localctx = new Tipo_CharacterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(142);
				match(CHAR);
				}
				break;
			case ID:
				_localctx = new Tipo_StructContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(143);
				match(ID);
				}
				break;
			case CORCHETEIZQ:
				_localctx = new Tipo_VectorContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(144);
				match(CORCHETEIZQ);
				setState(145);
				tipo();
				setState(146);
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
			setState(150);
			match(IF);
			setState(151);
			condicion(0);
			setState(152);
			match(LLAVEIZQ);
			setState(153);
			l_sentencias();
			setState(154);
			match(LLAVEDER);
			setState(163);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(155);
				match(ELSE);
				setState(161);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case IF:
					{
					setState(156);
					if_sentencia();
					}
					break;
				case LLAVEIZQ:
					{
					setState(157);
					match(LLAVEIZQ);
					setState(158);
					l_sentencias();
					setState(159);
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
			setState(165);
			match(GUARD);
			setState(166);
			condicion(0);
			setState(167);
			match(ELSE);
			setState(168);
			match(LLAVEIZQ);
			setState(169);
			l_sentencias();
			setState(170);
			trans_sentencia();
			setState(171);
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
			setState(173);
			match(SWITCH);
			setState(174);
			e(0);
			setState(175);
			match(LLAVEIZQ);
			setState(179);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE) {
				{
				{
				setState(176);
				lcasos();
				}
				}
				setState(181);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(183);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT) {
				{
				setState(182);
				cdefault();
				}
			}

			setState(185);
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
			setState(187);
			match(CASE);
			setState(188);
			e(0);
			setState(189);
			match(DOSPT);
			setState(190);
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
			setState(192);
			match(DEFAULT);
			setState(193);
			match(DOSPT);
			setState(194);
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
			setState(196);
			match(WHILE);
			setState(197);
			condicion(0);
			setState(198);
			match(LLAVEIZQ);
			setState(199);
			l_sentencias();
			setState(200);
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
			setState(220);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				_localctx = new ForIntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(202);
				match(FOR);
				setState(203);
				match(ID);
				setState(204);
				match(IN);
				setState(205);
				e(0);
				setState(206);
				match(RANGO);
				setState(207);
				e(0);
				setState(208);
				match(LLAVEIZQ);
				setState(209);
				l_sentencias();
				setState(210);
				match(LLAVEDER);
				}
				break;
			case 2:
				_localctx = new ForListContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(212);
				match(FOR);
				setState(213);
				match(ID);
				setState(214);
				match(IN);
				setState(215);
				e(0);
				setState(216);
				match(LLAVEIZQ);
				setState(217);
				l_sentencias();
				setState(218);
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
			setState(240);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				_localctx = new Declaracion_VectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(222);
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
				setState(223);
				match(ID);
				setState(224);
				match(DOSPT);
				setState(225);
				match(CORCHETEIZQ);
				setState(226);
				tipo();
				setState(227);
				match(CORCHETEDER);
				setState(228);
				match(IGUAL);
				setState(229);
				def_vector();
				}
				break;
			case 2:
				_localctx = new Declaracion_AlternaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(231);
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
				setState(232);
				match(ID);
				setState(233);
				match(IGUAL);
				setState(234);
				match(CORCHETEIZQ);
				setState(235);
				tipo();
				setState(236);
				match(CORCHETEDER);
				setState(237);
				match(PARIZQ);
				setState(238);
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
			setState(256);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				_localctx = new Def_VectorContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(242);
				match(CORCHETEIZQ);
				setState(243);
				e(0);
				setState(248);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMA) {
					{
					{
					setState(244);
					match(COMA);
					setState(245);
					e(0);
					}
					}
					setState(250);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(251);
				match(CORCHETEDER);
				}
				break;
			case 2:
				_localctx = new Def_Vector_VacioContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(253);
				match(CORCHETEIZQ);
				setState(254);
				match(CORCHETEDER);
				}
				break;
			case 3:
				_localctx = new Def_Vector_IdContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(255);
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
		enterRule(_localctx, 34, RULE_func_vector);
		try {
			setState(279);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				_localctx = new Func_Vector_AppendContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(258);
				match(ID);
				setState(259);
				match(PUNTO);
				setState(260);
				match(APPEND);
				setState(261);
				match(PARIZQ);
				setState(262);
				e(0);
				setState(263);
				match(PARDER);
				}
				break;
			case 2:
				_localctx = new Func_Vector_RemoveLastContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(265);
				match(ID);
				setState(266);
				match(PUNTO);
				setState(267);
				match(REMOVELAST);
				setState(268);
				match(PARIZQ);
				setState(269);
				match(PARDER);
				}
				break;
			case 3:
				_localctx = new Func_Vector_RemoveContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(270);
				match(ID);
				setState(271);
				match(PUNTO);
				setState(272);
				match(REMOVE);
				setState(273);
				match(PARIZQ);
				setState(274);
				match(AT);
				setState(275);
				match(DOSPT);
				setState(276);
				e(0);
				setState(277);
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
	public static class Dec_matrizContext extends ParserRuleContext {
		public Dec_matrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dec_matriz; }
	 
		public Dec_matrizContext() { }
		public void copyFrom(Dec_matrizContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Declaracion_MatrizContext extends Dec_matrizContext {
		public Token tipod;
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode IGUAL() { return getToken(Tswift_GrammarNParser.IGUAL, 0); }
		public Def_matrizContext def_matriz() {
			return getRuleContext(Def_matrizContext.class,0);
		}
		public TerminalNode VAR() { return getToken(Tswift_GrammarNParser.VAR, 0); }
		public TerminalNode LET() { return getToken(Tswift_GrammarNParser.LET, 0); }
		public TerminalNode DOSPT() { return getToken(Tswift_GrammarNParser.DOSPT, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public Declaracion_MatrizContext(Dec_matrizContext ctx) { copyFrom(ctx); }
	}

	public final Dec_matrizContext dec_matriz() throws RecognitionException {
		Dec_matrizContext _localctx = new Dec_matrizContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_dec_matriz);
		int _la;
		try {
			_localctx = new Declaracion_MatrizContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(281);
			((Declaracion_MatrizContext)_localctx).tipod = _input.LT(1);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
				((Declaracion_MatrizContext)_localctx).tipod = (Token)_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(282);
			match(ID);
			setState(285);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DOSPT) {
				{
				setState(283);
				match(DOSPT);
				setState(284);
				tipo();
				}
			}

			setState(287);
			match(IGUAL);
			setState(288);
			def_matriz();
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
	public static class Def_matrizContext extends ParserRuleContext {
		public Def_matrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_def_matriz; }
	 
		public Def_matrizContext() { }
		public void copyFrom(Def_matrizContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Def_Matriz_ListaContext extends Def_matrizContext {
		public Listavalores_matrizContext listavalores_matriz() {
			return getRuleContext(Listavalores_matrizContext.class,0);
		}
		public Def_Matriz_ListaContext(Def_matrizContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Def_Matriz_SimpleContext extends Def_matrizContext {
		public Simple_matrizContext simple_matriz() {
			return getRuleContext(Simple_matrizContext.class,0);
		}
		public Def_Matriz_SimpleContext(Def_matrizContext ctx) { copyFrom(ctx); }
	}

	public final Def_matrizContext def_matriz() throws RecognitionException {
		Def_matrizContext _localctx = new Def_matrizContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_def_matriz);
		try {
			setState(292);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				_localctx = new Def_Matriz_ListaContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(290);
				listavalores_matriz();
				}
				break;
			case 2:
				_localctx = new Def_Matriz_SimpleContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(291);
				simple_matriz();
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
	public static class Listavalores_matrizContext extends ParserRuleContext {
		public Listavalores_matrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listavalores_matriz; }
	 
		public Listavalores_matrizContext() { }
		public void copyFrom(Listavalores_matrizContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Def_Matriz_ExpresionContext extends Listavalores_matrizContext {
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
		public Def_Matriz_ExpresionContext(Listavalores_matrizContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Def_Matriz_Lista_ValoresContext extends Listavalores_matrizContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public List<Listavalores_matrizContext> listavalores_matriz() {
			return getRuleContexts(Listavalores_matrizContext.class);
		}
		public Listavalores_matrizContext listavalores_matriz(int i) {
			return getRuleContext(Listavalores_matrizContext.class,i);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public List<TerminalNode> COMA() { return getTokens(Tswift_GrammarNParser.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(Tswift_GrammarNParser.COMA, i);
		}
		public Def_Matriz_Lista_ValoresContext(Listavalores_matrizContext ctx) { copyFrom(ctx); }
	}

	public final Listavalores_matrizContext listavalores_matriz() throws RecognitionException {
		Listavalores_matrizContext _localctx = new Listavalores_matrizContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_listavalores_matriz);
		int _la;
		try {
			setState(316);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				_localctx = new Def_Matriz_Lista_ValoresContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(294);
				match(CORCHETEIZQ);
				setState(295);
				listavalores_matriz();
				setState(300);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMA) {
					{
					{
					setState(296);
					match(COMA);
					setState(297);
					listavalores_matriz();
					}
					}
					setState(302);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(303);
				match(CORCHETEDER);
				}
				break;
			case 2:
				_localctx = new Def_Matriz_ExpresionContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(305);
				match(CORCHETEIZQ);
				setState(306);
				e(0);
				setState(311);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMA) {
					{
					{
					setState(307);
					match(COMA);
					setState(308);
					e(0);
					}
					}
					setState(313);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(314);
				match(CORCHETEDER);
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
	public static class Simple_matrizContext extends ParserRuleContext {
		public Simple_matrizContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_simple_matriz; }
	 
		public Simple_matrizContext() { }
		public void copyFrom(Simple_matrizContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Def_Matriz_Simple3Context extends Simple_matrizContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public TerminalNode REPEATING() { return getToken(Tswift_GrammarNParser.REPEATING, 0); }
		public List<TerminalNode> DOSPT() { return getTokens(Tswift_GrammarNParser.DOSPT); }
		public TerminalNode DOSPT(int i) {
			return getToken(Tswift_GrammarNParser.DOSPT, i);
		}
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public TerminalNode COMA() { return getToken(Tswift_GrammarNParser.COMA, 0); }
		public TerminalNode COUNT() { return getToken(Tswift_GrammarNParser.COUNT, 0); }
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public Def_Matriz_Simple3Context(Simple_matrizContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Def_Matriz_Simple2Context extends Simple_matrizContext {
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public TipoContext tipo() {
			return getRuleContext(TipoContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public TerminalNode PARIZQ() { return getToken(Tswift_GrammarNParser.PARIZQ, 0); }
		public TerminalNode REPEATING() { return getToken(Tswift_GrammarNParser.REPEATING, 0); }
		public List<TerminalNode> DOSPT() { return getTokens(Tswift_GrammarNParser.DOSPT); }
		public TerminalNode DOSPT(int i) {
			return getToken(Tswift_GrammarNParser.DOSPT, i);
		}
		public Simple_matrizContext simple_matriz() {
			return getRuleContext(Simple_matrizContext.class,0);
		}
		public TerminalNode COMA() { return getToken(Tswift_GrammarNParser.COMA, 0); }
		public TerminalNode COUNT() { return getToken(Tswift_GrammarNParser.COUNT, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Tswift_GrammarNParser.PARDER, 0); }
		public Def_Matriz_Simple2Context(Simple_matrizContext ctx) { copyFrom(ctx); }
	}

	public final Simple_matrizContext simple_matriz() throws RecognitionException {
		Simple_matrizContext _localctx = new Simple_matrizContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_simple_matriz);
		try {
			setState(344);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				_localctx = new Def_Matriz_Simple2Context(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(318);
				match(CORCHETEIZQ);
				setState(319);
				tipo();
				setState(320);
				match(CORCHETEDER);
				setState(321);
				match(PARIZQ);
				setState(322);
				match(REPEATING);
				setState(323);
				match(DOSPT);
				setState(324);
				simple_matriz();
				setState(325);
				match(COMA);
				setState(326);
				match(COUNT);
				setState(327);
				match(DOSPT);
				setState(328);
				e(0);
				setState(329);
				match(PARDER);
				}
				break;
			case 2:
				_localctx = new Def_Matriz_Simple3Context(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(331);
				match(CORCHETEIZQ);
				setState(332);
				tipo();
				setState(333);
				match(CORCHETEDER);
				setState(334);
				match(PARIZQ);
				setState(335);
				match(REPEATING);
				setState(336);
				match(DOSPT);
				setState(337);
				e(0);
				setState(338);
				match(COMA);
				setState(339);
				match(COUNT);
				setState(340);
				match(DOSPT);
				setState(341);
				e(0);
				setState(342);
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
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_condicion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(358);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				{
				_localctx = new Cond_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(347);
				((Cond_NegContext)_localctx).op = match(NOT);
				setState(348);
				((Cond_NegContext)_localctx).c = condicion(5);
				}
				break;
			case 2:
				{
				_localctx = new Cond_RelContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(349);
				((Cond_RelContext)_localctx).e1 = e(0);
				setState(350);
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
				setState(351);
				((Cond_RelContext)_localctx).e2 = e(0);
				}
				break;
			case 3:
				{
				_localctx = new Cond_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(353);
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
				setState(354);
				match(PARIZQ);
				setState(355);
				((Cond_ParContext)_localctx).c = condicion(0);
				setState(356);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(365);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,31,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Cond_LogicaContext(new CondicionContext(_parentctx, _parentState));
					((Cond_LogicaContext)_localctx).c1 = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_condicion);
					setState(360);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(361);
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
					setState(362);
					((Cond_LogicaContext)_localctx).c2 = condicion(4);
					}
					} 
				}
				setState(367);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,31,_ctx);
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
	public static class Expr_MatrizContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public List<TerminalNode> CORCHETEIZQ() { return getTokens(Tswift_GrammarNParser.CORCHETEIZQ); }
		public TerminalNode CORCHETEIZQ(int i) {
			return getToken(Tswift_GrammarNParser.CORCHETEIZQ, i);
		}
		public List<EContext> e() {
			return getRuleContexts(EContext.class);
		}
		public EContext e(int i) {
			return getRuleContext(EContext.class,i);
		}
		public List<TerminalNode> CORCHETEDER() { return getTokens(Tswift_GrammarNParser.CORCHETEDER); }
		public TerminalNode CORCHETEDER(int i) {
			return getToken(Tswift_GrammarNParser.CORCHETEDER, i);
		}
		public Expr_MatrizContext(EContext ctx) { copyFrom(ctx); }
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
	public static class Expr_VectorContext extends EContext {
		public TerminalNode ID() { return getToken(Tswift_GrammarNParser.ID, 0); }
		public TerminalNode CORCHETEIZQ() { return getToken(Tswift_GrammarNParser.CORCHETEIZQ, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode CORCHETEDER() { return getToken(Tswift_GrammarNParser.CORCHETEDER, 0); }
		public Expr_VectorContext(EContext ctx) { copyFrom(ctx); }
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
		int _startState = 46;
		enterRecursionRule(_localctx, 46, RULE_e, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				_localctx = new Expr_NegContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(369);
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
				setState(370);
				((Expr_NegContext)_localctx).e1 = e(18);
				}
				break;
			case 2:
				{
				_localctx = new Expr_BooleanoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(371);
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
				setState(372);
				((Expr_NilContext)_localctx).n = match(NIL);
				}
				break;
			case 4:
				{
				_localctx = new Expr_VectorContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(373);
				match(ID);
				setState(374);
				match(CORCHETEIZQ);
				setState(375);
				e(0);
				setState(376);
				match(CORCHETEDER);
				}
				break;
			case 5:
				{
				_localctx = new Expr_MatrizContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(378);
				match(ID);
				setState(379);
				match(CORCHETEIZQ);
				setState(380);
				e(0);
				setState(381);
				match(CORCHETEDER);
				setState(386); 
				_errHandler.sync(this);
				_alt = 1;
				do {
					switch (_alt) {
					case 1:
						{
						{
						setState(382);
						match(CORCHETEIZQ);
						setState(383);
						e(0);
						setState(384);
						match(CORCHETEDER);
						}
						}
						break;
					default:
						throw new NoViableAltException(this);
					}
					setState(388); 
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,32,_ctx);
				} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
				}
				break;
			case 6:
				{
				_localctx = new Expr_IsEmptyContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(390);
				match(ID);
				setState(391);
				match(PUNTO);
				setState(392);
				match(ISEMPTY);
				}
				break;
			case 7:
				{
				_localctx = new Expr_CountContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(393);
				match(ID);
				setState(394);
				match(PUNTO);
				setState(395);
				match(COUNT);
				}
				break;
			case 8:
				{
				_localctx = new Expr_IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(396);
				((Expr_IdContext)_localctx).id = match(ID);
				}
				break;
			case 9:
				{
				_localctx = new Expr_EnteroContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(397);
				((Expr_EnteroContext)_localctx).n = match(ENTERO);
				}
				break;
			case 10:
				{
				_localctx = new Expr_DecimalContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(398);
				((Expr_DecimalContext)_localctx).n = match(DECIMAL);
				}
				break;
			case 11:
				{
				_localctx = new Expr_CadenaContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(399);
				((Expr_CadenaContext)_localctx).n = match(CADENA);
				}
				break;
			case 12:
				{
				_localctx = new Expr_CaracterContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(400);
				((Expr_CaracterContext)_localctx).n = match(CARACTER);
				}
				break;
			case 13:
				{
				_localctx = new Expr_ParContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(401);
				match(PARIZQ);
				setState(402);
				((Expr_ParContext)_localctx).e1 = e(0);
				setState(403);
				match(PARDER);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(424);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(422);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_MulDivContext(new EContext(_parentctx, _parentState));
						((Expr_MulDivContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(407);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(408);
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
						setState(409);
						((Expr_MulDivContext)_localctx).e2 = e(18);
						}
						break;
					case 2:
						{
						_localctx = new Expr_SumResContext(new EContext(_parentctx, _parentState));
						((Expr_SumResContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(410);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(411);
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
						setState(412);
						((Expr_SumResContext)_localctx).e2 = e(17);
						}
						break;
					case 3:
						{
						_localctx = new Expr_ModContext(new EContext(_parentctx, _parentState));
						((Expr_ModContext)_localctx).e1 = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(413);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(414);
						((Expr_ModContext)_localctx).op = match(MOD);
						setState(415);
						((Expr_ModContext)_localctx).e2 = e(16);
						}
						break;
					case 4:
						{
						_localctx = new Expr_RelContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(416);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(417);
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
						setState(418);
						e(15);
						}
						break;
					case 5:
						{
						_localctx = new Expr_LogicaContext(new EContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_e);
						setState(419);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(420);
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
						setState(421);
						e(14);
						}
						break;
					}
					} 
				}
				setState(426);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
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
		case 22:
			return condicion_sempred((CondicionContext)_localctx, predIndex);
		case 23:
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
			return precpred(_ctx, 17);
		case 2:
			return precpred(_ctx, 16);
		case 3:
			return precpred(_ctx, 15);
		case 4:
			return precpred(_ctx, 14);
		case 5:
			return precpred(_ctx, 13);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001M\u01ac\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0005\u00014\b\u0001\n\u0001\f\u00017\t\u0001\u0001\u0002"+
		"\u0001\u0002\u0003\u0002;\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"?\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002C\b\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002J\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002P\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002T\b\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002X\b\u0002\u0003\u0002Z\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0003\u0003`\b\u0003\u0003\u0003b\b\u0003\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0005\u0004i\b\u0004\n\u0004"+
		"\f\u0004l\t\u0004\u0001\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u0081\b\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u0089"+
		"\b\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u0095"+
		"\b\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0003\b\u00a2\b\b\u0003\b\u00a4\b\b\u0001\t"+
		"\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\t\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0005\n\u00b2\b\n\n\n\f\n\u00b5\t\n\u0001\n\u0003\n"+
		"\u00b8\b\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0003\u000e\u00dd\b\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00f1\b\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u00f7\b\u0010\n"+
		"\u0010\f\u0010\u00fa\t\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0003\u0010\u0101\b\u0010\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0003\u0011\u0118\b\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0003\u0012\u011e\b\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0013\u0001\u0013\u0003\u0013\u0125\b\u0013\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0005\u0014\u012b\b\u0014\n\u0014\f\u0014\u012e\t\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0005\u0014\u0136\b\u0014\n\u0014\f\u0014\u0139\t\u0014\u0001\u0014\u0001"+
		"\u0014\u0003\u0014\u013d\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u0159"+
		"\b\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0003\u0016\u0167\b\u0016\u0001\u0016\u0001\u0016\u0001\u0016\u0005"+
		"\u0016\u016c\b\u0016\n\u0016\f\u0016\u016f\t\u0016\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0004\u0017\u0183\b\u0017"+
		"\u000b\u0017\f\u0017\u0184\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017"+
		"\u0196\b\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0005\u0017\u01a7\b\u0017"+
		"\n\u0017\f\u0017\u01aa\t\u0017\u0001\u0017\u0000\u0002,.\u0018\u0000\u0002"+
		"\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e"+
		" \"$&(*,.\u0000\b\u0001\u0000 !\u0001\u0000\u000e\u000f\u0001\u0000\u0016"+
		"\u001b\u0001\u0000\"#\u0001\u0000\u001c\u001d\u0002\u0000\u0013\u0013"+
		"\u001e\u001e\u0002\u0000\u0011\u0011\u0015\u0015\u0001\u0000\u0013\u0014"+
		"\u01da\u00000\u0001\u0000\u0000\u0000\u00025\u0001\u0000\u0000\u0000\u0004"+
		"Y\u0001\u0000\u0000\u0000\u0006a\u0001\u0000\u0000\u0000\bc\u0001\u0000"+
		"\u0000\u0000\n\u0080\u0001\u0000\u0000\u0000\f\u0088\u0001\u0000\u0000"+
		"\u0000\u000e\u0094\u0001\u0000\u0000\u0000\u0010\u0096\u0001\u0000\u0000"+
		"\u0000\u0012\u00a5\u0001\u0000\u0000\u0000\u0014\u00ad\u0001\u0000\u0000"+
		"\u0000\u0016\u00bb\u0001\u0000\u0000\u0000\u0018\u00c0\u0001\u0000\u0000"+
		"\u0000\u001a\u00c4\u0001\u0000\u0000\u0000\u001c\u00dc\u0001\u0000\u0000"+
		"\u0000\u001e\u00f0\u0001\u0000\u0000\u0000 \u0100\u0001\u0000\u0000\u0000"+
		"\"\u0117\u0001\u0000\u0000\u0000$\u0119\u0001\u0000\u0000\u0000&\u0124"+
		"\u0001\u0000\u0000\u0000(\u013c\u0001\u0000\u0000\u0000*\u0158\u0001\u0000"+
		"\u0000\u0000,\u0166\u0001\u0000\u0000\u0000.\u0195\u0001\u0000\u0000\u0000"+
		"01\u0003\u0002\u0001\u00001\u0001\u0001\u0000\u0000\u000024\u0003\u0004"+
		"\u0002\u000032\u0001\u0000\u0000\u000047\u0001\u0000\u0000\u000053\u0001"+
		"\u0000\u0000\u000056\u0001\u0000\u0000\u00006\u0003\u0001\u0000\u0000"+
		"\u000075\u0001\u0000\u0000\u00008:\u0003\b\u0004\u00009;\u0005\b\u0000"+
		"\u0000:9\u0001\u0000\u0000\u0000:;\u0001\u0000\u0000\u0000;Z\u0001\u0000"+
		"\u0000\u0000<>\u0003\n\u0005\u0000=?\u0005\b\u0000\u0000>=\u0001\u0000"+
		"\u0000\u0000>?\u0001\u0000\u0000\u0000?Z\u0001\u0000\u0000\u0000@B\u0003"+
		"\f\u0006\u0000AC\u0005\b\u0000\u0000BA\u0001\u0000\u0000\u0000BC\u0001"+
		"\u0000\u0000\u0000CZ\u0001\u0000\u0000\u0000DZ\u0003\u0010\b\u0000EZ\u0003"+
		"\u0014\n\u0000FZ\u0003\u0012\t\u0000GI\u0003\u0006\u0003\u0000HJ\u0005"+
		"\b\u0000\u0000IH\u0001\u0000\u0000\u0000IJ\u0001\u0000\u0000\u0000JZ\u0001"+
		"\u0000\u0000\u0000KZ\u0003\u001a\r\u0000LZ\u0003\u001c\u000e\u0000MO\u0003"+
		"\u001e\u000f\u0000NP\u0005\b\u0000\u0000ON\u0001\u0000\u0000\u0000OP\u0001"+
		"\u0000\u0000\u0000PZ\u0001\u0000\u0000\u0000QS\u0003\"\u0011\u0000RT\u0005"+
		"\b\u0000\u0000SR\u0001\u0000\u0000\u0000ST\u0001\u0000\u0000\u0000TZ\u0001"+
		"\u0000\u0000\u0000UW\u0003$\u0012\u0000VX\u0005\b\u0000\u0000WV\u0001"+
		"\u0000\u0000\u0000WX\u0001\u0000\u0000\u0000XZ\u0001\u0000\u0000\u0000"+
		"Y8\u0001\u0000\u0000\u0000Y<\u0001\u0000\u0000\u0000Y@\u0001\u0000\u0000"+
		"\u0000YD\u0001\u0000\u0000\u0000YE\u0001\u0000\u0000\u0000YF\u0001\u0000"+
		"\u0000\u0000YG\u0001\u0000\u0000\u0000YK\u0001\u0000\u0000\u0000YL\u0001"+
		"\u0000\u0000\u0000YM\u0001\u0000\u0000\u0000YQ\u0001\u0000\u0000\u0000"+
		"YU\u0001\u0000\u0000\u0000Z\u0005\u0001\u0000\u0000\u0000[b\u00050\u0000"+
		"\u0000\\b\u0005.\u0000\u0000]_\u0005/\u0000\u0000^`\u0003.\u0017\u0000"+
		"_^\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`b\u0001\u0000\u0000"+
		"\u0000a[\u0001\u0000\u0000\u0000a\\\u0001\u0000\u0000\u0000a]\u0001\u0000"+
		"\u0000\u0000b\u0007\u0001\u0000\u0000\u0000cd\u0005\u001f\u0000\u0000"+
		"de\u0005\u0002\u0000\u0000ej\u0003.\u0017\u0000fg\u0005\n\u0000\u0000"+
		"gi\u0003.\u0017\u0000hf\u0001\u0000\u0000\u0000il\u0001\u0000\u0000\u0000"+
		"jh\u0001\u0000\u0000\u0000jk\u0001\u0000\u0000\u0000km\u0001\u0000\u0000"+
		"\u0000lj\u0001\u0000\u0000\u0000mn\u0005\u0001\u0000\u0000n\t\u0001\u0000"+
		"\u0000\u0000op\u0007\u0000\u0000\u0000pq\u0005J\u0000\u0000qr\u0005\u0007"+
		"\u0000\u0000rs\u0003\u000e\u0007\u0000st\u0005\u0010\u0000\u0000tu\u0003"+
		".\u0017\u0000u\u0081\u0001\u0000\u0000\u0000vw\u0007\u0000\u0000\u0000"+
		"wx\u0005J\u0000\u0000xy\u0005\u0010\u0000\u0000y\u0081\u0003.\u0017\u0000"+
		"z{\u0007\u0000\u0000\u0000{|\u0005J\u0000\u0000|}\u0005\u0007\u0000\u0000"+
		"}~\u0003\u000e\u0007\u0000~\u007f\u0005\t\u0000\u0000\u007f\u0081\u0001"+
		"\u0000\u0000\u0000\u0080o\u0001\u0000\u0000\u0000\u0080v\u0001\u0000\u0000"+
		"\u0000\u0080z\u0001\u0000\u0000\u0000\u0081\u000b\u0001\u0000\u0000\u0000"+
		"\u0082\u0083\u0005J\u0000\u0000\u0083\u0084\u0005\u0010\u0000\u0000\u0084"+
		"\u0089\u0003.\u0017\u0000\u0085\u0086\u0005J\u0000\u0000\u0086\u0087\u0007"+
		"\u0001\u0000\u0000\u0087\u0089\u0003.\u0017\u0000\u0088\u0082\u0001\u0000"+
		"\u0000\u0000\u0088\u0085\u0001\u0000\u0000\u0000\u0089\r\u0001\u0000\u0000"+
		"\u0000\u008a\u0095\u0005@\u0000\u0000\u008b\u0095\u0005A\u0000\u0000\u008c"+
		"\u0095\u0005D\u0000\u0000\u008d\u0095\u0005B\u0000\u0000\u008e\u0095\u0005"+
		"C\u0000\u0000\u008f\u0095\u0005J\u0000\u0000\u0090\u0091\u0005\u0005\u0000"+
		"\u0000\u0091\u0092\u0003\u000e\u0007\u0000\u0092\u0093\u0005\u0006\u0000"+
		"\u0000\u0093\u0095\u0001\u0000\u0000\u0000\u0094\u008a\u0001\u0000\u0000"+
		"\u0000\u0094\u008b\u0001\u0000\u0000\u0000\u0094\u008c\u0001\u0000\u0000"+
		"\u0000\u0094\u008d\u0001\u0000\u0000\u0000\u0094\u008e\u0001\u0000\u0000"+
		"\u0000\u0094\u008f\u0001\u0000\u0000\u0000\u0094\u0090\u0001\u0000\u0000"+
		"\u0000\u0095\u000f\u0001\u0000\u0000\u0000\u0096\u0097\u0005$\u0000\u0000"+
		"\u0097\u0098\u0003,\u0016\u0000\u0098\u0099\u0005\u0003\u0000\u0000\u0099"+
		"\u009a\u0003\u0002\u0001\u0000\u009a\u00a3\u0005\u0004\u0000\u0000\u009b"+
		"\u00a1\u0005%\u0000\u0000\u009c\u00a2\u0003\u0010\b\u0000\u009d\u009e"+
		"\u0005\u0003\u0000\u0000\u009e\u009f\u0003\u0002\u0001\u0000\u009f\u00a0"+
		"\u0005\u0004\u0000\u0000\u00a0\u00a2\u0001\u0000\u0000\u0000\u00a1\u009c"+
		"\u0001\u0000\u0000\u0000\u00a1\u009d\u0001\u0000\u0000\u0000\u00a2\u00a4"+
		"\u0001\u0000\u0000\u0000\u00a3\u009b\u0001\u0000\u0000\u0000\u00a3\u00a4"+
		"\u0001\u0000\u0000\u0000\u00a4\u0011\u0001\u0000\u0000\u0000\u00a5\u00a6"+
		"\u0005-\u0000\u0000\u00a6\u00a7\u0003,\u0016\u0000\u00a7\u00a8\u0005%"+
		"\u0000\u0000\u00a8\u00a9\u0005\u0003\u0000\u0000\u00a9\u00aa\u0003\u0002"+
		"\u0001\u0000\u00aa\u00ab\u0003\u0006\u0003\u0000\u00ab\u00ac\u0005\u0004"+
		"\u0000\u0000\u00ac\u0013\u0001\u0000\u0000\u0000\u00ad\u00ae\u0005&\u0000"+
		"\u0000\u00ae\u00af\u0003.\u0017\u0000\u00af\u00b3\u0005\u0003\u0000\u0000"+
		"\u00b0\u00b2\u0003\u0016\u000b\u0000\u00b1\u00b0\u0001\u0000\u0000\u0000"+
		"\u00b2\u00b5\u0001\u0000\u0000\u0000\u00b3\u00b1\u0001\u0000\u0000\u0000"+
		"\u00b3\u00b4\u0001\u0000\u0000\u0000\u00b4\u00b7\u0001\u0000\u0000\u0000"+
		"\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b6\u00b8\u0003\u0018\f\u0000\u00b7"+
		"\u00b6\u0001\u0000\u0000\u0000\u00b7\u00b8\u0001\u0000\u0000\u0000\u00b8"+
		"\u00b9\u0001\u0000\u0000\u0000\u00b9\u00ba\u0005\u0004\u0000\u0000\u00ba"+
		"\u0015\u0001\u0000\u0000\u0000\u00bb\u00bc\u0005\'\u0000\u0000\u00bc\u00bd"+
		"\u0003.\u0017\u0000\u00bd\u00be\u0005\u0007\u0000\u0000\u00be\u00bf\u0003"+
		"\u0002\u0001\u0000\u00bf\u0017\u0001\u0000\u0000\u0000\u00c0\u00c1\u0005"+
		"(\u0000\u0000\u00c1\u00c2\u0005\u0007\u0000\u0000\u00c2\u00c3\u0003\u0002"+
		"\u0001\u0000\u00c3\u0019\u0001\u0000\u0000\u0000\u00c4\u00c5\u0005)\u0000"+
		"\u0000\u00c5\u00c6\u0003,\u0016\u0000\u00c6\u00c7\u0005\u0003\u0000\u0000"+
		"\u00c7\u00c8\u0003\u0002\u0001\u0000\u00c8\u00c9\u0005\u0004\u0000\u0000"+
		"\u00c9\u001b\u0001\u0000\u0000\u0000\u00ca\u00cb\u0005*\u0000\u0000\u00cb"+
		"\u00cc\u0005J\u0000\u0000\u00cc\u00cd\u0005+\u0000\u0000\u00cd\u00ce\u0003"+
		".\u0017\u0000\u00ce\u00cf\u0005,\u0000\u0000\u00cf\u00d0\u0003.\u0017"+
		"\u0000\u00d0\u00d1\u0005\u0003\u0000\u0000\u00d1\u00d2\u0003\u0002\u0001"+
		"\u0000\u00d2\u00d3\u0005\u0004\u0000\u0000\u00d3\u00dd\u0001\u0000\u0000"+
		"\u0000\u00d4\u00d5\u0005*\u0000\u0000\u00d5\u00d6\u0005J\u0000\u0000\u00d6"+
		"\u00d7\u0005+\u0000\u0000\u00d7\u00d8\u0003.\u0017\u0000\u00d8\u00d9\u0005"+
		"\u0003\u0000\u0000\u00d9\u00da\u0003\u0002\u0001\u0000\u00da\u00db\u0005"+
		"\u0004\u0000\u0000\u00db\u00dd\u0001\u0000\u0000\u0000\u00dc\u00ca\u0001"+
		"\u0000\u0000\u0000\u00dc\u00d4\u0001\u0000\u0000\u0000\u00dd\u001d\u0001"+
		"\u0000\u0000\u0000\u00de\u00df\u0007\u0000\u0000\u0000\u00df\u00e0\u0005"+
		"J\u0000\u0000\u00e0\u00e1\u0005\u0007\u0000\u0000\u00e1\u00e2\u0005\u0005"+
		"\u0000\u0000\u00e2\u00e3\u0003\u000e\u0007\u0000\u00e3\u00e4\u0005\u0006"+
		"\u0000\u0000\u00e4\u00e5\u0005\u0010\u0000\u0000\u00e5\u00e6\u0003 \u0010"+
		"\u0000\u00e6\u00f1\u0001\u0000\u0000\u0000\u00e7\u00e8\u0007\u0000\u0000"+
		"\u0000\u00e8\u00e9\u0005J\u0000\u0000\u00e9\u00ea\u0005\u0010\u0000\u0000"+
		"\u00ea\u00eb\u0005\u0005\u0000\u0000\u00eb\u00ec\u0003\u000e\u0007\u0000"+
		"\u00ec\u00ed\u0005\u0006\u0000\u0000\u00ed\u00ee\u0005\u0002\u0000\u0000"+
		"\u00ee\u00ef\u0005\u0001\u0000\u0000\u00ef\u00f1\u0001\u0000\u0000\u0000"+
		"\u00f0\u00de\u0001\u0000\u0000\u0000\u00f0\u00e7\u0001\u0000\u0000\u0000"+
		"\u00f1\u001f\u0001\u0000\u0000\u0000\u00f2\u00f3\u0005\u0005\u0000\u0000"+
		"\u00f3\u00f8\u0003.\u0017\u0000\u00f4\u00f5\u0005\n\u0000\u0000\u00f5"+
		"\u00f7\u0003.\u0017\u0000\u00f6\u00f4\u0001\u0000\u0000\u0000\u00f7\u00fa"+
		"\u0001\u0000\u0000\u0000\u00f8\u00f6\u0001\u0000\u0000\u0000\u00f8\u00f9"+
		"\u0001\u0000\u0000\u0000\u00f9\u00fb\u0001\u0000\u0000\u0000\u00fa\u00f8"+
		"\u0001\u0000\u0000\u0000\u00fb\u00fc\u0005\u0006\u0000\u0000\u00fc\u0101"+
		"\u0001\u0000\u0000\u0000\u00fd\u00fe\u0005\u0005\u0000\u0000\u00fe\u0101"+
		"\u0005\u0006\u0000\u0000\u00ff\u0101\u0005J\u0000\u0000\u0100\u00f2\u0001"+
		"\u0000\u0000\u0000\u0100\u00fd\u0001\u0000\u0000\u0000\u0100\u00ff\u0001"+
		"\u0000\u0000\u0000\u0101!\u0001\u0000\u0000\u0000\u0102\u0103\u0005J\u0000"+
		"\u0000\u0103\u0104\u0005\u000b\u0000\u0000\u0104\u0105\u00052\u0000\u0000"+
		"\u0105\u0106\u0005\u0002\u0000\u0000\u0106\u0107\u0003.\u0017\u0000\u0107"+
		"\u0108\u0005\u0001\u0000\u0000\u0108\u0118\u0001\u0000\u0000\u0000\u0109"+
		"\u010a\u0005J\u0000\u0000\u010a\u010b\u0005\u000b\u0000\u0000\u010b\u010c"+
		"\u00053\u0000\u0000\u010c\u010d\u0005\u0002\u0000\u0000\u010d\u0118\u0005"+
		"\u0001\u0000\u0000\u010e\u010f\u0005J\u0000\u0000\u010f\u0110\u0005\u000b"+
		"\u0000\u0000\u0110\u0111\u00054\u0000\u0000\u0111\u0112\u0005\u0002\u0000"+
		"\u0000\u0112\u0113\u00055\u0000\u0000\u0113\u0114\u0005\u0007\u0000\u0000"+
		"\u0114\u0115\u0003.\u0017\u0000\u0115\u0116\u0005\u0001\u0000\u0000\u0116"+
		"\u0118\u0001\u0000\u0000\u0000\u0117\u0102\u0001\u0000\u0000\u0000\u0117"+
		"\u0109\u0001\u0000\u0000\u0000\u0117\u010e\u0001\u0000\u0000\u0000\u0118"+
		"#\u0001\u0000\u0000\u0000\u0119\u011a\u0007\u0000\u0000\u0000\u011a\u011d"+
		"\u0005J\u0000\u0000\u011b\u011c\u0005\u0007\u0000\u0000\u011c\u011e\u0003"+
		"\u000e\u0007\u0000\u011d\u011b\u0001\u0000\u0000\u0000\u011d\u011e\u0001"+
		"\u0000\u0000\u0000\u011e\u011f\u0001\u0000\u0000\u0000\u011f\u0120\u0005"+
		"\u0010\u0000\u0000\u0120\u0121\u0003&\u0013\u0000\u0121%\u0001\u0000\u0000"+
		"\u0000\u0122\u0125\u0003(\u0014\u0000\u0123\u0125\u0003*\u0015\u0000\u0124"+
		"\u0122\u0001\u0000\u0000\u0000\u0124\u0123\u0001\u0000\u0000\u0000\u0125"+
		"\'\u0001\u0000\u0000\u0000\u0126\u0127\u0005\u0005\u0000\u0000\u0127\u012c"+
		"\u0003(\u0014\u0000\u0128\u0129\u0005\n\u0000\u0000\u0129\u012b\u0003"+
		"(\u0014\u0000\u012a\u0128\u0001\u0000\u0000\u0000\u012b\u012e\u0001\u0000"+
		"\u0000\u0000\u012c\u012a\u0001\u0000\u0000\u0000\u012c\u012d\u0001\u0000"+
		"\u0000\u0000\u012d\u012f\u0001\u0000\u0000\u0000\u012e\u012c\u0001\u0000"+
		"\u0000\u0000\u012f\u0130\u0005\u0006\u0000\u0000\u0130\u013d\u0001\u0000"+
		"\u0000\u0000\u0131\u0132\u0005\u0005\u0000\u0000\u0132\u0137\u0003.\u0017"+
		"\u0000\u0133\u0134\u0005\n\u0000\u0000\u0134\u0136\u0003.\u0017\u0000"+
		"\u0135\u0133\u0001\u0000\u0000\u0000\u0136\u0139\u0001\u0000\u0000\u0000"+
		"\u0137\u0135\u0001\u0000\u0000\u0000\u0137\u0138\u0001\u0000\u0000\u0000"+
		"\u0138\u013a\u0001\u0000\u0000\u0000\u0139\u0137\u0001\u0000\u0000\u0000"+
		"\u013a\u013b\u0005\u0006\u0000\u0000\u013b\u013d\u0001\u0000\u0000\u0000"+
		"\u013c\u0126\u0001\u0000\u0000\u0000\u013c\u0131\u0001\u0000\u0000\u0000"+
		"\u013d)\u0001\u0000\u0000\u0000\u013e\u013f\u0005\u0005\u0000\u0000\u013f"+
		"\u0140\u0003\u000e\u0007\u0000\u0140\u0141\u0005\u0006\u0000\u0000\u0141"+
		"\u0142\u0005\u0002\u0000\u0000\u0142\u0143\u00059\u0000\u0000\u0143\u0144"+
		"\u0005\u0007\u0000\u0000\u0144\u0145\u0003*\u0015\u0000\u0145\u0146\u0005"+
		"\n\u0000\u0000\u0146\u0147\u00057\u0000\u0000\u0147\u0148\u0005\u0007"+
		"\u0000\u0000\u0148\u0149\u0003.\u0017\u0000\u0149\u014a\u0005\u0001\u0000"+
		"\u0000\u014a\u0159\u0001\u0000\u0000\u0000\u014b\u014c\u0005\u0005\u0000"+
		"\u0000\u014c\u014d\u0003\u000e\u0007\u0000\u014d\u014e\u0005\u0006\u0000"+
		"\u0000\u014e\u014f\u0005\u0002\u0000\u0000\u014f\u0150\u00059\u0000\u0000"+
		"\u0150\u0151\u0005\u0007\u0000\u0000\u0151\u0152\u0003.\u0017\u0000\u0152"+
		"\u0153\u0005\n\u0000\u0000\u0153\u0154\u00057\u0000\u0000\u0154\u0155"+
		"\u0005\u0007\u0000\u0000\u0155\u0156\u0003.\u0017\u0000\u0156\u0157\u0005"+
		"\u0001\u0000\u0000\u0157\u0159\u0001\u0000\u0000\u0000\u0158\u013e\u0001"+
		"\u0000\u0000\u0000\u0158\u014b\u0001\u0000\u0000\u0000\u0159+\u0001\u0000"+
		"\u0000\u0000\u015a\u015b\u0006\u0016\uffff\uffff\u0000\u015b\u015c\u0005"+
		"\u001e\u0000\u0000\u015c\u0167\u0003,\u0016\u0005\u015d\u015e\u0003.\u0017"+
		"\u0000\u015e\u015f\u0007\u0002\u0000\u0000\u015f\u0160\u0003.\u0017\u0000"+
		"\u0160\u0167\u0001\u0000\u0000\u0000\u0161\u0167\u0007\u0003\u0000\u0000"+
		"\u0162\u0163\u0005\u0002\u0000\u0000\u0163\u0164\u0003,\u0016\u0000\u0164"+
		"\u0165\u0005\u0001\u0000\u0000\u0165\u0167\u0001\u0000\u0000\u0000\u0166"+
		"\u015a\u0001\u0000\u0000\u0000\u0166\u015d\u0001\u0000\u0000\u0000\u0166"+
		"\u0161\u0001\u0000\u0000\u0000\u0166\u0162\u0001\u0000\u0000\u0000\u0167"+
		"\u016d\u0001\u0000\u0000\u0000\u0168\u0169\n\u0003\u0000\u0000\u0169\u016a"+
		"\u0007\u0004\u0000\u0000\u016a\u016c\u0003,\u0016\u0004\u016b\u0168\u0001"+
		"\u0000\u0000\u0000\u016c\u016f\u0001\u0000\u0000\u0000\u016d\u016b\u0001"+
		"\u0000\u0000\u0000\u016d\u016e\u0001\u0000\u0000\u0000\u016e-\u0001\u0000"+
		"\u0000\u0000\u016f\u016d\u0001\u0000\u0000\u0000\u0170\u0171\u0006\u0017"+
		"\uffff\uffff\u0000\u0171\u0172\u0007\u0005\u0000\u0000\u0172\u0196\u0003"+
		".\u0017\u0012\u0173\u0196\u0007\u0003\u0000\u0000\u0174\u0196\u00051\u0000"+
		"\u0000\u0175\u0176\u0005J\u0000\u0000\u0176\u0177\u0005\u0005\u0000\u0000"+
		"\u0177\u0178\u0003.\u0017\u0000\u0178\u0179\u0005\u0006\u0000\u0000\u0179"+
		"\u0196\u0001\u0000\u0000\u0000\u017a\u017b\u0005J\u0000\u0000\u017b\u017c"+
		"\u0005\u0005\u0000\u0000\u017c\u017d\u0003.\u0017\u0000\u017d\u0182\u0005"+
		"\u0006\u0000\u0000\u017e\u017f\u0005\u0005\u0000\u0000\u017f\u0180\u0003"+
		".\u0017\u0000\u0180\u0181\u0005\u0006\u0000\u0000\u0181\u0183\u0001\u0000"+
		"\u0000\u0000\u0182\u017e\u0001\u0000\u0000\u0000\u0183\u0184\u0001\u0000"+
		"\u0000\u0000\u0184\u0182\u0001\u0000\u0000\u0000\u0184\u0185\u0001\u0000"+
		"\u0000\u0000\u0185\u0196\u0001\u0000\u0000\u0000\u0186\u0187\u0005J\u0000"+
		"\u0000\u0187\u0188\u0005\u000b\u0000\u0000\u0188\u0196\u00056\u0000\u0000"+
		"\u0189\u018a\u0005J\u0000\u0000\u018a\u018b\u0005\u000b\u0000\u0000\u018b"+
		"\u0196\u00057\u0000\u0000\u018c\u0196\u0005J\u0000\u0000\u018d\u0196\u0005"+
		"F\u0000\u0000\u018e\u0196\u0005G\u0000\u0000\u018f\u0196\u0005I\u0000"+
		"\u0000\u0190\u0196\u0005H\u0000\u0000\u0191\u0192\u0005\u0002\u0000\u0000"+
		"\u0192\u0193\u0003.\u0017\u0000\u0193\u0194\u0005\u0001\u0000\u0000\u0194"+
		"\u0196\u0001\u0000\u0000\u0000\u0195\u0170\u0001\u0000\u0000\u0000\u0195"+
		"\u0173\u0001\u0000\u0000\u0000\u0195\u0174\u0001\u0000\u0000\u0000\u0195"+
		"\u0175\u0001\u0000\u0000\u0000\u0195\u017a\u0001\u0000\u0000\u0000\u0195"+
		"\u0186\u0001\u0000\u0000\u0000\u0195\u0189\u0001\u0000\u0000\u0000\u0195"+
		"\u018c\u0001\u0000\u0000\u0000\u0195\u018d\u0001\u0000\u0000\u0000\u0195"+
		"\u018e\u0001\u0000\u0000\u0000\u0195\u018f\u0001\u0000\u0000\u0000\u0195"+
		"\u0190\u0001\u0000\u0000\u0000\u0195\u0191\u0001\u0000\u0000\u0000\u0196"+
		"\u01a8\u0001\u0000\u0000\u0000\u0197\u0198\n\u0011\u0000\u0000\u0198\u0199"+
		"\u0007\u0006\u0000\u0000\u0199\u01a7\u0003.\u0017\u0012\u019a\u019b\n"+
		"\u0010\u0000\u0000\u019b\u019c\u0007\u0007\u0000\u0000\u019c\u01a7\u0003"+
		".\u0017\u0011\u019d\u019e\n\u000f\u0000\u0000\u019e\u019f\u0005\u0012"+
		"\u0000\u0000\u019f\u01a7\u0003.\u0017\u0010\u01a0\u01a1\n\u000e\u0000"+
		"\u0000\u01a1\u01a2\u0007\u0002\u0000\u0000\u01a2\u01a7\u0003.\u0017\u000f"+
		"\u01a3\u01a4\n\r\u0000\u0000\u01a4\u01a5\u0007\u0004\u0000\u0000\u01a5"+
		"\u01a7\u0003.\u0017\u000e\u01a6\u0197\u0001\u0000\u0000\u0000\u01a6\u019a"+
		"\u0001\u0000\u0000\u0000\u01a6\u019d\u0001\u0000\u0000\u0000\u01a6\u01a0"+
		"\u0001\u0000\u0000\u0000\u01a6\u01a3\u0001\u0000\u0000\u0000\u01a7\u01aa"+
		"\u0001\u0000\u0000\u0000\u01a8\u01a6\u0001\u0000\u0000\u0000\u01a8\u01a9"+
		"\u0001\u0000\u0000\u0000\u01a9/\u0001\u0000\u0000\u0000\u01aa\u01a8\u0001"+
		"\u0000\u0000\u0000$5:>BIOSWY_aj\u0080\u0088\u0094\u00a1\u00a3\u00b3\u00b7"+
		"\u00dc\u00f0\u00f8\u0100\u0117\u011d\u0124\u012c\u0137\u013c\u0158\u0166"+
		"\u016d\u0184\u0195\u01a6\u01a8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}