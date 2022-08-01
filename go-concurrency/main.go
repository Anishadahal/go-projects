package main

import (
	"fmt"
	"time"
)

//Begin
// func main(){
// 	go count("sheep") //runs in background
// 	go count("fish")

// 	// time.Sleep(time.Second * 2) //stop after 2 sec
// 	// fmt.Scanln() //stop after user presses enter

// 	//practically we use wait group
// }
//wait group
// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(1) //we have 1 go routine to wait for

// 	go func() {
// 		count("Sheep")
// 		wg.Done() //decrement counter by 1
// 	}()

// 	wg.Wait() //block until counter is zero
// }

//channels

// func main() {
// 	c := make(chan string) //make channel
// 	go count("Sheep", c)
// 	// for {
// 	// 	msg, open := <-c //receive msg from channel
// 	// 	if !open {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(msg)
// 	// }

// 	//alternatively
// 	for msg := range c {
// 		fmt.Println(msg)
// 	}
// }

//constraints
// func main() {
// 	c := make(chan string, 3)
// 	c <- "hello"
// 	c <- "hello"
// 	c <- "hello"

// 	msg := <-c
// 	fmt.Println(msg)
// 	msg = <-c
// 	fmt.Println(msg)
// 	msg = <-c
// 	fmt.Println(msg)
// }

//select statement

// func main() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)
// 	go func() {
// 		for {
// 			c1 <- "Every 500ms"
// 			time.Sleep(time.Millisecond * 500)
// 		}
// 	}()

// 	go func() {
// 		for {
// 			c2 <- "Every 2 seconds"
// 			time.Sleep(time.Second * 2)
// 		}
// 	}()

// 	for {

// 		//waits for other despite ready
// 		fmt.Println(<-c1)
// 		fmt.Println(<-c2)

// 		for {
// 			select {
// 			case msg1 := <-c1:
// 				fmt.Println(msg1)
// 			case msg2 := <-c2:
// 				fmt.Println(msg2)
// 			}
// 		}
// 	}
// }

//fibonacci -- WORKER POOL PATTERN

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 0; i <= 100; i++ {
		jobs <- i
	}
	close(jobs)
	for j := 0; j <= 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) { //<-chan == receive | chan<- == send
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

//channels are the way for go routines to communicate with each other
//It is like a pipe through which you can send a message or receive a message
//channels have a type as well -> string/channel/int....

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing //send value/message to channel
		// fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
	close(c) //otherwise deadlock //sender closes channel
}
