package data

import (
	"fmt"
	"time"
)

var channel chan int = make(chan int, 1)

func EP_chan() {

	ticker := time.NewTicker(time.Second)

	defer ticker.Stop()

	go func() {
		time.Sleep(10 * time.Second)
		channel <- 1
	}()

	for {
		select {
		case <-channel:
			fmt.Println("Done!")
			return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}

}
