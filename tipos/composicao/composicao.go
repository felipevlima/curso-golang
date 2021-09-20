package main

import "fmt"

type exportivo interface {
	ligarTurbo()
}

type luxuoso interface {
	fazerBaliza()
}

type exportivoLuxuoso interface {
	exportivo
	luxuoso
}

type bwm7 struct{}

func (b bwm7) ligarTurbo() {
	fmt.Println("Turbo...")
}

func (b bwm7) fazerBaliza() {
	fmt.Println("Baliza...")
}

func main() {
	var b exportivoLuxuoso = bwm7{}
	b.ligarTurbo()
	b.fazerBaliza()
}
