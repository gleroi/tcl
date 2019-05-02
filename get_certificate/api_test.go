package main

import (
	"testing"
)

func TestApiRequestSubscriptionCertificateKO(t *testing.T) {
	server := NewHttpStub(errorHtmlOutput)
	defer server.Close()

	api := NewApi(server.URL)
	err := api.RequestSubscriptionCertificate("023333", "plop@test.com", Birthday(1985, 3, 22), 2019, 5)
	if err.Error() != "Un numéro de carte correct est obligatoire\n" {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestApiRequestSubscriptionCertificateOK(t *testing.T) {
	server := NewHttpStub(okHtmlOutput)
	defer server.Close()

	api := NewApi(server.URL)
	err := api.RequestSubscriptionCertificate("023333", "plop@test.com", Birthday(1985, 3, 22), 2019, 5)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
