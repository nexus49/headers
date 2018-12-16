package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/url"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8001", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	for key, value := range r.Header {
		w.Write([]byte(fmt.Sprintf("Key: %s\nValue: %+v\n\n", key, value)))
	}

	if r.Header["Ssl-Client-Cert"] != nil {
		certStr, _ := url.PathUnescape(r.Header["Ssl-Client-Cert"][0])
		pemStr, _ := pem.Decode([]byte(certStr))
		cert, _ := x509.ParseCertificates(pemStr.Bytes)

		w.Write([]byte("Client Certificate\n"))
		w.Write([]byte(fmt.Sprintf("SignatureAlgorithm: %+v\n", cert[0].SignatureAlgorithm)))
		w.Write([]byte(fmt.Sprintf("PublicKeyAlgorithm: %+v\n", cert[0].PublicKeyAlgorithm)))
		subj := cert[0].Subject
		w.Write([]byte(fmt.Sprintf("Subject: %+v\n", subj)))
		w.Write([]byte(fmt.Sprintf("  Country: %+v\n", subj.Country)))
		w.Write([]byte(fmt.Sprintf("  Organization: %+v\n", subj.Organization)))
		w.Write([]byte(fmt.Sprintf("  OrganizationalUnit: %+v\n", subj.OrganizationalUnit)))
		w.Write([]byte(fmt.Sprintf("  Locality: %+v\n", subj.Locality)))
		w.Write([]byte(fmt.Sprintf("  Province: %+v\n", subj.Province)))
		w.Write([]byte(fmt.Sprintf("  CommonName: %+v\n", subj.CommonName)))

		w.Write([]byte(fmt.Sprintf("Issuer: %+v\n", cert[0].Issuer)))
		w.Write([]byte(fmt.Sprintf("Raw Certificate: %+v\n\n", cert[0])))
	}

}
