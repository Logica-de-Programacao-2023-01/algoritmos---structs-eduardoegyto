package main

import (
	"fmt"
	"time"
)

type Funcionário struct {
	nome   string
	salário float64
	idade  int
}

func (f *Funcionário) aumentarSalário(porcentagem float64) {
	f.salário += f.salário * (porcentagem / 100)
}

func (f *Funcionário) diminuirSalário(porcentagem float64) {
	f.salário -= f.salário * (porcentagem / 100)
}

func (f *Funcionário) tempoServiço() int {
	idadeAtual := time.Now().Year() - f.idade
	tempoServiço := idadeAtual - 18
	return tempoServiço
}

func main() {
	funcionário := Funcionário{
		nome:   "João",
		salário: 2000.0,
		idade:  1985,
	}

	fmt.Println("Salário inicial:", funcionário.salário)
	funcionário.aumentarSalário(10)
	fmt.Println("Salário após aumento:", funcionário.salário)
	funcionário.diminuirSalário(5)
	fmt.Println("Salário após redução:", funcionário.salário)

	tempoServiço := funcionário.tempoServiço()
	fmt.Println("Tempo de serviço:", tempoServiço, "anos")
}
