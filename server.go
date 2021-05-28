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
	for i := 0; i < len(p); i++ {
		p[i] = byte(rand.Intn(95) + 32)
	}

	myString := string(p)

	fmt.Println(p)
	fmt.Println(myString)
	return
}

func main() {
	// test
	s := []byte("this is a test")
	createDNA(s)
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
