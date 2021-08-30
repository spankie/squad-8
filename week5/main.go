package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func Bar() {
	fmt.Println("called to bar")
}

func Foo() {
	fmt.Println("foo was called")
	wg.Done()
}
/*
Parallelism is running more than one thing at exactly the same time.
scheduler: time sharing (switches)

2006: intel made the first commercially available dual core cpu

concurrency: pattern on how code is structured to make it run independently
 */


func main() {
	/*
	fmt.Println("CPUs:", runtime.NumCPU())
	wg.Add(1)

	go Foo()
	Bar()
	Bar()
	wg.Wait()
	// main says i am done

	// Go routines are cheap threads
	//         Foo ---------------------
	//            /
	//main ----------------------------------
	*/


	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second) // wait for a second
			runtime.Gosched()
			v++ // 1
			// processor waits
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	//

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}

