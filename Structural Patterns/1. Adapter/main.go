package main

import "fmt"

func main() {
	cli := &Client{}
	ser := &Service{}
	adp := &Adapter{ser: *ser}	
	fmt.Println(cli.CalculateInt(*adp))
}
