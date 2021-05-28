package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DNA struct {
	Phrase  []byte
	Fitness float32
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

	// convert back to string just to test
	//myString := string(dna.Phrase)

	//fmt.Println(dummy)
	//fmt.Println(myString)
	//fmt.Println("Fitness : ", dna.Fitness)
	return
}

// need to measure the fitness
func measureFitness(p []byte, dummy []byte) (fitness float32) {
	score := 0
	for i := 0; i < len(p); i++ {
		if p[i] == dummy[i] {
			score++
		}
	}

	//fmt.Println("score : ", score)

	fitness = float32(score) / float32(len(p))
	//fmt.Println("fitness : ", fitness)

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
// TODO: Might be bugged, remember size of pool is > 100
// TODO: NOTE that 0 DNA objects make it through into the pool -keep them out
func generateMatingPool(p []byte) {
	population := createPopulation(p)
	for i := 0; i < len(population); i++ {
		fmt.Println(population[i].Fitness)
	}

	matingPool := make([]DNA, len(population))

	for j := 0; j < len(population); j++ {
		// find percentage score
		n := int(population[j].Fitness * 100)
		// add this dna N number of times
		for k := 0; k < n; k++ {
			matingPool = append(matingPool, population[j])
		}
	}

	// mating pool
	/*
		for i := 0; i < len(matingPool); i++ {
			fmt.Println(matingPool[i])
		}
	*/

	fmt.Println("Size of matingPool : ", len(matingPool))

	// parents selection : 2 parents to mimic human reproduction
	one := rand.Intn(len(matingPool))
	two := rand.Intn(len(matingPool))
	fmt.Println("PARENTS 1 & 2 : ", one, two)

	//TODO: Make sure parents are NOT the same values

	parent1 := matingPool[one]
	parent2 := matingPool[two]

	//fmt.Println("PARENTS 1 & 2 : ", parent1, parent2)
	child := crossover(parent1, parent2)
	fmt.Println("Child DNA : ", child)
}

// crossover to generate the child DNA
func crossover(p1 DNA, p2 DNA) (child DNA) {
	fmt.Println(p1, p2)
	// child

	// get random midpoint
	midpoint := rand.Intn(len(p1.Phrase))
	fmt.Println("midpoint ", midpoint)

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

func main() {
	// for constant random numbers
	rand.Seed(time.Now().UTC().UnixNano())

	// test
	//s := []byte("to be or not to be")

	s := []byte("w,qI8Te'$/Z'{&>d98")
	generateMatingPool(s)
	//createDNA(s)

	/*
		t := []byte("to be dr vot do be")
		fmt.Println(len(s), len(t))
		test_fitness := measureFitness(s, t)
		fmt.Println("test_fitness : ", test_fitness)
	*/

}
