package main

import "fmt"

func main() {
	for i := 0; i <= 512; i++ {
		fmt.Print(i, "\t", string(rune(i)), "\t")
		if i%16 == 0 {
			fmt.Println()
		}
	}
}
