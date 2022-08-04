package bdcst

import (
	"log"
	"sync"
)

type IBroadcaster[T any] interface {
	Send(data T)
	NumListeners() int
	AddListener(l Listener[T])
	RemoveListener(l Listener[T])
}

type Broadcaster[T any] struct {
	mutex     sync.Mutex
	listeners []Listener[T]
}

func NewBroadcaster[T any]() *Broadcaster[T] {
	return &Broadcaster[T]{}
}

func (broadcaster *Broadcaster[T]) NumListeners() int {
	return len(broadcaster.listeners)
}

func (broadcaster *Broadcaster[T]) Send(data T) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	for _, l := range broadcaster.listeners {
		l.notify(data)
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

	log.Panicf("Cannot remove unregistered listener: %v", l)
}
