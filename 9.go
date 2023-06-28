package main

import (
	"fmt"
)

type Aluno struct {
	nome  string
	idade int
	notas []float64
}

func (a *Aluno) adicionarNota(nota float64) {
	a.notas = append(a.notas, nota)
}

func (a *Aluno) removerNota(index int) {
	if index >= 0 && index < len(a.notas) {
		a.notas = append(a.notas[:index], a.notas[index+1:]...)
	}
}

func (a *Aluno) calcularMédia() float64 {
	total := 0.0
	for _, nota := range a.notas {
		total += nota
	}
	média := total / float64(len(a.notas))
	return média
}

func (a *Aluno) imprimirInformações() {
	fmt.Println("Nome:", a.nome)
	fmt.Println("Idade:", a.idade)
	fmt.Println("Média:", a.calcularMédia())
}

func main() {
	aluno := Aluno{
		nome:  "Maria",
		idade: 20,
		notas: []float64{7.5, 8.0, 9.5},
	}

	aluno.adicionarNota(8.5)
	aluno.imprimirInformações()

	aluno.removerNota(1)
	aluno.imprimirInformações()
}
