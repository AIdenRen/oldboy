package test_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

var a sync.WaitGroup

func TestA(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}

func TestB(t *testing.T) {
	for i := 0; i < 5; i++ {
		go f()
	}
	a.Wait()
	time.Sleep(time.Second * 5)
}

func f() {
	defer a.Done()
	a.Add(1)
	fmt.Println(1)
}

func Test_gomaxprocs(t *testing.T) {
	fmt.Print(runtime.NumCPU())
	a.Add(2)
	go a1()
	go b()
	a.Wait()

}

func a1() {
	defer a.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}

func b() {
	defer a.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}
