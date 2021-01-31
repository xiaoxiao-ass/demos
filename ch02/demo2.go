package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//变量简短声明
	i:=10
	bf:=false
	fmt.Println(i,"\t",bf)

	pi:=&i
	//&i i的内存地址
	fmt.Println("i address:",&i)
	//pi内存地址
	fmt.Println("pi address:",pi)
	//pi内存地址上的值
	fmt.Println("pi value:",*pi)

	//常量 const
	const name="xiao"
	fmt.Println(name)

	//iota常量生成器，初始化相似规则的常量，避免重复的初始化（const中才使用）
	const(
		a=1
		b=2
		c=3
		d=4
	)

	const(
		one=iota+1   //one=(0)+1
		two          //two=(0+1)+1
		three 		 //three=(0+1+1)+1
		four         //four=(0+1+1+1)+1
	)
	fmt.Println(one,two,three,four)

	//字符串数字互转
	//转string itoa
	i2s:=strconv.Itoa(a)
	fmt.Println(i2s)

	//转int atoi
	is:="4"
    i3s,err:=strconv.Atoi(is)
	if err == nil {
		fmt.Printf("%T \t %v",i3s,i3s)
	}

	//int float互转
	f64:=22.222222
	i2f:=float64(i)
	f2i:=int(f64)
	fmt.Println()
	fmt.Println(i2f,f2i)

	str:="sssdeb"
	//字符串工具类strings
	fmt.Println(strings.HasPrefix(is,"i"))
	//查找指定字符串，没有-1 有则返回第一次出现的索引位置 b:5
	fmt.Println(strings.Index(str,"b"))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.Contains(str,"hh"))

}
