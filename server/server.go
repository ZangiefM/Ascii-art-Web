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
	http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("../statics/"))))
	fmt.Println("Starting server at port 8080")
	fmt.Println("URL : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8082", nil))
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
	bannerData, _ := utils.HandleBanner(banner)
	asciiArt := utils.PrintAscii(text, bannerData)
	data := PageData{Result: asciiArt}

	tmpl, err := template.ParseFiles("../statics/input.html/{{.Result}}")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
