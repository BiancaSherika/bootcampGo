package exercicio3

type Manutencao struct {
	Nome  string
	Preco float64
}

func SomarManutencao(manutencoes []Manutencao) float64 {
	var precoTotal float64
	for _, produto := range manutencoes {
		precoTotal += produto.Preco
	}
	return precoTotal
}
