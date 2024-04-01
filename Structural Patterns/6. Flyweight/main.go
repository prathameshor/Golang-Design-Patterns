package main

import "fmt"

func main() {
	product := Product{}
	product.AddPizza("Dominos", "Regular", "Cheese", "Thin")
	product.AddPizza("Dominos", "Regular", "Cheese", "Thin")
	product.AddPizza("Pizza Hut", "Medium", "Mushroom", "Thin")

	fmt.Println(product)
}
