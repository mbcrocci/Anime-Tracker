package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

func RunWeb() {
	log.Println("Starting server on http://localhost:3000")

	// Root Hanlder
	http.HandleFunc("/", IndexHandler)

	// Add new anime handler
	http.HandleFunc("/addAnime", AddHandler)

	// Increment/Remove handler
	http.HandleFunc("/action", ActionHandler)

	http.ListenAndServe(":300", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Load html file
	index, err := ioutil.ReadFile("./template/index.html")
	if err != nil {
		log.Println("Can't read index.html")
		os.Exit(2)
	}
	// Generate template
	var templ = template.Must(template.New("index").Parse(string(index[:])))

	// Serve template with animeList
	templ.Execute(w, animeList)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if err := addAnime(r.Form["title"][0], r.Form["episode"][0]); err != nil {
		log.Println("Could not add beacuse: %v\n", err)
	}
	log.Println("Adding anime", r.Form)
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func ActionHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Both do nothing. They just print to the console what they shourd do
	if r.FormValue("Increment") != "" {
		log.Println("Trying to increment")
		// if err := Increment( ... ); err != nil { ...

	} else if r.FormValue("Remove") != "" {
		log.Println("Trying to remove")
		// if err := Remove( ... ); err != nil { ...
	}
	// Redirect to root
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
