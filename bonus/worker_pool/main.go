package main

import "fmt"

func main() {
	// Create new tasks
	tasks := []Task{
		&EmailTask{Email: "caique@gmail.com", Subject: "Hello", MessageBody: "Hello, how are you?"},
		&ImageProcessingTask{ImageUrl: "http://example.com/horse.jpg"},
		&EmailTask{Email: "caique@gmail.com", Subject: "Test", MessageBody: "Just testing"},
		&ImageProcessingTask{ImageUrl: "http://example.com/spider-man.png"},
		&EmailTask{Email: "caique@gmail.com", Subject: "Monkeys", MessageBody: "I hate monkeys"},
		&ImageProcessingTask{ImageUrl: "http://example.com/macaco.jpg"},
		&EmailTask{Email: "caique@gmail.com", Subject: "Give ma money", MessageBody: "Gimme ma fucking money!"},
		&ImageProcessingTask{ImageUrl: "http://example.com/michael_jackson_in_rj.jpg"},
		&EmailTask{Email: "caique@gmail.com", Subject: "Good bye", MessageBody: "Good bye bitches"},
		&ImageProcessingTask{ImageUrl: "http://example.com/image1.jpg"},
	}

	// Create new worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, // Number of workers that can run at a time
	}

	// run the poool
	wp.Run()
	fmt.Println("All tasks are done!")
}
