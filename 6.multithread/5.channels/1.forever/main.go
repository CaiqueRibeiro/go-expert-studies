package main

// Forever is a problem that happens when a channel never returns because it is waiting for a value that will never be sent.
// To avoid this problem, we can use a goroutine to send the value to the channel.
func main() {
	forever := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		// forever <- true // This will resolve the deadlock forever problem
	}()

	<-forever
}
