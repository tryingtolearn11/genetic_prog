package main

import (
	"fmt"
	"math/rand"
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
func generateMatingPool(p []byte) {
	population := createPopulation(p)
	for i := 0; i < len(population); i++ {
		fmt.Println(population[i].Fitness)
	}

	matingPool := make([]DNA, len(population))

	for j := 0; j < len(population); j++ {
		// find percentage score
		n := int(population[j].Fitness * 100)
		fmt.Print(" N : ", n)
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

}

func main() {
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
