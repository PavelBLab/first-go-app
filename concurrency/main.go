package main

import (
	"fmt"
	"time"
)

func slowGreet(doneChannel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello World!")
	doneChannel <- true
}

func main() {
	done := make(chan bool) // channel for signaling completion
	go slowGreet(done)      // start goroutine
	<-done                  // wait for goroutine to finish
}
