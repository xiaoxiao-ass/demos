package waitgroup

import (
	"demo/ch09/mutex"
	"demo/ch09/rwmutex"
	"fmt"
	"time"
)

//改造前 使用time
func Run(){
	for i:=0; i<100;i++  {
		go mutex.Add3(10)
	}

	for i:=0; i<10;i++  {
		go fmt.Println(rwmutex.ReadSum())
	}
	time.Sleep(2*time.Second)
}
