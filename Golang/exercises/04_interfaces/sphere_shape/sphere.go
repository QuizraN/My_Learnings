package sphere

import (
	"fmt"
	"math"
)

type Sphere struct {
	Radius float64
}

func (s Sphere) Volume() float64 {
	return (4 * math.Pi * math.Pow(s.Radius, float64(3))) / 3
}

func (s Sphere) PrintStructure() string {
	return "The Radius of the sphere is: " + fmt.Sprintf("%f", s.Radius) + "\nThe volume of the sphere is: " + fmt.Sprintf("%f", (4*math.Pi*math.Pow(s.Radius, float64(3)))/3)
}
