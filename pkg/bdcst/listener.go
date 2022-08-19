package bdcst

// A Listener can be notified of data sent by an Emitter.Notify call.
type Listener[T any] interface {
	// Called when any Emitter that holds this Listener sends data by a
	// call to its Emitter.Notify.
	OnNotify(data T)
}
