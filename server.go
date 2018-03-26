package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
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

/*
UrlLigne schema
0	"code_titan"
1	"ligne"
2	"sens"
3	"indice"
4	"infos"
5	"libelle"
6	"ut"
7	"couleur"
8	"gid"
9	"last_update"
10	"last_update_fme"
*/
const UrlLigne = "https://download.data.grandlyon.com/ws/rdata/tcl_sytral.tcllignebus/all.json"

/*
UrlPassageArret schema
0	"id" : ID arret
1	"ligne"
2	"direction"
3	"delaipassage"
4	"type"
5	"heurepassage"
6	"idtarretdestination"
7	"coursetheorique"
8	"gid"
9	"last_update_fme"
*/
const UrlPassageArret = "https://download.data.grandlyon.com/ws/rdata/tcl_sytral.tclpassagearret/all.json"

type PassageArret struct {
	ID           string
	Ligne        string
	Direction    string
	DelaiPassage string
}

var TclLogin = ""
var TclPassword = ""

func apiPassageArret(w http.ResponseWriter, r *http.Request) {
	query := r.URL.RawQuery
	req, err := http.NewRequest("GET", UrlPassageArret+"?"+query, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	arrets, err := apiArret()
	log.Println(arrets)
	log.Println(err)
	req.SetBasicAuth(TclLogin, TclPassword)
	clt := &http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error: %s\n  l/p:%s/%s\n  %s\n", req.URL, TclLogin, TclPassword, err)
		fmt.Fprintf(w, "Error: %s\n  l/p:%s/%s\n  %s\n", req.URL, TclLogin, TclPassword, err)
		return
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

/*
UrlArret schema
0	"id"
1	"nom"
2	"desserte"
3	"pmr"
4	"ascenseur"
5	"escalator"
6	"gid"
7	"last_update"
8	"last_update_fme"
*/
const UrlArret = "https://download.data.grandlyon.com/ws/rdata/tcl_sytral.tclarret/all.json"

type Arret struct {
	ID        string
	Nom       string
	Desserte  string
	PMR       string
	Escalator string
	GID       string
}

type ArretResult struct {
	Values []Arret
}

func apiArret() (ArretResult, error) {
	var result ArretResult
	req, err := http.NewRequest("GET", UrlArret, nil)
	if err != nil {
		return result, err
	}
	req.SetBasicAuth(TclLogin, TclPassword)
	clt := &http.Client{}
	resp, err := clt.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	jsDecoder := json.NewDecoder(resp.Body)
	err = jsDecoder.Decode(&result)
	return result, err
}
