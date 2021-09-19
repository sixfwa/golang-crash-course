package main

import (
	"fmt"
	"time"
)

func giveMeNumber(intChannel chan int) {
	defer close(intChannel)
	for i := 0; i <= 10; i++ {
		intChannel <- i
	}
}

func doubleNumber(intChannel chan int) {
	for number := range intChannel {
		// number, isOpen := <-intChannel
		// if !isOpen {
		// 	break
		// }
		fmt.Println(number, "doubled is:", number*2)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	intChannel := make(chan int)
	go giveMeNumber(intChannel)
	doubleNumber(intChannel)
}
