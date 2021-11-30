package main

import (
	"fmt"
	"guess-the-name/student/main.go/github/Ascii-art-web/asciiart"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func download(w http.ResponseWriter, r *http.Request) {

	userInput := r.FormValue("text")
	userBanner := r.FormValue("banner")

	//If package error, Delete text on row 18 and write it again. Then save & import should fix itself automatically.
	s := asciiart.Generate(userInput, userBanner)
	w.Header().Set("Content-Disposition", "attachment; filename=ascii.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(s)))

	fmt.Fprintf(w, s)
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ascii-art/download", download)
	fmt.Println("Server is running @ localhost:8080")
	if http.ListenAndServe(":8080", nil) != nil {
		log.Fatalf("%v - Internal Server Error", http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("templates/ascii.html"))
	if r.URL.Path != "/" {
		http.Error(w, "404 - Not Found", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		template.Execute(w, nil)
	case "POST":
		text, banner := r.FormValue("text"), r.FormValue("banner")
		output := asciiart.Generate(text, banner)
		err := template.Execute(w, output)
		if err != nil {
			http.Error(w, "400 - Bad Request", http.StatusBadRequest)
			return
		}
	}
}
