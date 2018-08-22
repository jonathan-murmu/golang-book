package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	// Program to convert Fahrenheit into Celsius
	fmt.Println("Enter Fahrenheit:")
	var temperature int
	fmt.Scanf("%d", &temperature)

	fmt.Println((temperature - 32) * 5/9)

}