package exercicio2

import (
	"fmt"
)

type Produto struct {
	Tipo  string
	Nome  string
	Preço float64
}

type Loja struct {
	Produtos []Produto
}

type ProdutoInterface interface {
	CalcularCusto() float64
}

type EcommerceInterface interface {
	Total() float64
	Adicionar(Produto)
}

func (p Produto) CalcularCusto() float64 {
	switch p.Tipo {
	case "Pequeno":
		return p.Preço
	case "Medio":
		return p.Preço + (p.Preço * 0.03)
	case "Grande":
		return p.Preço + (p.Preço * 0.06) + 2500
	}
	return 0
}

func (l Loja) Total() float64 {
	var total float64
	for _, produto := range l.Produtos {
		total += produto.CalcularCusto()
	}
	return total
}

func (l *Loja) Adicionar(produto Produto) {
	l.Produtos = append(l.Produtos, produto)
}

func novoProduto(tipo, nome string, preco float64) Produto {
	return Produto{
		Tipo:  tipo,
		Nome:  nome,
		Preço: preco,
	}
}

func novaLoja() EcommerceInterface {
	return &Loja{}
}

func Run() {
	loja := novaLoja()
	loja.Adicionar(novoProduto("Pequeno", "Lapis", 2.50))
	loja.Adicionar(novoProduto("Medio", "Notebook", 4000))

	fmt.Println("Exercicio 2:")
	fmt.Println("Total da loja: R$", loja.Total())
}
