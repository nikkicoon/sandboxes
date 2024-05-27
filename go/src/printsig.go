package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc)
	go func() {
		for {
			s := <-sigc
			fmt.Printf("Signal %v received\n", s)
		}
	}()
	select {}
}

