package exercicio2

import "fmt"

type usuario struct {
	Nome      string
	Sobrenome string
	Email     string
	Produtos  []produto
}

type produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

type Produto interface {
	novoProduto(nome string, preco float64) produto
	adicionarProduto(usuario *usuario, produto produto, quantidade int)
	deletarProduto(usuario *usuario)
}

func novoProduto(nome string, preco float64) produto {
	return produto{
		Nome:  nome,
		Preco: preco,
	}
}

func (usuario *usuario) adicionarProduto(produto produto, quantidade int) {
	usuario.Produtos = append(usuario.Produtos, produto)
}

func (usuario *usuario) deletarProduto() {
	usuario.Produtos = nil
}

func Run() {
	fmt.Println("Exercicio 2")
	userTeste := usuario{"Bianca", "Claro", "bianca@email.com", nil}
	var p *usuario = &userTeste
	printarInfos(p.Produtos)

	produtoTeste := novoProduto("abacaxi", 7.50)
	p.adicionarProduto(produtoTeste, 2)
	printarInfos(p.Produtos)

	p.deletarProduto()
	printarInfos(p.Produtos)
}

func printarInfos(produtos []produto) {
	fmt.Println("Inicio da execução")
	for _, produto := range produtos {
		fmt.Println(produto.Nome, &produto.Nome)
		fmt.Println(produto.Preco, &produto.Preco)
		fmt.Println(produto.Quantidade, &produto.Quantidade)
	}
}
