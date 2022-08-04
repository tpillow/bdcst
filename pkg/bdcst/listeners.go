package bdcst

type Listener[T any] interface {
	notify(data T)
}

type CallbackListener[T any] func(data T)

func (callback CallbackListener[T]) notify(data T) {
	callback(data)
}

type ChannelListener[T any] chan<- T

func (channel ChannelListener[T]) notify(data T) {
	channel <- data
}
