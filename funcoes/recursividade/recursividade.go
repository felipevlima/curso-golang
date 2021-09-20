package main

import "fmt"

func factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		fatorialAnterior, _ := factorial(n - 1)
		return n * fatorialAnterior, nil
	}
}

func main() {
	resultado, _ := factorial(5)
	fmt.Println(resultado)

	_, err := factorial(-4)
	if err != nil {
		fmt.Println(err)
	}
}
