package main

import "fmt"

func main() {
	age := 18

	if age > 18 {
		fmt.Println("Adult")
	} else if age <= 18 && age > 12 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Child")
	}
}
