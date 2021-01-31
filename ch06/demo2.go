package main

import (
 "demo/ch06/stuct"
 "fmt"
)

func main() {
  s:=stuct.Person{Id: 1, Name: "xiao"}
  //fmt.Println(s)

    ass := stuct.Ass{
        Demo: "s",
        Addr: stuct.Address{
            City:"湘",
        },
    }
   // fmt.Print(ass.Address.City)

   printString(s)
    printString(ass.Addr)
    p:=stuct.NewPerson("xiaoss")
    fmt.Println(p.Name)
    stuct.NewError("errors")
    //
    bss := stuct.Bss{Id: 1, Address: stuct.Address{City: "ddddd"}}
    fmt.Println(bss.City)
}

func printString(stringer fmt.Stringer){
    fmt.Println( stringer.String())
}
