package main

import (
	"fmt"
	"ga/vistwitch/monkey"
	"html/template"
	"log"
	"net/http"
)

// templates
var monkey_tmpl = template.Must(template.New("tmpl").ParseFiles("templates/form.html", "templates/home.html", "templates/basictemplate.html"))
var picture_tmpl = template.Must(template.New("tmpl").ParseFiles("templates/picture.html"))

// part two
func input_picture(w http.ResponseWriter, r *http.Request) {
	if err := picture_tmpl.ExecuteTemplate(w, "picture.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// part one : Monkey
func input_monkey(w http.ResponseWriter, r *http.Request) {
	if err := monkey_tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
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
	mux.HandleFunc("/", input_monkey)
	mux.HandleFunc("/picture", input_picture)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static", http.StripPrefix("/static", fileServer))
	log.Println("Starting Server on :5000")

	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)

}
