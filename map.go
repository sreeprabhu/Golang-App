package main

import "fmt"

func main() {
	// different ways to define map
	colors1 := map[string]string{
		"red":   "#ff000",
		"green": "#4b000",
		"white": "#fffff",
	}

	var colors2 map[string]string

	colors3 := make(map[string]string)

	// delete keys from a map
	delete(colors1, "red")

	printMap(colors1)
	fmt.Println(colors2, colors3)
}

func printMap(m map[string]string) {
	for color, hex := range m {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
