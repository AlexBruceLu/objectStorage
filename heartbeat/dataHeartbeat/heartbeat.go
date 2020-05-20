package dataHeartbeat

import (
	"os"
	"test/objectStorage/rabbitmq"
	"time"
)

func StartHeartBeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER"))
	defer q.Close()
	for { // send heartbeat every 5 seconds
		q.Publish("apiServers", os.Getenv("LISTEN_ADDRESS"))
		time.Sleep(time.Second * 5)
	}
}
