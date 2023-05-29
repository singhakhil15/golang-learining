package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":    "#ff09",
		"green":  "#nf45",
		"white":  "#34fvb",
		"yellow": "#46fgh",
	}

	fmt.Println(colors)
	printMap(colors)
	delete(colors, "red")
	fmt.Println(colors)

	var colorsA map[string]string
	fmt.Println(colorsA)

	colorsB := make(map[string]string)

	colorsB["white"] = "#hu45"
	//colorsA["black"] = "#r456v"

	fmt.Println(colorsA)
	fmt.Println(colorsB)
}

func printMap(c map[string]string) {
	for key, val := range c {
		fmt.Println(key + " ===== " + val)
	}
}
