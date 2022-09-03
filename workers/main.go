package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	concurrency := 5
	in := make(chan int)
	done := make(chan []byte)

	go func() {
		i := 0

		for {
			in <- i
			i++
		}
	}()

	for x := 0; x < concurrency; x++ {
		go ProcessWorkers(in, x)
	}

	<-done
}

func ProcessWorkers(in chan int, worker int) {
	for x := range in {
		timer := time.Duration(rand.Intn(4) * int(time.Second))
		time.Sleep(timer)
		fmt.Println("Worker ", worker, " :", int(x))
	}
}
