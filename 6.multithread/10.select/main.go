package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Id  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	go func() { // RabbitMQ
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 1)
			msg := Message{i, "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	go func() { // Kafka
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{i, "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select { // looks for the first channel that has data. It's good to find which one is faster and ignore others
		case msg := <-c1:
			fmt.Printf("received from RabbitMQ: ID: %d - %s\n", msg.Id, msg.Msg)

		case msg := <-c2:
			fmt.Printf("received from Kafka: ID: %d - %s\n", msg.Id, msg.Msg)

		case <-time.After(time.Second * 3): // if after 3 seconds none of channels is filled with data, returns timeout
			println("timeout")

			/*
				default: // used to avoid blocking and if do not want to wait for any channel
					println("default")
			*/
		}
	}
}
