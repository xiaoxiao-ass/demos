package waitgroup

import (
	"demo/ch09/mutex"
	"demo/ch09/rwmutex"
	"fmt"
	"sync"
)

//等待所有协程执行完才结束，可用于管理多个协程同时执行一个任务   必须多协程下载一个文件
//改造后使用waitgroup
func Run2(){
	var wg sync.WaitGroup
	//监控110个协程，所以设置110
	wg.Add(110)
	for i:=0; i<100;i++  {
		go func() {
			//计数器减1
			defer wg.Done()
			mutex.Add3(10)
		}()

	}

	for i:=0; i<10;i++  {
		go func() {
			//计数器减1
			defer wg.Done()
			fmt.Println(rwmutex.ReadSum())
		}()
	}

	//一直等待，直到计数器为0
	wg.Wait()
}
