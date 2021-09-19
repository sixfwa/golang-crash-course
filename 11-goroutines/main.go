package main

import (
	"fmt"
	"time"
)

func doSomething(thingTodo string) {
	for i := 0; i <= 10; i++ {
		fmt.Println("Step:", i, "I am:", thingTodo)
		time.Sleep(time.Millisecond * 500)
	}
}

// implicit goroutine - thread 1
func main() {
	go doSomething("eating breakfast") // an explicit goroutine - thread 2
	go doSomething("watching netflix") // an explicit goroutine - thread 3
	go doSomething("drinking coffee")

	// time.Sleep(time.Millisecond * 3000)
	fmt.Scanln()
}
