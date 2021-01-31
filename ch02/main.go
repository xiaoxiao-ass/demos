package main

import "fmt"

func main() {
	//1有符号整型
	//有符号整型，值可以是正数，负数，0
	//带类型
	var a1 int =-11
	fmt.Println(a1)
	//不带类型，根据值自动判断
	var a2=2
	fmt.Println(a2)

	//无符号整型，值为0和正数
	var a3 uint=1
	fmt.Println(a3)

	//同时声明多个变量
	var (
		a=0
		b="s"
	)
	fmt.Println("a:",a,"\t b:",b)

	//浮点型
	var f32 float32=2.2
	fmt.Println(f32)
	var f64 float64=10.3456
	fmt.Println(f64)

	//布尔型
	var bf bool=false
	var bt bool=true
	fmt.Println("bf:",bf,"bt:",bt)

	//字符串
	var s1="1"
	var s2="2"
	s1+=s2
	fmt.Print("s1+s2:",s1)
}
