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
	br.AddListener(NewCallbackListener(func(data int) {
		complete <- data
	}))
	br.Send(20)
	assert.Equal(t, 20, <-complete)
}

func TestBroadcasterChannelListener(t *testing.T) {
	br := NewBroadcaster[int]()
	br.Send(10)
	complete := make(chan int, 1)
	br.AddListener(NewChannelListener(complete))
	br.Send(20)
	assert.Equal(t, 20, <-complete)
}

func TestBroadcasterCallbackAndChannelListener(t *testing.T) {
	br := NewBroadcaster[int]()
	fComplete := make(chan int, 1)
	cComplete := make(chan int, 1)
	br.AddListener(NewCallbackListener(func(data int) {
		fComplete <- data
	}))
	br.AddListener(NewChannelListener(cComplete))
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
		br.RemoveListener(NewCallbackListener(func(data int) {}))
	})
}

func TestBroadcasterRemoveCallbackListener(t *testing.T) {
	br := NewBroadcaster[int]()
	complete := make(chan int, 1)
	cbl := NewCallbackListener(func(data int) {
		complete <- data
	})
	br.AddListener(cbl)
	assert.Equal(t, 1, br.NumListeners())
	br.Send(10)
	assert.Equal(t, 10, <-complete)

	br.RemoveListener(cbl)
	assert.Equal(t, 0, br.NumListeners())
}

func TestBroadcasterTwoListenersRemoveChannelListener(t *testing.T) {
	br := NewBroadcaster[int]()
	complete := make(chan int, 1)
	complete2 := make(chan int, 1)
	listener := NewChannelListener(complete)
	br.AddListener(listener)
	assert.Equal(t, 1, len(br.listeners))
	br.AddListener(NewChannelListener(complete2))
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
