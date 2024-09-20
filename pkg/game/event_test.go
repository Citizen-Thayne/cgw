package game

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEventDispatcher(t *testing.T) {
	dispatcher := NewEventDispatcher()

	channel1 := make(chan int, 2)
	channel2 := make(chan int, 2)

	dispatcher.Subscribe(channel1)
	dispatcher.Subscribe(channel2)

	assert.Equal(t, 2, len(dispatcher.channels))

	var wg sync.WaitGroup
	done := make(chan struct{})

	wg.Add(1)
	go func() {
		defer wg.Done()
		assert.Equal(t, NewGeneration, <-channel1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		assert.Equal(t, NewGeneration, <-channel2)
	}()

	dispatcher.Dispatch(NewGeneration)

	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Error("Test timeout")
	}
}
