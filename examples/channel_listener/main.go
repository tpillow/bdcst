package main

import (
	"fmt"

	"github.com/tpillow/bdcst/pkg/bdcst"
)

func main() {
	broadcaster := bdcst.NewBroadcaster[string]()
	myChan := make(chan string, 1)
	broadcaster.AddListener(bdcst.NewChannelListener(myChan))

	broadcaster.Send("Hello")
	fmt.Printf("Channel received: %v\n", <-myChan)
	broadcaster.Send("World!")
	fmt.Printf("Channel received: %v\n", <-myChan)
}
