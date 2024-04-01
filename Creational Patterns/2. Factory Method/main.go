package main

import (
	"fmt"

	guns "github.com/prathameshor/design-patterns/factory_method/Guns"
)

func main() {
	ar := guns.NewAssaultRifle()
	mg := guns.NewMachineGun()

	fmt.Printf("Name of gun: %s", ar)
	fmt.Printf("Name of gun: %s", mg)
}