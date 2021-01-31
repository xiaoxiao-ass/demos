package syncmap

import (
	"fmt"
	"sync"
)

func TestSyncMap(){
	s:=sync.Map{}
	s.Store("b","bb")
	s.Store("a","ssss")
	s.Store("c","cc")
	fmt.Println(s.Load("a"))
	fmt.Println(s.Load("b"))
	fmt.Println(s.Load("d"))
	//s.Delete("a")
	//fmt.Println(s.Load("a"))

	//LoadOrStore 判断key是否存在，存在为true并返回原来值，不存在为false且保存value,并且返回
	fmt.Println(s.LoadOrStore("a","bbs"))
	fmt.Println(s.LoadOrStore("d","sd"))
	s.Range(func(key, value interface{}) bool {
		fmt.Println(key,value)
		return true
	})

}