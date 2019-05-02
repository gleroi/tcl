package main

import (
	"fmt"
	"bytes"
	"time"
	"net/http"
	"mime/multipart"

	"github.com/PuerkitoBio/goquery"
)

const TCLBaseURL = "https://www.tcl.fr"
const TCLSubscriptionCertificateService = "/Mon-TCL/Service-Attestation-d-abonnement2"

type api struct {
	clt     *http.Client
	baseURL string
}

func NewApi(url string) *api {
	return &api{
		clt:     &http.Client{},
		baseURL: url,
	}
}

func (api *api) RequestSubscriptionCertificate(cardID string, email string, birthday BirthDate, month int, year int) error {
	// Form fields
	form := bytes.NewBuffer(nil)
	w := multipart.NewWriter(form)
	// - classe_id(hidden): 132
	w.WriteField("classe_id", "132")
	// - titre(hidden): "Demande_attestationanonymous$day$month$year$hour$minutes"
	w.WriteField("titre", fmt.Sprintf("Demande_attestationanonymous%s", time.Now().Format("020120061504")))
	// - envoyer(submit)
	w.WriteField("envoyer", "envoyer")
	// - email(text)
	w.WriteField("email", email)
	// - jour_date_naissance(select): 1-31 (%02d)
	w.WriteField("jour_date_naissance", fmt.Sprintf("%02d", birthday.day()))
	// - mois_date_naissance(select): 1-12 (%02d)
	w.WriteField("mois_date_naissance", fmt.Sprintf("%02d", birthday.month()))
	// - annee_date_naissance(select): 1900-2019 (%04d)
	w.WriteField("annee_date_naissance", fmt.Sprintf("%d", birthday.year()))
	// - type_carte(select): 0 (carte tecely)
	w.WriteField("type_carte", "0")
	// - num_tecely(text)
	w.WriteField("num_tecely", cardID)
	// - mois_date_debut(select): 1-12 (%02d)
	w.WriteField("mois_date_debut", fmt.Sprintf("%02d", month))
	// - annee_date_debut(select): 1900-2019 (%d)
	w.WriteField("annee_date_debut", fmt.Sprintf("%d", year))
	// - mois_date_fin(select): 1-12 (%02d)
	w.WriteField("mois_date_fin", fmt.Sprintf("%02d", month))
	// - annee_date_fin(select): 1900-2019 (%d)
	w.WriteField("annee_date_fin", fmt.Sprintf("%d", year))
	w.Close()

	req, err := http.NewRequest("POST", api.baseURL + TCLSubscriptionCertificateService, form)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", w.FormDataContentType())
	resp, err := api.clt.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	doc.Find(".BLOC-alerte").Each(func (_ int, sel *goquery.Selection) {
		msg := sel.Text()
		fmt.Fprintln(buf, msg)
	})
	if buf.Len() > 0 {
		return fmt.Errorf(buf.String())
	}
	return nil
}