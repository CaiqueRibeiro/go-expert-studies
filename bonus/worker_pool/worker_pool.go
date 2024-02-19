package main

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task interface {
	Process()
}

// Email task definition
type EmailTask struct {
	Email       string
	Subject     string
	MessageBody string
}

// Process method for EmailTask
func (t *EmailTask) Process() {
	fmt.Printf("Sending email to %s\n", t.Email)
	// Simulate time consuming process
	time.Sleep(2 * time.Second)
}

type ImageProcessingTask struct {
	ImageUrl string
}

func (t *ImageProcessingTask) Process() {
	fmt.Printf("Processing the image %s\n", t.ImageUrl)
	time.Sleep(4 * time.Second)
}

// WorkerPool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	tasksChan   chan Task
	wg          sync.WaitGroup
}

// Functions to execute the worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.tasksChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	// Initialize the tasks channel
	wp.tasksChan = make(chan Task, len(wp.Tasks)) // If 20 tasks, the channel will have a buffer of 20

	for i := 0; i < wp.concurrency; i++ {
		go wp.worker() // if concurrency is 5, 5 workers will be created (in other words, 5 threads)
	}

	// Send tasks to the task channel
	wp.wg.Add(len(wp.Tasks))        // have to wait for all tasks to finish (20)
	for _, task := range wp.Tasks { // gets tasks one by one and send it to the channel
		wp.tasksChan <- task // wont be blocked because the buffer is the size of the number of tasks
	}
	close(wp.tasksChan)

	// Wait for all tasks to finish
	wp.wg.Wait()
}
