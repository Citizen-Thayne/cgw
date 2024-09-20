package game

type mockEventDispatcher struct {
	dispatchEvents []int
	subscribers    []chan int
}

func createMockEventDispatcher() *mockEventDispatcher {
	dispatcher := &mockEventDispatcher{
		dispatchEvents: make([]int, 0),
		subscribers:    make([]chan int, 0),
	}
	return dispatcher
}

func (m *mockEventDispatcher) Dispatch(event int) {
	m.dispatchEvents = append(m.dispatchEvents, event)
}

func (m *mockEventDispatcher) Subscribe(c chan int) {
	m.subscribers = append(m.subscribers, c)
}

func (m *mockEventDispatcher) reset() {
	m.dispatchEvents = make([]int, 0)
	m.subscribers = make([]chan int, 0)
}
