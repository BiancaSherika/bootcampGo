package exercicio1

import (
	"errors"
	"fmt"
	"os"
)

type produto struct {
	id         int
	quantidade int
	preco      float64
}

func (p produto) criarProduto() string {
	return fmt.Sprintf("%d,%d,%.2f\n", p.id, p.quantidade, p.preco)
}

func (p produto) criarCabecalho() string {
	return "id,quantidade,preço\n"
}

func criarArquivoProdutosComprados(caminho string, produtos []produto) error {
	if len(produtos) == 0 {
		return errors.New("Erro, nenhum produto encontrado")
	}

	file, err := os.OpenFile(caminho, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("Erro ao abrir arquivo: %w", err)
	}

	defer file.Close()
	p := produtos[0]

	if _, err = file.WriteString(p.criarCabecalho()); err != nil {
		return fmt.Errorf("erro ao gerar cabeçalho: %w", err)
	}

	for _, p = range produtos {

		if _, err = file.WriteString(p.criarProduto()); err != nil {
			return fmt.Errorf("erro ao salvar produto: %w", err)
		}
	}

	return nil
}

func Run() {
	produtos := []produto{{1, 10, 5}, {2, 50, 20.99}}
	criarArquivoProdutosComprados("produtos.csv", produtos)
}
