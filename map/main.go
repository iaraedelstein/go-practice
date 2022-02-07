package main

import "fmt"

func main() {

	colors := map[string]string {
		"red": "#ff000",
		"green": "#4bf745",
	}

	//zero values
	//var colors map[string]string

	// colors:=make(map[string]string)
	// colors["red"]="#ff000"
	// delete(colors, "red")

	printMap(colors)
}


func printMap(c map[string]string){
	for k,v := range c {
		fmt.Println("Hex color for", k, "is", v)
	}
}