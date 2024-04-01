package main

type TomatoToppings struct{
	pizza IPizza
}

func(t *TomatoToppings) getPrice() int{
	return t.pizza.getPrice() + 5
}