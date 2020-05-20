package apiObject

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"test/objectStorage/heartbeat/apiHeartbeat"
	"test/objectStorage/locate/apiLocate"
	"test/objectStorage/streams/apiObjectStream"
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

func put(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	c, e := storageObject(r.Body, object)
	if e != nil {
		log.Println("[ERROR]:", e)
	}
	w.WriteHeader(c)
}

func storageObject(r io.Reader, object string) (int, error) {
	stream, e := putStream(object)
	if e != nil {
		return http.StatusServiceUnavailable, e
	}
	io.Copy(stream, r)
	e = stream.Close()
	if e != nil {
		return http.StatusInternalServerError, e
	}
	return http.StatusOK, nil
}

func putStream(object string) (*apiObjectStream.PutStream, error) {
	server := apiHeartbeat.ChooseRandomDataServer()
	if server == "" {
		return nil, fmt.Errorf("can not find any dataServer")
	}
	return apiObjectStream.NewPutStream(server, object), nil
}

//get
func get(w http.ResponseWriter, r *http.Request) {
	object := strings.Split(r.URL.EscapedPath(), "/")[2]
	stream, e := getStream(object)
	if e != nil {
		log.Println("[ERROR]: ", e)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	io.Copy(w, stream)
}

func getStream(object string) (io.Reader, error) {
	server := apiLocate.Locate(object)
	if server == "" {
		return nil, fmt.Errorf("object locate %s fail", object)
	}
	return apiObjectStream.NewGetStream(server, object)
}
