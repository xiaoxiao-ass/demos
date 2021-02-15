package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//intToFloat()

//
	//pointerCalculation()

	//normalSetValue()
	sizeOf()
}

func sizeOf(){
	p:=make([]int,5,10)
	fmt.Println(len(p))
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(0)))
	fmt.Println(unsafe.Sizeof(int16(10)))
	fmt.Println(unsafe.Sizeof(int32(10000000)))
	fmt.Println(unsafe.Sizeof(int64(10000000000000)))
	fmt.Println(unsafe.Sizeof(int(10000000000000000)))
	fmt.Println(unsafe.Sizeof(string("飞雪无情")))
	fmt.Println(unsafe.Sizeof([]string{"飞雪无情"}))
	fmt.Println(unsafe.Sizeof(person{}))
 ps:=new(person)
 fmt.Println(unsafe.Alignof(ps.Id))
	fmt.Println(unsafe.Alignof(ps.I))
	fmt.Println(unsafe.Alignof(ps.Name))

	fmt.Println(unsafe.Offsetof(ps.Id))
	fmt.Println(unsafe.Offsetof(ps.I))
	fmt.Println(unsafe.Offsetof(ps.Name))


fmt.Println(&ps.Id)
	fmt.Println(&ps.Name)
}




//正常设置值
func normalSetValue(){
	p:=new(person)
	p.Name="sss"
	p.Id=1
	fmt.Println(*p)
}

//指针偏移量计算
func pointerCalculation(){
	p:=new(person)
	//因为id为结构体的第一个字段，所以不需要进行位置偏移即可赋值
	pId:=(*int)(unsafe.Pointer(p))
	*pId=100000


	//Name不是第一个字段，赋值需要计算偏移量
	pName:=(*string)(unsafe.Pointer(uintptr(unsafe.Pointer(p))+unsafe.Offsetof(p.Name)))

	*pName="哈哈哈哈"
	fmt.Println(*p)


}



//int类型指针转换为float指针     *int -> unsafe.Pointer ->*float
func intToFloat(){

	   i:=10
	   p :=(* float64)(unsafe.Pointer(&i))
	   *p*=2
	   fmt.Println(i)

}



type person struct {
    I int8
	Id int
	Name string
}


/*func (person person) String()string{
	return fmt.Sprintf(person.Name)
}
*/


func testPointer(){
	a:=1
	f:=(* float32)(unsafe.Pointer(&a))
	*f*=2
	fmt.Print()
}

