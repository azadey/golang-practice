package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// websites := map[string]string{
	// 	"Google": "google.com",
	// 	"AWS":    "aws.com",
	// }
	// fmt.Println(websites)
	// fmt.Println(websites["Google"])

	// websites["Facebook"] = "facebook.com"
	// fmt.Println(websites)

	// delete(websites, "AWS")
	// fmt.Println(websites)

	userNames := make([]string, 2, 5)
	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Xin")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 2)
	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.2
	courseRatings["node"] = 3.8

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("U :", value, index)
	}

	// Maps for

	for k, val := range courseRatings {
		fmt.Println("C: ", k, val)
	}

}
