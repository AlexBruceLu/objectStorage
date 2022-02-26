package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get(w, r)
	case http.MethodPut:
		put(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func get(w http.ResponseWriter, r *http.Request) {
	path := fmt.Sprintf("%s/objects/%s", os.Getenv("STORAGE_ROOT"), strings.Split(r.URL.EscapedPath(), "/")[2])
	f, e := os.Open(path)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func put(w http.ResponseWriter, r *http.Request) {
	path := fmt.Sprintf("%s/objects/%s", os.Getenv("STORAGE_ROOT"), strings.Split(r.URL.EscapedPath(), "/")[2])
	f, e := os.Create(path)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	io.Copy(f, r.Body)
}
