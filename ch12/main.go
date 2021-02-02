package  main

import (
	"demo/ch12/futures"
	"fmt"
	"time"
)

func main()  {

	//管道
	//pipeline.Test2()

	//无缓冲，不能存数据，读取写入需同步否则阻塞
	/*out := make(chan string)
	go func() {
		out<-"sss"
	}()

	fmt.Println(<-out)*/

	/*select {
   		case v:=<-out:
   	     fmt.Println(v)
   }*/

	//futures模式  分散任务,不需要结果之前，忽略其他goroutine执行状态   需要直接获取返回结果，如果该goroutine还未执行完，阻塞等待执行完
	a:=futures.WashingVegetables()
	b:=futures.BoilWater()
	fmt.Println("休息")
	time.Sleep(2*time.Second)
	v:=<-a
	v2:=<-b

	fmt.Println(v,v2)





}

