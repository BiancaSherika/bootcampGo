package exercicio1

import (
	"fmt"
)

type Alune struct {
	Nome           string
	Sobrenome      string
	RG             string
	DataDeAdmissao string
}

func (alune Alune) aluneInfo() {
	var alune1 Alune
	alune1.Nome = alune.Nome
	alune1.Sobrenome = alune.Sobrenome
	alune1.RG = alune.RG
	alune1.DataDeAdmissao = alune.DataDeAdmissao

	fmt.Println("Informações de alunes: ", alune1)
}

func Run() {
	fmt.Println("Exercicio 1:")
	alune := Alune{"Bianca", "Claro", "12345678X", "26/05/2023"}
	alune.aluneInfo()
}
