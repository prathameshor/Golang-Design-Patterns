package main

type Box struct {
	components []Component
	id int
}

func (b *Box) totalPrice() int {
	var price int
	for _, i := range b.components {
		price += i.totalPrice()
	}
	return price
}

func (b *Box) add(c Component) {
	b.components = append(b.components, c)
}
