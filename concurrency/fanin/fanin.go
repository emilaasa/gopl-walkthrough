package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Demonstrating some concurrency patterns from Rob Pike
func main() {
	// simplest()
	// twoSyncedGenerators()
	// multiplexing()
	// mySelect()
	timeoutUsingSelect()

}
func timeoutUsingSelect() {
	c := generator("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("too slow")
			return
		}
	}

}

// blocks until a channel can proceed
// if you add a default it is nonblocking which means that
// it runs the default if
func mySelect() {
	c := selectFanIn(generator("channel1: "), generator("channel2: "))
	for i := 0; i < 85; i++ {
		fmt.Println(<-c)
	}

}

func multiplexing() {
	c := fanIn(generator("channel1: "), generator("channel2: "))
	for i := 0; i < 50; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("done")
}

func selectFanIn(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-in1:
				c <- s
			case s := <-in2:
				c <- s
			}
		}
	}()
	return c
}

// Would be called merge in ReactiveX
func fanIn(in1, in2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-in1
		}
	}()
	go func() {
		for {
			c <- <-in2
		}
	}()
	return c

}

func twoSyncedGenerators() {
	c1 := generator("First channel")
	c2 := generator("Second channel")
	for i := 0; i < 5; i++ {
		fmt.Printf("Getting values in forced sync: %q\n", <-c1)
		fmt.Printf("Getting values in forced sync: %q\n", <-c2)
	}
}
func simplest() {
	c := generator("Hello!")
	n := 5
	for i := 0; i < n; i++ {
		fmt.Printf("Getting values for %d iterations: %q\n", n, <-c)
	}
}

// Rob defines a generator as a function that returns a channel
func generator(s string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ { // while true with i incrementation
			c <- fmt.Sprintf("%s %d", s, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}
	}()
	return c
}
