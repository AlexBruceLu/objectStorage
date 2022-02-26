package heartbeat

import (
	"os"
	"time"

	"object.storage/rabbitmq"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBIT_SERVER"))
	defer q.Close()

	for {
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(time.Second * 5)
	}
}
