package main

import "fmt"

//constants cannot be changed later
const age = 19

var ID = 1

func main() {

	fmt.Println(age)
	fmt.Println(ID)

	const name = "Gurvansh"

	fmt.Println(name)

	//constant group

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(port, host)

}
