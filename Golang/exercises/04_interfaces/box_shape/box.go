package box

import "fmt"

type Box struct {
	Length float64
	Width  float64
	Height float64
}

func (b Box) Volume() float64 {
	return b.Length * b.Width * b.Height
}

func (b Box) PrintStructure() string {
	return "The length, width, height of the box is: " + fmt.Sprintf("%f, %f,%f", b.Length, b.Width, b.Height) + "\nThe volume of the box is: " + fmt.Sprintf("%f", b.Length*b.Width*b.Height)
}
