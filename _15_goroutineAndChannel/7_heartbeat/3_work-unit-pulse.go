/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"fmt"
	"math/rand"
)

// before  the unit work  always have  heartbeat
func main() {
	doWork := func(done <-chan interface{}) (<-chan interface{}, <-chan int) {
		heartbeatStream := make(chan interface{}, 1) // <1>
		workStream := make(chan int)
		go func() {
			defer close(heartbeatStream)
			defer close(workStream)

			for i := 0; i < 10; i++ {
				select { // <2>
				case heartbeatStream <- struct{}{}:
				default: // <3>
				}

				select {
				case <-done:
					return
				case workStream <- rand.Intn(10):
				}
			}
		}()

		return heartbeatStream, workStream
	}

	done := make(chan interface{})
	defer close(done)

	heartbeat, results := doWork(done)
	for {
		select {
		case _, ok := <-heartbeat:
			if ok {
				fmt.Println("pulse")
			} else {
				return
			}
		case r, ok := <-results:
			if ok {
				fmt.Printf("results %v\n", r)
			} else {
				return
			}
		}
	}
}