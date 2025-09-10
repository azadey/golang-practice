package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1
	hobbies := [3]string{"Cooking", "Gardening", "Foodie"}
	fmt.Println("Hobbies:", hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	//3
	hobbiesSlice := hobbies[:2]
	fmt.Println("Hobby Slice: ", hobbiesSlice)

	//4
	fmt.Println(cap(hobbiesSlice))
	hobbiesSlice = hobbiesSlice[1:3]
	fmt.Println(hobbiesSlice)

	//5
	goals := []string{"Complete it soon!", "Watch movies"}
	fmt.Println("Goals", goals)

	//6
	goals[1] = "Learn all courses!"
	goals = append(goals, "Sleep")
	fmt.Println("Goal 6: ", goals)

	//7
	products := []Product{
		{"1a", "first product", 12.99},
		{"2b", "second product", 8.9},
	}
	fmt.Println("Products Count: ", len(products))
	fmt.Println("Products: ", products)

	products = append(products, Product{"3b", "third product", 1.9})
	fmt.Println("3rd product", products)
}
