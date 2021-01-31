package main

import (
	"errors"
	"fmt"
)

/**
  多维数组，带参函数，多个返回值，返回值带名称
 */
func main(){

	//多维数组
	aa:=[3][2]int{{1,2},{3,4},{6,8}}
	for i,v:=range aa{
		fmt.Println(i,v)
	}
	fmt.Println(sum(1,2))
	r,v:=sum3(1,6)
	fmt.Println(r,v)
	fmt.Println(sum4(-1,1))
	fmt.Println(test(1,3))
}

//单个返回值
func sum(a int,b int) int{
	return a+b
}

//参数类型一样,前面参数类型可省略
func sum2(a ,b int) int{
	return a+b
}

//多个返回值,只有类型，没有名称
func sum3(a int,b int)(int,error){
	if a>0&&b>0{
		return a+b,nil
	}
	return 0,errors.New("参数不大于0")
}

//多个返回值,带类型带名称
func sum4(a int,b int)(r int,e error){
	if a>0&&b>0{
		r=a+b
	}else{
		r=-1
		e=errors.New("参数不大于0")
	}
	return
}

//可变参数
func test(a ...int)int{
	var sum int
	for _,v:=range a{
		sum+=v

	}
	return sum

}

