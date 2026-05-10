package main

import "fmt"

// func makeCoffee(kind string, isSugar bool) {
// 	fmt.Printf("Making %s Coffee .....\n", kind)
// 	fmt.Println("Sugar added : ", isSugar)
// }

func makeCoffee(kind string) string {
	coffee := fmt.Sprintf("%s Coffee!", kind)
	return coffee
}

func main() {
	// makeCoffee("Espresso", false)
	// makeCoffee("Latte", true)

	myCoffee := makeCoffee("Cappuccino")
	fmt.Println("I am having", myCoffee)
}
