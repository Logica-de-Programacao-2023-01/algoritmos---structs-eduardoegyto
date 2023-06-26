package main

import "fmt"

type Musica struct {
	titulo   string
	artista  string
	duracao  int
}

type Playlist struct {
	nome    string
	musicas []Musica
}

func lista(p Playlist) {
	fmt.Println("Nome da Playlist:", p.nome)
	totalDuracao := 0
	for _, musica := range p.musicas {
		fmt.Println("Título:", musica.titulo)
		fmt.Println("Artista:", musica.artista)
		fmt.Println("Duração:", musica.duracao, "segundos")
		totalDuracao += musica.duracao
	}
	fmt.Println("Tempo Total:", totalDuracao, "segundos")
}

func main() {
	p := Playlist{
		nome: "Minha Playlist",
		musicas: []Musica{
			{titulo: "Música 1", artista: "Artista 1", duracao: 180},
			{titulo: "Música 2", artista: "Artista 2", duracao: 240},
			{titulo: "Música 3", artista: "Artista 3", duracao: 200},
		},
	}
	lista(p)
}
