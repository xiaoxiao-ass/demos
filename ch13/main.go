package main

import "fmt"

type person struct {
	id int
}

func a()*person{
	return &person{1}
}

func b()*person{

	a:=new(person)
	a.id=2
	return a
}



func main() {
	fmt.Println(a())
	fmt.Println(b())
}
