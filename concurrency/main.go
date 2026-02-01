package main

import (
	"fmt"
	"time"
)

func greet(doneChannel chan bool) {
	fmt.Println("Hello World!")
	doneChannel <- true
}

func slowGreet(doneChannel chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Slow Hello World!")
	doneChannel <- true
	close(doneChannel)
}

func main() {
	//dones := make([]chan bool, 4)
	done := make(chan bool) // channel for signaling completion

	//dones[0] = make(chan bool)
	go greet(done) // start goroutine
	//dones[1] = make(chan bool)
	go greet(done) // start goroutine
	//dones[2] = make(chan bool)
	go slowGreet(done) // start goroutine
	//dones[3] = make(chan bool)
	go greet(done) // start goroutine

	for range done {
		//fmt.Println(doneChan)
	}

	//for _, done := range dones {
	//	<-done
	//}

	//<-done // wait for goroutine to finish
	//<-done // wait for goroutine to finish
	//<-done // wait for goroutine to finish
	//<-done // wait for goroutine to finish

}
