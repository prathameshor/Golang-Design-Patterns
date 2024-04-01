package main

type CheeseTopings struct{
	pizza IPizza
}

func(c *CheeseTopings) getPrice() int{
	price := c.pizza.getPrice() 
	return price + 7
}

