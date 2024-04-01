package main

import "fmt"

type Pizza struct {
	size     string
	toppings string
	crust    string
}

type PizzaBrand struct {
	brand string
	pizza *Pizza
}

type Product struct {
	pizzaList []PizzaBrand
}

var pizzaList = make([]Pizza, 0)

func GetPizza(size, toppings, crust string) *Pizza {
	for _, item := range pizzaList {
		if item.size == size && item.toppings == toppings && item.crust == crust {
			fmt.Println("Pizza already available.")
			return &item
		}
	}
	fmt.Println("Making fresh pizza.")
	newPizza := Pizza{size, toppings, crust}
	pizzaList = append(pizzaList, newPizza)
	return &newPizza
}

func (p *Product) AddPizza(brand, size, toppings, crust string) {
	pizza := GetPizza(size, toppings, crust)
	newPizza := &PizzaBrand{
		brand: brand,
		pizza: pizza,
	}
	p.pizzaList = append(p.pizzaList, *newPizza)
}
