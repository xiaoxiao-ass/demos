package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main()  {
	p:=[2]string{"ddd","ssd"}
    a1:=p[0:1]
    a2:=p[:]

    fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&a1)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&a2)).Data)


    //自定义结构体接收，一般不建议
	sh1:=(*slice)(unsafe.Pointer(&a1))
	fmt.Println(sh1.Data,sh1.Len,sh1.Cap)


	a1s:=[2]string{"飞雪无情","张三"}
	fmt.Printf("函数main数组指针：%p\n",&a1s)
	arrayF(a1s)
	s1:=a1s[0:1]
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	sliceF(s1)



	stringToByteSlices1()



	stringToByteSlices2()


}

//零拷贝
func stringToByteSlices2(){
	s:="小"
	//by:=[]byte(s)
	by:=(*reflect.SliceHeader)(unsafe.Pointer(&s))
	by.Cap=by.Len
	ss:=*(*[]byte)(unsafe.Pointer(by))
	fmt.Println(ss)git

}

//零拷贝
func stringToByteSlices1(){
	s:="小"
	by:=[]byte(s)
	//ss:=string(by)
	ss:=*(*string)(unsafe.Pointer(&by))
	fmt.Println(ss)
}

//强制转换
func stringTobyteSlice() {


	s:="飞雪无情"
	fmt.Printf("s的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b:=[]byte(s)
	fmt.Printf("b的内存地址：%d\n",(*reflect.SliceHeader)(unsafe.Pointer(&b)).Data)
	s3:=string(b)
	fmt.Printf("s3的内存地址：%d\n", (*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
}

func arrayF(a [2]string){
	fmt.Printf("函数arrayF数组指针：%p\n",&a)
}
func sliceF(s []string){
	fmt.Printf("函数sliceF Data：%d\n",(*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)
}

type slice struct {
	Data uintptr
	Len  int
	Cap  int
}

type person struct {
	Id int
	Name string
}
