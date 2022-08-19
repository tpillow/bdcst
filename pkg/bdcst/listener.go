package bdcst

type Listener[T any] interface {
	OnNotify(data T)
}
