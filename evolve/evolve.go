//package evolve

package main

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

// mona1 : 150 x 225
var number_of_polygons = 100
var mutationRate = 0.01
var PopulationSize = 100
var Poolsize = 15

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
	X int
	Y int
}
type Polygon struct {
	PointOne   Point
	PointTwo   Point
	PointThree Point
	Color      []float64
}

// Need to define an entity
// will be composed of a slice of circles
type Entity struct {
	Polygons []Polygon
	Fitness  int64
	DNA      *image.RGBA
}

// Output data
type Data struct {
	Time       string
	Fitness    string
	Peak       string
	Population string
	Generation string
	SizePool   string
}

func generatePolygon(width int, height int) (polygon Polygon) {
	r := float64(rand.Intn(255))
	g := float64(rand.Intn(255))
	b := float64(rand.Intn(255))
	a := float64(rand.Intn(255))
	p1 := Point{X: rand.Intn(width), Y: rand.Intn(height)}
	//p2 := Point{X: rand.Intn(width), Y: rand.Intn(height)}
	//p3 := Point{X: rand.Intn(width), Y: rand.Intn(height)}
	p2 := Point{X: p1.X + (rand.Intn(100) - 15), Y: p1.Y + (rand.Intn(100) - 15)}
	p3 := Point{X: p1.X + (rand.Intn(100) - 15), Y: p1.Y + (rand.Intn(100) - 15)}
	polygon = Polygon{
		PointOne:   p1,
		PointTwo:   p2,
		PointThree: p3,
		Color:      []float64{r, g, b, a},
	}
	return
}

func display(width int, height int, pa []Polygon) *image.RGBA {
	end := image.NewRGBA(image.Rect(0, 0, width, height))
	dc := gg.NewContextForRGBA(end)
	for _, poly := range pa {
		dc.MoveTo(float64(poly.PointOne.X), float64(poly.PointOne.Y))
		dc.LineTo(float64(poly.PointTwo.X), float64(poly.PointTwo.Y))
		dc.LineTo(float64(poly.PointThree.X), float64(poly.PointThree.Y))
		dc.Push()
		dc.SetRGBA(poly.Color[0], poly.Color[1], poly.Color[2], poly.Color[3])
		dc.SetLineWidth(1)
		dc.StrokePreserve()
		dc.Fill()
		dc.Pop()
	}

	return end
}

// Prints the Data to Screen
func displayData(width int, height int, data Data) {
	const s = 250
	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Time : "+data.Time, s/2, 25, 0.5, 0.5)
	dc.DrawStringAnchored("Generation : "+data.Generation, s/2, 50, 0.5, 0.5)
	dc.DrawStringAnchored("Fitness : "+data.Fitness, s/2, 75, 0.5, 0.5)
	dc.DrawStringAnchored("Peak : "+data.Peak, s/2, 100, 0.5, 0.5)
	dc.DrawStringAnchored("Pool Size : "+data.SizePool, s/2, 125, 0.5, 0.5)
	dc.DrawStringAnchored("Population : "+data.Population, s/2, 150, 0.5, 0.5)
	dc.SavePNG("../static/pictures/" + "data.png")
}

// Now to create the entity
// An Entity is composed of an array of Polygons
func generateEntity(i *image.RGBA) (entity Entity) {
	polygon_array := make([]Polygon, number_of_polygons)
	for k := 0; k < number_of_polygons; k++ {
		polygon_array[k] = generatePolygon(i.Rect.Dx(), i.Rect.Dy())
	}

	entity_image := display(i.Rect.Dx(), i.Rect.Dy(), polygon_array)
	entity = Entity{
		Polygons: polygon_array,
		Fitness:  0,
		DNA:      entity_image,
	}
	entity.calculateFitness(i)
	return
}

func (e *Entity) calculateFitness(a *image.RGBA) {
	fitness := difference(e.DNA, a)
	if fitness == 0 {
		e.Fitness = 1
	} else {
		e.Fitness = fitness
	}
}

