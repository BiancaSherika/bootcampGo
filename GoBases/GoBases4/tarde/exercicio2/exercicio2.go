package exercicio2

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

func gerarID() {
	var id any
	id = rand.Intn(1000)
	if id == nil {
		msg := erroPersonalizado("gerarID")
		panic(msg)
	}
	fmt.Println("Id gerado:", id)
}

func verificarCliente() {
	_, erro := os.Open("costumers.txt")
	if erro != nil {
		msg := erroPersonalizado("verificarCliente")
		panic(msg)
	}
}

func recuperarPanico() {
	if erro := recover(); erro != nil {
		fmt.Println(erro)
	}
}

func erroPersonalizado(tipo string) error {
	switch tipo {
	case "verificarCliente":
		return fmt.Errorf("Erro: o arquivo indicado não foi encontrado ou está danificado")
	case "gerarID":
		return fmt.Errorf("Erro ao gerar ID")
	default:
		return errors.New("Erro genérico")
	}
}

func Run() {
	defer recuperarPanico()
	fmt.Print("Exercicio 2\n")
	gerarID()
	defer verificarCliente()
	fmt.Println("função ainda em execução")
}
