package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15462.78,
			"Gugu Pereira":   87545.9,
		},
		"J": {
			"João José": 11234.98,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)

		for nome, salario := range funcs {
			fmt.Println(nome, salario)
		}
	}
}
