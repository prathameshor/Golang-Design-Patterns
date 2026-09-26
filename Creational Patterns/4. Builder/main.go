package main

import "fmt"

func main() {
	advBike := getBikeBuilder("adventure")
	sportsBike := getBikeBuilder("sports")

	director := newDirector(advBike)
	newAdv := director.buildBike()
	fmt.Println(newAdv)

	director = newDirector(sportsBike)
	newSports := director.buildBike()
	fmt.Println(newSports)
}
