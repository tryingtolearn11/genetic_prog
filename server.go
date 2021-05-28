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
	myString := string(dna.Phrase)

	//fmt.Println(dummy)
	fmt.Println(myString)
	fmt.Println("Fitness : ", dna.Fitness)
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

func createPopulation(p []byte) {
	population := []DNA{}
	for i := 0; i < 100; i++ {
		organism := createDNA(p)
		population = append(population, organism)
	}

	fmt.Println("Size of Population : ", len(population))

}

// Build Mating Pool

func main() {
	// test
	//s := []byte("to be or not to be")
	s := []byte("w,qI8Te'$/Z'{&>d98")
	//createDNA(s)
	createPopulation(s)

	/*
		t := []byte("to be dr vot do be")
		fmt.Println(len(s), len(t))
		test_fitness := measureFitness(s, t)
		fmt.Println("test_fitness : ", test_fitness)
	*/

}
