package locate

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"object.storage/rabbitmq"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	info := Locate(strings.Split(r.URL.EscapedPath(), "/")[2])
	if len(info) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	b, _ := json.Marshal(info)
	w.Write(b)
}

func Locate(name string) string {
	q := rabbitmq.New(os.Getenv("RABBIT_SERVER"))
	q.Publish("dataServers", name)
	c := q.Consume()

	go func() {
		time.Sleep(time.Second)
		q.Close()
	}()

	msg := <-c
	s, _ := strconv.Unquote(string(msg.Body))
	return s
}

func Exist(name string) bool {
	return Locate(name) != ""
}
