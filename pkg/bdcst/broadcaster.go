package bdcst

import (
	"log"
	"sync"
)

// A basic Emitter implementation whose operations are synchronized using a
// sync.Mutex.
type Broadcaster[T any] struct {
	// The mutex used to synchronize operations.
	mutex sync.Mutex
	// The list of added Listeners.
	listeners []Listener[T]
}

// Returns a new Broadcaster[T] with the provided listeners pre-added.
func NewBroadcaster[T any](listeners ...Listener[T]) *Broadcaster[T] {
	return &Broadcaster[T]{listeners: listeners}
}

// Calls the Listener.OnNotify function for every added listener with data.
func (broadcaster *Broadcaster[T]) Notify(data T) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()
	broadcaster.notifyNoLock(data)
}

func (broadcaster *Broadcaster[T]) notifyNoLock(data T) {
	for _, l := range broadcaster.listeners {
		l.OnNotify(data)
	}
}

// Adds a Listener to be notified by subsequent Broadcaster.Notify calls.
func (broadcaster *Broadcaster[T]) AddListener(l Listener[T]) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	broadcaster.listeners = append(broadcaster.listeners, l)
}

// Removes a Listener, which will no longer be notified by subsequent
// Broadcaster.Notify calls. Panics if the listener to remove has not been
// added.
func (broadcaster *Broadcaster[T]) RemoveListener(l Listener[T]) {
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
