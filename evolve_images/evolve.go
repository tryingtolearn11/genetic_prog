package main

import (
	"fmt"
	"github.com/llgcode/draw2d/draw2dimg"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"time"
)

// constants
var number_of_circles = 130
var circleSize = 15

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

func display(width int, height int, circle_array []Circle) (i *image.RGBA) {
	end := image.NewRGBA(image.Rect(0, 0, width, height))
	gc := draw2dimg.NewGraphicContext(end)

	for _, c := range circle_array {
		gc.SetFillColor(c.Color)
		gc.MoveTo(float64(c.X), float64(c.Y))
		gc.ArcTo(float64(c.X), float64(c.Y), float64(c.Radius), float64(c.Radius), 0, 6.283185)
		gc.Close()
		gc.Fill()
	}
	return end
}

// load the parent image
func loadImg(filePath string) *image.RGBA {
	img, err := os.Open(filePath)
	defer img.Close()
	if err != nil {
		fmt.Println("Error reading File")
	}

	pic, _, err := image.Decode(img)
	if err != nil {
		fmt.Println("err decoding file")
	}
	return pic.(*image.RGBA)
}

// where to save generated image
func saveImg(filePath string, rgba *image.RGBA) {
	img, err := os.Create(filePath)
	defer img.Close()
	if err != nil {
		fmt.Println("Err creating File", err)
	}
	png.Encode(img, rgba.SubImage(rgba.Rect))
}

// going to pass randomized values for x, y and r here
func generateRandomizedCircle(width int, height int, radius int) (circle Circle) {
	// TODO: try to aim for semi-transparent
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	a := uint8(rand.Intn(255))
	circle = Circle{
		X:      width,
		Y:      height,
		Radius: radius,
		Color:  color.RGBA{r, g, b, a},
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
		circle_array[k] = generateRandomizedCircle(width, height, r)
	}

	entity = Entity{
		Circles: circle_array,
		Fitness: 0,
		DNA:     display(i.Rect.Dx(), i.Rect.Dy(), circle_array),
	}

	return
}

// TODO : try to create an entity image and render it to webpage

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("this is a test :D ")

	img := loadImg("./test_imgs/clown.png")
	test_img := generateEntity(img)
	saveImg("../static/pictures/"+"clown2.png", test_img.DNA)

}
