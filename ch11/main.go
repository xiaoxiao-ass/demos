package main

import (
	"context"
	"demo/ch11/forselect"
	"fmt"
	"sync"
	"time"
)



	func a(ctx context.Context,ch1 chan int){

		a:=[]int{1,2}
		for _,v:= range a{
			select {
			case <-ctx.Done():
				return
			case ch1<-v:

			}
		}
	}

	func tests(){
		var ch1  = make(chan int, 100)
		var ch2  = make(chan int, 100)
		ctx,_:=context.WithCancel(context.Background())
		ch2<-34
		ch2<-45
		close(ch2)

		go func(ch chan int) {
			a(ctx,ch)
			close(ch)

		}(ch1)

		//closee()

		for i := range ch1 {
			fmt.Print(i,"\t")
		}
	}

	func main() {
		//超时自动取消context
		ctx,_:=context.WithTimeout(context.Background(),2*time.Second)
		forselect.InfiniteFor2(ctx,"hah")

	}
































func LimitedFor(ctx context.Context,ch chan int){
	a:=[]int{1,2}
	for _,v:= range a{
		select {
		case <-ctx.Done():
			return
		case ch<-v:

		}
	}



}


func test(){
	var chs=make(chan int,5)
	var wg sync.WaitGroup
	wg.Add(1)
	ctx,closee:=context.WithCancel(context.Background())

	//var once sync.Once
	go func(ch chan int) {
		defer wg.Done()
		/*defer once.Do(func() {
			fmt.Println(len(ch))
			close(ch)
		})*/
		LimitedFor(ctx,ch)
		//close(ch)

	}(chs)

	closee()
	//wg.Wait()

	time.Sleep(3*time.Second)
	for data:=range chs{
		fmt.Println(data)
	}


	/*for{
		if data,ok:=<-chs;ok{
			fmt.Println(data)
		}
	}*/



}



func test1(){
	var wg sync.WaitGroup
	wg.Add(2)
	ctx,close:=context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		forselect.InfiniteFor(ctx,"ssssss")
	}()

	go func() {
		defer wg.Done()
		forselect.InfiniteFor(ctx,"ed")
	}()
	time.Sleep(3*time.Second)

	close()
	wg.Wait()
}
