package main

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

func main() {
	//typeValue()
	//method()
//	updateValue()
	//kind()

	//types()
	//convert()
	//getTag()
	//structToString()
    //fmt.Print(reflect.TypeOf(( * fmt.Stringer)(nil)).Elem())
/*	p:=person{1,"xiao"}
	fmt.Print(reflect.TypeOf(p))*/


	/*p:=person{1,"xiao"}

	t:=reflect.ValueOf(p)
	//args:= []reflect.Value{reflect.ValueOf("对应参数类型")}  方法有参数时添加
	fmt.Print(t.MethodByName("As").Call(nil))*/








}

func structToString(){
	p:=person{1,"xiao"}
	t:=reflect.TypeOf(p)  //main.person
	v:=reflect.ValueOf(p)  //{1,"xiao"}
	//fmt.Print(v)

	build:=strings.Builder{}

	build.WriteString("{")

	for i:=0;i<t.NumField();i++{
		//tag:=t.Field(i).Tag.Get("json")
		tag:=t.Field(i).Name
		value:=v.Field(i)
		build.WriteString("\""+tag+"\"")
		build.WriteString(":")
		build.WriteString(fmt.Sprintf("%v",value))

		if i<t.NumField()-1{
			build.WriteString(",")
		}
	}
	build.WriteString("}")

	fmt.Print(build.String())






}




func getTag(){
	p:=person{1,"xiao"}
	t:=reflect.TypeOf(p)

	for i:=0;i<t.NumField();i++{
		sf:=t.Field(i)
		fmt.Println(sf.Tag.Get("yaml"),sf.Tag.Get("json") )
	}
}

//convert
func convert(){
	p:=person{1,"xiao"}
	//struct->json string
	v,err:=json.Marshal(p)
	if err==nil{
		fmt.Println(string(v))
	}


	//json string->struct
	s:=`{"id":1}`
	json.Unmarshal([]byte(s),&p)
	fmt.Println(p)

}

//
func types(){
	p:=person{1,"xiao"}
	t:=reflect.TypeOf(p) // main.person 包+类名
	/*fmt.Println(t)
     fmt.Println(t.NumField())
	for i:=0;i<t.NumField();i++ {
		fmt.Println("字段：",t.Field(i).Name)
	}

	for i:=0;i<t.NumMethod();i++ {
		fmt.Println("方法：",t.Method(i).Name)
	}*/

	stringerType:=reflect.TypeOf((* fmt.Stringer)(nil)).Elem()
	fmt.Println(stringerType)
	writerType:=reflect.TypeOf((* io.Writer)(nil)).Elem()

	fmt.Println("是否实现了fmt.Stringer：\",t.Implements(stringerType))\n\tfmt.Println(\"是否实现了io.Writer：",t.Implements(writerType))



}





//kind获取对应的底层类型
func kind(){
	p:=person{1,"xiao"}
	ppv:=reflect.ValueOf(&p)
	fmt.Println(ppv.Kind())  //ptr
	pp:=reflect.ValueOf(p)
	fmt.Println(pp.Kind())   //struct
}




//需要修改的值必须为可export,首字母大写
type person struct {
	Id int   `json:"id" yaml:"ss"` //tag
	name string

}

func (p person) As() string{
	return fmt.Sprintf("id is %d",p.Id)
}
/*
func (p person) String() string{
	return fmt.Sprintf("id is %d",p.Id)
}*/

func (p person) Sa() string{
	return fmt.Sprintf("id is %d",p.Id)
}

//修改值传递指针
func updateValue(){
	/*i:=3
	//reflect.ValueOf返回的是值的拷贝
	iv:=reflect.ValueOf(&i)
	fmt.Println(iv.Kind())
	//Elem()  获得指针指向的值
    iv.Elem().SetInt(1)
	fmt.Println(i)*/

	//
	p:=person{1,"xiao"}
	s:=reflect.ValueOf(&p)
	s.Elem().Field(0).SetInt(2)
	//s.Elem().Field(1).SetString("ssss")   error name不能修改，私有的
	fmt.Println(p)
	fmt.Println(s.Elem().Field(0).CanSet())
	fmt.Println(s.Elem().Field(1).CanSet())

}



//reflect.Value 转int             reflect.Value->interface{}->int
func method(){
	i:="ssss"
	//int to reflect.Value
	iv:=reflect.ValueOf(i)

	//reflect.Value to int
	a,ok:=iv.Interface().(string)
	if ok{
		fmt.Println(a)
	}


}



func typeValue(){
	i:=3
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(reflect.ValueOf(i))
}


