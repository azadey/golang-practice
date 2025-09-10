package main

import (
	"fmt"
)

func main() {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99

	addedPrices := append(prices, 5.99)
	addedPrices = prices[1:]
	fmt.Println(addedPrices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	addedPrices = append(addedPrices, discountPrices...)

	fmt.Println("Discount Prices: ", addedPrices)
}

// func main() {
// 	var productNames [4]string = [4]string{"Item 1 Item 2"}
// 	productNames[2] = "A Carpet"
// 	prices := [4]float64{10.99, 9, 45, 20.00}

// 	fmt.Println(prices[3])
// 	fmt.Println(productNames)

// 	featuredPrices := prices[1:]
// 	highlightedPrices := featuredPrices[:1]

// 	fmt.Println("Featured Prices", featuredPrices)
// 	fmt.Println("Highlighted ", highlightedPrices)
// 	fmt.Println("Highlighted len", len(highlightedPrices), cap(highlightedPrices))
// }
