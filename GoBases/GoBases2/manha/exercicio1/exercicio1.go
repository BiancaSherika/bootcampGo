package exercicio1

import (
	"fmt"
)

func calcularImposto(salario float64) float64 {
	if salario >= 50000 && salario < 150000 {
		return salario * 0.17
	} else if salario >= 150000 {
		return salario * 0.27
	}
	return 0
}

func Run() {
	fmt.Println("Exercicio 1:")
	fmt.Println(calcularImposto(50000))
	fmt.Println(calcularImposto(150000))
	fmt.Println(calcularImposto(25000))
}
