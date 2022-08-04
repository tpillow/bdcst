package bdcst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBroadcasterNoListeners(t *testing.T) {
	br := NewBroadcaster[int]()
	assert.Equal(t, 0, len(br.listeners))
	br.Send(10)
}

func TestBroadcasterCallbackListener(t *testing.T) {
	br := NewBroadcaster[int]()
	br.Send(10)
	complete := make(chan int, 1)
	br.AddListener(CallbackListener[int](func(data int) {
		complete <- data
	}))
	br.Send(20)
	assert.Equal(t, 20, <-complete)
}

func TestBroadcasterChannelListener(t *testing.T) {
	br := NewBroadcaster[int]()
	br.Send(10)
	complete := make(chan int, 1)
	br.AddListener(ChannelListener[int](complete))
	br.Send(20)
	assert.Equal(t, 20, <-complete)
}

func TestBroadcasterCallbackAndChannelListener(t *testing.T) {
	br := NewBroadcaster[int]()
	fComplete := make(chan int, 1)
	cComplete := make(chan int, 1)
	br.AddListener(CallbackListener[int](func(data int) {
		fComplete <- data
	}))
	br.AddListener(ChannelListener[int](cComplete))
	br.Send(10)
	assert.Equal(t, 10, <-fComplete)
	assert.Equal(t, 10, <-cComplete)
	br.Send(20)
	assert.Equal(t, 20, <-fComplete)
	assert.Equal(t, 20, <-cComplete)
}

func TestBroadcasterRemoveInvalidListener(t *testing.T) {
	br := NewBroadcaster[int]()
	assert.Panics(t, func() {
		br.RemoveListener(CallbackListener[int](func(data int) {}))
	})
}

func TestBroadcasterTwoListenersRemoveOne(t *testing.T) {
	br := NewBroadcaster[int]()
	complete := make(chan int, 1)
	complete2 := make(chan int, 1)
	listener := ChannelListener[int](complete)
	br.AddListener(listener)
	assert.Equal(t, 1, len(br.listeners))
	br.AddListener(ChannelListener[int](complete2))
	assert.Equal(t, 2, len(br.listeners))
	br.Send(10)
	assert.Equal(t, 10, <-complete)
	assert.Equal(t, 10, <-complete2)

	br.RemoveListener(listener)
	assert.Equal(t, 1, len(br.listeners))
	br.Send(20)
	assert.Equal(t, 20, <-complete2)
	select {
	case <-complete:
		assert.Fail(t, "Unexpected data sent to removed listener")
	default:
	}
}
