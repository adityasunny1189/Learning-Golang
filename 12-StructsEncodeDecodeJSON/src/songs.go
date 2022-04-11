package src

import "fmt"

type Song struct {
	title  string
	artist string
}

type Playlist struct {
	genre string
	songs []Song
}

func SongCollection() {
	rock := Playlist{
		genre: "indie rock",
		songs: []Song{
			{"badmash dil", "Jubin nautiyal"},
			{"dil mange more", "Abhijit singh"},
		},
	}

	fmt.Println(rock)
}
