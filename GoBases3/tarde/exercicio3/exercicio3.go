package exercicio3

import "fmt"

func Run() {
	fmt.Println("Exercicio 3")
	produtos := []Produto{
		{"Café da manhã", 10.20, 1},
		{"Almoço", 20.50, 10},
		{"Fruta", 5.00, 20},
	}

	servicos := []Servico{
		{"Cozinheiro", 40.00, 4000},
		{"Garcom", 30.0, 3000},
		{"Chef", 100.00, 20},
	}

	manutencao := []Manutencao{
		{"Limpeza", 200},
		{"Atendimento", 400},
	}

	totalProdutos := SomarProdutos(produtos)
	totalServicos := SomarServicos(servicos)
	totalManutencao := SomarManutencao(manutencao)
	precoTotal := totalProdutos + totalServicos + totalManutencao

	fmt.Println("total gasto: ", precoTotal)
}
