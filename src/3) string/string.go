package main

import "fmt"

var teste = "teste"

func main() {
	x := 10
	y := "Bom dia!"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	x = 20
	fmt.Printf("x: %v, %T\n", x, x)

	fmt.Printf("teste: %v, %T", teste, teste)
}