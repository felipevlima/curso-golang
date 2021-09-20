package main

import "fmt"

func main() {
	type Pessoa struct {
		Nome string
	}

	p1 := Pessoa{"Felipe"}
	p2 := Pessoa{"Felipe"}
	fmt.Println("Pessoas:", p1 == p2)

}
