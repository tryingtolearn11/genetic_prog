package main

import (
	"fmt"
	"ga/vistwitch/monkey"
	"log"
	"net/http"
)

func input_handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		input := r.FormValue("Phrase")
		fmt.Fprintf(w, "Phrase = %s\n", input)
		s := []byte(input)
		monkey.Run_phrase(w, r, s)
	default:
		fmt.Fprintf(w, "Only GET and POST")
	}
}

func main() {

	http.HandleFunc("/", input_handler)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
