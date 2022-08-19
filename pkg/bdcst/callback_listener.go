package bdcst

// A Listener that will call a callback function when CallbackListener.OnNotify
// is called.
type CallbackListener[T any] struct {
	// The callback function to call when CallbackListener.OnNotify is called.
	callback func(data T)
}

// Returns a new CallbackListener[T] instance with the provided callback.
func NewCallbackListener[T any](callback func(data T)) *CallbackListener[T] {
	return &CallbackListener[T]{callback}
}

// Calls the callback function with data.
func (l CallbackListener[T]) OnNotify(data T) {
	l.callback(data)
}
