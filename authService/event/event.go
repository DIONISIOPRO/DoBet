package event

type EventManger struct{
	EventPublisher
	IncomingEventProcessor
	EventQueuecreator
	EventSubscriber
}


func newEventManager(publisher EventPublisher, processor IncomingEventProcessor, creator EventQueuecreator, subscriber EventSubscriber) EventManger{
	return EventManger{
		EventPublisher: publisher,
		IncomingEventProcessor: processor,
		EventQueuecreator: creator,
		EventSubscriber: subscriber,
	}
}