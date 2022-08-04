package main

import (
	"fmt"

	"github.com/tpillow/bdcst/pkg/bdcst"
)

func myCallbackFunc1(data string) {
	fmt.Printf("Listener 1 received data: %v\n", data)
}

func myCallbackFunc2(data string) {
	fmt.Printf("Listener 2 received data: %v\n", data)
}

func main() {
	broadcaster := bdcst.NewBroadcaster[string]()
	broadcaster.AddListener(bdcst.NewCallbackListener(myCallbackFunc1))
	listener2 := bdcst.NewCallbackListener(myCallbackFunc2)

	broadcaster.AddListener(listener2)
	broadcaster.Send("Hello")
	broadcaster.RemoveListener(listener2)
	broadcaster.Send("World!")
}
