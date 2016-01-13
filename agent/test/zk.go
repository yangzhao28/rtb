package main

import (
	"fmt"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	children, stat, ch, err := c.ChildrenW("/a")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v %+v\n", children, stat)

	for {
		select {
		case e := <-ch:
			fmt.Printf("%+v\n", e)
		}
	}
	fmt.Println(ch)
}
