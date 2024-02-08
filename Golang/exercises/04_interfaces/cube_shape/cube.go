package cube

import "fmt"

type Cube struct {
	Length float64
}

func (c Cube) Volume() float64 {
	return c.Length * c.Length * c.Length
}

func (c Cube) PrintStructure() string {
	return "The length of the Cube is: " + fmt.Sprintf("%f", c.Length) + "\nThe volume of the Cube is: " + fmt.Sprintf("%f", c.Length*c.Length*c.Length)
}
