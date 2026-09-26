package main

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

// func (d *Director) setBuilder(b IBuilder) {
//     d.builder = b
// }

func (d *Director) buildBike() Bike {
	d.builder.setBikeName()
	d.builder.setBikeType()
	d.builder.setEngineType()
	return d.builder.getBike()
}
