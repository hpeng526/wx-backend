package main

import (
	"github.com/hpeng526/wx-backend/server"
	"time"
)

func main() {
	wxServer := server.NewRedisMqServer("wx9482dd922f2b6ba4", "8bce95b05895f73ed6a7cb912c6edb51", "127.0.0.1:6379")
	tenSec := 10 * time.Second
	for {
		val := wxServer.Mq.Poll("key", tenSec)
		wxServer.HandleMessage(val)
	}
}
