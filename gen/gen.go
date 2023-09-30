package gen

import "fmt"

var tmp = 0
var lb = 0

func Gen(out string) {
	fmt.Println(out)
}

func NewTemp() string {
	tmp++
	return fmt.Sprintf("t%d", tmp)
}

func NewEtq() string {
	lb++
	return fmt.Sprintf("L%d", lb)
}

func ImprimirEtq(etiquetas []string) {
	for _, e := range etiquetas {
		Gen(e + ":")
	}
}

func Unir(etq1 []string, etq2 []string) []string {
	etiquetas := append(etq1, etq2...)
	return etiquetas
}
