package calculator

import (
	"fmt"

	"github.com/AdityaSoni128/tablebuilderwithgo"
)

func Add(a, b int) int {
	return a + b
}

func Subtract(a, b int) int {
	return a - b
}

func Multiply(a, b int) int {
	return a * b
}

func Divide(a, b int) int {
	return a / b
}

func Tablegetter(a int) {
	fmt.Println("TAble getter called")
	tablebuilderwithgo.GetTable(a)
}
