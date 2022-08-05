package bdcst

// IListener represents an entity that can notify something of sent data of type T.
type IListener[T any] interface {
	notify(data T)
}

// CallbackListener will call the callback function when data of type T is sent.
type CallbackListener[T any] struct {
	callback func(data T)
}

// NewCallbackListener creates a CallbackListener with the provided callback function.
func NewCallbackListener[T any](callback func(data T)) *CallbackListener[T] {
	return &CallbackListener[T]{callback}
}

func (callbackListener CallbackListener[T]) notify(data T) {
	callbackListener.callback(data)
}

// ChannelListener will send data on the channel when data of type T is sent.
type ChannelListener[T any] chan<- T

// NewChannelListener creates a ChannelListener with the provided channel.
func NewChannelListener[T any](channel chan<- T) ChannelListener[T] {
	return ChannelListener[T](channel)
}

func (channel ChannelListener[T]) notify(data T) {
	channel <- data
}
