package main

import (
	"net/http"
	"log"

	"github.com/gorilla/mux"
	"github.com/steffen25/go-templates-test/templates"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/about", About)
	r.HandleFunc("/contact", Contact)

	log.Println("Listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["path"] = r.URL.Path
	err := templates.Render(w, "home.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["path"] = r.URL.Path
	err := templates.Render(w, "about.html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Contact(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["path"] = r.URL.Path
	err := templates.Render(w, "contact.html", data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
