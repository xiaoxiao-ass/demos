package main

import(
	"fmt"
	"sync"
)
//开启goroutine将0~20的数发送到ch1中
//开启goroutine从ch1中接收值，将该值的平方发送到ch2中，最后把ch2中的所有值打印出来

var wg sync.WaitGroup

var once sync.Once

func f(ch1, ch2 chan int) {
	wg.Done()
	for n := range ch1 {
		ch2 <- n * n
	}
	once.Do(func(){ close(ch2) })
	// close(ch2)   //报错 ，一个goroutine关闭，另一个再关闭就会报错
}

func main() {
	var ch1 chan int = make(chan int, 100)
	var ch2 chan int = make(chan int, 100)

	// wg.Add(1)    //这里没必要加，因为下面这个goroutine不会阻塞
	go func(ch chan int) {
		for i:=0;i<20;i++{
			ch1 <- i
		}
		close(ch1)     // 放完之后必须要关闭（这样接收端就知道我接收完就退出），不关闭接收端不知道该阻塞还是该退出
	}(ch1)

	wg.Add(2)
	go f(ch1, ch2)
	go f(ch1, ch2)

	wg.Wait()

	for i := range ch2 {
		fmt.Println(i)
	}
}