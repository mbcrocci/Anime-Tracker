package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

func RunWeb() error {
	log.Println("Starting server on http://localhost:3000")

	// Where the source code is located
	path := os.Getenv("GOPATH") + "/src/github.com/mbcrocci/Anime-Tracker/"

	// Make css stylesheets work
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(path+"static"))))

	r := mux.NewRouter()

	// Root Hanlder
	r.HandleFunc("/", IndexHandler)

	// Add new anime handler
	r.HandleFunc("/addAnime", AddHandler)

	// Increment handler
	r.HandleFunc("/increment", IncrementHandler)

	// Remove handler
	r.HandleFunc("/remove", RemoveHandler)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)

	return nil
}

// Everything thing is done here.
// Each action (ie. adding a new anime) calls another handler each one then redirects to root
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// Load html file
	path := os.Getenv("GOPATH") + "/src/github.com/mbcrocci/Anime-Tracker/"
	index, err := ioutil.ReadFile(path + "templates/index.html")
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

// Reads the form on top of the page and inserts the anime into the database
func AddHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if err := addAnime(r.Form["title"][0], r.Form["episode"][0]); err != nil {
		log.Println("Could not add beacuse: %v\n", err)
	}
	log.Println("Adding anime", r.Form)

	// Redirect to root
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// Called when the button "Increment" is pressed
// It reads a hidden field containing the title and updates the episode.
func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.FormValue("Increment") == "" {
		log.Println("Incrementing: ", r.Form)
		if err := Increment(r.Form["Title"][0]); err != nil {
			log.Println("Can't increment: %v", err)
		}

	}

	// Redirect to root
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

// Called when the button "Remove" is pressed
// It also reads a hidden field containing the title and removes the anime from the database
func RemoveHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.FormValue("Remove") == "" {
		log.Println("Removing ", r.Form)
		if err := Remove(r.Form["Title"][0]); err != nil {
			log.Println("Can't remove: %v", err)
		}
	}

	// Redirect to root
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
