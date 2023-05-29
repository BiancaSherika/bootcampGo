package main

import (
	"fmt"
)

var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
var sum int

func main() {
	fmt.Println("A idade de Benjamin é", employees["Benjamin"])

	for _, a := range employees {
		if a > 21 {
			sum++
		}
	}

	fmt.Println(sum, "funcionários tem mais de 21 anos")

	employees["Federico"] = 25
	delete(employees, "Pedro")
}
