package exercicio3

type Servico struct {
	Nome               string
	Preco              float64
	MinutosTrabalhados float64
}

func SomarServicos(servicos []Servico) float64 {
	var precoTotal float64
	for _, servico := range servicos {
		if servico.MinutosTrabalhados < 30 {
			servico.MinutosTrabalhados = 30
		}
		precoTotal += servico.Preco * (servico.MinutosTrabalhados / 60)
	}
	return precoTotal
}
