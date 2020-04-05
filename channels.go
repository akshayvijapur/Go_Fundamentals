package main

import "fmt"

// Channels are used to send and receive data.
// data := <-ch get the data from channel
// ch  <- data Send the data to channel
func main() {
	// Create a channel
	ch := make(chan string)
	go getMeData(ch) //Create a routine and send the channel.
	data := <-ch     // data := <-ch Get the data.
	fmt.Println(data)

}

func getMeData(ch chan string) {
	data := "Hello World"
	ch <- data //send the data to ch. ch <- data.
}
