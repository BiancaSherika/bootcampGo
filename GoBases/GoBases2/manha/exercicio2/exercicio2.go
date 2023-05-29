package exercicio2

import (
	"errors"
	"fmt"
)

func Run() {
	fmt.Println("Exercicio 2:")
	fmt.Println(calcularMedia())
	fmt.Println(calcularMedia(5))
	fmt.Println(calcularMedia(3, 5, 8))
}

func calcularMedia(numeros ...float64) (float64, error) {
	var erro error

	if len(numeros) <= 0 {
		erro = errors.New("Erro, tente novamente")
		return -1, erro
	}

	sum := numeros[0]

	for _, numero := range numeros {
		sum += numero
	}

	return sum / float64(len(numeros)), erro
}
