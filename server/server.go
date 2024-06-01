package main

import (
	utils "ascii/utlis"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Result string
}

func main() {
	http.HandleFunc("/", handleInput)
	http.HandleFunc("/utlis", handleAsciiArt)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("Starting server at port 8080")
	fmt.Println("URL : http://AScii-art-web:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleInput(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../statics/input.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func handleAsciiArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	asciiArt := utils.PrintLine(text, banner)
	data := PageData{Result: asciiArt}

	tmpl, err := template.ParseFiles("../statics/output.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
