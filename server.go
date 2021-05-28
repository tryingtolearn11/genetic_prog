package main

import (
	"fmt"
	"math/rand"
)

type DNA struct {
	Phrase []byte
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
		Phrase: dummy,
	}

	// convert back to string just to test
	myString := string(dummy)

	//fmt.Println(dummy)
	fmt.Println(myString)
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

func main() {
	// test
	s := []byte("this is a test")
	//createDNA(s)
	createPopulation(s)
}

/*
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
*/
