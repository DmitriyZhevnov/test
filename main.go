package main

import "fmt"

func main() {
	const (
		name        string = "Димас"
		dateOfBirth        = "08.11.2001"
	)
	age := 22
	var married bool
	balance := 18.59
	var unused string
	_ = unused
	fmt.Println("name:", name, "\n", "dateOfBirth:", dateOfBirth, "\n", "age:", age, "\n", "married:", married, "\n", "balance:", balance)

}
