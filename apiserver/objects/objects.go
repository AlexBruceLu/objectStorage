package objects

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"object.storage/apiserver/heartbeat"
	"object.storage/apiserver/locate"
	"object.storage/apiserver/objectstream"
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

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, e := storagObject(r.Body, object)
	if e != nil {
		log.Println(e)
	}

	w.WriteHeader(c)
}

func storagObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}

	io.Copy(stream, r)

	if e := stream.Close(); e != nil {
		return http.StatusInternalServerError, e
	}

	return http.StatusOK, nil
}

func putStream(object string) (objectstream.PutStream, error) {
	server := heartbeat.ChooseRadomDataServer()
	if server == "" {
		return nil, fmt.Errorf("cannot find any data server")
	}

	return objectstream.NewPutStream(server, object), nil
}

func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]

	stream, e := getStream(object)
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	io.Copy(w, stream)
}

func getStream(object string) (io.Reader, error) {
	server := locate.Locate(object)

	if server == "" {
		return nil, fmt.Errorf("object %s locate fail", object)
	}

	return objectstream.NewGetStream(server, object)
}
