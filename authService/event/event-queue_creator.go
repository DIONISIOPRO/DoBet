package event

import "github.com/streadway/amqp"

type EventQueuecreator struct {
	PublishingChannel *amqp.Channel
}

func NewQueueCreator(PublishingChannel *amqp.Channel) EventQueuecreator {
	return EventQueuecreator{
		PublishingChannel: PublishingChannel,
	}
}

func (manager *EventQueuecreator) CreateQueues(queues []string) (err error) {
	for _, queueName := range queues {
		_, err = manager.PublishingChannel.QueueDeclare(
			queueName,
			true,
			false,
			false,
			false,
			nil,
		)

	}
	return err
}
