package futures

import (
	"fmt"
	"time"
)

func WashingVegetables() <-chan string{
	fmt.Print("((((((((((((((((9")
	 ch:=make(chan  string)

	 go func() {
	 	time.Sleep(3*time.Second)
	 	ch<-"洗菜"
	 }()
	return ch
}

func BoilWater() <-chan string{
	ch:=make(chan  string)

	go func() {
		time.Sleep(5*time.Second)
		ch<-"烧水"
	}()
	return ch
}




