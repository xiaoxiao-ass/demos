package forselect

import (
	"fmt"
	"time"
)

//一直循环
func WatchDog1(name string){
	//开启for select循环，一直后台监控
	for{
		select  {
		default:
			fmt.Println(name,"正在监控……")
		}
		time.Sleep(1*time.Second)
	}
}

//一直循环到接收到停止指令
func WatchDog2(ch chan bool,name string){

	for {
		select {
		  case <-ch:
		  	fmt.Println(name,"停止指令收到，正在停止")
			  return
		default:
			fmt.Println(name,"监控中")
		}
		time.Sleep(time.Second)
	}

}