func difference(a *image.RGBA, b *image.RGBA) (p int64) {
	p = 0
	for i := 0; i < len(a.Pix); i++ {
		p += int64(squareD(a.Pix[i], b.Pix[i]))
	}
	return int64(math.Sqrt(float64(p)))
}

func squareD(a, b uint8) uint64 {
	k := uint64(a) - uint64(b)
	return k * k
}

// Create a Population of x Entitys
func generatePopulation(i *image.RGBA) (population []Entity) {
	population = make([]Entity, PopulationSize)
	for k := 0; k < PopulationSize; k++ {
		population[k] = generateEntity(i)
	}
	return population
}

func generateMatingPool(population []Entity, t *image.RGBA) (pool []Entity) {
	pool = make([]Entity, 0)

	// sort the population by fitness (the lower the fitness the better)
	sort.SliceStable(population, func(i, j int) bool {
		return population[i].Fitness < population[j].Fitness
	})

	if Poolsize > PopulationSize {
		Poolsize = PopulationSize - 1
	}
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

		parentA := one
		parentB := two

		child := crossover(parentA, parentB)
		child.mutation()
		child.calculateFitness(t)
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
	chance := rand.Intn(100)
	for i := 0; i < len(parentA.Polygons); i++ {
		//		chance := rand.Intn(100)
		if chance > 50 {
			child.Polygons[i] = parentB.Polygons[i]
		} else {
			child.Polygons[i] = parentA.Polygons[i]
		}
	}
	/*
		midpoint := rand.Intn(len(parentA.Polygons))
		for i := 0; i < len(parentA.Polygons); i++ {
			if i > midpoint {
				child.Polygons[i] = parentA.Polygons[i]
			} else {
				child.Polygons[i] = parentB.Polygons[i]
			}
		}
	*/

	child.DNA = display(parentA.DNA.Rect.Dx(), parentA.DNA.Rect.Dy(), child.Polygons)
	return

}

// mutate the []polygon
func (e *Entity) mutation() {
	for j := 0; j < len(e.Polygons); j++ {
		chance := rand.Float64()
		if chance < mutationRate {
			e.Polygons[j] = generatePolygon(int(e.DNA.Rect.Dx()), int(e.DNA.Rect.Dy()))
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
	img := loadImg("./test_imgs/mona1.png")

	test_img := generateEntity(img)
	population := generatePopulation(test_img.DNA)
	generation := 0
	prev_best := test_img
	peakEntity := test_img
	prev_best.Fitness = int64(9999999)
	d := Data{}

	for !match {
		generation++
		best := successor(population)
		// tracking the peak fitness
		if best.Fitness < peakEntity.Fitness {
			peakEntity = best
		}

		if best.Fitness < 20 {
			match = true
		} else {
			pool := generateMatingPool(population, img)
			population = generateNextGeneration(pool, population, img)
			// store the best fitness before looping
			prev_best = best
			time_taken := time.Since(start)
			gg.SavePNG("../static/pictures/"+"fogbranch.png", peakEntity.DNA)

			d = Data{Time: fmt.Sprint(time_taken), Fitness: fmt.Sprint(best.Fitness), Peak: fmt.Sprint(peakEntity.Fitness), Generation: fmt.Sprint(generation),
				Population: fmt.Sprint(PopulationSize), SizePool: fmt.Sprint(len(pool))}

			// Save Points
			if generation%10 == 0 {
				//fmt.Printf("\rTime : %s | Generation: %d | Fitness: %d | PoolSize: %d | Peak: %d |", time_taken, generation, best.Fitness, len(pool), peakEntity.Fitness)
				gg.SavePNG("../static/pictures/"+"fogbranch.png", peakEntity.DNA)
				displayData(250, 350, d)
				//fmt.Println("\nStats : ", d)
			}
		}
	}

}
