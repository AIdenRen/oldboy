package test_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var c1 chan int = make(chan int, 100)
var c2 chan int = make(chan int, 100)

var msg sync.WaitGroup

var verification chan bool = make(chan bool, 10)

func Test(t *testing.T) {
	msg.Add(1)
	go read(c1)

	for val := range c1 {
		fmt.Println(val)
	}

	msg.Wait()

}

func read(ch1 chan int) {
	defer msg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}

}

func get(ch1 chan int, ch2 chan int) {
	for val := range ch1 {
		ch2 <- val * val
	}
}
