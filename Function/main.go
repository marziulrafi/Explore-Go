package main

import "fmt"

func makeCoffee(kind string, isSugar bool) {
	fmt.Printf("Making %s Coffee .....\n", kind)
	fmt.Println("Sugar added : ", isSugar)
}

func main() {
	makeCoffee("Espresso", false)
	makeCoffee("Latte", true)
}
