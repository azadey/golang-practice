package main

import "fmt"

func main() {
	age := 32
	agePointer := &age

	fmt.Println("Age:", *agePointer)

	getAdultYerar(agePointer)
	fmt.Println("Adult in years:", *agePointer)
	fmt.Println("Age:", age)
}

func getAdultYerar(age *int) {
	*age = *age - 18
}
