package main

import (
	"ga/vistwitch/monkey"
	/*
		"html/template"
		"io"
		"log"
		"net/http"
	*/)

/*
func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "sorry, something went wrong", http.StatusInternalServerError)
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "sorry, something went wrong", http.StatusInternalServerError)
	}
}

func input_handler(w http.ResponseWriter, r *http.Request) {

	input := r.FormValue("Phrase")
	io.WriteString(w, input)

	//fmt.Fprintf(w, input, r.URL.Path[1:])
	// run phrase func from pkg monkey
	//	monkey.Run_phrase(s)
	render(w, "form.html", nil)

}
*/

func main() {
	monkey.Run_phrase()

	/*
		http.HandleFunc("/", input_handler)
		log.Fatal(http.ListenAndServe(":5000", nil))
	*/

}
