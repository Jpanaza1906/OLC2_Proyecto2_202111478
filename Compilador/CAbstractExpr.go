package compilador

type CAbstractExpr interface {
	Compilar(ctx *Contexto) *Atributos
}
