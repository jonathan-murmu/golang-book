package main

import "fmt"

func generateEven() func() uint {
	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	next_even := generateEven()
	fmt.Println(next_even())
	fmt.Println(next_even())
	fmt.Println(next_even())
}