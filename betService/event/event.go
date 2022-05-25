package event

type Listenner interface {
	Listenning(done <-chan bool)
}
type EventListennerManager struct {
	publisher  EventPublisher
	Listenners []Listenner
}

func (e EventListennerManager) Listenning(done <-chan bool) {
	for _, l := range e.Listenners {
		go l.Listenning(done)
	}
}

func (e *EventListennerManager) AddListenner(listenner Listenner) {
	e.Listenners = append(e.Listenners, listenner)
}

func NewEventListennersManager(publisher  EventPublisher) *EventListennerManager {
	return &EventListennerManager{
		publisher: publisher,
		Listenners: []Listenner{},
	}
}
