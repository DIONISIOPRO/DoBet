package event

type EventManger struct{
	EventPublisher
	IncomingEventProcessor
	EventSubscriber
}


func newEventManager(publisher EventPublisher, processor IncomingEventProcessor,subscriber EventSubscriber) EventManger{
	return EventManger{
		EventPublisher: publisher,
		IncomingEventProcessor: processor,
		EventSubscriber: subscriber,
	}
}