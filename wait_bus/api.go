package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

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
	ID        string `json:"id"`
	Nom       string `json:"nom"`
	Desserte  string `json:"desserte"`
	PMR       string `json:"pmr"`
	Escalator string `json:"escalator"`
	GID       string `json:"gid"`
}

type ArretResult struct {
	Next   string  `json:"next"`
	Values []Arret `json:"values"`
}

func get(url string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(TclLogin, TclPassword)
	return req, err
}

func GetArrets() (ArretResult, error) {
	var result ArretResult
	var query string
	for {
		page, err := GetArretsForQuery(query)
		if err != nil {
			return result, err
		}
		result.Values = append(result.Values, page.Values...)
		if page.Next == "" {
			return result, nil
		}
		nextURL, err := url.Parse(page.Next)
		if err != nil {
			return result, err
		}
		query = nextURL.RawQuery
	}
}

func GetArretsForQuery(query string) (ArretResult, error) {
	var result ArretResult
	req, err := get(UrlArret + "?" + query)
	if err != nil {
		return result, err
	}
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
	GID          string `json:"gid"`
	ID           string `json:"id"`
	Nom          string `json:"nom"`
	Ligne        string `json:"ligne"`
	Direction    string `json:"direction"`
	DelaiPassage string `json:"delaipassage"`
}

type PassageArretResult struct {
	Values []PassageArret `json:"values"`
}

func GetPassagesForQuery(query string) (PassageArretResult, error) {
	var result PassageArretResult
	req, err := get(UrlPassageArret + "?" + query)
	if err != nil {
		return result, err
	}
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
