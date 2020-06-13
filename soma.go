package main

import "fmt"

func Soma(a int, b int) (x int) {
	x = a + b
	return
}

func main() {
	fmt.Println(Soma(5,5))
}