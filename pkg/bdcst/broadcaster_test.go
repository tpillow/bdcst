package bdcst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBroadcasterNoListenersSendDoesntBlock(t *testing.T) {
	br := NewBroadcaster[int]()
	br.Send(10)
}

func TestBroadcasterAddingListenersIncreasesNumListeners(t *testing.T) {
	br := NewBroadcaster[int]()
	assert.Equal(t, 0, br.NumListeners())
	br.AddListener(NewChannelListener(make(chan int)))
	assert.Equal(t, 1, br.NumListeners())
	br.AddListener(NewChannelListener(make(chan int)))
	assert.Equal(t, 2, br.NumListeners())
}

func TestBroadcasterCallbackListenerReceivesData(t *testing.T) {
	br := NewBroadcaster[int]()
	complete := make(chan int, 1)
	br.AddListener(NewCallbackListener(func(data int) {
		complete <- data
	}))
	br.Send(10)
	assert.Equal(t, 10, <-complete)
}

func TestBroadcasterChannelListenerReceivesData(t *testing.T) {
	br := NewBroadcaster[int]()
	complete := make(chan int, 1)
	br.AddListener(NewChannelListener(complete))
	br.Send(10)
	assert.Equal(t, 10, <-complete)
}

func TestBroadcasterCallbackAndChannelListenerReceiveData(t *testing.T) {
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

func TestBroadcasterRemoveUnaddedListener(t *testing.T) {
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

	br.Send(10)
	assert.Equal(t, 10, <-complete)

	br.RemoveListener(cbl)
	assert.Equal(t, 0, br.NumListeners())

	br.Send(20)
	select {
	case <-complete:
		assert.Fail(t, "Unexpected data received by a listener")
	default:
	}
}

func TestBroadcasterRemoveChannelListener(t *testing.T) {
	br := NewBroadcaster[int]()
	complete := make(chan int, 1)
	cbl := NewChannelListener(complete)
	br.AddListener(cbl)

	br.Send(10)
	assert.Equal(t, 10, <-complete)

	br.RemoveListener(cbl)
	assert.Equal(t, 0, br.NumListeners())

	br.Send(20)
	select {
	case <-complete:
		assert.Fail(t, "Unexpected data received by a listener")
	default:
	}
}

func TestBroadcasterTwoListenersRemoveCallbackListener(t *testing.T) {
	br := NewBroadcaster[int]()
	complete := make(chan int, 1)
	cbl := NewCallbackListener(func(data int) {
		complete <- data
	})
	br.AddListener(cbl)
	complete2 := make(chan int, 1)
	cbl2 := NewChannelListener(complete2)
	br.AddListener(cbl2)

	br.Send(10)
	assert.Equal(t, 10, <-complete)
	assert.Equal(t, 10, <-complete2)

	br.RemoveListener(cbl)

	br.Send(20)
	assert.Equal(t, 20, <-complete2)
	select {
	case <-complete:
		assert.Fail(t, "Unexpected data received by a listener")
	default:
	}
}

func TestBroadcasterTwoListenersRemoveChannelListener(t *testing.T) {
	br := NewBroadcaster[int]()
	complete := make(chan int, 1)
	cbl := NewCallbackListener(func(data int) {
		complete <- data
	})
	br.AddListener(cbl)
	complete2 := make(chan int, 1)
	cbl2 := NewChannelListener(complete2)
	br.AddListener(cbl2)

	br.Send(10)
	assert.Equal(t, 10, <-complete)
	assert.Equal(t, 10, <-complete2)

	br.RemoveListener(cbl2)

	br.Send(20)
	assert.Equal(t, 20, <-complete)
	select {
	case <-complete2:
		assert.Fail(t, "Unexpected data received by a listener")
	default:
	}
}
