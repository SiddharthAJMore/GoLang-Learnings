package main

import "fmt"

/*
func processNum(messageChan chan int) {
	// value := <-messageChan
	for num := range messageChan { // in for loop we dont need <- for receiving value handled automatically
		fmt.Println(num)
		time.Sleep(time.Second)
	}
}
*/

/*
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}
*/

/*
// using channels instead of waitGroups, explained in subpart
// goRoutine synchronizer
func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing.....")
}
*/

/*
func emailSender(emailChan chan string, doneChan chan bool) {
	defer func() { doneChan <- true }()

	for email := range emailChan {
		fmt.Println("Sending mail to", email)
	}
}
*/

/*
func emailSender(emailChan <-chan string, doneChan chan<- bool) {
	defer func() { doneChan <- true }()

	//valueDone := <-doneChan // not possible because we have scoped the doneChan to only sender by doing `doneChan chan<- bool`

	// emailChan <- "Some message" // not possible since emailChan param is made as receive type only

	for email := range emailChan {
		fmt.Println("Sending mail to", email)
	}
}
*/

func main() {
	/*
		messageChan := make(chan int)
		messageChan <- 1  // push/send value into channel
		value := <-messageChan  // get value from the channel

		fmt.Println(value)

		This throws a error `all goroutines are asleep - deadlock!` because the messageChan <- 1 is a blocking call, which waits
			till the value passed into channel is received by the other side by someone.
		So we need to move this push and read to separate goroutine/threads, so that sender is sending and other is receiving in other goroutine
	*/

	/*
		messageChan := make(chan int)

		go processNum(messageChan)

		// messageChan <- 1

		for {
			messageChan <- rand.Intn(100)
		}
		// time.Sleep(time.Second * 2)
	*/

	/*
		resultChan := make(chan int)
		go sum(resultChan, 4, 5)

		result := <-resultChan  // data received

		fmt.Println(result)
	*/

	/*
		// Using channels as a goRoutine synchronizer
		// using channels as waitGroup. Since waitGroup is used to keep track of the GoRoutines and to hold the program till all
		//	GoRoutines are completed, we also know that sending and receiving messages via a channel is also blocking.
		// So we can pass the channel in GoRoutine and then add a receiver outside the GoRoutine, which will then wait till the GoRoutine is completed
		done := make(chan bool)

		go task(done)

		<-done

		// question is when to use Channel for waiting and when to use WaitGroup for waiting.
		// When there is only 1 goRoutine, then you can use the channel,
		// for multiple goRoutines waitGroup is preferred (since it has that add, done, wait method, and obviously we don't want to create a channel for each goRoutine)

	*/

	/*
		UnBuffered channels - Above channel implementations were unbuffered channels, what this means that only one message is pushed to the
			channel at a time until the message is consumed the next message is not pushed.

		Buffered channels - channels where we can send the data without waiting for it to be consumed
	*/

	/*
			//Buffered channels
			emailChannel := make(chan string, 100)
			done := make(chan bool)

			go emailSender(emailChannel, done)

			for i := range 10 {
				emailChannel <- fmt.Sprintf("%d@gmail.com", i)
			}

			close(emailChannel)

				//Closing a channel is required when we use for loops, so that the receiver is told that the messages are done sending.
				// without closing the channel, the receiver never knows the messages are finished and it will enter deadlock waiting infinitely

		<-done // goRoutine synchronizer
	*/

	// Listening to data from multiple channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 1
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("channel 1 value received", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("channel 2 value received", chan2Val)
		}
	}
}
