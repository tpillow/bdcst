package bdcst

import (
	"log"
	"sync"
)

// IBroadcaster provides an interface to add and remove listeners and send data to listeners.
type IBroadcaster[T any] interface {
	Send(data T)
	NumListeners() int
	AddListener(l IListener[T])
	RemoveListener(l IListener[T])
}

// Broadcaster is the most basic IBroadcaster implementation. Use Send to notify all added
// listeners of data. All Broadcaster operations lock a mutex before operation.
type Broadcaster[T any] struct {
	mutex     sync.Mutex
	listeners []IListener[T]
}

// NewBroadcaster creates a Broadcaster for use.
func NewBroadcaster[T any]() *Broadcaster[T] {
	return &Broadcaster[T]{}
}

// NumListeners returns the number of added listeners.
func (broadcaster *Broadcaster[T]) NumListeners() int {
	return len(broadcaster.listeners)
}

// Send calls notify of every Listener with data in the order that the
// listeners were added.
func (broadcaster *Broadcaster[T]) Send(data T) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	for _, l := range broadcaster.listeners {
		l.notify(data)
	}
}

// AddListener adds the listener l to the broadcaster, and will then be
// notified of future Send events.
func (broadcaster *Broadcaster[T]) AddListener(l IListener[T]) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	broadcaster.listeners = append(broadcaster.listeners, l)
}

// RemoveListener removes the listener l from the broadcaster, and will no longer be
// notified of future Send events. If the listener l is not added, a panic will occur.
func (broadcaster *Broadcaster[T]) RemoveListener(l IListener[T]) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	for idx, check := range broadcaster.listeners {
		if check == l {
			broadcaster.listeners = append(broadcaster.listeners[:idx], broadcaster.listeners[idx+1:]...)
			return
		}
	}

	log.Panicf("Cannot remove unadded listener: %v", l)
}
