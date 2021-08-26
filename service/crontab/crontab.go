package main

import (
	"time"
)

func main() {
	var ch chan int
	ticker := time.NewTicker(time.Hour * 168)
	go func() {
		for range ticker.C {

		}
		ch <- 1
	}()
	<-ch
}
