package exercicio3

import "fmt"

func Run() {
	fmt.Println("Exercicio 3:")
	fmt.Println(calcularSalario(100, ""))
	fmt.Println(calcularSalario(100, "A"))
	fmt.Println(calcularSalario(100, "B"))
	fmt.Println(calcularSalario(100, "C"))
	fmt.Println(calcularSalario(172, "A"))
	fmt.Println(calcularSalario(176, "B"))
	fmt.Println(calcularSalario(162, "C"))
	fmt.Println(calcularSalario(200, "A"))
	fmt.Println(calcularSalario(200, "B"))
	fmt.Println(calcularSalario(200, "C"))
}

func calcularSalario(horasTrabalhadas float64, categoria string) float64 {
	maisDe160HorasTrabalhadas := horasTrabalhadas > 160
	var multiplicador float64
	adicional := 1.0

	if categoria == "A" {
		multiplicador = 3000

		if maisDe160HorasTrabalhadas {
			adicional = 1.5
		}
	}

	if categoria == "B" {
		multiplicador = 1500

		if maisDe160HorasTrabalhadas {
			adicional = 1.2
		}
	}

	if categoria == "C" {
		multiplicador = 1000
	}

	return horasTrabalhadas * multiplicador * adicional
}
