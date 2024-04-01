package main

type Product struct{
	price int
}

func(p *Product) totalPrice() int{
	return p.price
}