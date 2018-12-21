package main

import (
	"../console/model"
	"fmt"
)

func main() {

	s := model.Song{
		Title: "Magix in the Hamptopns",
		Genre: "HipHop",
	}

	for true {
		fmt.Print("\n wil je de title, het genre of beiden ? \n")
		var result string
		fmt.Scan(&result)
		fmt.Print(analyser(result, s))
		fmt.Print("\n")
	}

}

func analyser(result string, s model.Song) string {
	switch result {
	case "title":
		return s.GetTitle()
	case "genre":
		return s.GetGenre()
	case "beiden":
		return s.ToString()
	default:
		return s.ToString()
	}
}
