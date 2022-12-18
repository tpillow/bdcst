package bdcst

type ValBroadcaster[T any] struct {
	Broadcaster[T]
	value T
}

// Returns a new Broadcaster[T] with the provided listeners pre-added.
func NewValBroadcaster[T any](value T, listeners ...Listener[T]) *ValBroadcaster[T] {
	return &ValBroadcaster[T]{
		Broadcaster[T]{listeners: listeners},
		value,
	}
}

func (vb *ValBroadcaster[T]) SetValue(value T) {
	vb.mutex.Lock()
	defer vb.mutex.Unlock()
	vb.value = value
	vb.notifyNoLock(vb.value)
}

func (vb *ValBroadcaster[T]) GetValue() T {
	vb.mutex.Lock()
	defer vb.mutex.Unlock()
	return vb.value
}
