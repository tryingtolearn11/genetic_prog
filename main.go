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
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	input := r.FormValue("Phrase")
	fmt.Fprintf(w, "Phrase = %s\n", input)
	s := []byte(input)
	monkey.Run_phrase(w, r, s)

	//default:
	//	fmt.Fprintf(w, "Only GET and POST")
	//}
}

func main() {

	http.HandleFunc("/", input_handler)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
