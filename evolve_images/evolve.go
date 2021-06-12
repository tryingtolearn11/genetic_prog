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

var number_of_polygons = 90
var mutationRate = 0.001
var PopulationSize = 60

//var sidesNum = rand.Intn(6-3) + 3

var sidesNum = 3

//const S = 50

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
	for k := 0; k < PopulationSize; k++ {
		organism := generateEntity(i)
		population = append(population, organism)
	}
	return population
}

// TODO: Review This Function - The fitness keeps increasing
func generateMatingPool(population []Entity) (pool []Entity) {
	// sort the population by fitness (the lower the fitness the better)
	sort.SliceStable(population, func(i, j int) bool {
		return population[i].Fitness < population[j].Fitness
	})
	Poolsize := 10
	top := population[0 : Poolsize+1]
	if top[len(top)-1].Fitness-top[0].Fitness == 0 {
		pool = population
		return
	}

	for i := 0; i < len(top)-1; i++ {
		num := int((top[Poolsize].Fitness - top[i].Fitness))
		for n := 0; n < num; n++ {
			pool = append(pool, top[i])
		}
	}

	return
}

func generateNextGeneration(pool []Entity, population []Entity, t *image.RGBA) []Entity {
	next_gen := make([]Entity, len(population))
	// make the next generation
	for i := 0; i < len(population); i++ {
		one := pool[rand.Intn(len(pool))]
		two := pool[rand.Intn(len(pool))]

		// make ParentA the dominant Parent
		// take the least random value and that will be
		// parentA.
		// OR PERHAPS : Compare the fitnesses and make the least the dominant
		// parent!
		var parentA Entity
		var parentB Entity
		if one.Fitness < two.Fitness {
			parentA = one
			parentB = two
		} else {
			parentA = two
			parentB = one
		}

		child := crossover(parentA, parentB)
		child.mutation()
		child.Fitness = calculateFitness(t, child.DNA)

		next_gen[i] = child
	}

	return next_gen
}

//TODO : BUG was here. Reminder to review
func crossover(parentA Entity, parentB Entity) (child Entity) {
	child = Entity{
		Polygons: make([]Polygon, len(parentA.Polygons)),
		//Fitness:  0,
	}

	mid := len(parentA.Polygons) / 2
	for j := 0; j < len(parentA.Polygons); j++ {
		if j <= mid {
			child.Polygons[j] = parentA.Polygons[j]
		} else {
			child.Polygons[j] = parentB.Polygons[j]
		}
	}

	child.DNA = display(parentA.DNA.Rect.Dx(), parentA.DNA.Rect.Dy(), child.Polygons)

	return child
}

// mutate the []polygon
func (e *Entity) mutation() {
	for j := 0; j < len(e.Polygons); j++ {
		chance := rand.Float64()
		if chance < mutationRate {
			r := float64(rand.Intn(100))
			e.Polygons[j] = generatePolygon(sidesNum, float64(e.DNA.Rect.Dx()), float64(e.DNA.Rect.Dy()), r)
		}
	}
	e.DNA = display(e.DNA.Rect.Dx(), e.DNA.Rect.Dy(), e.Polygons)
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
		//		fmt.Println("Generation : ", generation)
		//		fmt.Println("Best Match : ", best.Fitness)

		if best.Fitness < 8000 {
			match = true
		} else {
			pool := generateMatingPool(population)
			population = generateNextGeneration(pool, population, img)
			time_taken := time.Since(start)
			if generation%20 == 0 {
				fmt.Printf("\nTime : %s | Generation: %d | Fitness: %f | PoolSize: %d ", time_taken, generation, best.Fitness, len(pool))
				saveImg("../static/pictures/"+"dna.png", test_img.DNA)
			}
		}
	}

}
