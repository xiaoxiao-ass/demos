package main

import (
	"fmt"
)

func main() {

	//if
	// if后紧跟的{不能独占一行 ，else前的}也不能独占一行
	i:=1
	if i>10 {
		fmt.Println("i>10")
	} else{
		fmt.Println("i<=10")
	}

    //在条件表达式前声明变量用;隔开
	if a:=6;a>1{
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}
	
	
	//swith,从上至下判断，满足就跳出，不使用break
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("default")
	}

	//多重选择
	switch a:=0;a {
	case 0:
		fallthrough
	case 1:
		fmt.Println("aaaa")
	default:
		fmt.Println("defalut")
	}


	//for循环
	sum:=0
	for i:=1;i<=100;i++{
		sum+=i
	}
	fmt.Println("sum:",sum)

	//go没有while,模拟while
	sum=0
	i=1
	for i<=100{
		sum+=i
		i++
	}
	fmt.Println(sum)
		}
