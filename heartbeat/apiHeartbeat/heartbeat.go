package apiHeartbeat

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"test/objectStorage/rabbitmq"
	"time"
)

var dataServers = make(map[string]time.Time)
var mutex sync.Mutex

func ListenHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	q.Bind("apiServers")
	c := q.Consumer()
	go removeExpiredDataServer()
	for msg := range c {
		dataServer, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			panic(e)
		}
		mutex.Lock()
		dataServers[dataServer] = time.Now()
		mutex.Unlock()
	}
}

func removeExpiredDataServer() { // timeout 10s
	for {
		time.Sleep(time.Second * 5)
		mutex.Lock()
		for s, t := range dataServers {
			if t.Add(10 * time.Second).Before(time.Now()) {
				delete(dataServers, s)
			}
		}
		mutex.Unlock()
	}
}

func GetDataServers() []string {
	mutex.Lock()
	defer mutex.Unlock()
	ds := make([]string, 0)
	for s, _ := range dataServers {
		ds = append(ds, s)
	}
	return ds
}

func ChooseRandomDataServer() string {
	ds := GetDataServers()
	for i, v := range ds {
		fmt.Println(i, v)
	}

	n := len(ds)
	if n == 0 {
		return ""
	}
	return ds[rand.Intn(n)]
}
