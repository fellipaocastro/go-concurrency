package main

import (
	"fmt"
	"time"
)

func main() {
	// Goroutines

	// go say("go")
	// go say("for it")
	//
	// var in string
	// fmt.Scanln(&in)

	// Channels

	// c := make(chan string, 2)
	// c <- "Hello"
	// c <- "World"
	//
	// v := <-c
	// fmt.Println(v)
	// fmt.Println(<-c)

	// Select
	done := make(chan bool)

	go func() {
		fmt.Println("Running first...")
		time.Sleep(250 * time.Millisecond)
		done <- true
	}()

	go func() {
		fmt.Println("Running second...")
		time.Sleep(1000 * time.Millisecond)
		done <- true
	}()

	// fmt.Println("Reading from the channel for the first time")
	// fmt.Println(<-done)
	// fmt.Println("Reading from the channel for the second time")
	// fmt.Println(<-done)

	waiton(done)
	waiton(done)
}

// Goroutines

func say(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Going to sleep")
		// a lot of work
		time.Sleep(100 * time.Millisecond)
		fmt.Println(message)
	}
}

// Select

func waiton(done chan bool) {
	select {
	case <-done:
		fmt.Println("Finished")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout")
	}
}
