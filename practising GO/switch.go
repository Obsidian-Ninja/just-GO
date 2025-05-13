package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch statement
	i := 2

	switch i {
	case 1:
		fmt.Sprintln(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("default")
	}

	//multiple condition switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("Workday")
	}

	//type switch

	// see the notes

}
