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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// Root Hanlder
	http.HandleFunc("/", IndexHandler)

	// Add new anime handler
	http.HandleFunc("/addAnime", AddHandler)

	// Increment handler
	http.HandleFunc("/increment", IncrementHandler)

	// Remove handler
	http.HandleFunc("/remove", RemoveHandler)

	http.ListenAndServe(":3000", nil)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Load html file
	index, err := ioutil.ReadFile("./templates/index.html")
	if err != nil {
		log.Println("Can't read index.html")
		os.Exit(2)
	}
	// Generate template
	var templ = template.Must(template.New("index").Parse(string(index[:])))

	// Update animeList
	if err := db.Find(nil).All(&animeList); err != nil {
		log.Println("Can't find any animes")
	}

	// Serve template with animeList
	templ.Execute(w, animeList)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if err := addAnime(r.Form["title"][0], r.Form["episode"][0]); err != nil {
		log.Println("Could not add beacuse: %v\n", err)
	}
	log.Println("Adding anime", r.Form)

	// Redirect to root
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.FormValue("Increment") == "" {
		log.Println("Incrementing: " + r.Form["Title"][0])
		if err := Increment(r.Form["Title"][0]); err != nil {
			log.Println("Can't increment: %v", err)
		}

	}
	// Redirect to root
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.FormValue("Remove") == "" {
		log.Println("Removing " + r.Form["Title"][0])
		if err := Remove(r.Form["Title"][0]); err != nil {
			log.Println("Can't remove: %v", err)
		}
	}
	// Redirect to root
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
