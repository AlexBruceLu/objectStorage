package main

import (
	"log"
	"net/http"
	"os"

	"object.storage/dataserver/heartbeat"
	"object.storage/dataserver/locate"
	"object.storage/dataserver/objects"
)

func main() {
	go heartbeat.StartHeartbeat()
	go locate.StartLocate()
	http.HandleFunc("/objects/", objects.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))

}
