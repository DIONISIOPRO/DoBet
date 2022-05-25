package event

func NewAuthEventListenner() *AuthEventListnner {
	return &AuthEventListnner{
		Listenners: []Listenner{},
	}
}

func NewUserCreatedEventListenner(subsccriber EventSubscriber, service CreateUserEventProcessor) *UserCreatedListenner {
	return &UserCreatedListenner{
		Subscriber: subsccriber,
		service:    service,
	}
}

func NewUseruUpdatedEventListenner(subsccriber EventSubscriber, service UpdateUserEventProcessor) *UserUpdatedListenner {
	return &UserUpdatedListenner{
		Subscriber: subsccriber,
		service:    service,
	}
}

func NewUseruDeletedEventListenner(subsccriber EventSubscriber, service DeleteUserEventProcessor) *UserDeletedListenner {
	return &UserDeletedListenner{
		Subscriber: subsccriber,
		service:    service,
	}
}
