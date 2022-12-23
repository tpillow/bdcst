package bdcst

type ValBroadcaster[T comparable] struct {
	Broadcaster[T]
	value                  T
	notifySettingSameValue bool
}

// Returns a new Broadcaster[T] with the provided listeners pre-added.
func NewValBroadcaster[T comparable](value T, notifySettingSameValue bool, listeners ...Listener[T]) *ValBroadcaster[T] {
	return &ValBroadcaster[T]{
		Broadcaster[T]{listeners: listeners},
		value,
		notifySettingSameValue,
	}
}

func (vb *ValBroadcaster[T]) SetValue(value T) {
	vb.mutex.Lock()
	defer vb.mutex.Unlock()
	if !vb.notifySettingSameValue && vb.value == value {
		return
	}
	vb.value = value
	vb.notifyNoLock(vb.value)
}

func (vb *ValBroadcaster[T]) GetValue() T {
	vb.mutex.Lock()
	defer vb.mutex.Unlock()
	return vb.value
}
