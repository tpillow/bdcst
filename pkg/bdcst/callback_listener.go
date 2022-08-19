package bdcst

type CallbackListener[T any] struct {
	callback func(data T)
}

func NewCallbackListener[T any](callback func(data T)) *CallbackListener[T] {
	return &CallbackListener[T]{callback}
}

func (l CallbackListener[T]) OnNotify(data T) {
	l.callback(data)
}
