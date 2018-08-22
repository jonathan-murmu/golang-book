package main

import "fmt"

func first() {
	fmt.Println("1")
}

func second() {
	fmt.Println("2")
}

func third() {
	fmt.Println("3")
	defer second()
}
func main() {
	defer third()
	first()
}