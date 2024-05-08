package main

import (
	"fmt"

	"github.com/CaiqueRibeiro/go-expert-studies/queues/src/queue"
)

func main() {
	myQueue := queue.NewQueue[string]()
	myQueue.Enqueue("first")
	myQueue.Enqueue("second")
	myQueue.Enqueue("third")
	valueFromQueue, err := myQueue.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valueFromQueue)
}
