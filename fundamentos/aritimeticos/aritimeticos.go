package main

import "fmt"

func main() {
	a := 3
	b := 2

	fmt.Println("Soma =>", a+b)

	// bitwise
	fmt.Println("E =>", a&b)   // 11 & 10 = 10
	fmt.Println("Ou =>", a|b)  // 11 | 10 = 11
	fmt.Println("Xor =>", a^b) // 11 ^ 10 = 01
}
