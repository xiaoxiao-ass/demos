package main

import "fmt"

func main() {
  a:=Age(1)
  a.String()
  a.String2()

  b:=Age(2)
	(&b).String()

  /*a.String3()
  fmt.Print(a)*/


 /* //方法赋值给变量，方法表达式
  s:=Age.String // =  var s fun(Age)=Age.String
  //通过变量，要传一个接收者进行调用，也就是age
  s(a)*/


  s :=(*Age).String4f // var s fun(*Age,int)=(*Age).String
  s(&a,3) //不管方法有没有参数，第一个参数都是接收者


}

type Age uint
func (age Age) String(){
	fmt.Println(age)
}

func (age *Age)String2(){
	fmt.Println(*age+1)
}

func (age *Age)String3(){
	*age=Age(20)
}

func (age *Age)String4f(a int){
	fmt.Println(a)
}


