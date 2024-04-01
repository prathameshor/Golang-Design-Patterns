package main

type IChair interface {
	setFurniture(kind string)
	setType(style string)
	getFurniture() string
	getType() string
}

type Chair struct {
	Furniture string
	Type      string
}

func(c *Chair) setFurniture(kind string){
	c.Furniture = kind
}

func(c *Chair) setType(style string){
	c.Type = style
}

func(c *Chair) getFurniture() string{
	return c.Furniture
}

func(c *Chair) getType() string{
	return c.Type
}
