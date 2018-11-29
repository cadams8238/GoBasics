package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}

	// SAME OUTPUT below, initializes map to zero value
	// var colors map[string]string
	// colors := make(map[string]string)

	// colors["white"] = "#fff" // always use []!! this assigns value to key
	// delete(colors, "white")  // deletes key/value pair, pass in like: delete(map, key to delete)

	// fmt.Printf("%v", colors)
	printMap(colors)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hexcode for", color, "is", hex)
	}
}
