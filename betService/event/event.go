package event

type Listenner interface {
	Listenning()
}
type EventManager struct {
	publisher  EventPublisher
	Listenners []Listenner
}

func (e EventManager) Listenning() {
	for _, l := range e.Listenners {
		go l.Listenning()
	}
}

func (e *EventManager) AddListenner(listenner Listenner) {
	e.Listenners = append(e.Listenners, listenner)
}

func newEventManager(publisher EventPublisher) *EventManager {
	return &EventManager{
		publisher: publisher,
	}
}
