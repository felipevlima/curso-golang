package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritimetica de ponteiros

	var p *int = nil
	p = &i // Pegando endereço da variavel
	*p++   // desreferenciando (pegando o valor)
	i++
	fmt.Println(p, *p, i, &i)
}
