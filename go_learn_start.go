package main

import "fmt"

func main() {
	fmt.Println("Wee!")
	for i := 1; i <= 10; i++ {
		for j := 1; j<=10; j++ {
			fmt.Print(i * j, "\t")
		}
		fmt.Println()
	}
}
