package main

import (
	"characters"
	"encoding/json"
	"log"
	"myThings"
	"net/http"
	"os"
	"spells"
)

type charComp func(r *http.Request) ([]byte, error)

func magicHandler(f charComp) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := f(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("There was an error loading : %s", err)
			return
		}
		if _, err := w.Write(b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("There was an error writing: %s", err)
		}
	}
}

func main() {

	http.HandleFunc("/api/spells/", magicHandler(spells.Dispatcher))
	http.HandleFunc("/api/character/", magicHandler(characters.Dispatcher))
	PORT := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
