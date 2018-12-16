package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

}
