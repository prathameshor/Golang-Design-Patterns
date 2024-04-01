package main

import "fmt"

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Printing from windows")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}
