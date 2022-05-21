package event

type Listenner interface {
	Listenning()
}
type EventListennerManager struct {
	publisher  EventPublisher
	Listenners []Listenner
}

func (e EventListennerManager) Listenning() {
	for _, l := range e.Listenners {
		go l.Listenning()
	}
}

func (e *EventListennerManager) AddListenner(listenner Listenner) {
	e.Listenners = append(e.Listenners, listenner)
}

func NewEventListennersManager() *EventListennerManager {
	return &EventListennerManager{}
}
