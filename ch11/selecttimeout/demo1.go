package selecttimeout

import (
	"fmt"
	"time"
)

func Fortimeout(){
		result := make(chan string)
		go func() {
			//模拟网络访问
			time.Sleep(3 * time.Second)
			//time.Sleep(8 * time.Second)
			result <- "服务端结果"
		}()
		select {
		case v := <-result:
			fmt.Println(v)

		//time.After  设置超时时间,防止因为异常造成 select 语句的无限等待
		case <-time.After(5 * time.Second):
			fmt.Println("网络访问超时了")
		}




}
