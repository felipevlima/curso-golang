package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      12325.45,
		"Gabriela Silva": 15467.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.00
	delete(funcsESalarios, "Inexistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

}
