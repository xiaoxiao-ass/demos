package main

import (
	"fmt"
	"time"
)

/**
  goroutine,channel

 */
func main() {
	//test_go()
	//test_channel()
	//bufferChannel()
	//test_close()
	test_select()
}



//select+channel 多路复用
func test_select(){
	ch1:=make(chan string)
	ch2:=make(chan string)
	ch3:=make(chan string)

	go func() {
		ch1<-download("file1")
	}()

	go func() {
		ch1<-download("file2")
	}()

	go func() {
		ch1<-download("file3")
	}()


	select{
 		case p:=<-ch1:
 			fmt.Print(p)
 		case p:=<-ch2:
			fmt.Print(p)
		case p:=<-ch3:
		  fmt.Print(p)
	//default:


	}
}

func download( file string)string{
	time.Sleep(time.Second)
	return fmt.Sprintf("file: %v",file)
}

//函数参数out只能用来发送数据不能接收
func test_func_param_channel(out chan <- int){

	out<-2
	//s:=<-out 报错
}

//单向channel
func singleChannel(){
	//只用来发送消息
	//onlySend:=make(chan <- string)
	//只用来接收
	//onlyReceive:=make(<- chan string)
}

func test_close(){
	ch:=make(chan int,5)
	ch<-2
	ch<-3
	//close(ch)  //不能向一个关闭的channel中发送数据，否则报错panic
	ch<-4
	fmt.Print(<-ch)
}

//有缓冲channel
func bufferChannel(){
	ch:=make(chan int,5)
	ch<-2
	ch<-3
	fmt.Printf("cache大小为：%v,元素个数%v",cap(ch),len(ch))
}




//通过go关键字启动一个协程
func test_go(){
	go fmt.Println("xiao")
	fmt.Println("sssss")
	time.Sleep(time.Second)
}


//声明一个channel
func test_channel(){
  ch:=make(chan string)
  go func() {
  	fmt.Println("x")
  	ch<-"s"
  }()
  fmt.Println("ssss")
  v:=<-ch
  fmt.Println(v)
}