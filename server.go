package main

import (
	"fmt"
	"net/http"
)

type DNA struct {
	Phrase []byte
}

// create a DNA
//func createDNA(p []byte) (dna DNA) {}

/*
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":5000", nil)
}
*/
