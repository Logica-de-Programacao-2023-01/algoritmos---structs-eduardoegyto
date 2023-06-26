package main

import "fmt"

type Triangulo struct {
	base   float64
	altura float64
}

func area(t Triangulo) float64 {
	return t.base * t.altura / 2
}

func main() {
	t := Triangulo{
		base:   4.0,
		altura: 5.0,
	}
	area := area(t)
	fmt.Printf("Área do triângulo: %.2f\n", area)
}
