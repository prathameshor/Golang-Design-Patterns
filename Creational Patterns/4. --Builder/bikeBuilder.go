package main

type IBuilder interface {
	setBikeName()
	setBikeType()
	setEngineType()
	getBike() Bike
}

func getBikeBuilder(builderType string) IBuilder {
	if builderType == "adventure" {
		return newAdventureBuilder()
	} 
	
	if builderType == "sports" {
		return newSportsBuilder()
	}

	return nil
}

