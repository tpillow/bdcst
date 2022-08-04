package bdcst

type Listener[T any] interface {
	notify(data T)
}

type CallbackListener[T any] struct {
	callback func(data T)
}

func NewCallbackListener[T any](callback func(data T)) *CallbackListener[T] {
	return &CallbackListener[T]{callback}
}

func (callbackListener CallbackListener[T]) notify(data T) {
	callbackListener.callback(data)
}

type ChannelListener[T any] chan<- T

func NewChannelListener[T any](channel chan<- T) ChannelListener[T] {
	return ChannelListener[T](channel)
}

func (channel ChannelListener[T]) notify(data T) {
	channel <- data
}
