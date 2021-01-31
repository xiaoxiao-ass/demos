package rwmutex

import (
	mutex2 "demo/ch09/mutex"
	"sync"
)

var mutex  sync.RWMutex


func ReadSum()int{
	//只获取读锁
	mutex.RLock()
	defer mutex.RUnlock()
	b:=mutex2.Sum
	return b
}
