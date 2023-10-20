package main

import "fmt"

var x int = 30
var y string = "nome"

func main(){

	tempA := 64.09
	tempB := "sessenta e quatro virgula zero nove"
	tempC := 60
	fmt.Printf("O valor e tipo da tempA é: %v (%T). O valor e tipo da tempB é: %v (%T). O valor e tipo da tempC é: %v (%T).", tempA, tempA, tempB, tempB, tempC, tempC)
}