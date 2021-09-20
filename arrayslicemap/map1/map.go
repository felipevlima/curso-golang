package main

import "fmt"

func main() {
	// Maps devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789878] = "Maria"
	aprovados[987654321000] = "Pedro"
	aprovados[986428139874] = "Ana"

	fmt.Println(aprovados)

	for cpf, name := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(aprovados[987654321000])
	delete(aprovados, 987654321000)
	fmt.Println(aprovados[987654321000])
}
