package main

import (
	"fmt"
)

func main() {
if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Please enter a word.")
		return
	}

	prevR := '\n'
	for _, r := range os.Args[1] {
		if prevR == r {
			fmt.Println("Fits")
			return
		}
		prevR = r
	}

	fmt.Println("Does not fit")
}
