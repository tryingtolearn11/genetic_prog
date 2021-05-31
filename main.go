package main

import (
	"fmt"
	"ga/vistwitch/monkey"
	"log"
	"net/http"
)

/*

func input_handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		render(w, "form.html", nil)

	case "POST":
		input := r.FormValue("Phrase")

		fmt.Fprintf(w, "Phrase = %s\n", input)

		// run phrase func from pkg monkey
		//	monkey.Run_phrase(s)
		render(w, "form.html", nil)
	}

}

*/

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
		evolved_phrase := monkey.Run_phrase(w, r, s)
		fmt.Fprintln(w, "Best Phrase = ", evolved_phrase)
	default:
		fmt.Fprintf(w, "sorry only GET and POST methods")
	}
}

func main() {

	http.HandleFunc("/", input_handler)
	log.Fatal(http.ListenAndServe(":5000", nil))

}
