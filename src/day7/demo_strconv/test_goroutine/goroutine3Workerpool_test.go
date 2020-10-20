package test_goroutine

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var jobChan chan int
var resultChan chan int64

func TestAB(t *testing.T) {
	jobChan = make(chan int, 100)
	resultChan = make(chan int64, 100)

	for i := 0; i < 100; i++ {
		generate()
	}
	close(jobChan)

	working2(jobChan, resultChan)

	//for i := 0; i < 4; i++ {
	//	go
	//}

	for true {
		i, err := <-resultChan
		if !err {
			fmt.Println("退出")
		}
		fmt.Printf("%d\n%v%v", i, "正常", err)
	}

	//for i := 0; i < 4; i++ {
	//	go working()
	//}

	//for true {
	//	vals,ok := <- jobChan
	//	if !ok{
	//		fmt.Println("完毕")
	//		break;
	//	}
	//	fmt.Println(vals)
	//}

}

func working2(job <-chan int, result chan<- int64) {
	for true {
		vals, ok := <-job
		if !ok {
			break
		}
		formatInt := strconv.FormatInt(int64(vals), 10)
		var i int64
		for _, value := range []rune(formatInt) {
			fmt.Println(value)
			sprintf := fmt.Sprintf("%d", value)
			parseInt, err := strconv.ParseInt(sprintf, 10, 64)
			if err != nil {
				fmt.Println(err, "58")
			}
			i = i + parseInt

		}

		result <- i

	}
}

//生成随机数
func generate() {
	rand.Seed(time.Now().UnixNano())
	intn := rand.Intn(1000)
	jobChan <- intn
}

//取出一个并转成切片
func getNum() []rune {
	i := <-jobChan
	s := strconv.FormatInt(int64(i), 10)
	runes := []rune(s)
	return runes
}

//求和
func sum(val []rune) int64 {
	var i int64
	for _, value := range val {
		a := fmt.Sprintf("%d", value)
		var err error
		i, err = strconv.ParseInt(a, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

	}
	return i
}

//写入
func writer(i int64) {
	resultChan <- i
}

//working
func working() {
	//read
	for i := 0; i < 100; i++ {
		writer(sum(getNum()))
	}
}
