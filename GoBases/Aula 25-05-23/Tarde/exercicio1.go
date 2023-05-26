package main

import (
	"fmt"
	"strings"
)

var palavra = "teste"

func main() {
	fmt.Println(len(palavra))
	for _, slice := range strings.Split(palavra, "") {
		fmt.Println(slice)
	}
}
