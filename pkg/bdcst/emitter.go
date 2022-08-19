package bdcst

// An Emitter will call Listener.OnNotify for every listener that has been
// added to it when Emitter.Notify is called.
type Emitter[T any] interface {
	// Will call the Listener.OnNotify function for every added listener with
	// data.
	Notify(data T)
	// Adds a Listener to be notified by subsequent Emitter.Notify calls.
	AddListener(l Listener[T])
	// Removes a Listener, which will no longer be notified by subsequent
	// Emitter.Notify calls.
	RemoveListener(l Listener[T])
}
