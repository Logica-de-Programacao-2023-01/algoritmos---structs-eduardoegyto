package main

import (
	"fmt"
)

type Filme struct {
	título     string
	diretor    string
	ano        int
	avaliações []int
}

func (f *Filme) adicionarAvaliação(avaliação int) {
	f.avaliações = append(f.avaliações, avaliação)
}

func (f *Filme) removerAvaliação(index int) {
	if index >= 0 && index < len(f.avaliações) {
		f.avaliações = append(f.avaliações[:index], f.avaliações[index+1:]...)
	}
}

func (f *Filme) calcularMédiaAvaliações() float64 {
	total := 0
	for _, avaliação := range f.avaliações {
		total += avaliação
	}
	média := float64(total) / float64(len(f.avaliações))
	return média
}

func (f *Filme) imprimirInformações() {
	fmt.Println("Título:", f.título)
	fmt.Println("Diretor:", f.diretor)
	fmt.Println("Ano:", f.ano)
	fmt.Println("Média de avaliações:", f.calcularMédiaAvaliações())
}

func main() {
	filme := Filme{
		título:     "Matrix",
		diretor:    "Lana Wachowski",
		ano:        1999,
		avaliações: []int{5, 4, 4, 3, 5},
	}

	filme.adicionarAvaliação(4)
	filme.imprimirInformações()

	filme.removerAvaliação(2)
	filme.imprimirInformações()
}
