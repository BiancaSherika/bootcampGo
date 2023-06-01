package exercicio2

import (
	"errors"
	"fmt"
)

func verificarSalario(salario int) (int, error) {
	if salario < 15000 {
		return fmt.Println("salario:", salario, errors.New("- Erro: O salário digitado não está dentro do valor mínimo"))
	}
	return fmt.Printf("salario: %d - Necessário pagamento de imposto\n", salario)
}

func Run() {
	fmt.Println("exercicio 2")
	verificarSalario(10000)
	verificarSalario(20000)
}
