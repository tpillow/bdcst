package main

import (
	"fmt"

	"github.com/tpillow/bdcst/pkg/bdcst"
)

// A function used as a CallbackListener.
func myCallbackFunc1(data string) {
	fmt.Printf("Listener 1 received data: %v\n", data)
}

// A function used as a CallbackListener.
func myCallbackFunc2(data string) {
	fmt.Printf("Listener 2 received data: %v\n", data)
}

func main() {
	// Create a new broadcaster that can notify string data.
	broadcaster := bdcst.NewBroadcaster[string](
		// Begin with a CallbackListener for myCallbackFunc1 already added.
		bdcst.NewCallbackListener(myCallbackFunc1),
	)

	// Create another CallbackListener for myCallbackFunc2.
	listener2 := bdcst.NewCallbackListener(myCallbackFunc2)
	// Add listener2 so it will be notified of data.
	broadcaster.AddListener(listener2)

	// Notify listeners
	broadcaster.Notify("Hello")
	// Remove listener2; it will no longer be updated by broadcaster.Notify.
	broadcaster.RemoveListener(listener2)
	// Notify listeners
	broadcaster.Notify("World!")
}
