package main

import (
	"fmt"
	"ga/vistwitch/monkey"
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.New("tmpl").ParseFiles("form.html"))

func input_handler(w http.ResponseWriter, r *http.Request) {

	if err := tmpl.ExecuteTemplate(w, "form.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//switch r.Method {
	//case "POST":
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

	http.HandleFunc("/", input_handler)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
