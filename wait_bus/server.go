package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.StringVar(&TclLogin, "l", "", "login to connect to TCL service")
	flag.StringVar(&TclPassword, "p", "", "password to connect to TCL service")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("dist/")))

	http.HandleFunc("/api/passagearret", apiPassageArret)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

var TclLogin = ""
var TclPassword = ""

var arrets ArretResult

func apiPassageArret(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	passages, err := GetPassagesForQuery(query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	if arrets.Values == nil {
		arrets, err = GetArrets()
		if err != nil {
			log.Printf("ERROR: retrieving arrets: %s\n", err)
		}
	}

	for p := range passages.Values {
		arret, ok := findArret(arrets.Values, passages.Values[p].ID)
		if ok {
			passages.Values[p].Nom = arret.Nom
		}
	}
	jsPassages, err := json.Marshal(passages)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(jsPassages)
}

func findArret(arrets []Arret, id string) (Arret, bool) {
	for _, arret := range arrets {
		if arret.ID == id {
			return arret, true
		}
	}
	return Arret{}, false
}
