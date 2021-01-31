package main

import (
	"demo/ch07/demo"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	id int
	name string
}

func (person *Person) Walk()string{
	fmt.Println("person Walk is start")
	return person.name
}

func (person *Person) Run()  {
	fmt.Sprintf("person run is start")
}


func test(run demo.WalkRun)string{
	return run.Walk()
}

/**
 error panic deferred
  defer 在一个方法内可以有多个，执行顺序为从后往前，从下到上，遵循后进先出的原则 defer对应一个调用栈
 */
func main() {
	p:=Person{1,"xiao"}
   s:=test(&p)
   fmt.Println(s)
   v,e:=add(1,1) //error不处理可用_忽略

   if e!=nil{
   	fmt.Println(e)
   }
   fmt.Println(v)
	test_err()
	test_deferred()
	//test_panic()
	test_recover()
}

//从panic中恢复
func test_recover(){
	defer func() {
		//recover()的值=panic中传递的参数
		if p:=recover();p!=nil{
			fmt.Println("ddd",p)
		}
	}()
	a(-1)
}

//一般情况下不使用panic,使用error处理异常
func test_panic(){
	a(-2)
}

func a(i int){
	if i<0{
		panic("i<0")
	}
}



func test_deferred(){
	a, e:=ReadFile("C:\\Users\\weixiao\\Pictures\\3.png")
	fmt.Println(a,e)
}

func ReadFile(filename string)([]byte,error){
	f,e:=os.Open(filename)
	if e!=nil{
		return nil,errors.New("read error")
	}
	defer f.Close() //不管有没有错最后都会执行，一般成对操作：开启关闭 连接断开 加锁释放锁
	return ioutil.ReadAll(f)
}



func test_err(){
	e:=errors.New("error...........")
	e2:=fmt.Errorf("sss,%w",e)//只支持一个%w
	fmt.Println(e2)
	//获取原始error
	fmt.Println(errors.Unwrap(e2))
	//判断两个error是否为一个
	fmt.Println(errors.Is(errors.Unwrap(e2),e))
	//var cm *commonError
	//errors.As(errors.Unwrap(e2),&cm)
}


func add(a,b int)(int,error){
	if a>0&&b>0{
		return a+b,nil
	}
	return 0,errors.New("is null")
}
