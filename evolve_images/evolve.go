package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"time"
)

// constants
var number_of_circles = 120
var circleSize = 10

// Consider : Maybe try using semi transparent colors
type Circle struct {
	X      int
	Y      int
	Radius int
	Color  color.Color
}

// Need to define an entity
// will be composed of a slice of circles
type Entity struct {
	Circles []Circle
	Fitness float64
	DNA     *image.RGBA
}

// going to pass randomized values for x, y and r here
func (c *Circle) generateRandomizedCircle(width int, height int, radius int) (circle Circle) {
	circle = Circle{
		X:      width,
		Y:      height,
		Radius: radius,
	}
	return
}

// Now to create the entity
func generateEntity(i *image.RGBA) (entity Entity) {
	circle_array := make([]Circle, number_of_circles)

	for k := 0; k < number_of_circles; k++ {
		width := rand.Intn(i.Rect.Dx())
		height := rand.Intn(i.Rect.Dy())
		r := rand.Intn(circleSize)
		circle_array[k].generateRandomizedCircle(width, height, r)
	}

	return
}

// TODO : try to create an entity image and render it to webpage

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("this is a test :D ")
}
