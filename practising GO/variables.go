package main

import "fmt"

func main() {
	var name string = "golang"
	fmt.Println(name)

	//if unused, go gives warning
	//shortcut

	var name2 = "golang"
	fmt.Println(name2)

	//bool
	var isAdult = true
	fmt.Println(isAdult)

	//shorthand syntax
	price := 40
	fmt.Println(price)

	//only declaration

	var name3 string
	name3 = "golang"
	fmt.Println(name3)
}
