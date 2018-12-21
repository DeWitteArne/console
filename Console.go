package main

import (
	"fmt"
	"strings"
)

func main() {
	var response string

	for true {
		fmt.Print("\n Hey, alles goed ?")

		fmt.Scan(&response)

		var result = checker(response)

		fmt.Print(result)
	}

}

func checker(response string) string {
	if strings.ToLower(response) == "ja" {
		return "Dat is heel goed manneke"
	} else {
		return "Dat is niet zo goed manneke"
	}
}
