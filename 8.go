package main

import (
	"fmt"
)

type Viagem struct {
	origem  string
	destino string
	data    string
	preço   float64
}

func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	maiorPreço := 0.0
	var viagemMaisCara Viagem

	for _, viagem := range viagens {
		if viagem.preço > maiorPreço {
			maiorPreço = viagem.preço
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{
			origem:  "São Paulo",
			destino: "Rio de Janeiro",
			data:    "2023-07-10",
			preço:   250.0,
		},
		{
			origem:  "São Paulo",
			destino: "Curitiba",
			data:    "2023-08-15",
			preço:   180.0,
		},
		{
			origem:  "São Paulo",
			destino: "Salvador",
			data:    "2023-09-20",
			preço:   350.0,
		},
	}

	viagemMaisCara := encontrarViagemMaisCara(viagens)
	fmt.Println("Viagem mais cara:")
	fmt.Println("Origem:", viagemMaisCara.origem)
	fmt.Println("Destino:", viagemMaisCara.destino)
	fmt.Println("Data:", viagemMaisCara.data)
	fmt.Println("Preço:", viagemMaisCara.preço)
}
