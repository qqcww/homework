package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	messages := make(chan int, 10)
	done1 := make(chan bool)
	done2 := make(chan bool)

	defer close(messages)
	go func() {
		rand.Seed(time.Now().UnixNano())
		ticker := time.NewTicker(1 * time.Second)

		for _ = range ticker.C {
			select {
			case <-done1:
				fmt.Println("End of production...")
				return
			default:
				messages <- rand.Intn(100)
			}
		}
	}()
	go func() {
		rand.Seed(time.Now().UnixNano())
		ticker := time.NewTicker(1 * time.Second)
		for _ = range ticker.C {
			select {
			case <-done2:
				fmt.Println("Waiting timeout...")
				return
			default:
				fmt.Println(<-messages)
			}
		}
	}()
	time.Sleep(15 * time.Second)
	close(done1)
	time.Sleep(11 * time.Second)
	close(done2)
	time.Sleep(1 * time.Second)

}
