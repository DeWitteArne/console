package main

import (
	"../console/model"
	"fmt"
)

var songs []model.Song

func main() {
	s := model.Song{
		Genre: "HipHop",
		Title: "Magix in the Hamptons",
	}

	songs = append(songs, s)

	for true {
		fmt.Print("get, delete of post ?")
		var answer string
		fmt.Scan(&answer)
		analyse(answer)
	}
}

func analyse(answer string) {
	switch answer {
	case "post":
		poster()
	case "get":
		getter()
	case "delete":
		deleter()
	default:
		getter()

	}
}

func poster() {
	fmt.Print("\n vul een genre in: \n")
	var genre string
	fmt.Scan(&genre)
	fmt.Print("\n Vul een title in: \n")
	var title string
	fmt.Scan(&title)
	s2 := model.New(title, genre)
	songs = append(songs, s2)
	fmt.Print("Gelukt! \n")
}

func getter() {
	fmt.Println(songs)
}

func deleter() {
	var to_delete string
	fmt.Print("\n Vul de title van het te verwijderen liedje in: \n")
	fmt.Scan(&to_delete)

	for i, s := range songs {
		if s.GetTitle() == to_delete {
			songs = append(songs[:i], songs[i+1:]...)
		}
	}

	fmt.Print("Delete successful \n")

}
