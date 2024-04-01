package main

import "fmt"

func main() {
	epsonPrinter := &Epson{}
	hpPrinter := &Hp{}

	win := &Windows{}

	win.SetPrinter(epsonPrinter)
	win.Print()
	fmt.Println()
	win.SetPrinter(hpPrinter)
	win.Print()
	fmt.Println()

	mac := &Mac{}

	mac.SetPrinter(epsonPrinter)
	mac.Print()
	fmt.Println()
	mac.SetPrinter(hpPrinter)
	mac.Print()
	fmt.Println()
}
