package exercicio3

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func SomarProdutos(produtos []Produto) float64 {
	var precoTotal float64
	for _, produto := range produtos {
		precoTotal += produto.Preco * float64(produto.Quantidade)
	}
	return precoTotal
}
