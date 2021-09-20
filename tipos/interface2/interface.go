package main

import "fmt"

type esportivo interface {
	ligarTurbo()
}

type ferrari struct {
	modelo          string
	turboLigado     bool
	valocidadeAtual int
}

func (f *ferrari) ligarTurbo() {
	f.turboLigado = true
	f.valocidadeAtual = int(float64(f.valocidadeAtual) * 1.3)
}

func main() {
	carro1 := ferrari{"F40", false, 50}
	carro1.ligarTurbo()

	var carro2 esportivo = &ferrari{"F40", false, 0}
	carro2.ligarTurbo()

	fmt.Println(carro1, carro2)
}
