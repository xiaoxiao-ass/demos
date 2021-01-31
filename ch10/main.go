package main

import (
	"demo/ch10/context"
	"demo/ch10/forselect"
	"sync"
	"time"
)

func main() {
	//test1()
	//test2()
	context.TestContext()
}

func test2(){
	//syncmap.TestSyncMap()
	var wg sync.WaitGroup
	ch:=make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		forselect.WatchDog2(ch,"【监控狗1】")
	}()
    time.Sleep(5*time.Second)
	ch<-true
	wg.Wait()
}

func test1(){
	//syncmap.TestSyncMap()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		forselect.WatchDog1("【监控狗1】")
	}()
	wg.Wait()
}
