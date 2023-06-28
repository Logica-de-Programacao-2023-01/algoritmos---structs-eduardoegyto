package main

import (
	"fmt"
)

type Animal struct {
	nome   string
	espécie string
	idade  int
	som    string
}

func (a *Animal) modificarSom(novoSom string) {
	a.som = novoSom
}

func (a *Animal) imprimirInformações() {
	fmt.Println("Nome:", a.nome)
	fmt.Println("Espécie:", a.espécie)
	fmt.Println("Idade:", a.idade)
	fmt.Println("Som:", a.som)
}

func main() {
	animal := Animal{
		nome:   "Rex",
		espécie: "Cachorro",
		idade:  5,
		som:    "Latido",
	}

	animal.modificarSom("Au au")
	animal.imprimirInformações()
}
