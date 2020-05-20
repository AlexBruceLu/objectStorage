package dataObjects

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	m := r.Method
	switch m {
	case http.MethodPut:
		put(w, r)
		return
	case http.MethodGet:
		get(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// get object by object_name
func get(w http.ResponseWriter, r *http.Request) {
	// url split by "/" three strings,"" "/objects/" "<object_name>"
	f, e := os.Open(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])
	if e != nil {
		log.Println("[ERROR]: ", e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	defer f.Close()
	_, _ = io.Copy(w, f)
}

// put objects
func put(w http.ResponseWriter, r *http.Request) {
	f, e := os.Create(os.Getenv("STORAGE_ROOT") + "/objects/" + strings.Split(r.URL.EscapedPath(), "/")[2])

	if e != nil {
		log.Println("[ERROR]: ", e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	_, _ = io.Copy(f, r.Body)
}
