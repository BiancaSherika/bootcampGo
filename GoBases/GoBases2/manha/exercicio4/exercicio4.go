package exercicio4

import (
	"errors"
	"fmt"
	"math"
)

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func minFunc(values ...float64) float64 {
	result := math.Inf(1)

	for _, value := range values {
		if value < result {
			result = value
		}
	}

	return result
}

func averageFunc(values ...float64) float64 {
	accumulator := 0.0

	for _, value := range values {
		accumulator += value
	}

	return accumulator / float64(len(values))
}

func maxFunc(values ...float64) float64 {
	result := math.Inf(-1)

	for _, value := range values {
		if value > result {
			result = value
		}
	}

	return result
}

func operation(operation string) (func(values ...float64) float64, error) {
	if operation == minimum {
		return minFunc, nil
	}

	if operation == average {
		return averageFunc, nil
	}

	if operation == maximum {
		return maxFunc, nil
	}

	return nil, errors.New("Erro, operação não definida.")
}

func Run() {
	minValue, _ := operation(minimum)
	avgValue, _ := operation(average)
	maxValue, _ := operation(maximum)

	fmt.Println("Exercicio 4:")
	fmt.Println(minValue(2, 3, 3, 4, 10, 2, 4, 5))
	fmt.Println(avgValue(2, 3, 3, 4, 10, 2, 4, 5))
	fmt.Println(maxValue(2, 3, 3, 4, 10, 2, 4, 5))

	undefined, err := operation("palavra")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		undefined(2, 3, 4)
	}
}
