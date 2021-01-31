package mutex

import (
	"fmt"
	"sync"
	"time"
)
var(
	Sum int
	//互斥锁  sync.Mutex
	mutex sync.Mutex
)



func Test3(){
	for i:=0; i<100;i++  {
		go Add3(10)
	}

	for i:=0; i<10;i++  {
		go fmt.Println(read2())
	}
	time.Sleep(2*time.Second)
}

//加锁，读的时候不能写，共用互斥锁mutex  ,但会有性能问题
func read2()int{
	mutex.Lock()
	defer mutex.Unlock()
	b:=Sum
	return b
}

func Test2(){
	for i:=0; i<100;i++  {
		go Add3(10)
	}

	for i:=0; i<10;i++  {
		go fmt.Println(read())
	}
	time.Sleep(2*time.Second)
}


//改造前，没加锁，可能会读取到脏数据 ，读的时候在写数据
func read()int{
	return Sum
}



//Lock,UnLock一对存在，  加锁后必须释放锁
func Add3(i int){
	mutex.Lock()
	defer  mutex.Unlock() // 不管结果，最后都解锁，防止忘记解锁
	Sum+=i  //临界区

}

//Lock,UnLock一对存在，  加锁后必须释放锁
func add2(i int){
	mutex.Lock()
	Sum+=i  //临界区
	mutex.Unlock()
}




//线程不安全，共享资源会出现多种结果 理想1000 可能出现 990 980的情况 ，交叉执行
func test1(){
	for i:=0;i<100 ;i++  {
		go add(10)

	}
	time.Sleep(2*time.Second)
	fmt.Print(Sum)
}

func add(i int){
	Sum+=i

}
