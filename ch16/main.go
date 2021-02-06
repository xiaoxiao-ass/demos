package main

import (
	"fmt"
	"unsafe"
)

func main() {
	
}



func testPointer(){
	a:=1
	f:=(* float32)(unsafe.Pointer(&a))
	*f*=2
	fmt.Print()
}

