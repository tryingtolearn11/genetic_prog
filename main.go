package main

import (
	"fmt"
	"ga/vistwitch/evolve"
	"ga/vistwitch/monkey"
	"html/template"
	"image"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// load the parent image
func loadImg(filePath string) *image.RGBA {
	img, err := os.Open(filePath)
	defer img.Close()
	if err != nil {
		fmt.Println("Error reading File")
	}

	pic, _, err := image.Decode(img)
	if err != nil {
		fmt.Println("err decoding file")
	}
	return pic.(*image.RGBA)
}

// templates
var monkey_tmpl = template.Must(template.New("tmpl").ParseFiles("templates/form.html", "templates/home.html", "templates/basictemplate.html"))
var picture_tmpl = template.Must(template.New("tmpl").ParseFiles("templates/picture.html"))

// part two
func input_picture(w http.ResponseWriter, r *http.Request) {
	if err := picture_tmpl.ExecuteTemplate(w, "picture.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	img := loadImg("./evolve/test_imgs/mona1.png")
	evolve.StartEvolution(w, r, img)

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

// -- Test : Upload Files

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File upload endpoint hit")

	// 10 << 20 specifies a max upload of 10 mb
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File : %+v\n", handler.Filename)
	fmt.Printf("File size : %+v\n", handler.Size)
	fmt.Printf("MIME header : %+v\n", handler.Header)

	// create temp file
	tempFile, err := ioutil.TempFile("./static/temp-images", "upload-*.jpg")
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)

	fmt.Fprintf(w, "successful upload")

}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", input_monkey)
	mux.HandleFunc("/picture", input_picture)
	mux.HandleFunc("/upload", uploadFile)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Starting Server on :5000")

	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)

}
