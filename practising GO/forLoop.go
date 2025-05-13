package main

import "fmt"

func main() {
	//only construct in go for loops

	//while loop
	i := 3
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	//infinite loop

	// for {

	// }

	// classic for loop

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	//range

	for i := range 3 {
		fmt.Println(i)
	}
}
