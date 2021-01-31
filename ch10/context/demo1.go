package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func TestContext(){
	var wg sync.WaitGroup
	wg.Add(1)
	ctx,stop:=context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		WatchDog3(ctx,"监控狗")
	}()

	time.Sleep(3*time.Second)
	stop()
	wg.Wait()


}

//一直循环到接收到停止指令
func WatchDog3(context context.Context,name string){

	for {
		select {
		case <-context.Done():
			fmt.Println(name,"停止指令收到，正在停止")
			return
		default:
			fmt.Println(name,"监控中")
		}
		time.Sleep(time.Second)
	}

}