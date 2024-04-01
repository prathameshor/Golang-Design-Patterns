package main

import "fmt"

type FurnitureFactory interface {
	createChair() IChair
	createSofa() ISofa
}

func GetFurnitureFactory(furniture string) (FurnitureFactory, error) {
	if furniture == "modern" {
		return &Modern{}, nil
	}

	if furniture == "victorian" {
		return &Victorian{}, nil
	}

	return nil, fmt.Errorf("wrong furniture type passed")
}
