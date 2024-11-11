package main

import "fmt"

func main() {
	//this is one way to declare a map

	// colors := map[string]string{
	// 	"red": "#ff0000",
	// 	"green": "#4bf745",
	//	"white": "ffffff"
	// }

	//this is a way to declare an empty map
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"



	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}