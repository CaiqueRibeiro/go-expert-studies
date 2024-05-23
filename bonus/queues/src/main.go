package main

import (
	"fmt"

	"github.com/CaiqueRibeiro/go-expert-studies/queues/src/queue"
	"github.com/CaiqueRibeiro/go-expert-studies/queues/src/queue_from_scratch"
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

	// Queue from scratch using no library
	qfs := queue_from_scratch.NewQueue[string]()
	qfs.Enqueue("first")
	qfs.Enqueue("second")
	qfs.Enqueue("third")
	vfq, err := qfs.Dequeue()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(vfq)
}
