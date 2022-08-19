package bdcst

import (
	"log"
	"sync"
)

type Broadcaster[T any] struct {
	mutex     sync.Mutex
	listeners []Listener[T]
}

func NewBroadcaster[T any](listeners ...Listener[T]) *Broadcaster[T] {
	return &Broadcaster[T]{listeners: listeners}
}

func (broadcaster *Broadcaster[T]) Notify(data T) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	for _, l := range broadcaster.listeners {
		l.OnNotify(data)
	}
}

func (broadcaster *Broadcaster[T]) AddListener(l Listener[T]) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	broadcaster.listeners = append(broadcaster.listeners, l)
}

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
