package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
// don't communicate by sharing memory, share memory by communicating

func WaitForTaskUnBufferChannel() {
	ch := make(chan string) // unbuffered channel

	// send a message to a channel
	// -> ch

	// recieving from the channel
	// <- ch


	go func() {
		p := <-ch // block/wait for a message

		// Employee performs work here.
		fmt.Println(p)
		// Employee is done and free to go.
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	ch <- "paper"
}
// pattern: worker
// fan-out and fan-in

func FanOut() {
	// buffered channel
	emps := 20
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func() {
			// do some work that takes some time
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- fmt.Sprintf("paper %d", e+1) // submit result in a paper
		}()
	}

	// manager
	for emps > 0 {
		p := <-ch // receiving
		fmt.Println(p)
		emps--
	}
}

func selectDrop() {
	const cap = 5
	ch := make(chan string, cap)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for p := range ch {
			fmt.Println("employee : received :", p)
		}
		wg.Done()
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager : send ack")
		default:
			fmt.Println("manager : drop")
		}
	}
	close(ch)
	ch <- "closed channel"
	wg.Wait()
}

func main() {
	// channels
	// three types of channels
	// buffered channels
	// unbuffered channels
	// nil channels
	//var ch chan string // nil channel


	// sending ch <-
	// receiving <- ch
	//var wg sync.WaitGroup
	//wg.Add(1)
	//ch := make(chan string) // unbuffered channel
	//go func() {
	//	p := <-ch // block/wait for a message
	//	// Employee performs work here.
	//	fmt.Println(p)
	//	// Employee is done and free to go.
	//	wg.Done()
	//}()
	////time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	//time.Sleep(4 * time.Second)
	//ch <- "paper"
	//fmt.Println("i am waiting on the go routine to finish before i stop this program")
	//wg.Wait()

	// fan-out approach
	//emps := 20
	//ch := make(chan string, emps)
	//var e int
	//for e = 0; e < emps; e++ {
	//	go func(val int) {
	//		// do some work that takes some time
	//		// time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	//		ch <- fmt.Sprintf("paper %d", val) // submit result in a paper
	//	}(e)
	//}
	//
	//// manager
	//for emps > 0 {
	//	p := <-ch // receiving
	//	fmt.Println(p)
	//	emps--
	//}


selectDrop()
}