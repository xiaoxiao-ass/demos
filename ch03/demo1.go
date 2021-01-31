package main

import "fmt"

func main() {

	//数组
	//指定长度，初始化值
	array:=[5]string{"1","2","3","ss","hello"}

	//设置长度，为指定索引位置设置值,没有设值的位置为该类型的默认值，int 0, bool false,不设置长度则为后面最大索引位置+1
    arrs:=[4]int{1:1,3:4}

    //不设置长度，以元素个数为数组大小
	arrays:=[...]int{1,2,3}

	for i,value:=range arrs{
		fmt.Println(i,"\t",value)
	}
	fmt.Println("oooooooooooooo")
	for i,value:=range arrays{
		fmt.Println(i,"\t",value)
	}
	fmt.Println("0000000000000")
	for i:=0;i< len(array);i++{
		fmt.Println(array[i])
	}
	fmt.Println("______________--")

	//for index,val:=range array{}
	for index,value:=range array{
		fmt.Println("index:",index,"\t value:",value)
	}

	fmt.Println("**********************")
	// 丢弃索引
	for _,value:=range array{
		fmt.Println(value)
	}

	//基于数组生成切片 slice:=array[start index:end index] 包前不包后
	slice:=array[2:5]
	for index,value:=range slice{
	fmt.Println("index:",index,"\t value:",value)
	}
	// 1 array[:4] 等价于array[0:4]  2 array[0:] 等价于array[0:5] 3 array[:] 等价于array[0:5]

	//修改切片值,
	slice[1]="f"
	//原数组值被改变
	//基于数组的切片，底层数组还是原来的数组,只是元素索引改变
	fmt.Println(array)

	//直接声明一个切片 ,make方式
	//指定长度
	slice1:=make([]string,4)
	//指定长度和容量,切片的容量不能比长度小
	slice2:=make([]int,4,8)
	slice2[0]=1

	//直接声明一个切片 ,不使用make方式
	slice3:=[]int{1,3,2}
	fmt.Println(slice3)

	//append会追加到空闲的位置，不在指定长度中操作  第一个追加的元素5下标为4
	slice2=append(slice2, 5,6)
	fmt.Println(slice2)
	fmt.Println(slice1)


   //map
   //创建key为string类型，value为int类型
   maps:=make(map[string]int)
   maps["name"]=1
   maps["as"]=4
   as,bs:=maps["names"]
   fmt.Println(as,bs)
	//delete(maps, "as")
   fmt.Println(maps)

   for key,value:=range maps{
   	fmt.Println(key,value)
   }


   a:=[]string{"1","2","5"}
   a[0]="1"
   b:=make([]int,5,10)
   b[4]=1
   fmt.Println(b)




   //追加一个切片
   sliceA :=[]int{1,2,34,56,7}
   sliceB :=[]int{7,8,9,9,0}
	sliceA =append(sliceA, sliceB...)
fmt.Println(sliceA)


s:="Aahello非恶意s"
for a,b :=range s{

	fmt.Println("___________")
	fmt.Println(a,b)
}

//utf8编码下 一个中文3个字节
fmt.Println(len(s))



}
