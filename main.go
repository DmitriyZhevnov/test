package main

import "fmt"

func main() {
	const (
		name        = "Roman"
		dateOfBirth = "12.03.2008"
	)
	age := 16
	var married bool = false
	balance := 100.204
	unused := "variable"
	_ = unused
	fmt.Println(" name:", name, "\n", "dateOfBirth:", dateOfBirth, "\n", "age:", age, "\n", "married:", married, "\n", "balance:", balance)

	// fmt.Println("name:", name)
	// fmt.Println("dateOfBirth:", dateOfBirth)
	// fmt.Println("age:", age)
	// fmt.Println("married:", married)
	// fmt.Println("balance:", balance)
}
