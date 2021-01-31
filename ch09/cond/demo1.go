package cond

import (
	"fmt"
	"sync"
	"time"
)

func TestCond() {
	//sync.NewCond 函数生成一个 *sync.Cond，用于阻塞和唤醒协程；
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	//10个赛跑，1个发号施令
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(s int) {
			defer wg.Done()
			fmt.Println(s, "号就位")
			cond.L.Lock()
			cond.Wait()
			fmt.Println(s, "号开始跑")
			cond.L.Unlock()
		}(i)
	}

	//等待所有协程进入等待状态
	time.Sleep(2 * time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判好了")
		fmt.Println("开始")
		cond.Signal()
		cond.Broadcast() //发令枪响
	}()

	wg.Wait()
}
