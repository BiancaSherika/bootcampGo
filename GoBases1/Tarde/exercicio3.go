package main

import (
	"fmt"
)

var mes = 7

func main() {
	switch {
	case mes == 1:
		fmt.Println("1 de janeiro")
	case mes == 2:
		fmt.Println("2 de fevereiro")
	case mes == 3:
		fmt.Println("3 de março")
	case mes == 4:
		fmt.Println("4 de abril")
	case mes == 5:
		fmt.Println("5 de maio")
	case mes == 6:
		fmt.Println("6 de junho")
	case mes == 7:
		fmt.Println("7 de julho")
	case mes == 8:
		fmt.Println("8 de agosto")
	case mes == 9:
		fmt.Println("9 de setembro")
	case mes == 10:
		fmt.Println("10 de outubro")
	case mes == 11:
		fmt.Println("11 de novembro")
	case mes == 12:
		fmt.Println("12 de dezembro")
	default:
		fmt.Println("Não identificado.")
	}
}
