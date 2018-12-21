package main

import "fmt"

var musicList []string

func main() {

	for true {
		var antwoord string
		fmt.Print("druk get of post: \n")
		fmt.Scan(&antwoord)
		var result []string
		result = handler(antwoord)
		fmt.Println(result)
		fmt.Print("\n")
	}

}

func handler(antwoord string) []string {
	switch antwoord {
	case "get":
		return get()
	case "post":
		return post()
	default:
		return get()

	}
}

func get() []string {
	return musicList
}

func post() []string {
	fmt.Print("\n vul een liedje in: \n")
	var liedje string
	fmt.Scan(&liedje)
	musicList = append(musicList, liedje)
	return musicList
}
