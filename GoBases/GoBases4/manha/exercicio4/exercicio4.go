package exercicio4

import (
	"errors"
	"fmt"
)

func verificarSalario(horasTrabalhadasMensal float64, valorHora float64) (int, error) {
	salarioTotal := horasTrabalhadasMensal * valorHora

	if salarioTotal >= 20000 {
		salarioTotal = salarioTotal - (salarioTotal * 0.10)
	} else if horasTrabalhadasMensal < 80 {
		return fmt.Println("Horas trabalhadas:", horasTrabalhadasMensal, errors.New("- Erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês"))
	}

	erro := fmt.Errorf("O salário mensal será: %v", salarioTotal)
	return fmt.Println(erro)
}

func Run() {
	fmt.Println("exercicio 4")
	verificarSalario(180, 150)
	verificarSalario(180, 100)
	verificarSalario(70, 100)
}
