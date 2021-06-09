package main

import (
	"fmt"
	"image"
	"image/color"
	"math/rand"
	"os"
	"time"

	"github.com/fogleman/gg"
)

/*
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


// returns an array of Points
func Polygon(number_of_sides int) []Point {
	result := make([]Point, number_of_sides)
	for i := 0; i < number_of_sides; i++ {
		a := float64(i)*2*math.Pi/float64(number_of_sides) - math.Pi/2
		result[i] = Point{math.Cos(a), math.Sin(a)}
	}
	return result
}

func display(width int, height int, circle_array []Circle) (i *image.RGBA) {
	end := image.NewRGBA(image.Rect(0, 0, width, height))
	dc := gg.NewContext(width, height)

	return end
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
*/
// TODO : try to create an entity image and render it to webpage

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

type Point struct {
	X float64
	Y float64
}

// Need to define an entity
// will be composed of a slice of circles
type Entity struct {
	Polygons []Polygon
	Fitness  float64
	DNA      *image.RGBA
}

type Polygon struct {
	Number_of_sides int
	Width           float64
	Height          float64
	Radius          float64
	Color           []float64
}

func generatePolygon(n int, width float64, height float64, radius float64) (polygon Polygon) {
	r := float64(rand.Intn(255))
	g := float64(rand.Intn(255))
	b := float64(rand.Intn(255))
	//a := float64(rand.Intn(255))
	polygon = Polygon{
		Number_of_sides: n,
		Width:           width,
		Height:          height,
		Radius:          radius,
		Color:           []float64{r, g, b, 0.5},
	}
	return
}

func display(width int, height int) {
	const S = 50
	const W = 500
	const H = 500
	const number_of_polygons = 120
	//end := image.NewRGBA(image.Rect(0, 0, width, height))
	dc := gg.NewContext(width, height)
	for k := 0; k < number_of_polygons; k++ {
		//sidesNum := rand.Intn((5 - 3) + 3)
		x_pos := float64(rand.Intn(W))
		y_pos := float64(rand.Intn(H))
		radius := float64(rand.Intn(100))
		rotation := float64(rand.Intn(360))

		dc.DrawRegularPolygon(3, x_pos, y_pos, radius, rotation)
		dc.Push()
		dc.SetRGBA(r, g, b, 0.5)
		dc.SetLineWidth(1)
		//dc.SetHexColor("#FFCC00")
		dc.StrokePreserve()
		//dc.SetHexColor("#FFE43A")
		dc.Fill()
		dc.Pop()

	}

	dc.SavePNG("../static/pictures/" + "output.png")
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("this is a test :D ")
	img := loadImg("./test_imgs/clown.png")
	display(img.Rect.Dx(), img.Rect.Dy())
	//test_img := generateEntity(img)
	//saveImg("../static/pictures/"+"result.png", result)
}
