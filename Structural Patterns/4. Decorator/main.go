package main

import "fmt"

func main() {
	pizza := &VeggiMania{}

	pizzaWithTomato := &TomatoToppings{pizza: pizza}

	pizzaWithCheese := &CheeseTopings{pizza: pizza}

	pizzaWithCheeseAndTomato := &CheeseTopings{pizza: pizzaWithTomato}

	fmt.Println(pizzaWithTomato.getPrice())
	fmt.Println(pizzaWithCheese.getPrice())
	fmt.Println(pizzaWithCheeseAndTomato.getPrice())
}
