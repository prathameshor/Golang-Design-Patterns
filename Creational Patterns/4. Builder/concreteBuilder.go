package main

type NewAdventureBuilder struct {
	bikeType   string
	engineType string
	bikeName   string
}

func newAdventureBuilder() *NewAdventureBuilder {
	return &NewAdventureBuilder{}
}

func (n *NewAdventureBuilder) setBikeName() {
	n.bikeName = "Honda 500nx"
}

func (n *NewAdventureBuilder) setBikeType() {
	n.bikeType = "ADV"
}

func (n *NewAdventureBuilder) setEngineType() {
	n.engineType = "Parallel Twin"
}

func (n *NewAdventureBuilder) getBike() Bike {
	return Bike{
		bikeName:   n.bikeName,
		bikeType:   n.bikeType,
		engineType: n.engineType,
	}
}

/*-----*/

type NewSportsBuilder struct {
	bikeType   string
	engineType string
	bikeName   string
}

func newSportsBuilder() *NewSportsBuilder {
	return &NewSportsBuilder{}
}

func (n *NewSportsBuilder) setBikeName() {
	n.bikeName = "Honda CBR250R"
}

func (n *NewSportsBuilder) setBikeType() {
	n.bikeType = "Sports"
}

func (n *NewSportsBuilder) setEngineType() {
	n.engineType = "Single Cylinder"
}

func (n *NewSportsBuilder) getBike() Bike {
	return Bike{
		bikeName:   n.bikeName,
		bikeType:   n.bikeType,
		engineType: n.engineType,
	}
}
