package exercicio3

import "fmt"

func verificarSalario(salario int) {
	if salario < 15000 {
		erro := fmt.Errorf("Erro: o mínimo tributável é 15.000 e o salário informado é: %v", salario)
		fmt.Println(erro)
	} else {
		fmt.Printf("salario: %v - Necessário pagamento de imposto\n", salario)
	}
}

func Run() {
	fmt.Println("exercicio 3")
	verificarSalario(10000)
	verificarSalario(20000)
}
