package interfaces

import (
	"fmt"
	"math"
)

type cube struct {
	length float64
}

type box struct {
	length float64
	width  float64
	height float64
}

type sphere struct {
	radius float64
}

type ofStructure interface {
	volume() float64
}
type stringer interface {
	printStructure() string
}

func (c cube) volume() float64 {
	return c.length * c.length * c.length
}

func (s sphere) volume() float64 {
	return (4 * math.Pi * math.Pow(s.radius, float64(3))) / 3
}

func (b box) volume() float64 {
	return b.length * b.width * b.height
}
func (c cube) printStructure() string {
	return "The length of the cube is: " + fmt.Sprintf("%f", c.length) + "\nThe volume of the cube is: " + fmt.Sprintf("%f", c.length*c.length*c.length)
}

func (s sphere) printStructure() string {
	return "The radius of the sphere is: " + fmt.Sprintf("%f", s.radius) + "\nThe volume of the sphere is: " + fmt.Sprintf("%f", (4*math.Pi*math.Pow(s.radius, float64(3)))/3)
}

func (b box) printStructure() string {
	return "The length, width, height of the box is: " + fmt.Sprintf("%f, %f,%f", b.length, b.width, b.height) + "\nThe volume of the box is: " + fmt.Sprintf("%f", b.length*b.width*b.height)
}

func calculateVolume(kind ofStructure, called string) {
	fmt.Printf("The Volume calculated for our %s is: %f \n", called, kind.volume())
}
func displayStructure(kind stringer, called string) {
	fmt.Printf("The Structure of %s is:\n", called)
	fmt.Println(kind.printStructure())
}

func Display() {

	c := cube{
		length: 7,
	}

	b := box{
		length: 5.5,
		width:  5.5,
		height: 7.7,
	}

	s := sphere{
		radius: 7.14,
	}

	calculateVolume(c, "Cube")
	calculateVolume(b, "Box")
	calculateVolume(s, "Sphere")
	displayStructure(c, "Cube")
	displayStructure(b, "Box")
	displayStructure(s, "Sphere")
}
