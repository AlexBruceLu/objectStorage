package main

import (
	"log"
	"net/http"
	"os"
	"test/objectStorage/heartbeat/apiHeartbeat"
	"test/objectStorage/locate/apiLocate"
	"test/objectStorage/objects/apiObject"
)

func main() {
	go apiHeartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", apiObject.Handler)
	http.HandleFunc("/locate/", apiLocate.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
