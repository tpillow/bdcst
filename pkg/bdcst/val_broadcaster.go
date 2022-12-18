package bdcst

type ValBroadcaster[T any] struct {
	broadcaster *Broadcaster[T]
	value       T
}

// Returns a new Broadcaster[T] with the provided listeners pre-added.
func NewValBroadcaster[T any](val T, listeners ...Listener[T]) *ValBroadcaster[T] {
	return &ValBroadcaster[T]{
		broadcaster: NewBroadcaster[T](listeners...),
		value:       val,
	}
}

func (vb *ValBroadcaster[T]) SetValue(value T) {
	vb.broadcaster.mutex.Lock()
	defer vb.broadcaster.mutex.Unlock()
	vb.value = value
	vb.broadcaster.notifyNoLock(vb.value)
}

func (vb *ValBroadcaster[T]) GetValue() T {
	vb.broadcaster.mutex.Lock()
	defer vb.broadcaster.mutex.Unlock()
	return vb.value
}

func (vb *ValBroadcaster[T]) GetBroadcaster() *Broadcaster[T] {
	return vb.broadcaster
}
