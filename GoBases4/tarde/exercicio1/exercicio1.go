package exercicio1

import (
	"fmt"
	"os"
)

func recuperarPanico() {
	if erro := recover(); erro != nil {
		fmt.Println(erro)
	}
}

func Run() {
	defer recuperarPanico()
	fmt.Println("Exercicio 1")
	_, erro := os.Open("customers.txt")

	if erro != nil {
		panic("Erro: o arquivo indicado não foi encontrado ou está danificado")
	}
}
