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

var number_of_polygons = 80
var mutationRate = 0.0001
var PopulationSize = 100

//var sidesNum = rand.Intn(6-3) + 3
var sidesNum = 3

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
	Fitness  int64
	DNA      *image.RGBA
}

func generatePolygon(n int, width float64, height float64, radius float64) (polygon Polygon) {
	r := float64(rand.Intn(255))
	g := float64(rand.Intn(255))
	b := float64(rand.Intn(255))
	//a := float64(rand.Intn(255))
	a := 0.5
	polygon = Polygon{
		Number_of_sides: n,
		Width:           width,
		Height:          height,
		Radius:          radius,
		Color:           []float64{r, g, b, a},
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
func calculateFitness(a *image.RGBA, b *image.RGBA) (fitness int64) {
	//fmt.Println("Len(a.Pix) : ", len(a.Pix))
	// go thru the pixels and find the difference
	var p float64
	for x := 0; x < len(a.Pix); x++ {
		p += math.Pow(float64(int64(a.Pix[x])-int64(b.Pix[x])), 2)
	}

	fitness = int64(math.Sqrt(p))

	//	fmt.Println("FITNESS :", fitness)
	if fitness == 0 {
		return 1
	} else {
		return fitness
	}
}

// Create a Population of x Entitys
func generatePopulation(i *image.RGBA) (population []Entity) {
	for k := 0; k < PopulationSize; k++ {
		organism := generateEntity(i)
		population = append(population, organism)
	}
	return population
}

func generateMatingPool(population []Entity, t *image.RGBA) (pool []Entity) {
	pool = make([]Entity, 0)

	// sort the population by fitness (the lower the fitness the better)
	sort.SliceStable(population, func(i, j int) bool {
		return population[i].Fitness < population[j].Fitness
	})
	Poolsize := 25
	top := population[0 : Poolsize+1]
	if top[len(top)-1].Fitness-top[0].Fitness == 0 {
		pool = population
		return
	}

	for i := 0; i < len(top)-1; i++ {
		num := int((top[Poolsize].Fitness - top[i].Fitness))
		//fmt.Println("number of times added : ", num)
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

		parentA := one
		parentB := two

		child := crossover(parentA, parentB)
		child.mutation()
		child.Fitness = calculateFitness(t, child.DNA)
		next_gen[i] = child
	}
	return next_gen
}

func crossover(parentA Entity, parentB Entity) (child Entity) {
	child = Entity{
		Polygons: make([]Polygon, len(parentA.Polygons)),
		Fitness:  0,
	}

	// 50% chance to come from either parent
	for i := 0; i < len(parentA.Polygons); i++ {
		chance := rand.Intn(100)
		if chance > 50 {
			child.Polygons[i] = parentA.Polygons[i]
		} else {
			child.Polygons[i] = parentB.Polygons[i]
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

// doesnt return the least fitness
func successor(p []Entity) (e Entity) {
	// just sort
	sort.SliceStable(p, func(i, j int) bool {
		return p[i].Fitness < p[j].Fitness
	})
	return p[0]
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
			pool := generateMatingPool(population, img)
			population = generateNextGeneration(pool, population, img)
			time_taken := time.Since(start)
			if generation%100 == 0 {
				fmt.Printf("\nTime : %s | Generation: %d | Fitness: %d | PoolSize: %d ", time_taken, generation, best.Fitness, len(pool))
				saveImg("../static/pictures/"+"dna.png", best.DNA)
			}
		}
	}

}
