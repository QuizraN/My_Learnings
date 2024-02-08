package interfaces

import (
	box "exercises/04_interfaces/box_shape"
	cube "exercises/04_interfaces/cube_shape"
	sphere "exercises/04_interfaces/sphere_shape"
	"fmt"
)

type ofStructure interface {
	Volume() float64
}
type stringer interface {
	PrintStructure() string
}

func calculateVolume(kind ofStructure, called string) {
	fmt.Printf("The Volume calculated for our %s is: %f \n", called, kind.Volume())
}
func displayStructure(kind stringer, called string) {
	fmt.Printf("The Structure of %s is:\n", called)
	fmt.Println(kind.PrintStructure())
}

func Display() {

	c := cube.Cube{
		Length: 7,
	}

	b := box.Box{
		Length: 5.5,
		Width:  5.5,
		Height: 7.7,
	}

	s := sphere.Sphere{
		Radius: 7.14,
	}

	calculateVolume(c, "Cube")
	calculateVolume(b, "Box")
	calculateVolume(s, "Sphere")
	displayStructure(c, "Cube")
	displayStructure(b, "Box")
	displayStructure(s, "Sphere")
}
