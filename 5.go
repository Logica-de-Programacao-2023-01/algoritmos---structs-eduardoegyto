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

func listarPlaylistsPorTitulo(titulo string, playlists []Playlist) []Playlist {
	var playlistsEncontradas []Playlist
	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == titulo {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}
	return playlistsEncontradas
}

func main() {
	playlists := []Playlist{
		{
			nome: "Playlist 1",
			musicas: []Musica{
				{titulo: "Música 1", artista: "Artista 1", duracao: 180},
				{titulo: "Música 2", artista: "Artista 2", duracao: 240},
				{titulo: "Música 3", artista: "Artista 3", duracao: 200},
			},
		},
		{
			nome: "Playlist 2",
			musicas: []Musica{
				{titulo: "Música 4", artista: "Artista 4", duracao: 220},
				{titulo: "Música 5", artista: "Artista 5", duracao: 260},
			},
		},
		{
			nome: "Playlist 3",
			musicas: []Musica{
				{titulo: "Música 2", artista: "Artista 2", duracao: 240},
				{titulo: "Música 6", artista: "Artista 6", duracao: 190},
			},
		},
	}

	titulo := "Música 2"
	playlistsEncontradas := listarPlaylistsPorTitulo(titulo, playlists)

	if len(playlistsEncontradas) > 0 {
		fmt.Printf("Playlists com o título '%s':\n", titulo)
		for _, playlist := range playlistsEncontradas {
			fmt.Println(playlist.nome)
		}
	} else {
		fmt.Println("Nenhuma playlist encontrada com o título", titulo)
	}
}
