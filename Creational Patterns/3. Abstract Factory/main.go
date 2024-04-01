package main

import (
	"fmt"
)

func main() {
	modernFactory, _ := GetFurnitureFactory("modern")
	victorianFactory, _ := GetFurnitureFactory("victorian")

	modernChair := modernFactory.createChair()
	modernSofa := modernFactory.createSofa()

	victorianChair := victorianFactory.createChair()
	victorianSofa := victorianFactory.createSofa()

	printChairDetails(modernChair)
	printChairDetails(victorianChair)

	printSofaDetails(modernSofa)
	printSofaDetails(victorianSofa)
}

func printChairDetails(s IChair) {
	fmt.Printf("Furniture: %s", s.getFurniture())
	fmt.Println()
	fmt.Printf("Type: %s", s.getType())
	fmt.Println()
}

func printSofaDetails(s ISofa) {
	fmt.Printf("Furniture: %s", s.getFurniture())
	fmt.Println()
	fmt.Printf("Type: %s", s.getType())
	fmt.Println()
}
