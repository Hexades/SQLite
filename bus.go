package sqlite

var localbus = sqliteBus{}

type sqliteBus struct {
	repositoryListenerChannels []chan Event
}

func addRepositoryListener(EventListener EventListener) {
	listenerChannel := make(chan Event, 10)
	localbus.repositoryListenerChannels = append(localbus.repositoryListenerChannels, listenerChannel)
	go EventListener.onEvent(listenerChannel)
}

func sendEvent(Event Event) {
	for _, channel := range localbus.repositoryListenerChannels {
		channel <- Event
	}
}

type EventListener interface {
	onEvent(Event <-chan Event)
}
