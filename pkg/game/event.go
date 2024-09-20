package game

const (
	NewGeneration = iota
	GamePaused
	GameResumed
)

type EventDispatcher interface {
	Dispatch(event int)
	Subscribe(c chan int)
}

type eventDispatcher struct {
	channels []chan int
}

func (e *eventDispatcher) Subscribe(c chan int) {
	e.channels = append(e.channels, c)
}

func NewEventDispatcher() *eventDispatcher {
	return &eventDispatcher{
		channels: make([]chan int, 0),
	}
}

func (d *eventDispatcher) Dispatch(event int) {
	for _, ch := range d.channels {
		ch <- event
	}
}
