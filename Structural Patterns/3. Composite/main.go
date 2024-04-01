package main

import "fmt"

func main() {
	product1 := &Product{price: 200}
	product2 := &Product{price: 300}
	product3 := &Product{price: 500}

	box1 := &Box{
		id: 1,
	}
	box1.add(product1)

	box2 := &Box{
		id: 2,
	}
	box2.add(product2)
	box2.add(product3)
	box1.add(box2)

	fmt.Println(box1.totalPrice())
}
