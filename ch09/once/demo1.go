package once

import (
	"fmt"
	"sync"
)

//sync.Once 用于某个对象的单例，只加载一次的资源  只执行一次的场景
func DoOnce(){
	var once sync.Once
	onceBody:= func() {
		fmt.Println("only once")
	}

	ch:=make(chan  bool)

	for i:=0;i<10;i++{
		go func() {
			once.Do(onceBody)
			ch<-true
		}()
	}

	for i:=0;i<10;i++{
		fmt.Println(<- ch)
	}
}

