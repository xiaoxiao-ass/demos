package  main

import "fmt"

func main()  {

 s:= buy(10)

 p:=build(s)

 pack:=pack(p)

 for v:=range pack{
 	fmt.Println(v)
 }

 /*for i:=0;i<10;i++{
	 select {
 		case v:=<-s:
 			fmt.Println(v)
	}
 }*/
}

func buy(num int) <-chan string{
	ch:=make(chan string)

	go func() {
		defer close(ch)
		for i:=1;i<num;i++{
            ch<-fmt.Sprint(i,"号")
		}
	}()
	return ch
}

//工序2组装
func build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

//工序3打包
func pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}