package main

import (
	"fmt"
	"time"
	// "sync"
)

func main() {
	// var wg sync.WaitGroup
	// wg.Add(1)
	// "go" runs a goroutine
	// anonymous function
	//go func() {
	//	count("sheep")
	//	wg.Done()
	//}()
	//
	//wg.Wait()
	c := make(chan string)
	go count("sheep", c)
	for msg := range c {
	// for {
		// msg, open := <- c
		//if !open {
		//	break
		//}
		fmt.Println(msg)
	}
}

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		// fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}
