package main

type ISofa interface {
	setFurniture(kind string)
	setType(style string)
	getFurniture() string
	getType() string
}

type Sofa struct {
	Furniture string
	Type      string
}

func(s *Sofa) setFurniture(kind string) {
	s.Furniture = kind
}

func(s *Sofa) setType(style string) {
	s.Type = style
}

func(s *Sofa) getFurniture() string{
	return s.Furniture
}

func(s *Sofa) getType() string{
	return s.Type
}