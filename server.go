package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

type DNA struct {
	Phrase  []byte
	Fitness float64
}

// create a DNA
func createDNA(p []byte) (dna DNA) {

	// random ascii values into a string
	dummy := make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		dummy[i] = byte(rand.Intn(95) + 32)
	}

	// set new dna from the random string
	dna = DNA{
		Phrase:  dummy,
		Fitness: measureFitness(p, dummy),
	}

	return
}

// need to measure the fitness
func measureFitness(p []byte, dummy []byte) (fitness float64) {
	score := 0
	for i := 0; i < len(p); i++ {
		if p[i] == dummy[i] {
			score++
		}
	}

	fitness = float64(score) / float64(len(p))

	return
}

func createPopulation(p []byte) (population []DNA) {
	for i := 0; i < 100; i++ {
		organism := createDNA(p)
		population = append(population, organism)
	}

	fmt.Println("Size of Population : ", len(population))
	return population
}

// Build Mating Pool
func generateMatingPool(p []byte, population []DNA, fit float64) []DNA {

	var matingPool []DNA
	count := 0
	for j := 0; j < len(population); j++ {
		// find percentage score
		n := int((population[j].Fitness / fit) * 100)
		count += n
		// add this dna N number of times
		for k := 0; k < n; k++ {
			matingPool = append(matingPool, population[j])
		}
	}

	newPool := make([]DNA, len(population))

	for i := 0; i < len(population); i++ {

		// parents selection : 2 parents to mimic human reproduction
		one := rand.Intn(len(matingPool))
		two := rand.Intn(len(matingPool))

		//TODO: Make sure parents are NOT the same values

		parent1 := matingPool[one]
		parent2 := matingPool[two]

		child := crossover(parent1, parent2)
		child.mutate()
		// just to measure the child's fitness
		child.Fitness = measureFitness(p, child.Phrase)
		newPool[i] = child
	}
	return newPool

}

// crossover to generate the child DNA
func crossover(p1 DNA, p2 DNA) (child DNA) {
	// get random midpoint
	midpoint := rand.Intn(len(p1.Phrase))
	//fmt.Println("midpoint ", midpoint)

	// initialize child phrase size
	child.Phrase = make([]byte, len(p1.Phrase))

	// perform crossover
	for i := 0; i < len(p1.Phrase); i++ {
		if i > midpoint {
			child.Phrase[i] = p1.Phrase[i]

		} else {
			child.Phrase[i] = p2.Phrase[i]
		}
	}

	return
}

func (child *DNA) mutate() {
	mutateRate := 0.01
	for i := 0; i < len(child.Phrase); i++ {

		if rand.Float64() < mutateRate {
			child.Phrase[i] = byte(rand.Intn(95) + 32)

		}
	}
}

func successor(population []DNA) DNA {
	position := 0
	model := 0.0
	for i := 0; i < len(population); i++ {
		if population[i].Fitness > model {
			model = population[i].Fitness
			position = i
		}
	}
	return population[position]
}

func main() {
	// for constant random numbers
	rand.Seed(time.Now().UTC().UnixNano())
	match := false
	s := []byte("to be or not to be")
	population := createPopulation(s)

	for !match {
		best := successor(population)
		fmt.Printf("\r best fitness : %2f", best.Fitness)
		fmt.Println("Phrase : ", string(best.Phrase))

		if bytes.Compare(best.Phrase, s) == 0 {
			match = true
		} else {
			peak := best.Fitness
			population = generateMatingPool(s, population, peak)

		}
	}
}
