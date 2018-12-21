package main

import "fmt"

func main() {

	for true {

		var getal1 int
		var getal2 int
		var operator string
		fmt.Print("vul het eerste getal in \n")
		fmt.Scan(&getal1)
		fmt.Print("vul het tweede getal in \n")
		fmt.Scan(&getal2)
		fmt.Print("kies een bewerking (+ - * /) \n")
		fmt.Scan(&operator)
		fmt.Printf("De uitkomst is: ")

		fmt.Print(controler(getal1, getal2, operator))
		fmt.Print("\n")
	}

}

func controler(getal1 int, getal2 int, operator string) int {
	var result int
	switch operator {
	case "+":
		result = plus(getal1, getal2)
	case "-":
		result = min(getal1, getal2)
	case "*":
		result = maal(getal1, getal2)
	case "/":
		result = delen(getal1, getal2)
	}

	return result
}

func plus(getal1 int, getal2 int) int {
	return getal1 + getal2
}

func min(getal1 int, getal2 int) int {
	return getal1 - getal2
}

func maal(getal1 int, getal2 int) int {
	return getal1 * getal2
}

func delen(getal1 int, getal2 int) int {
	return getal1 / getal2
}
