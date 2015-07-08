package main

import (
	"fmt"
	"time"
)

func main() {
	go say("go")
	go say("for it")

	var in string
	fmt.Scanln(&in)
}

func say(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println("Going to I sleep")
		// a lot of work
		time.Sleep(100 * time.Millisecond)
		fmt.Println(message)
	}
}
