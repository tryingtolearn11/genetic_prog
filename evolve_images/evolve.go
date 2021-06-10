package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
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


// going to pass randomized values for x, y and r here
func generateRandomizedCircle(width int, height int, radius int) (circle Circle) {
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

*/

// where to save generated image
func saveImg(filePath string, rgba *image.RGBA) {
	img, err := os.Create(filePath)
	defer img.Close()
	if err != nil {
		fmt.Println("Err creating File", err)
	}
	png.Encode(img, rgba.SubImage(rgba.Rect))
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

var number_of_polygons = 50
var mutationRate = 0.021

const S = 50
const W = 250
const H = 281

type Point struct {
	X float64
	Y float64
}

type Polygon struct {
	Number_of_sides int
	Width           float64
	Height          float64
	Radius          float64
	Color           []float64
}

// Need to define an entity
// will be composed of a slice of circles
type Entity struct {
	Polygons []Polygon
	Fitness  float64
	DNA      *image.RGBA
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

// Now to create the entity
// An Entity is composed of an array of Polygons
func generateEntity(i *image.RGBA) (entity Entity) {
	polygon_array := make([]Polygon, number_of_polygons)

	for k := 0; k < number_of_polygons; k++ {
		width := rand.Intn(i.Rect.Dx())
		height := rand.Intn(i.Rect.Dy())
		r := float64(rand.Intn(100))
		sidesNum := rand.Intn(6-3) + 3
		polygon_array[k] = generatePolygon(sidesNum, float64(width), float64(height), r)

	}

	entity_image := display(i.Rect.Dx(), i.Rect.Dy(), polygon_array)
	entity = Entity{
		Polygons: polygon_array,
		Fitness:  calculateFitness(i, entity_image),
		DNA:      entity_image,
	}

	return
}

func display(width int, height int, pa []Polygon) *image.RGBA {
	//const number_of_polygons = 120
	end := image.NewRGBA(image.Rect(0, 0, width, height))
	dc := gg.NewContextForRGBA(end)
	for _, poly := range pa {
		rotation := float64(rand.Intn(360))
		dc.DrawRegularPolygon(poly.Number_of_sides, poly.Width, poly.Height, poly.Radius, rotation)
		dc.Push()
		dc.SetRGBA(poly.Color[0], poly.Color[1], poly.Color[2], poly.Color[3])
		dc.SetLineWidth(1)
		dc.StrokePreserve()
		dc.Fill()
		dc.Pop()

	}

	return end
}

// 2 images are different = fitness of len(a.Pix),
// 2 images are same = fitness of 0
func calculateFitness(a *image.RGBA, b *image.RGBA) (fitness float64) {
	//fmt.Println("Len(a.Pix) : ", len(a.Pix))
	// go thru the pixels and find the difference
	var p float64
	for x := 0; x < len(a.Pix); x++ {
		p += math.Pow(float64(a.Pix[x])-float64(b.Pix[x]), 2)
	}

	fitness = math.Sqrt(p)

	//	fmt.Println("FITNESS :", fitness)
	return fitness
}

// Create a Population of 100 Entitys
func generatePopulation(i *image.RGBA) (population []Entity) {
	for k := 0; k < 100; k++ {
		organism := generateEntity(i)
		population = append(population, organism)
	}
	return population
}

// Create the mating pool
// sort the population by their fitness and find the difference between the
// best and the worst entity. The difference will the size of the pool and we
// will take the 'difference' amount of entites starting from the top going
// down
// TODO: Review This Function - The fitness keeps increasing
func generateMatingPool(population []Entity) (pool []Entity) {
	// sort the population by fitness (the lower the fitness the better)
	sort.SliceStable(population, func(i, j int) bool {
		return population[i].Fitness < population[j].Fitness
	})

	best_fitness := population[0].Fitness
	worst_fitness := population[(len(population) - 1)].Fitness

	difference := ((math.Abs(best_fitness - worst_fitness)) / 281000) * 100

	poolSize := int(difference)
	if poolSize >= len(population) {
		poolSize = 100
	}

	// lets try a pool of top 30 + the difference amount as the extra organisms
	for j := 0; j < poolSize+10; j++ {
		pool = append(pool, population[j])
	}

	/*

		fmt.Println("Best Fitness , Worst Fitness  : ", best_fitness, "  ", worst_fitness)
		fmt.Println("Difference : ", difference)
		fmt.Println("Mating Pool : ", len(pool))
			for i := 0; i < len(pool); i++ {
				fmt.Println(pool[i].Fitness)
			}
	*/

	return
}

func generateNextGeneration(pool []Entity, population []Entity, t *image.RGBA) []Entity {
	next_gen := make([]Entity, len(population))
	// make the next generation
	for i := 0; i < len(population); i++ {
		one := rand.Intn(len(pool))
		two := rand.Intn(len(pool))

		// make ParentA the dominant Parent
		// take the least random value and that will be
		// parentA
		/*
			var parentA Entity
			var parentB Entity
				if one < two {
					parentA = pool[one]
					parentB = pool[two]
				} else {
					parentA = pool[two]
					parentB = pool[one]
				}
		*/

		child := crossover(pool[one], pool[two])
		child.mutation()
		child.Fitness = calculateFitness(t, child.DNA)

		next_gen[i] = child
	}

	return next_gen
}

func crossover(parentA Entity, parentB Entity) (child Entity) {
	child = Entity{
		DNA: &image.RGBA{
			Pix:    make([]uint8, len(parentA.DNA.Pix)),
			Stride: parentA.DNA.Stride,
			Rect:   parentA.DNA.Rect,
		},
		Fitness: 0,
	}

	mid := rand.Intn(len(parentA.DNA.Pix))
	for j := 0; j < len(parentA.DNA.Pix); j++ {
		if j > mid {
			child.DNA.Pix[j] = parentA.DNA.Pix[j]
		} else {
			child.DNA.Pix[j] = parentB.DNA.Pix[j]
		}
	}

	return
}

func (e *Entity) mutation() {
	for j := 0; j < len(e.DNA.Pix); j++ {
		chance := rand.Float64()
		if chance <= mutationRate {
			e.DNA.Pix[j] = uint8(rand.Intn(255))
		}
	}
}

func successor(p []Entity) (e Entity) {
	model := float64(0)
	position := 0
	for i := 0; i < len(p); i++ {
		if p[i].Fitness > model {
			position = i
			model = p[i].Fitness
		}
	}
	return p[position]
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	start := time.Now()
	fmt.Println("Running evolve_pictures")
	match := false
	img := loadImg("./test_imgs/resized_clown.png")
	test_img := generateEntity(img)
	population := generatePopulation(test_img.DNA)
	generation := 0

	for !match {
		generation++
		best := successor(population)
		fmt.Println("Generation : ", generation)
		fmt.Println("Best Match : ", best.Fitness)

		if best.Fitness < 8000 {
			match = true
		} else {
			pool := generateMatingPool(population)
			population = generateNextGeneration(pool, population, img)
			time_taken := time.Since(start)
			if generation%10 == 0 {
				fmt.Printf("\nTime : %s | Generation: %d | Fitness: %f | PoolSize: %d ", time_taken, generation, best.Fitness, len(pool))
				saveImg("../static/pictures/"+"dna.png", test_img.DNA)
			}
		}
	}

	/*

		fmt.Println("population : ", len(population))
			for i := 0; i < len(nextGeneration); i++ {
				fmt.Println(nextGeneration[i].Fitness)
				//fmt.Println(len(nextGeneration[i].DNA.Pix))
			}
	*/

	// print tests
	//fmt.Println("ENTITY's FITNESS : ", test_img.Fitness)

}
