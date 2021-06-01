package main

import (
	"fmt"
	"ga/vistwitch/monkey"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/home" {
		http.NotFound(w, r)
		return
	}

	if err := tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var tmpl = template.Must(template.New("tmpl").ParseFiles("templates/form.html", "templates/home.html"))

var data monkey.Output

func input_handler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "form.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Parse the Input String
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	input := r.FormValue("Phrase")
	if len(input) == 0 {
		fmt.Fprintf(w, "Please enter your phrase")
	} else {
		fmt.Fprintf(w, "Phrase = %s\n", input)
		s := []byte(input)
		// Run the genetic program
		monkey.Run_phrase(w, r, s)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", input_handler)
	mux.HandleFunc("/home", home)
	log.Println("Starting Server on :5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)

}
