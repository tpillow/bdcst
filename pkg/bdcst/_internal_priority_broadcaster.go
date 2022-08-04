// TODO: this is a work in progress, hence internal
package bdcst

import (
	"log"
	"sort"
	"sync"
)

type priorityListener[P any, T any] struct {
	priority P
	listener Listener[T]
}

func newPriorityListener[T any, P any](priority P, listener Listener[T]) *priorityListener[P, T] {
	return &priorityListener[P, T]{priority, listener}
}

type PriorityBroadcaster[P any, T any] struct {
	mutex     sync.Mutex
	listeners []*priorityListener[P, T]
	comparer  func(a P, b P) bool
}

func NewPriorityBroadcaster[P any, T any](comparer func(a P, b P) bool) *PriorityBroadcaster[P, T] {
	return &PriorityBroadcaster[P, T]{
		comparer: comparer,
	}
}

func NewIntPriBroadcaster[T any]() *PriorityBroadcaster[int, T] {
	return NewPriorityBroadcaster[int, T](func(a int, b int) bool {
		return a < b
	})
}

func (broadcaster *PriorityBroadcaster[P, T]) NumListeners() int {
	return len(broadcaster.listeners)
}

func (broadcaster *PriorityBroadcaster[P, T]) Send(data T) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	for _, l := range broadcaster.listeners {
		l.listener.notify(data)
	}
}

func (broadcaster *PriorityBroadcaster[P, T]) sortListeners() {
	sort.Slice(broadcaster.listeners, func(a int, b int) bool {
		return broadcaster.comparer(broadcaster.listeners[a].priority, broadcaster.listeners[b].priority)
	})
}

func (broadcaster *PriorityBroadcaster[P, T]) AddListener(priority P, l Listener[T]) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	broadcaster.listeners = append(broadcaster.listeners, newPriorityListener(priority, l))
	broadcaster.sortListeners()
}

func (broadcaster *PriorityBroadcaster[P, T]) RemoveListener(l Listener[T]) {
	broadcaster.mutex.Lock()
	defer broadcaster.mutex.Unlock()

	for idx, check := range broadcaster.listeners {
		if check.listener == l {
			// TODO: validate this
			broadcaster.listeners = append(broadcaster.listeners[:idx], broadcaster.listeners[idx+1:]...)
			broadcaster.sortListeners()
			return
		}
	}

	log.Panicf("Cannot remove unregistered listener: %v", l)
}
