package main

import "fmt"

type Endereco struct {
	rua     string
	numero  int
	cidade  string
	estado  string
}

type Pessoa struct {
	nome     string
	idade    int
	endereco Endereco
}

func endereco(p Pessoa) {
	fmt.Printf("Endereço de %s:\n", p.nome)
	fmt.Printf("Rua: %s, %d\n", p.endereco.rua, p.endereco.numero)
	fmt.Printf("Cidade: %s, Estado: %s\n", p.endereco.cidade, p.endereco.estado)
}

func main() {
	p := Pessoa{
		nome:  "João",
		idade: 30,
		endereco: Endereco{
			rua:     "Rua A",
			numero:  123,
			cidade:  "São Paulo",
			estado:  "SP",
		},
	}
	endereco(p)
}


