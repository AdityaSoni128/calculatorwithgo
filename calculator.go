package Calculator

import (
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
	tablebuilderwithgo.GetTable(a)
}
