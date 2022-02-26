package main

import (
	"log"
	"net/http"

	"object.storage/objects"
)

var LISTEN_ADDRESS = ":9999"

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(LISTEN_ADDRESS, nil))
}
