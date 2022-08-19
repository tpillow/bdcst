package bdcst_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpillow/bdcst/pkg/bdcst"
)

func TestBroadcasterNoListeners(t *testing.T) {
	b := bdcst.NewBroadcaster[int]()
	b.Notify(5)
}

func TestBroadcaster1CallbackListener0Added(t *testing.T) {
	received := 0
	b := bdcst.NewBroadcaster[int](bdcst.NewCallbackListener(func(data int) {
		received = data
	}))

	b.Notify(5)
	assert.Equal(t, 5, received)
	b.Notify(6)
	assert.Equal(t, 6, received)
}

func TestBroadcaster2CallbackListeners1Added(t *testing.T) {
	received0 := 0
	received1 := 0
	b := bdcst.NewBroadcaster[int](bdcst.NewCallbackListener(func(data int) {
		received0 = data
	}))
	b.AddListener(bdcst.NewCallbackListener(func(data int) {
		received1 = data
	}))

	b.Notify(5)
	assert.Equal(t, 5, received0)
	assert.Equal(t, 5, received1)
	b.Notify(6)
	assert.Equal(t, 6, received0)
	assert.Equal(t, 6, received1)
}

func TestBroadcasterAddRemove2CallbackListeners(t *testing.T) {
	received0 := 0
	received1 := 0

	l0 := bdcst.NewCallbackListener(func(data int) {
		received0 = data
	})
	l1 := bdcst.NewCallbackListener(func(data int) {
		received1 = data
	})

	b := bdcst.NewBroadcaster[int]()
	b.AddListener(l0)
	b.AddListener(l1)

	b.Notify(5)
	assert.Equal(t, 5, received0)
	assert.Equal(t, 5, received1)

	b.RemoveListener(l0)

	b.Notify(6)
	assert.Equal(t, 5, received0)
	assert.Equal(t, 6, received1)

	b.RemoveListener(l1)

	b.Notify(7)
	assert.Equal(t, 5, received0)
	assert.Equal(t, 6, received1)

	b.AddListener(l0)

	b.Notify(7)
	assert.Equal(t, 7, received0)
	assert.Equal(t, 6, received1)
}
