package main

import "fmt"

/**
  闭包 匿名函数
 */
func main() {
    sum7:=func (a,b int)int{
    	return a+b
	}
	a:=sum7(1,2)
	fmt.Println(a)


	as:=colsure() //i=1 func()
	fmt.Println(as())//i=2
	fmt.Println(as())//i=3

  /* fmt.Println(colsure())
	fmt.Println(colsure())*/
}

func colsure() func() int{
	i:=0
	i+=1
	return func() int {
		i++
		return i
	}
}
