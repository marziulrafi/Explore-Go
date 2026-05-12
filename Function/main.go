package main

import "fmt"

// func makeCoffee(kind string, isSugar bool) {
// 	fmt.Printf("Making %s Coffee .....\n", kind)
// 	fmt.Println("Sugar added : ", isSugar)
// }

func makeCoffee(kind string) (string, int) {
	coffee := fmt.Sprintf("%s Coffee!", kind)
	price := 30
	return coffee, price
}

func main() {
	// makeCoffee("Espresso", false)
	// makeCoffee("Latte", true)

	myCoffee, mybill := makeCoffee("Cappuccino")
	fmt.Printf("I am having a %d$ %s \n", mybill, myCoffee)
}
