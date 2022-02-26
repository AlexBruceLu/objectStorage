package locate

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"object.storage/rabbitmq"
)

func Locate(name string) bool {
	_, e := os.Stat(name)
	return !os.IsNotExist(e)
}

func StartLocate() {
	q := rabbitmq.New(os.Getenv("RABBIT_SERVER"))
	defer q.Close()

	q.Bind("dataServers")
	c := q.Consume()
	for msg := range c {
		obj, e := strconv.Unquote(string(msg.Body))
		if e != nil {
			log.Fatal(e)
		}
		path := fmt.Sprintf("%s/objects/%s", os.Getenv("STORAGE_ROOT"), obj)
		if Locate(path) {
			q.Send(msg.ReplyTo, os.Getenv("LISTEN_ADDRESS"))
		}
	}

}
