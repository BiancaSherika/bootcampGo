package main

import (
	"fmt"
)

var (
	idade            = 26
	mesesTrabalhando = 18
	empregado        = true
	salario          = 150000
)

func main() {
	switch {
	case idade <= 22:
		fmt.Println("Não concedemos emprestimos a clientes com menos de 22 anos")
	case empregado == false:
		fmt.Println("Não concedemos emprestimos a clientes desempregados")
	case mesesTrabalhando <= 12:
		fmt.Println("Não concedemos emprestimos a clientes com menos de um ano de atividade")
	case salario <= 100000:
		fmt.Println("Empréstimo concedido com juros")
	default:
		fmt.Println("Empréstimo concedido sem juros")
	}
}
