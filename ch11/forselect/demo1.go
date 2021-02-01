package forselect

import (
	"context"
	"fmt"
)

//无限for select
func InfiniteFor(ctx context.Context,name string){
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(name,"。。。。。。")

		}
	}
}

//有限for select


func InfiniteFor2(ctx context.Context,name string){
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(name,"。。。。。。")

		}
	}
}
