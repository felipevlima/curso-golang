package main

import (
	"fmt"
	"time"
)

func fale(pessoa, texto string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", pessoa, texto, i+1)
	}
}

func main() {
	// fale("Maria", "Pq você não fala cmg?", 3)
	// fale("Joao", "Só posso falar depois de você!", 1)

	go fale("Maria", "Pq você não fala cmg?", 500)
	go fale("Joao", "Só posso falar depois de você!", 500)
	time.Sleep(time.Second * 20)
}
