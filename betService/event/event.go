package event

type EventManager struct{
	EventPublisher
	IncomingEventProcessor
	EventSubscriber
}


func newEventManager(publisher EventPublisher, processor IncomingEventProcessor,subscriber EventSubscriber) *EventManager{
	return &EventManager{
		EventPublisher: publisher,
		IncomingEventProcessor: processor,
		EventSubscriber: subscriber,
	}
}