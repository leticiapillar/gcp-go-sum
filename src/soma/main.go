package main

import "fmt"

func main() {
	fmt.Println("Resultado da soma é:", sum(5, 5))
}

func sum(a int, b int) int {
	return a + b
}