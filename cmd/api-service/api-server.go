package main

import (
	"log"
	"net/http"
	"os"

	"object.storage/apiserver/heartbeat"
	"object.storage/apiserver/locate"
	"object.storage/apiserver/objects"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
