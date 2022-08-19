package bdcst

type Emitter[T any] interface {
	Notify(data T)
	AddListener(l Listener[T])
	RemoveListener(l Listener[T])
}
