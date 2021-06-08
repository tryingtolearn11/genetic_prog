package main

import (
	"fmt"
	"image"
	"image/color"
)

// Position
/*
type Point struct {
	X int
	Y int
}
*/

type Circle struct {
	X      int
	Y      int
	Radius int
	Color  color.Color
}

// Need to define an entity
type Entity struct {
	Circles []Circle
	Fitness float64
	DNA     *image.RGBA
}

func main() {
	fmt.Println("this is a test :D ")
}
