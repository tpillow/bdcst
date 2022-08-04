package bdcst

import "sort"

// PriorityBroadcaster will sort listeners before sending data to them.
type PriorityBroadcaster[P any, T any] struct {
	broadcaster Broadcaster[T]
	priorities  []P
	comparer    func(a P, b P) bool
}

// NewPriorityBroadcaster creates a new PriorityBroadcaster with the provided comparison function.
// Listeners with lower (<) priorities will receive data first.
func NewPriorityBroadcaster[P any, T any](comparer func(a P, b P) bool) *PriorityBroadcaster[P, T] {
	return &PriorityBroadcaster[P, T]{
		comparer: comparer,
	}
}

// NumListeners returns the number of added listeners.
func (broadcaster *PriorityBroadcaster[P, T]) NumListeners() int {
	return broadcaster.broadcaster.NumListeners()
}

// Send calls notify of every Listener with data in the order of priority
// specified for each listener. Listeners with lower (<) priorities will
// receive data first.
func (broadcaster *PriorityBroadcaster[P, T]) Send(data T) {
	broadcaster.broadcaster.Send(data)
}

// AddListener adds the listener l to the broadcaster, and will then be
// notified of future Send events.
func (broadcaster *PriorityBroadcaster[P, T]) AddListener(priority P, l Listener[T]) {
	broadcaster.broadcaster.AddListener(l)

	broadcaster.broadcaster.mutex.Lock()
	defer broadcaster.broadcaster.mutex.Unlock()

	broadcaster.priorities = append(broadcaster.priorities, priority)
	broadcaster.sortListeners()
}

// RemoveListener removes the listener l from the broadcaster, and will no longer be
// notified of future Send events. If the listener l is not added, a panic will occur.
func (broadcaster *PriorityBroadcaster[P, T]) RemoveListener(l Listener[T]) {
	broadcaster.broadcaster.RemoveListener(l)

	broadcaster.broadcaster.mutex.Lock()
	defer broadcaster.broadcaster.mutex.Unlock()

	// todo...
	broadcaster.priorities = append(broadcaster.priorities, priority)
	broadcaster.priorities = append(broadcaster.priorities[:idx], broadcaster.priorities[idx+1:]...)
	broadcaster.sortListeners()
}

// sortListeners expects the mutex to already be locked
func (broadcaster *PriorityBroadcaster[P, T]) sortListeners() {
	sort.Slice(broadcaster.broadcaster.listeners, func(a int, b int) bool {
		return broadcaster.comparer()
	})
}
