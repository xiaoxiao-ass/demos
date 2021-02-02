package pipeline

import (
	"fmt"
	"sync"
)

func Test2(){
	s:= buy(100)

	p1:=build(s)
	p2:=build(s)
	p3:=build(s)


	m:=merge(p1,p2,p3)
	/*for v:=range m{
		fmt.Println(v)
	}*/


	pack:=pack(m)

	for v:=range pack{
		fmt.Println(v)
	}
}

//扇入函数（组件），把多个chanel中的数据发送到一个channel中
func merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	//把一个channel中的数据发送到out中
	p:=func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	wg.Add(len(ins))
	//扇入，需要启动多个goroutine用于处于多个channel中的数据
	for _,cs:=range ins{
		go p(cs)
	}
	//等待所有输入的数据ins处理完，再关闭输出out
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}


func Test1(){
	s:= buy(10)

	p:=build(s)

	pack:=pack(p)

	for v:=range pack{
		fmt.Println(v)
	}
}

func buy(num int) <-chan string{
	ch:=make(chan string)

	go func() {
		defer close(ch)
		for i:=1;i<num;i++{
			ch<-fmt.Sprint(i,"号")
		}
	}()
	return ch
}

//工序2组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

//工序3打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}


func Test3(){
	chs:=make(chan int,100)

	ch:=make(chan int,100)
	ch<-1
	ch<-2
	ch<-1
	close(ch)
	ch2:=make(chan int,100)
	ch2<-1
	ch2<-2
	ch2<-1
	close(ch2)
	ch3:=make(chan int,100)
	ch3<-1
	ch3<-2
	ch3<-1
	close(ch3)
	for v:=range ch{
		chs<-v
	}
	for v:=range ch2{
		chs<-v
	}
	for v:=range ch3{
		chs<-v
	}
	close(chs)
	for v:=range chs{
		fmt.Println(v)
	}
}