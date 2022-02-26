package heartbeat

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"object.storage/rabbitmq"
)

var (
	dataServers = make(map[string]time.Time)
	mtx         sync.RWMutex
)

func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBIT_SERVER"))
	defer q.Close()

	q.Bind("apiServers")
	c := q.Consume()

	go removeExpiredDataServer()

	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			log.Fatal(e)
		}
		mtx.Lock()
		dataServers[dataServer] = time.Now()
		mtx.Unlock()
	}
}

func removeExpiredDataServer() {
	for {
		time.Sleep(time.Second * 5)
		mtx.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mtx.Unlock()
	}
}

func GetDataServers() []string {
	mtx.RLock()
	ds := make([]string, 0)
	for s := range dataServers {
		ds = append(ds, s)
	}
	mtx.RUnlock()
	return ds
}

func ChooseRadomDataServer() string {
	ds := GetDataServers()
	n := len(ds)
	if n == 0 {
		return ""
	}

	return ds[rand.Intn(n)]
}
